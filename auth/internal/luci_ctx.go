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

package internal

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/context/ctxhttp"
	"golang.org/x/oauth2"

	"github.com/TriggerMail/luci-go/auth/integration/localauth/rpcs"
	"github.com/TriggerMail/luci-go/common/logging"
	"github.com/TriggerMail/luci-go/common/retry/transient"
	"github.com/TriggerMail/luci-go/lucictx"
)

type luciContextTokenProvider struct {
	localAuth *lucictx.LocalAuth
	email     string // an email or NoEmail
	scopes    []string
	transport http.RoundTripper
	cacheKey  CacheKey // used only for in-memory cache
}

// NewLUCIContextTokenProvider returns TokenProvider that knows how to use a
// local auth server to mint tokens.
//
// It requires LUCI_CONTEXT["local_auth"] to be present in the 'ctx'. It's a
// description of how to locate and contact the local auth server.
//
// See auth/integration/localauth package for the implementation of the server.
func NewLUCIContextTokenProvider(ctx context.Context, scopes []string, transport http.RoundTripper) (TokenProvider, error) {
	localAuth := lucictx.GetLocalAuth(ctx)
	switch {
	case localAuth == nil:
		return nil, fmt.Errorf(`no "local_auth" in LUCI_CONTEXT`)
	case localAuth.DefaultAccountID == "":
		return nil, fmt.Errorf(`no "default_account_id" in LUCI_CONTEXT["local_auth"]`)
	}

	// Grab an email associated with default account, if any.
	email := NoEmail
	for _, account := range localAuth.Accounts {
		if account.ID == localAuth.DefaultAccountID {
			// Previous protocol version didn't expose the email, so keep the value
			// as NoEmail in this case. This should be rare.
			if account.Email != "" {
				email = account.Email
			}
			break
		}
	}

	// All authenticators share singleton in-process token cache, see
	// ProcTokenCache variable in proc_cache.go.
	//
	// It is possible (though very unusual), for a single process to use multiple
	// local auth servers (e.g if it enters a subcontext with another "local_auth"
	// value).
	//
	// For these reasons we use a digest of localAuth parameters as a cache key.
	// It is used only in the process-local cache, the token never ends up in
	// the disk cache, as indicated by Lightweight() returning true (tokens from
	// such providers aren't cached on disk by Authenticator).
	blob, err := json.Marshal(localAuth)
	if err != nil {
		return nil, err
	}
	digest := sha256.Sum256(blob)

	return &luciContextTokenProvider{
		localAuth: localAuth,
		email:     email,
		scopes:    scopes,
		transport: transport,
		cacheKey: CacheKey{
			Key:    fmt.Sprintf("luci_ctx/%s", hex.EncodeToString(digest[:])),
			Scopes: scopes,
		},
	}, nil
}

func (p *luciContextTokenProvider) RequiresInteraction() bool {
	return false
}

func (p *luciContextTokenProvider) Lightweight() bool {
	return true
}

func (p *luciContextTokenProvider) Email() string {
	return p.email
}

func (p *luciContextTokenProvider) CacheKey(ctx context.Context) (*CacheKey, error) {
	return &p.cacheKey, nil
}

func (p *luciContextTokenProvider) MintToken(ctx context.Context, base *Token) (*Token, error) {
	// Note: deadlines and retries are implemented by Authenticator. MintToken
	// should just make a single attempt, and mark an error as transient to
	// trigger a retry, if necessary.
	request := rpcs.GetOAuthTokenRequest{
		Scopes:    p.scopes,
		Secret:    p.localAuth.Secret,
		AccountID: p.localAuth.DefaultAccountID,
	}
	if err := request.Validate(); err != nil {
		return nil, err // should not really happen
	}
	body, err := json.Marshal(&request)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("http://127.0.0.1:%d/rpc/LuciLocalAuthService.GetOAuthToken", p.localAuth.RPCPort)
	logging.Debugf(ctx, "POST %s", url)
	httpReq, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	httpResp, err := ctxhttp.Do(ctx, &http.Client{Transport: p.transport}, httpReq)
	if err != nil {
		return nil, transient.Tag.Apply(err)
	}
	defer httpResp.Body.Close()
	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, transient.Tag.Apply(err)
	}

	if httpResp.StatusCode != 200 {
		err := fmt.Errorf("local auth - HTTP %d: %s", httpResp.StatusCode, strings.TrimSpace(string(respBody)))
		if httpResp.StatusCode >= 500 {
			return nil, transient.Tag.Apply(err)
		}
		return nil, err
	}

	response := rpcs.GetOAuthTokenResponse{}
	if err := json.Unmarshal(respBody, &response); err != nil {
		return nil, err
	}
	if response.ErrorMessage != "" || response.ErrorCode != 0 {
		msg := response.ErrorMessage
		if msg == "" {
			msg = "unknown error"
		}
		return nil, fmt.Errorf("local auth - RPC code %d: %s", response.ErrorCode, msg)
	}

	return &Token{
		Token: oauth2.Token{
			AccessToken: response.AccessToken,
			Expiry:      time.Unix(response.Expiry, 0).UTC(),
			TokenType:   "Bearer",
		},
		Email: p.Email(),
	}, nil
}

func (p *luciContextTokenProvider) RefreshToken(ctx context.Context, prev, base *Token) (*Token, error) {
	// Minting and refreshing is the same thing: a call to local auth server.
	return p.MintToken(ctx, base)
}
