// Package customauth implements the API-key → session-token exchange and
// auto-refreshing token cache used by the top-level islo package.
//
// Users typically don't import this directly — call islo.New from the
// module root, which wires the provider into the generated client.
package customauth

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

// DefaultRefreshMargin is how far before expiry a token is proactively
// refreshed. Matches the python-sdk.
const DefaultRefreshMargin = 60 * time.Second

// providerKey scopes shared cache state to a (baseURL, accessKey) pair so
// multiple Islo() instances with the same key share one token.
type providerKey struct {
	baseURL   string
	accessKey string
}

type tokenState struct {
	mu        sync.Mutex
	token     string
	expiresAt time.Time
}

var (
	sharedMu     sync.Mutex
	sharedStates = map[providerKey]*tokenState{}
)

// Provider exchanges an Islo API key for a short-lived session JWT and
// caches it until shortly before expiry.
type Provider struct {
	baseURL       string
	accessKey     string
	refreshMargin time.Duration
	httpClient    *http.Client
	state         *tokenState
}

// NewProvider returns a Provider that shares cached state with any other
// Provider built from the same (baseURL, accessKey).
//
// httpClient is used for the /auth/token exchange call. Pass nil to use
// http.DefaultClient with a 10s timeout.
func NewProvider(baseURL, accessKey string, refreshMargin time.Duration, httpClient *http.Client) *Provider {
	if refreshMargin <= 0 {
		refreshMargin = DefaultRefreshMargin
	}
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 10 * time.Second}
	}

	sharedMu.Lock()
	defer sharedMu.Unlock()
	key := providerKey{baseURL: baseURL, accessKey: accessKey}
	state, ok := sharedStates[key]
	if !ok {
		state = &tokenState{}
		sharedStates[key] = state
	}
	return &Provider{
		baseURL:       baseURL,
		accessKey:     accessKey,
		refreshMargin: refreshMargin,
		httpClient:    httpClient,
		state:         state,
	}
}

// Token returns a cached session token, refreshing it if expired.
func (p *Provider) Token(ctx context.Context) (string, error) {
	p.state.mu.Lock()
	defer p.state.mu.Unlock()
	if p.state.token != "" && time.Now().Before(p.state.expiresAt) {
		return p.state.token, nil
	}
	return p.refresh(ctx)
}

type tokenResponse struct {
	SessionToken string `json:"session_token"`
	CookieMaxAge int    `json:"cookie_max_age"`
}

func (p *Provider) refresh(ctx context.Context) (string, error) {
	body, err := json.Marshal(map[string]string{"access_key": p.accessKey})
	if err != nil {
		return "", fmt.Errorf("islo: encode token request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, p.baseURL+"/auth/token", bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("islo: build token request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := p.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("islo: token exchange: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		snippet, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
		return "", fmt.Errorf("islo: token exchange failed (%d): %s", resp.StatusCode, bytes.TrimSpace(snippet))
	}

	var data tokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", fmt.Errorf("islo: decode token response: %w", err)
	}
	if data.SessionToken == "" {
		return "", errors.New("islo: token exchange response missing session_token")
	}

	ttl := time.Duration(data.CookieMaxAge)*time.Second - p.refreshMargin
	if ttl < 0 {
		ttl = 0
	}
	p.state.token = data.SessionToken
	p.state.expiresAt = time.Now().Add(ttl)
	return p.state.token, nil
}
