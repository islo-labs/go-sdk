package customauth

import (
	"net/http"
)

// Transport injects a fresh Authorization: Bearer <token> header on every
// request, sourcing the token from a Provider so it auto-refreshes.
type Transport struct {
	Base     http.RoundTripper
	Provider *Provider
}

// NewTransport wraps base with auth header injection. If base is nil,
// http.DefaultTransport is used.
func NewTransport(base http.RoundTripper, provider *Provider) *Transport {
	if base == nil {
		base = http.DefaultTransport
	}
	return &Transport{Base: base, Provider: provider}
}

// RoundTrip implements http.RoundTripper. It clones the request, fetches
// (or reuses cached) token, and sets the Authorization header before
// delegating to the base transport.
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	token, err := t.Provider.Token(req.Context())
	if err != nil {
		return nil, err
	}
	cloned := req.Clone(req.Context())
	cloned.Header.Set("Authorization", "Bearer "+token)
	return t.Base.RoundTrip(cloned)
}
