// Copyright 2015 The LUCI Authors.
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

package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/TriggerMail/luci-go/common/errors"
	"github.com/TriggerMail/luci-go/common/logging"
	"github.com/TriggerMail/luci-go/common/retry/transient"

	"github.com/TriggerMail/luci-go/auth/identity"
	"github.com/TriggerMail/luci-go/server/auth/delegation"
	"github.com/TriggerMail/luci-go/server/auth/signing"
	"github.com/TriggerMail/luci-go/server/router"
)

var (
	// ErrNotConfigured is returned by Authenticate if auth library wasn't
	// properly initialized (see SetConfig).
	ErrNotConfigured = errors.New("auth: the library is not properly configured")

	// ErrNoUsersAPI is returned by LoginURL and LogoutURL if none of
	// the authentication methods support UsersAPI.
	ErrNoUsersAPI = errors.New("auth: methods do not support login or logout URL")

	// ErrBadClientID is returned by Authenticate if caller is using
	// non-whitelisted OAuth2 client. More info is in the log.
	ErrBadClientID = errors.New("auth: OAuth client_id is not whitelisted")

	// ErrIPNotWhitelisted is returned when an account is restricted by an IP
	// whitelist and request's remote_addr is not in it.
	ErrIPNotWhitelisted = errors.New("auth: IP is not whitelisted")
)

// Method implements a particular kind of low-level authentication mechanism.
//
// It may also optionally implement UsersAPI (if the method support login and
// logout URLs).
//
// Methods are not usually used directly, but passed to Authenticator{...} that
// knows how to apply them.
type Method interface {
	// Authenticate extracts user information from the incoming request.
	//
	// It returns:
	//   * (*User, nil) on success.
	//   * (nil, nil) if the method is not applicable.
	//   * (nil, error) if the method is applicable, but credentials are invalid.
	Authenticate(context.Context, *http.Request) (*User, error)
}

// UsersAPI may be additionally implemented by Method if it supports login and
// logout URLs.
type UsersAPI interface {
	// LoginURL returns a URL that, when visited, prompts the user to sign in,
	// then redirects the user to the URL specified by dest.
	LoginURL(c context.Context, dest string) (string, error)

	// LogoutURL returns a URL that, when visited, signs the user out,
	// then redirects the user to the URL specified by dest.
	LogoutURL(c context.Context, dest string) (string, error)
}

// User represents identity and profile of a user.
type User struct {
	// Identity is identity string of the user (may be AnonymousIdentity).
	// If User is returned by Authenticate(...), Identity string is always present
	// and valid.
	Identity identity.Identity `json:"identity,omitempty"`

	// Superuser is true if the user is site-level administrator. For example, on
	// GAE this bit is set for GAE-level administrators. Optional, default false.
	Superuser bool `json:"superuser,omitempty"`

	// Email is email of the user. Optional, default "". Don't use it as a key
	// in various structures. Prefer to use Identity() instead (it is always
	// available).
	Email string `json:"email,omitempty"`

	// Name is full name of the user. Optional, default "".
	Name string `json:"name,omitempty"`

	// Picture is URL of the user avatar. Optional, default "".
	Picture string `json:"picture,omitempty"`

	// ClientID is the ID of the pre-registered OAuth2 client so its identity can
	// be verified. Used only by authentication methods based on OAuth2.
	// See https://developers.google.com/console/help/#generatingoauth2 for more.
	ClientID string `json:"client_id,omitempty"`
}

// Authenticator performs authentication of incoming requests.
//
// It is a stateless object configured with a list of methods to try when
// authenticating incoming requests. It implements Authenticate method that
// performs high-level authentication logic using the provided list of low-level
// auth methods.
//
// Note that most likely you don't need to instantiate this object directly.
// Use Authenticate middleware instead. Authenticator is exposed publicly only
// to be used in advanced cases, when you need to fine-tune authentication
// behavior.
type Authenticator struct {
	Methods []Method // a list of authentication methods to try
}

