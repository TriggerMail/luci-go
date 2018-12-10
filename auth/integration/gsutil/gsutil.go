// Copyright 2017 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package gsutil implements a hacky shim that makes gsutil use LUCI local auth.
//
// It constructs a special .boto config file that instructs gsutil to use local
// HTTP endpoint as token_uri (it's the one that exchanges OAuth2 refresh token
// for an access token). This endpoint is implemented on top of LUCI auth.
//
// Thus gsutil thinks it's using 3-legged OAuth2 flow, while in fact it is
// getting the token through LUCI protocols.
package gsutil

import (
	"bytes"
	"context"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"
	"time"

	"golang.org/x/oauth2"

	"github.com/TriggerMail/luci-go/common/clock"
	"github.com/TriggerMail/luci-go/common/data/rand/cryptorand"
	"github.com/TriggerMail/luci-go/common/errors"
	"github.com/TriggerMail/luci-go/common/logging"
	"github.com/TriggerMail/luci-go/common/retry/transient"
	"github.com/TriggerMail/luci-go/common/runtime/paniccatcher"

	"github.com/TriggerMail/luci-go/auth/integration/internal/localsrv"
)

// Server runs a local server that handles requests to token_uri.
//
// It also manages a directory with gsutil state, since part of the state is
// the cached OAuth2 token that we don't want to put into default global
// ~/.gsutil state directory.
type Server struct {
	// Source is used to obtain OAuth2 tokens.
	Source oauth2.TokenSource
	// StateDir is where to drop new .boto file and where to keep gsutil state.
	StateDir string
	// Port is a local TCP port to bind to or 0 to allow the OS to pick one.
	Port int

	srv localsrv.Server
}

// Start launches background goroutine with the serving loop and prepares .boto.
//
// Returns absolute path to new .boto file. It is always inside StateDir. Caller
// is responsible for creating StateDir (and later deleting it, if necessary).
//
// The provided context is used as base context for request handlers and for
// logging. The server must be eventually stopped with Stop().
func (s *Server) Start(ctx context.Context) (botoCfg string, err error) {
	// The secret will be used as fake refresh token, to verify clients of the
	// protocol have read access to .boto file (where this secret is stored).
	blob := make([]byte, 48)
	if _, err := cryptorand.Read(ctx, blob); err != nil {
		return "", errors.Annotate(err, "failed to read random bytes").Err()
	}
	secret := base64.RawStdEncoding.EncodeToString(blob)

	// Launch the server to get the port number.
	addr, err := s.srv.Start(ctx, "gsutil-auth", s.Port, func(c context.Context, l net.Listener, wg *sync.WaitGroup) error {
		return s.serve(c, l, wg, secret)
	})
	if err != nil {
		return "", errors.Annotate(err, "failed to start the server").Err()
	}
	defer func() {
		if err != nil {
			s.srv.Stop(ctx)
		}
	}()

	// Prepare new .boto in memory (to simplify IO error handling below).
	args := map[string]string{
		"StateDir":     s.StateDir,
		"RefreshToken": secret,
		"AuthURI":      fmt.Sprintf("http://%s/gsutil/authorization", addr),
		"TokenURI":     fmt.Sprintf("http://%s/gsutil/token", addr),
	}
	buf := bytes.Buffer{}
	if err := botoCfgTempl.Execute(&buf, args); err != nil {
		return "", errors.Annotate(err, "failed to generate .boto").Err()
	}

	// Drop it to the disk.
	botoCfg = filepath.Join(s.StateDir, ".boto")
	if err := ioutil.WriteFile(botoCfg, buf.Bytes(), 0600); err != nil {
		return "", errors.Annotate(err, "failed to write %s", botoCfg).Err()
	}

	// Make sure the credentials cache file is empty, otherwise it will grow
	// after each server launch, since it uses refresh_token (which we generate
	// randomly) as cache key. We don't really need this cache anyway.
	if err := os.Remove(filepath.Join(s.StateDir, "credstore")); err != nil && !os.IsNotExist(err) {
		return "", errors.Annotate(err, "failed to remove credstore").Err()
	}

	return botoCfg, nil
}

