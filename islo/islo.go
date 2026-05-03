// Package islo is a thin convenience wrapper around the Fern-generated
// client. It handles the API-key → session-token exchange and refresh so
// callers don't have to.
//
//	c, err := islo.New(islo.WithAPIKey("ak_..."))
//	if err != nil { ... }
//	sb, err := c.Sandboxes.CreateSandbox(ctx, &api.SandboxCreate{...})
//
// Request/response types live in the root package
// (`github.com/islo-labs/go-sdk`); import it as `api` alongside this one.
package islo

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/islo-labs/go-sdk/client"
	"github.com/islo-labs/go-sdk/customauth"
	"github.com/islo-labs/go-sdk/option"
)

const (
	// DefaultBaseURL is the production Islo API.
	DefaultBaseURL = "https://api.islo.dev"

	// EnvAPIKey is the environment variable consulted when WithAPIKey is unset.
	EnvAPIKey = "ISLO_API_KEY"

	// EnvBaseURL overrides DefaultBaseURL when WithBaseURL is unset.
	EnvBaseURL = "ISLO_BASE_URL"
)

// Client is the generated Fern client, re-exported for convenience.
type Client = client.Client

// Option configures New.
type Option func(*config)

type config struct {
	apiKey        string
	baseURL       string
	refreshMargin time.Duration
	httpClient    *http.Client
}

// WithAPIKey supplies the Islo API key. Falls back to ISLO_API_KEY.
func WithAPIKey(apiKey string) Option {
	return func(c *config) { c.apiKey = apiKey }
}

// WithBaseURL overrides the API base URL. Falls back to ISLO_BASE_URL,
// then DefaultBaseURL.
func WithBaseURL(baseURL string) Option {
	return func(c *config) { c.baseURL = baseURL }
}

// WithRefreshMargin sets how far before expiry the session token is
// proactively refreshed. Defaults to 60s.
func WithRefreshMargin(d time.Duration) Option {
	return func(c *config) { c.refreshMargin = d }
}

// WithHTTPClient supplies a custom *http.Client. Its Transport (if any) is
// wrapped so the Authorization header is injected on every request; other
// fields (Timeout, Jar, CheckRedirect) are preserved.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *config) { c.httpClient = httpClient }
}

// New constructs an Islo client that auto-exchanges your API key for a
// session JWT and refreshes it before expiry.
func New(opts ...Option) (*Client, error) {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}

	if cfg.apiKey == "" {
		cfg.apiKey = os.Getenv(EnvAPIKey)
	}
	if cfg.apiKey == "" {
		return nil, errors.New("islo: API key required (set " + EnvAPIKey + " or pass islo.WithAPIKey)")
	}

	if cfg.baseURL == "" {
		cfg.baseURL = os.Getenv(EnvBaseURL)
	}
	if cfg.baseURL == "" {
		cfg.baseURL = DefaultBaseURL
	}

	var baseTransport http.RoundTripper
	if cfg.httpClient != nil && cfg.httpClient.Transport != nil {
		baseTransport = cfg.httpClient.Transport
	}

	provider := customauth.NewProvider(cfg.baseURL, cfg.apiKey, cfg.refreshMargin, nil)
	authedClient := &http.Client{
		Transport: customauth.NewTransport(baseTransport, provider),
	}
	if cfg.httpClient != nil {
		authedClient.Timeout = cfg.httpClient.Timeout
		authedClient.Jar = cfg.httpClient.Jar
		authedClient.CheckRedirect = cfg.httpClient.CheckRedirect
	}

	return client.NewClient(
		option.WithBaseURL(cfg.baseURL),
		option.WithHTTPClient(authedClient),
	), nil
}