// GetMiddleware returns a middleware that uses this Authenticator for
// authentication.
//
// It uses a.Authenticate internally and handles errors appropriately.
func (a *Authenticator) GetMiddleware() router.Middleware {
	return func(c *router.Context, next router.Handler) {
		ctx, err := a.Authenticate(c.Context, c.Request)
		switch {
		case transient.Tag.In(err):
			replyError(c.Context, c.Writer, 500, "Transient error during authentication", err)
		case err != nil:
			replyError(c.Context, c.Writer, 401, "Authentication error", err)
		default:
			c.Context = ctx
			next(c)
		}
	}
}

// Authenticate authenticates the requests and adds State into the context.
//
// Returns an error if credentials are provided, but invalid. If no credentials
// are provided (i.e. the request is anonymous), finishes successfully, but in
// that case State.Identity() returns AnonymousIdentity.
func (a *Authenticator) Authenticate(c context.Context, r *http.Request) (context.Context, error) {
	report := durationReporter(c, authenticateDuration)

	// We will need working DB factory below to check IP whitelist.
	cfg := getConfig(c)
	if cfg == nil || cfg.DBProvider == nil || len(a.Methods) == 0 {
		report(ErrNotConfigured, "ERROR_NOT_CONFIGURED")
		return nil, ErrNotConfigured
	}

	// Pick first authentication method that applies.
	s := state{authenticator: a}
	for _, m := range a.Methods {
		var err error
		s.user, err = m.Authenticate(c, r)
		if err != nil {
			report(err, "ERROR_BROKEN_CREDS") // e.g. malformed OAuth token
			return nil, err
		}
		if s.user != nil {
			if err = s.user.Identity.Validate(); err != nil {
				report(err, "ERROR_BROKEN_IDENTITY") // a weird looking email address
				return nil, err
			}
			s.method = m
			break
		}
	}

	// If no authentication method is applicable, default to anonymous identity.
	if s.method == nil {
		s.user = &User{Identity: identity.AnonymousIdentity}
	}

	var err error
	s.peerIP, err = parseRemoteIP(r.RemoteAddr)
	if err != nil {
		panic(fmt.Errorf("auth: bad remote_addr: %v", err))
	}

	// Grab a snapshot of auth DB to use consistently for the duration of this
	// request.
	s.db, err = cfg.DBProvider(c)
	if err != nil {
		report(ErrNotConfigured, "ERROR_NOT_CONFIGURED")
		return nil, ErrNotConfigured
	}

	// If using OAuth2, make sure ClientID is whitelisted.
	if s.user.ClientID != "" {
		valid, err := s.db.IsAllowedOAuthClientID(c, s.user.Email, s.user.ClientID)
		if err != nil {
			report(err, "ERROR_TRANSIENT_IN_OAUTH_WHITELIST")
			return nil, err
		}
		if !valid {
			logging.Warningf(
				c, "auth: %q is using client_id %q not in the whitelist",
				s.user.Email, s.user.ClientID)
			report(ErrBadClientID, "ERROR_FORBIDDEN_OAUTH_CLIENT")
			return nil, ErrBadClientID
		}
	}

	// Some callers may be constrained by an IP whitelist.
	switch ipWhitelist, err := s.db.GetWhitelistForIdentity(c, s.user.Identity); {
	case err != nil:
		report(err, "ERROR_TRANSIENT_IN_IP_WHITELIST")
		return nil, err
	case ipWhitelist != "":
		switch whitelisted, err := s.db.IsInWhitelist(c, s.peerIP, ipWhitelist); {
		case err != nil:
			report(err, "ERROR_TRANSIENT_IN_IP_WHITELIST")
			return nil, err
		case !whitelisted:
			report(ErrIPNotWhitelisted, "ERROR_FORBIDDEN_IP")
			return nil, ErrIPNotWhitelisted
		}
	}

	// peerIdent always matches the identity of a remote peer. It may be different
	// from s.user.Identity if the delegation is used (see below).
	s.peerIdent = s.user.Identity

	// Check the delegation token. This is LUCI-specific authentication protocol.
	// Delegation tokens are generated by the central auth service (see luci-py's
	// auth_service) and validated by checking their RSA signature using auth
	// server's public keys.
	delegationTok := r.Header.Get(delegation.HTTPHeaderName)
	if delegationTok != "" {
		// Log the token fingerprint (even before parsing the token), it can be used
		// to grab the info about the token from the token server logs.
		logging.Fields{
			"fingerprint": tokenFingerprint(delegationTok),
		}.Debugf(c, "auth: Received delegation token")

		// Need to grab our own identity to verify that the delegation token is
		// minted for consumption by us and not some other service.
		ownServiceIdentity, err := getOwnServiceIdentity(c, cfg.Signer)
		if err != nil {
			report(err, "ERROR_TRANSIENT_IN_OWN_IDENTITY")
			return nil, err
		}
		delegatedIdentity, err := delegation.CheckToken(c, delegation.CheckTokenParams{
			Token:                delegationTok,
			PeerID:               s.peerIdent,
			CertificatesProvider: s.db,
			GroupsChecker:        s.db,
			OwnServiceIdentity:   ownServiceIdentity,
		})
		if err != nil {
			if transient.Tag.In(err) {
				report(err, "ERROR_TRANSIENT_IN_TOKEN_CHECK")
			} else {
				report(err, "ERROR_BAD_DELEGATION_TOKEN")
			}
			return nil, err
		}

		// User profile information is not available when using delegation, so just
		// wipe it.
		s.user = &User{Identity: delegatedIdentity}

		// Log that 'peerIdent' is pretending to be 'delegatedIdentity'.
		logging.Fields{
			"peerID":      s.peerIdent,
			"delegatedID": delegatedIdentity,
		}.Debugf(c, "auth: Using delegation")
	}

	// Inject auth state.
	report(nil, "SUCCESS")
	return WithState(c, &s), nil
}

