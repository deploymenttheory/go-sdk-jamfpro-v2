package client

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newMockAuthServer(t *testing.T) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/v1/oauth/token":
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"access_token": "test-token",
				"expires_in":   3600,
			})
		case "/api/v1/auth/token":
			if r.Method != http.MethodPost {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"token":   "test-token",
				"expires": time.Now().Add(24 * time.Hour).Format(time.RFC3339),
			})
		case "/api/v1/auth/invalidate-token":
			w.WriteHeader(http.StatusOK)
		case "/api/v1/auth/keep-alive":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_ = json.NewEncoder(w).Encode(map[string]any{
				"token":   "new-token",
				"expires": time.Now().Add(24 * time.Hour).Format(time.RFC3339),
			})
		default:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"id":"1","name":"test"}`))
		}
	}))
}

func TestNewTransport(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()

	cfg := &AuthConfig{
		InstanceDomain: srv.URL,
		AuthMethod:     AuthMethodOAuth2,
		ClientID:       "cid",
		ClientSecret:   "secret",
	}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	require.NotNil(t, tr)
	assert.Equal(t, srv.URL, tr.BaseURL)
	assert.NotNil(t, tr.GetHTTPClient())
	assert.NotNil(t, tr.GetLogger())
	assert.NotNil(t, tr.RSQLBuilder())
}

func TestNewTransport_NilConfig(t *testing.T) {
	_, err := NewTransport(nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "auth config is required")
}

func TestNewTransport_EmptyInstanceDomain(t *testing.T) {
	cfg := &AuthConfig{InstanceDomain: "", AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	_, err := NewTransport(cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "instance domain")
}

func TestNewTransport_TrimTrailingSlash(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL + "/", AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	assert.Equal(t, srv.URL, tr.BaseURL)
}

func TestTransport_Get_Post_Put_Delete(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	var getResult map[string]string
	resp, err := tr.Get(ctx, "/api/test", nil, nil, &getResult)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", getResult["id"])

	var postResult map[string]string
	resp, err = tr.Post(ctx, "/api/test", map[string]string{"k": "v"}, nil, &postResult)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	resp, err = tr.Put(ctx, "/api/test", map[string]string{"k": "v"}, nil, &postResult)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	resp, err = tr.Delete(ctx, "/api/test", nil, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	var patchResult map[string]string
	resp, err = tr.Patch(ctx, "/api/test", map[string]string{"k": "v"}, nil, &patchResult)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	resp, err = tr.DeleteWithBody(ctx, "/api/test", map[string]string{"id": "1"}, nil, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestTransport_PostWithQuery_PostForm_GetBytes(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	var result map[string]string
	resp, err := tr.PostWithQuery(ctx, "/api/test", map[string]string{"q": "v"}, nil, nil, &result)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	resp, err = tr.PostForm(ctx, "/api/test", map[string]string{"form": "data"}, nil, &result)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	resp, body, err := tr.GetBytes(ctx, "/api/test", nil, nil)
	require.NoError(t, err)
	require.NotNil(t, body)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestTransport_PostMultipart(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	var result map[string]string
	resp, err := tr.PostMultipart(ctx, "/api/upload", "file", "test.txt", bytes.NewReader([]byte("content")), 7, map[string]string{"key": "val"}, nil, nil, &result)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)

	called := false
	resp, err = tr.PostMultipart(ctx, "/api/upload", "file", "a.txt", bytes.NewReader([]byte("x")), 1, nil, nil, func(_, _ string, written, total int64) { called = true }, &result)
	require.NoError(t, err)
	assert.True(t, called)
}

func TestTransport_GetBytes_ErrorResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/oauth/token" && r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"access_token":"t","expires_in":3600}`))
			return
		}
		if r.URL.Path == "/api/notfound" {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(`{"message":"not found"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	resp, body, err := tr.GetBytes(ctx, "/api/notfound", nil, nil)
	require.Error(t, err)
	assert.Nil(t, body)
	assert.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestTransport_ExecuteRequest_ConcurrencyLimit(t *testing.T) {
	block := make(chan struct{})
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/oauth/token" && r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"access_token":"t","expires_in":3600}`))
			return
		}
		if r.URL.Path == "/api/block" {
			<-block
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer srv.Close()
	close(block)
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg, WithMaxConcurrentRequests(1))
	require.NoError(t, err)

	block = make(chan struct{})
	ctx := context.Background()
	var result map[string]string
	go func() { _, _ = tr.Get(ctx, "/api/block", nil, nil, &result) }()
	time.Sleep(50 * time.Millisecond)
	ctx2, cancel2 := context.WithCancel(context.Background())
	cancel2()
	_, err = tr.Get(ctx2, "/api/test", nil, nil, &result)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "concurrency limit")
	close(block)
}

