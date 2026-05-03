// This file is hand-written and protected by .fernignore. Fern will not
// regenerate or overwrite it. The companion generated client lives in
// client.go (DO NOT EDIT).

package client

import (
	"net/http"
	"os"
	"time"

	"github.com/islo-labs/go-sdk/core"
	"github.com/islo-labs/go-sdk/customauth"
	"github.com/islo-labs/go-sdk/option"
)

// Default base URL and environment variable names for NewIslo.
const (
	defaultIsloBaseURL = "https://api.islo.dev"
	envIsloAPIKey      = "ISLO_API_KEY"
	envIsloBaseURL     = "ISLO_BASE_URL"
)

// NewIslo behaves like NewClient, but treats the value passed to
// option.WithAPIKey as an Islo API key (e.g. "ak_..."). It exchanges the
// key for a short-lived session JWT against /auth/token and refreshes
// the token before expiry. If WithAPIKey is not provided, the
// ISLO_API_KEY environment variable is used. The base URL falls back to
// ISLO_BASE_URL, then https://api.islo.dev.
//
// Most callers should use this rather than NewClient.
func NewIslo(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)

	apiKey := options.APIKey
	if apiKey == "" {
		apiKey = os.Getenv(envIsloAPIKey)
	}

	baseURL := options.BaseURL
	if baseURL == "" {
		baseURL = os.Getenv(envIsloBaseURL)
	}
	if baseURL == "" {
		baseURL = defaultIsloBaseURL
	}

	// If the caller supplied an *http.Client, preserve its Transport so
	// it sits below our auth-injecting RoundTripper.
	var baseTransport http.RoundTripper
	var baseTimeout time.Duration
	if hc, ok := options.HTTPClient.(*http.Client); ok && hc != nil {
		baseTransport = hc.Transport
		baseTimeout = hc.Timeout
	}

	provider := customauth.NewProvider(baseURL, apiKey, 0, nil)
	authedClient := &http.Client{
		Transport: customauth.NewTransport(baseTransport, provider),
		Timeout:   baseTimeout,
	}

	return NewClient(
		option.WithBaseURL(baseURL),
		option.WithHTTPClient(authedClient),
	)
}