// usersAPI returns implementation of UsersAPI by examining Methods.
//
// Returns nil if none of Methods implement UsersAPI.
func (a *Authenticator) usersAPI() UsersAPI {
	for _, m := range a.Methods {
		if api, ok := m.(UsersAPI); ok {
			return api
		}
	}
	return nil
}

// LoginURL returns a URL that, when visited, prompts the user to sign in,
// then redirects the user to the URL specified by dest.
//
// Returns ErrNoUsersAPI if none of the authentication methods support login
// URLs.
func (a *Authenticator) LoginURL(c context.Context, dest string) (string, error) {
	if api := a.usersAPI(); api != nil {
		return api.LoginURL(c, dest)
	}
	return "", ErrNoUsersAPI
}

// LogoutURL returns a URL that, when visited, signs the user out, then
// redirects the user to the URL specified by dest.
//
// Returns ErrNoUsersAPI if none of the authentication methods support login
// URLs.
func (a *Authenticator) LogoutURL(c context.Context, dest string) (string, error) {
	if api := a.usersAPI(); api != nil {
		return api.LogoutURL(c, dest)
	}
	return "", ErrNoUsersAPI
}

////

// replyError logs the error and writes it to ResponseWriter.
func replyError(c context.Context, rw http.ResponseWriter, code int, msg string, err error) {
	logging.WithError(err).Errorf(c, "HTTP %d: %s", code, msg)
	http.Error(rw, msg, code)
}

// getOwnServiceIdentity returns 'service:<appID>' identity of the current
// service.
func getOwnServiceIdentity(c context.Context, signer signing.Signer) (identity.Identity, error) {
	if signer == nil {
		return "", ErrNotConfigured
	}
	serviceInfo, err := signer.ServiceInfo(c)
	if err != nil {
		return "", err
	}
	return identity.MakeIdentity("service:" + serviceInfo.AppID)
}