func TestTransport_ValidateResponse_EmptyBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/oauth/token" && r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"access_token":"t","expires_in":3600}`))
			return
		}
		if r.URL.Path == "/api/empty" {
			w.Header().Set("Content-Length", "0")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()
	var result map[string]interface{}
	resp, err := tr.Get(ctx, "/api/empty", nil, nil, &result)
	require.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestTransport_ValidateResponse_NonJSONWarn(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/oauth/token" && r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"access_token":"t","expires_in":3600}`))
			return
		}
		if r.URL.Path == "/api/plain" {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("plain text"))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	var result []byte
	resp, err := tr.Get(ctx, "/api/plain", nil, nil, &result)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestTransport_WithTotalRetryDuration_Request(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg, WithTotalRetryDuration(30*time.Second))
	require.NoError(t, err)
	ctx := context.Background()
	var result map[string]string
	resp, err := tr.Get(ctx, "/api/test", nil, nil, &result)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestTransport_RequestWithMandatoryDelay(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg, WithMandatoryRequestDelay(time.Millisecond))
	require.NoError(t, err)
	ctx := context.Background()
	var result map[string]string
	resp, err := tr.Get(ctx, "/api/test", nil, nil, &result)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestTransport_ExecuteRequest_ContextCanceled(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	var result map[string]string
	resp, err := tr.Get(ctx, "/api/test", nil, nil, &result)
	require.Error(t, err)
	assert.NotNil(t, resp)
}

func TestTransport_AdaptiveDelay(t *testing.T) {
	first := true
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/oauth/token" && r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"access_token":"t","expires_in":3600}`))
			return
		}
		if r.URL.Path == "/api/slow" {
			if first {
				first = false
			} else {
				time.Sleep(150 * time.Millisecond)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{}`))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()
	var result map[string]interface{}
	_, err = tr.Get(ctx, "/api/slow", nil, nil, &result)
	require.NoError(t, err)
	_, err = tr.Get(ctx, "/api/slow", nil, nil, &result)
	require.NoError(t, err)
}

func TestTransport_ExecuteRequest_ServerError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/oauth/token" && r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"access_token":"t","expires_in":3600}`))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message":"internal"}`))
	}))
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()
	var result map[string]string
	resp, err := tr.Get(ctx, "/api/err", nil, nil, &result)
	require.Error(t, err)
	assert.Equal(t, 500, resp.StatusCode)
}

func TestTransport_DeprecationHeader(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/oauth/token" && r.Method == http.MethodPost {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"access_token":"t","expires_in":3600}`))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Deprecation", "true")
		w.Header().Set("Sunset", "2026-01-01")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	var result map[string]any
	resp, err := tr.Get(ctx, "/api/deprecated", nil, nil, &result)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestTransport_InvalidateToken(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	err = tr.InvalidateToken()
	require.NoError(t, err)
}

func TestTransport_KeepAliveToken(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	err = tr.KeepAliveToken()
	require.NoError(t, err)
}

func TestTransport_ApplyHeaders(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &AuthConfig{InstanceDomain: srv.URL, AuthMethod: AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg, WithGlobalHeader("X-Global", "global-val"))
	require.NoError(t, err)
	ctx := context.Background()

	var result map[string]string
	resp, err := tr.Get(ctx, "/api/test", nil, map[string]string{"X-Request": "request-val"}, &result)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}
