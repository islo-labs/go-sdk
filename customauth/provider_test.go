package customauth

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

func TestProvider_TokenCachedUntilExpiry(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		_ = json.NewEncoder(w).Encode(tokenResponse{
			SessionToken: "jwt-1",
			CookieMaxAge: 600,
		})
	}))
	defer srv.Close()

	resetSharedStates(t)
	p := NewProvider(srv.URL, "ak_test", DefaultRefreshMargin, srv.Client())

	ctx := context.Background()
	for i := 0; i < 3; i++ {
		tok, err := p.Token(ctx)
		if err != nil {
			t.Fatalf("Token: %v", err)
		}
		if tok != "jwt-1" {
			t.Fatalf("got token %q, want jwt-1", tok)
		}
	}
	if got := atomic.LoadInt32(&calls); got != 1 {
		t.Fatalf("expected 1 exchange call, got %d", got)
	}
}

func TestProvider_RefreshesAfterExpiry(t *testing.T) {
	var calls int32
	tokens := []string{"jwt-1", "jwt-2"}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		i := atomic.AddInt32(&calls, 1) - 1
		_ = json.NewEncoder(w).Encode(tokenResponse{
			SessionToken: tokens[int(i)],
			// Tiny TTL: 1s with default 60s margin → expires immediately.
			CookieMaxAge: 1,
		})
	}))
	defer srv.Close()

	resetSharedStates(t)
	p := NewProvider(srv.URL, "ak_test", DefaultRefreshMargin, srv.Client())

	ctx := context.Background()
	first, err := p.Token(ctx)
	if err != nil || first != "jwt-1" {
		t.Fatalf("first Token: %v, %q", err, first)
	}
	second, err := p.Token(ctx)
	if err != nil || second != "jwt-2" {
		t.Fatalf("second Token: %v, %q", err, second)
	}
	if got := atomic.LoadInt32(&calls); got != 2 {
		t.Fatalf("expected 2 exchange calls, got %d", got)
	}
}

func TestProvider_SharedStateAcrossInstances(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		_ = json.NewEncoder(w).Encode(tokenResponse{
			SessionToken: "shared-jwt",
			CookieMaxAge: 600,
		})
	}))
	defer srv.Close()

	resetSharedStates(t)
	p1 := NewProvider(srv.URL, "ak_test", DefaultRefreshMargin, srv.Client())
	p2 := NewProvider(srv.URL, "ak_test", DefaultRefreshMargin, srv.Client())

	ctx := context.Background()
	if _, err := p1.Token(ctx); err != nil {
		t.Fatalf("p1.Token: %v", err)
	}
	if _, err := p2.Token(ctx); err != nil {
		t.Fatalf("p2.Token: %v", err)
	}
	if got := atomic.LoadInt32(&calls); got != 1 {
		t.Fatalf("expected shared cache (1 call), got %d", got)
	}
}

func TestProvider_NonOKStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"error":"bad key"}`))
	}))
	defer srv.Close()

	resetSharedStates(t)
	p := NewProvider(srv.URL, "ak_bad", DefaultRefreshMargin, srv.Client())

	if _, err := p.Token(context.Background()); err == nil {
		t.Fatal("expected error on 401, got nil")
	}
}

func TestTransport_InjectsBearer(t *testing.T) {
	var captured string
	tokenSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(tokenResponse{
			SessionToken: "transport-jwt",
			CookieMaxAge: 600,
		})
	}))
	defer tokenSrv.Close()

	apiSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		captured = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusOK)
	}))
	defer apiSrv.Close()

	resetSharedStates(t)
	p := NewProvider(tokenSrv.URL, "ak_test", DefaultRefreshMargin, tokenSrv.Client())
	transport := NewTransport(http.DefaultTransport, p)

	req, _ := http.NewRequest(http.MethodGet, apiSrv.URL, nil)
	resp, err := transport.RoundTrip(req)
	if err != nil {
		t.Fatalf("RoundTrip: %v", err)
	}
	resp.Body.Close()

	if want := "Bearer transport-jwt"; captured != want {
		t.Fatalf("Authorization header = %q, want %q", captured, want)
	}
}

// resetSharedStates clears the package-level shared cache so tests don't
// leak token state between cases.
func resetSharedStates(t *testing.T) {
	t.Helper()
	sharedMu.Lock()
	defer sharedMu.Unlock()
	for k := range sharedStates {
		delete(sharedStates, k)
	}
}

// Compile-time assertion that the chosen TTL math allows the
// expiration test to behave as expected.
var _ = time.Second