// Stop closes the listening socket, notifies pending requests to abort and
// stops the internal serving goroutine.
//
// Safe to call multiple times. Once stopped, the server cannot be started again
// (make a new instance of Server instead).
//
// Uses the given context for the deadline when waiting for the serving loop
// to stop.
func (s *Server) Stop(ctx context.Context) error {
	return s.srv.Stop(ctx)
}

////////////////////////////////////////////////////////////////////////////////

var botoCfgTempl = template.Must(template.New(".boto").Parse(`# Autogenerated

[GSUtil]
software_update_check_period = 0
state_dir = {{.StateDir}}

[Credentials]
gs_oauth2_refresh_token = {{.RefreshToken}}

[OAuth2]
provider_label = LUCI Local
provider_authorization_uri = {{.AuthURI}}
provider_token_uri = {{.TokenURI}}
`))

// serve runs the serving loop.
func (s *Server) serve(ctx context.Context, l net.Listener, wg *sync.WaitGroup, secret string) error {
	mux := http.NewServeMux()

	mux.Handle("/gsutil/authorization", &handler{ctx, wg, func(rw http.ResponseWriter, r *http.Request) {
		// Authorization URI is normally used during interactive login to eventually
		// generate a refresh token. Since we pass the refresh token in .boto
		// already, it must never be called.
		rw.WriteHeader(http.StatusNotImplemented)
	}})

	mux.Handle("/gsutil/token", &handler{ctx, wg, func(rw http.ResponseWriter, r *http.Request) {
		err := s.handleTokenRequest(rw, r, secret)

		code := 0
		msg := ""
		if transient.Tag.In(err) {
			code = http.StatusInternalServerError
			msg = fmt.Sprintf("Transient error - %s", err)
		} else if err != nil {
			code = http.StatusBadRequest
			msg = fmt.Sprintf("Bad request - %s", err)
		}

		if code != 0 {
			logging.Errorf(ctx, "%s", msg)
			http.Error(rw, msg, code)
		}
	}})

	srv := http.Server{Handler: mux}
	return srv.Serve(l)
}

// handleTokenRequest handles /token call.
//
// The body of the request is documented here (among many other places):
//   https://developers.google.com/identity/protocols/OAuth2InstalledApp#offline
//
// We ignore client_id and client_secret, since we aren't really running OAuth2.
func (s *Server) handleTokenRequest(rw http.ResponseWriter, r *http.Request, secret string) error {
	ctx := r.Context()

	// We support only refreshing access token via 'refresh_token' grant.
	if r.PostFormValue("grant_type") != "refresh_token" {
		return fmt.Errorf("expecting 'refresh_token' grant type")
	}

	// The token must match whatever we passed to gsutil via .boto. Unfortunately,
	// gcloud's gsutil wrapper overrides the refresh token in the config unless
	// 'pass_credentials_to_gsutil' is set to false via
	//
	// $ gcloud config set pass_credentials_to_gsutil false
	//
	// So hint the user to set it.
	passedToken := r.PostFormValue("refresh_token")
	if subtle.ConstantTimeCompare([]byte(passedToken), []byte(secret)) != 1 {
		return fmt.Errorf("wrong refresh_token (if running via gcloud, set pass_credentials_to_gsutil = false)")
	}

	// Good enough. Grab an access token through the source and return it.
	tok, err := s.Source.Token()
	if err != nil {
		return err
	}
	rw.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(rw).Encode(map[string]interface{}{
		"access_token": tok.AccessToken,
		"expires_in":   clock.Until(ctx, tok.Expiry) / time.Second,
		"token_type":   "Bearer",
	})
}

// handler implements http.Handler by wrapping the given handler with some
// housekeeping stuff.
type handler struct {
	ctx     context.Context
	wg      *sync.WaitGroup
	handler http.HandlerFunc
}

func (h *handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.wg.Add(1)
	defer h.wg.Done()

	defer paniccatcher.Catch(func(p *paniccatcher.Panic) {
		logging.Fields{
			"panic.error": p.Reason,
		}.Errorf(h.ctx, "Caught panic during handling of %q: %s\n%s", r.RequestURI, p.Reason, p.Stack)
		http.Error(rw, "Internal Server Error. See logs.", http.StatusInternalServerError)
	})

	logging.Debugf(h.ctx, "Handling %s %s", r.Method, r.RequestURI)
	h.handler(rw, r.WithContext(h.ctx))
}
