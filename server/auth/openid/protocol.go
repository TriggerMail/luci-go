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

package openid

import (
	"context"
	"errors"
	"net/url"
	"strings"
	"time"

	"github.com/TriggerMail/luci-go/auth/identity"
	"github.com/TriggerMail/luci-go/server/auth"
	"github.com/TriggerMail/luci-go/server/auth/internal"
	"github.com/TriggerMail/luci-go/server/caching"
	"github.com/TriggerMail/luci-go/server/tokens"
)

var (
	discoveryDocCache = caching.RegisterLRUCache(8) // URL string => *discoveryDoc
	signingKeysCache  = caching.RegisterLRUCache(8) // URL string => *JSONWebKeySet
)

// openIDStateToken is used to generate `state` parameter used in OpenID flow to
// pass state between our app and authentication backend.
var openIDStateToken = tokens.TokenKind{
	Algo:       tokens.TokenAlgoHmacSHA256,
	Expiration: 30 * time.Minute,
	SecretKey:  "openid_state_token",
	Version:    1,
}

// discoveryDoc describes subset of OpenID Discovery JSON document.
// See https://developers.google.com/identity/protocols/OpenIDConnect#discovery.
type discoveryDoc struct {
	Issuer                string `json:"issuer"`
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	JwksURI               string `json:"jwks_uri"`
}

// signingKeys returns a JSON Web Key set fetched from the location specified
// in the discovery document.
//
// It fetches them on the first use and then keeps them cached in the process
// cache for 6h.
//
// May return both fatal and transient errors.
func (d *discoveryDoc) signingKeys(c context.Context) (*JSONWebKeySet, error) {
	fetcher := func() (interface{}, time.Duration, error) {
		raw := &JSONWebKeySetStruct{}
		req := internal.Request{
			Method: "GET",
			URL:    d.JwksURI,
			Out:    raw,
		}
		if err := req.Do(c); err != nil {
			return nil, 0, err
		}
		keys, err := NewJSONWebKeySet(raw)
		if err != nil {
			return nil, 0, err
		}
		return keys, time.Hour * 6, nil
	}

	cached, err := signingKeysCache.LRU(c).GetOrCreate(c, d.JwksURI, fetcher)
	if err != nil {
		return nil, err
	}
	return cached.(*JSONWebKeySet), nil
}

// fetchDiscoveryDoc fetches discovery document from given URL. It is cached in
// the process cache for 24 hours.
func fetchDiscoveryDoc(c context.Context, url string) (*discoveryDoc, error) {
	if url == "" {
		return nil, ErrNotConfigured
	}

	fetcher := func() (interface{}, time.Duration, error) {
		doc := &discoveryDoc{}
		req := internal.Request{
			Method: "GET",
			URL:    url,
			Out:    doc,
		}
		if err := req.Do(c); err != nil {
			return nil, 0, err
		}
		return doc, time.Hour * 24, nil
	}

	// Cache the document in the process cache.
	cached, err := discoveryDocCache.LRU(c).GetOrCreate(c, url, fetcher)
	if err != nil {
		return nil, err
	}
	return cached.(*discoveryDoc), nil
}

// authenticationURI returns an URI to redirect a user to in order to
// authenticate via OpenID.
//
// This is step 1 of the authentication flow. Generate authentication URL and
// redirect user's browser to it. After consent screen, redirect_uri will be
// called (via user's browser) with `state` and authorization code passed to it,
// eventually resulting in a call to 'handle_authorization_code'.
func authenticationURI(c context.Context, cfg *Settings, state map[string]string) (string, error) {
	if cfg.ClientID == "" || cfg.RedirectURI == "" {
		return "", ErrNotConfigured
	}

	// Grab authorization URL from discovery doc.
	discovery, err := fetchDiscoveryDoc(c, cfg.DiscoveryURL)
	if err != nil {
		return "", err
	}
	if discovery.AuthorizationEndpoint == "" {
		return "", errors.New("openid: bad discovery doc, empty authorization_endpoint")
	}

	// Wrap state into HMAC-protected token.
	stateTok, err := openIDStateToken.Generate(c, nil, state, 0)
	if err != nil {
		return "", err
	}

	// Generate final URL.
	v := url.Values{}
	v.Set("client_id", cfg.ClientID)
	v.Set("redirect_uri", cfg.RedirectURI)
	v.Set("response_type", "code")
	v.Set("scope", "openid email profile")
	v.Set("prompt", "select_account")
	v.Set("state", stateTok)
	return discovery.AuthorizationEndpoint + "?" + v.Encode(), nil
}

