// This file is hand-written and protected by .fernignore. Fern will not
// regenerate or overwrite it. The companion generated client lives in
// client.go (DO NOT EDIT).

package client

import (
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	api "github.com/islo-labs/go-sdk"
	"github.com/islo-labs/go-sdk/core"
	"github.com/islo-labs/go-sdk/customauth"
	"github.com/islo-labs/go-sdk/option"
)

// Default base URL and environment variable names for NewIslo.
const (
	envIsloAPIKey     = "ISLO_API_KEY"
	envIsloBaseURL    = "ISLO_BASE_URL"
	envIsloComputeURL = "ISLO_COMPUTE_URL"
)

// NewIslo behaves like NewClient, but treats the value passed to
// option.WithAPIKey as an Islo API key (e.g. "ak_..."). It exchanges the
// key for a short-lived session JWT against /auth/token and refreshes
// the token before expiry. If WithAPIKey is not provided, the
// ISLO_API_KEY environment variable is used. WithBaseURL configures the
// control-plane URL and falls back to ISLO_BASE_URL, then
// https://api.islo.dev. The compute-plane URL is read from
// ISLO_COMPUTE_URL and falls back to https://ca.compute.islo.dev.
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
		baseURL = api.Environments.Production.Control
	}

	computeURL := os.Getenv(envIsloComputeURL)
	if computeURL == "" {
		computeURL = api.Environments.Production.Compute
	}

	// If the caller supplied an *http.Client, preserve its Transport so
	// it sits below our auth-injecting RoundTripper.
	var baseTransport http.RoundTripper
	var baseTimeout time.Duration
	if hc, ok := options.HTTPClient.(*http.Client); ok && hc != nil {
		baseTransport = hc.Transport
		baseTimeout = hc.Timeout
	}

	rewriter := newEnvironmentRewriteTransport(baseTransport, baseURL, computeURL)
	provider := customauth.NewProvider(baseURL, apiKey, 0, nil)
	authedClient := &http.Client{
		Transport: customauth.NewTransport(rewriter, provider),
		Timeout:   baseTimeout,
	}

	clientOpts := []option.RequestOption{
		option.WithHTTPClient(authedClient),
	}
	if options.HTTPHeader != nil && len(options.HTTPHeader) > 0 {
		clientOpts = append(clientOpts, option.WithHTTPHeader(options.HTTPHeader))
	}
	if options.MaxAttempts > 0 {
		clientOpts = append(clientOpts, option.WithMaxAttempts(options.MaxAttempts))
	}
	if options.BodyProperties != nil && len(options.BodyProperties) > 0 {
		clientOpts = append(clientOpts, option.WithBodyProperties(options.BodyProperties))
	}
	if options.QueryParameters != nil && len(options.QueryParameters) > 0 {
		clientOpts = append(clientOpts, option.WithQueryParameters(options.QueryParameters))
	}

	return NewClient(clientOpts...)
}

type environmentRewriteTransport struct {
	base       http.RoundTripper
	controlURL *url.URL
	computeURL *url.URL
}

func newEnvironmentRewriteTransport(base http.RoundTripper, controlURL string, computeURL string) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	control, err := url.Parse(controlURL)
	if err != nil {
		control = mustParseURL(api.Environments.Production.Control)
	}
	compute, err := url.Parse(computeURL)
	if err != nil {
		compute = mustParseURL(api.Environments.Production.Compute)
	}
	return &environmentRewriteTransport{
		base:       base,
		controlURL: control,
		computeURL: compute,
	}
}

func (t *environmentRewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	target := t.targetURL(req.URL)
	if target == nil {
		return t.base.RoundTrip(req)
	}

	cloned := req.Clone(req.Context())
	cloned.URL = rewriteURL(req.URL, target)
	return t.base.RoundTrip(cloned)
}

func (t *environmentRewriteTransport) targetURL(requestURL *url.URL) *url.URL {
	switch requestURL.Host {
	case mustParseURL(api.Environments.Production.Control).Host:
		return t.controlURL
	case mustParseURL(api.Environments.Production.Compute).Host:
		return t.computeURL
	default:
		return nil
	}
}

func rewriteURL(original *url.URL, target *url.URL) *url.URL {
	rewritten := *original
	rewritten.Scheme = target.Scheme
	rewritten.Host = target.Host
	if target.Path != "" && target.Path != "/" {
		rewritten.Path = strings.TrimRight(target.Path, "/") + original.Path
	}
	return &rewritten
}

func mustParseURL(rawURL string) *url.URL {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	return parsed
}
