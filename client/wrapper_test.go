// This file is hand-written and protected by .fernignore.

package client

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync/atomic"
	"testing"

	api "github.com/islo-labs/go-sdk"
	"github.com/islo-labs/go-sdk/option"
)

func TestNewIslo_CreateSandboxUsesComputeHost(t *testing.T) {
	control := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/auth/token" {
			t.Errorf("unexpected control request: %s", r.URL.Path)
			http.NotFound(w, r)
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"session_token": "jwt-create", "cookie_max_age": 600,
		})
	}))
	defer control.Close()

	var computeHost string
	compute := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Host != computeHost {
			t.Errorf("compute Host = %q, want %q", r.Host, computeHost)
			http.Error(w, "403 Forbidden: wrong authority", http.StatusForbidden)
			return
		}
		if r.Method != http.MethodPost || r.URL.Path != "/sandboxes" {
			t.Errorf("unexpected compute request: %s %s", r.Method, r.URL.Path)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer jwt-create" {
			t.Errorf("Authorization = %q, want exchanged token", got)
		}
		var body map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body["image"] != "alpine:latest" {
			t.Errorf("unexpected create body: %v (decode error: %v)", body, err)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":"sandbox-id","name":"sandbox-name","status":"creating"}`))
	}))
	defer compute.Close()
	computeURL, _ := url.Parse(compute.URL)
	computeHost = computeURL.Host
	t.Setenv("ISLO_COMPUTE_URL", compute.URL)

	c := NewIslo(option.WithAPIKey("ak_create"), option.WithBaseURL(control.URL))
	image := "alpine:latest"
	sandbox, err := c.Sandboxes.CreateSandbox(context.Background(), &api.CreateSandboxRequest{Image: &image})
	if err != nil {
		t.Fatalf("create sandbox: %v", err)
	}
	if sandbox.ID != "sandbox-id" || sandbox.Name != "sandbox-name" {
		t.Fatalf("unexpected sandbox response: %+v", sandbox)
	}
}

// TestNewIslo_ExchangesAPIKeyAndAttachesBearer verifies that an API key
// passed via option.WithAPIKey gets exchanged at /auth/token and the
// resulting JWT is sent as a Bearer token on subsequent calls.
func TestNewIslo_ExchangesAPIKeyAndAttachesBearer(t *testing.T) {
	var (
		exchangeCalls int32
		seenAuth      string
	)

	mux := http.NewServeMux()
	mux.HandleFunc("/auth/token", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&exchangeCalls, 1)
		var body struct {
			AccessKey string `json:"access_key"`
		}
		_ = json.NewDecoder(r.Body).Decode(&body)
		if body.AccessKey != "ak_test" {
			t.Errorf("/auth/token got access_key %q, want ak_test", body.AccessKey)
		}
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"session_token":  "jwt-from-exchange",
			"cookie_max_age": 600,
		})
	})
	// Catch-all: capture Authorization header for any non-/auth/token request.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		seenAuth = r.Header.Get("Authorization")
		w.WriteHeader(http.StatusNotFound)
	})

	srv := httptest.NewServer(mux)
	defer srv.Close()

	c := NewIslo(
		option.WithAPIKey("ak_test"),
		option.WithBaseURL(srv.URL),
	)

	// Trigger any request — we don't care about the response body, only
	// that the auth flow ran. Use a method that sends an HTTP request.
	_, _ = c.Credits.GetCreditBalance(context.Background())

	if got := atomic.LoadInt32(&exchangeCalls); got != 1 {
		t.Fatalf("expected 1 /auth/token call, got %d", got)
	}
	if want := "Bearer jwt-from-exchange"; seenAuth != want {
		t.Fatalf("Authorization header = %q, want %q", seenAuth, want)
	}
}

// TestNewIslo_BaseURLFallback ensures the default base URL is used when
// nothing is configured. We don't actually hit the production API — we
// just verify the constructor doesn't panic and the client carries the
// expected base URL through to its components.
func TestNewIslo_BaseURLFallback(t *testing.T) {
	t.Setenv("ISLO_API_KEY", "")
	t.Setenv("ISLO_BASE_URL", "")
	t.Setenv("ISLO_COMPUTE_URL", "")

	c := NewIslo(option.WithAPIKey("ak_test"))
	if c == nil {
		t.Fatal("NewIslo returned nil")
	}
	// The generated Client doesn't expose its baseURL, but we can still
	// assert non-nil so the call signature stays sane.
	if c.Sandboxes == nil {
		t.Fatal("c.Sandboxes is nil")
	}
}

func TestNewIslo_RoutesControlURLAndComputeURLFromEnv(t *testing.T) {
	var (
		controlPath string
		computePath string
		controlHost string
		computeHost string
	)

	control := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/auth/token" {
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"session_token":  "jwt-routing",
				"cookie_max_age": 600,
			})
			return
		}
		controlPath = r.URL.Path
		controlHost = r.Host
		w.WriteHeader(http.StatusNotFound)
	}))
	defer control.Close()

	compute := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		computePath = r.URL.Path
		computeHost = r.Host
		w.WriteHeader(http.StatusNotFound)
	}))
	defer compute.Close()
	t.Setenv("ISLO_COMPUTE_URL", compute.URL)

	c := NewIslo(
		option.WithAPIKey("ak_route"),
		option.WithBaseURL(control.URL),
	)

	_, _ = c.Credits.GetCreditBalance(context.Background())
	_, _ = c.Sandboxes.ListSandboxes(context.Background(), nil)

	if controlPath != "/credits/balance" {
		t.Fatalf("control request path = %q, want /credits/balance", controlPath)
	}
	if computePath != "/sandboxes" {
		t.Fatalf("compute request path = %q, want /sandboxes", computePath)
	}
	controlURL, _ := url.Parse(control.URL)
	computeURL, _ := url.Parse(compute.URL)
	if controlHost != controlURL.Host || computeHost != computeURL.Host {
		t.Fatalf("request hosts = (%q, %q), want (%q, %q)", controlHost, computeHost, controlURL.Host, computeURL.Host)
	}
}

func TestNewIslo_UsesEnvironmentVariablesForURLs(t *testing.T) {
	var (
		controlPath string
		computePath string
	)

	control := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/auth/token" {
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"session_token":  "jwt-env",
				"cookie_max_age": 600,
			})
			return
		}
		controlPath = r.URL.Path
		w.WriteHeader(http.StatusNotFound)
	}))
	defer control.Close()

	compute := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		computePath = r.URL.Path
		w.WriteHeader(http.StatusNotFound)
	}))
	defer compute.Close()

	t.Setenv("ISLO_API_KEY", "ak_env")
	t.Setenv("ISLO_BASE_URL", control.URL)
	t.Setenv("ISLO_COMPUTE_URL", compute.URL)

	c := NewIslo()

	_, _ = c.Credits.GetCreditBalance(context.Background())
	_, _ = c.Sandboxes.ListSandboxes(context.Background(), nil)

	if controlPath != "/credits/balance" {
		t.Fatalf("control request path = %q, want /credits/balance", controlPath)
	}
	if computePath != "/sandboxes" {
		t.Fatalf("compute request path = %q, want /sandboxes", computePath)
	}
}

// Compile-time assurance that core types remain accessible.
var _ = url.Values{}