// validateStateToken validates 'state' token passed to redirect_uri. Returns
// whatever `state` was passed to authenticationURI.
func validateStateToken(c context.Context, stateTok string) (map[string]string, error) {
	return openIDStateToken.Validate(c, stateTok, nil)
}

// handleAuthorizationCode exchange `code` for user ID token and user profile.
func handleAuthorizationCode(c context.Context, cfg *Settings, code string) (uid string, u *auth.User, err error) {
	if cfg.ClientID == "" || cfg.ClientSecret == "" || cfg.RedirectURI == "" {
		return "", nil, ErrNotConfigured
	}

	// Validate the discover doc has necessary fields to proceed.
	discovery, err := fetchDiscoveryDoc(c, cfg.DiscoveryURL)
	switch {
	case err != nil:
		return "", nil, err
	case discovery.TokenEndpoint == "":
		return "", nil, errors.New("openid: bad discovery doc, empty token_endpoint")
	}

	// Prepare a request to exchange authorization code for the ID token.
	v := url.Values{}
	v.Set("code", code)
	v.Set("client_id", cfg.ClientID)
	v.Set("client_secret", cfg.ClientSecret)
	v.Set("redirect_uri", cfg.RedirectURI)
	v.Set("grant_type", "authorization_code")
	payload := v.Encode()

	// Send POST to the token endpoint with URL-encoded parameters to get back the
	// ID token. There's more stuff in the reply, we don't need it.
	var token struct {
		IDToken string `json:"id_token"`
	}
	req := internal.Request{
		Method: "POST",
		URL:    discovery.TokenEndpoint,
		Body:   []byte(payload),
		Headers: map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		},
		Out: &token,
	}
	if err := req.Do(c); err != nil {
		return "", nil, err
	}

	// Unpack the ID token to grab the user information from it.
	return userFromIDToken(c, token.IDToken, cfg, discovery)
}

// userFromIDToken validates the ID token and extracts user information from it.
func userFromIDToken(c context.Context, token string, cfg *Settings, discovery *discoveryDoc) (uid string, u *auth.User, err error) {
	// Validate the discovery doc has necessary fields to proceed.
	switch {
	case discovery.Issuer == "":
		return "", nil, errors.New("openid: bad discovery doc, empty issuer")
	case discovery.JwksURI == "":
		return "", nil, errors.New("openid: bad discovery doc, empty jwks_uri")
	}

	// Grab the signing keys needed to verify the token. This is almost always
	// hitting the local process cache and thus must be fast.
	signingKeys, err := discovery.signingKeys(c)
	if err != nil {
		return "", nil, err
	}

	// Unpack the ID token to grab the user information from it.
	verifiedToken, err := VerifyIDToken(c, token, signingKeys, discovery.Issuer, cfg.ClientID)
	if err != nil {
		return "", nil, err
	}

	// Ignore non https:// URLs for pictures. We serve all pages over HTTPS and
	// don't want to break this rule just for a pretty picture.
	picture := verifiedToken.Picture
	if picture != "" && !strings.HasPrefix(picture, "https://") {
		picture = ""
	}

	// Build the identity string from the email. This essentially validates it.
	id, err := identity.MakeIdentity("user:" + verifiedToken.Email)
	if err != nil {
		return "", nil, err
	}

	return verifiedToken.Sub, &auth.User{
		Identity: id,
		Email:    verifiedToken.Email,
		Name:     verifiedToken.Name,
		Picture:  picture,
	}, nil
}
