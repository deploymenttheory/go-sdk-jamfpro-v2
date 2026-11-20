package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/config"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
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

	cfg := &config.AuthConfig{
		InstanceDomain: srv.URL,
		AuthMethod:     constants.AuthMethodOAuth2,
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
	cfg := &config.AuthConfig{InstanceDomain: "", AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	_, err := NewTransport(cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "instance domain")
}

func TestNewTransport_TrimTrailingSlash(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL + "/", AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	assert.Equal(t, srv.URL, tr.BaseURL)
}

func TestTransport_NewRequest_Methods(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	var getResult map[string]string
	resp, err := tr.NewRequest(ctx).SetResult(&getResult).Get("/api/test")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", getResult["id"])

	var postResult map[string]string
	resp, err = tr.NewRequest(ctx).SetBody(map[string]string{"k": "v"}).SetResult(&postResult).Post("/api/test")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	resp, err = tr.NewRequest(ctx).SetBody(map[string]string{"k": "v"}).SetResult(&postResult).Put("/api/test")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	resp, err = tr.NewRequest(ctx).Delete("/api/test")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	var patchResult map[string]string
	resp, err = tr.NewRequest(ctx).SetBody(map[string]string{"k": "v"}).SetResult(&patchResult).Patch("/api/test")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	// DeleteWithBody equivalent: SetBody + Delete
	resp, err = tr.NewRequest(ctx).SetBody(map[string]string{"id": "1"}).Delete("/api/test")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_NewRequest_PostWithQueryAndGetBytes(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	// PostWithQuery equivalent: SetQueryParam + SetBody + Post
	var result map[string]string
	resp, err := tr.NewRequest(ctx).SetQueryParam("q", "v").SetResult(&result).Post("/api/test")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	// GetBytes
	resp, body, err := tr.NewRequest(ctx).GetBytes("/api/test")
	require.NoError(t, err)
	require.NotNil(t, body)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_NewRequest_PostMultipart(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	var result map[string]string
	resp, err := tr.NewRequest(ctx).
		SetMultipartFile("file", "test.txt", bytes.NewReader([]byte("content")), 7, nil).
		SetMultipartFormData(map[string]string{"key": "val"}).
		SetResult(&result).
		Post("/api/upload")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	called := false
	resp, err = tr.NewRequest(ctx).
		SetMultipartFile("file", "a.txt", bytes.NewReader([]byte("x")), 1, func(_, _ string, written, total int64) { called = true }).
		SetResult(&result).
		Post("/api/upload")
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
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	resp, body, err := tr.NewRequest(ctx).GetBytes("/api/notfound")
	require.Error(t, err)
	assert.Nil(t, body)
	assert.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
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
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg, func(s *TransportSettings) error { s.MaxConcurrentRequests = 1; return nil })
	require.NoError(t, err)

	block = make(chan struct{})
	ctx := context.Background()
	var result map[string]string
	go func() { _, _ = tr.NewRequest(ctx).SetResult(&result).Get("/api/block") }()
	time.Sleep(50 * time.Millisecond)
	ctx2, cancel2 := context.WithCancel(context.Background())
	cancel2()
	_, err = tr.NewRequest(ctx2).SetResult(&result).Get("/api/test")
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
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()
	var result map[string]any
	resp, err := tr.NewRequest(ctx).SetResult(&result).Get("/api/empty")
	require.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode())
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
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	var result []byte
	resp, err := tr.NewRequest(ctx).SetResult(&result).Get("/api/plain")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_WithTotalRetryDuration_Request(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg, func(s *TransportSettings) error { s.TotalRetryDuration = 30 * time.Second; return nil })
	require.NoError(t, err)
	ctx := context.Background()
	var result map[string]string
	resp, err := tr.NewRequest(ctx).SetResult(&result).Get("/api/test")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_RequestWithMandatoryDelay(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg, func(s *TransportSettings) error { s.MandatoryRequestDelay = time.Millisecond; return nil })
	require.NoError(t, err)
	ctx := context.Background()
	var result map[string]string
	resp, err := tr.NewRequest(ctx).SetResult(&result).Get("/api/test")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_ExecuteRequest_ContextCanceled(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	var result map[string]string
	resp, err := tr.NewRequest(ctx).SetResult(&result).Get("/api/test")
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
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()
	var result map[string]any
	_, err = tr.NewRequest(ctx).SetResult(&result).Get("/api/slow")
	require.NoError(t, err)
	_, err = tr.NewRequest(ctx).SetResult(&result).Get("/api/slow")
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
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()
	var result map[string]string
	resp, err := tr.NewRequest(ctx).SetResult(&result).Get("/api/err")
	require.Error(t, err)
	assert.Equal(t, 500, resp.StatusCode())
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
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	var result map[string]any
	resp, err := tr.NewRequest(ctx).SetResult(&result).Get("/api/deprecated")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_InvalidateToken(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	err = tr.InvalidateToken()
	require.NoError(t, err)
}

func TestTransport_KeepAliveToken(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	err = tr.KeepAliveToken()
	require.NoError(t, err)
}

func TestTransport_GlobalHeaders(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg, func(s *TransportSettings) error {
		s.GlobalHeaders = map[string]string{"X-Global": "global-val"}
		return nil
	})
	require.NoError(t, err)
	ctx := context.Background()

	// Global header from TransportSettings is applied automatically;
	// per-request header is set on the builder.
	var result map[string]string
	resp, err := tr.NewRequest(ctx).SetHeader("X-Request", "request-val").SetResult(&result).Get("/api/test")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_SetQueryParams_EmptyValuesSkipped(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	var result map[string]string
	resp, err := tr.NewRequest(ctx).
		SetQueryParams(map[string]string{"a": "", "b": "val"}).
		SetResult(&result).
		Get("/api/x")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_NilBody_NilQuery(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	var result map[string]string
	resp, err := tr.NewRequest(ctx).SetBody(nil).SetResult(&result).Post("/api/x")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	resp, err = tr.NewRequest(ctx).SetBody(nil).SetResult(&result).Put("/api/x")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	resp, err = tr.NewRequest(ctx).SetBody(nil).SetResult(&result).Patch("/api/x")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	resp, err = tr.NewRequest(ctx).Delete("/api/x")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_GetBytes_WithQueryParams(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	resp, body, err := tr.NewRequest(ctx).SetQueryParam("filter", `name=="test"`).GetBytes("/api/x")
	require.NoError(t, err)
	require.NotNil(t, body)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestTransport_PostMultipart_NoFile(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg)
	require.NoError(t, err)
	ctx := context.Background()

	var result map[string]string
	// SetMultipartFile with nil reader is a no-op; only form data is sent
	resp, err := tr.NewRequest(ctx).
		SetMultipartFile("", "", nil, 0, nil).
		SetMultipartFormData(map[string]string{"k": "v"}).
		SetResult(&result).
		Post("/api/x")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

// uploadRoundTripFunc lets the upload test return a response while its source is blocked.
type uploadRoundTripFunc func(*http.Request) (*http.Response, error)

func (f uploadRoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

type startedUploadReader struct {
	*io.PipeReader
	started chan struct{}
	once    sync.Once
}

func (r *startedUploadReader) Read(p []byte) (int, error) {
	r.once.Do(func() { close(r.started) })
	return r.PipeReader.Read(p)
}

func TestUnit_Transport_MultipartEarlyRejection(t *testing.T) {
	pipeReader, pipeWriter := io.Pipe()
	source := &startedUploadReader{PipeReader: pipeReader, started: make(chan struct{})}
	t.Cleanup(func() { _ = source.Close(); _ = pipeWriter.Close() })
	transport := uploadRoundTripFunc(func(req *http.Request) (*http.Response, error) {
		defer req.Body.Close()
		body := `{"access_token":"test-token","expires_in":3600}`
		status := http.StatusOK
		if req.URL.Path == "/api/upload" {
			reader, err := req.MultipartReader()
			if err != nil {
				return nil, err
			}
			if _, err := reader.NextPart(); err != nil {
				return nil, err
			}
			// Ensure the producer has entered Read before returning the rejection.
			select {
			case <-source.started:
			case <-req.Context().Done():
				return nil, req.Context().Err()
			}
			body = `{"code":"UPLOAD_TOO_LARGE","message":"Upload rejected"}`
			status = http.StatusRequestEntityTooLarge
		}
		return &http.Response{
			StatusCode: status, Status: http.StatusText(status),
			Header: http.Header{"Content-Type": []string{"application/json"}},
			Body:   io.NopCloser(bytes.NewBufferString(body)), Request: req,
		}, nil
	})
	cfg := &config.AuthConfig{InstanceDomain: "https://example.test", AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg, func(s *TransportSettings) error { s.HTTPTransport = transport; return nil })
	require.NoError(t, err)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	type outcome struct {
		status int
		err    error
	}
	done := make(chan outcome, 1)
	go func() {
		resp, err := tr.NewRequest(ctx).SetMultipartFile("file", "blocked.pkg", source, 1, nil).Post("/api/upload")
		status := 0
		if resp != nil {
			status = resp.StatusCode()
		}
		done <- outcome{status: status, err: err}
	}()
	select {
	case result := <-done:
		require.Equal(t, http.StatusRequestEntityTooLarge, result.status)
		var apiErr *APIError
		require.ErrorAs(t, result.err, &apiErr)
		require.Equal(t, "UPLOAD_TOO_LARGE", apiErr.Code)
	case <-ctx.Done():
		// Release the producer even when running against a dependency that hangs.
		_ = source.Close()
		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("multipart request did not stop after closing its reader")
		}
		t.Fatal("multipart rejection hung until the caller closed the source reader")
	}
}

func TestUnit_Transport_RequestDefaultsIsolation(t *testing.T) {
	srv := newMockAuthServer(t)
	defer srv.Close()
	cfg := &config.AuthConfig{InstanceDomain: srv.URL, AuthMethod: constants.AuthMethodOAuth2, ClientID: "c", ClientSecret: "s"}
	tr, err := NewTransport(cfg, func(s *TransportSettings) error {
		s.GlobalHeaders = map[string]string{"X-SDK-Default": "original"}
		return nil
	})
	require.NoError(t, err)
	tr.GetHTTPClient().SetQueryParam("sdk-default", "original")
	first, err := tr.NewRequest(context.Background()).Get("/api/test")
	require.NoError(t, err)
	// A caller may modify request metadata returned with a response. It must not
	// change the defaults or metadata used by subsequent requests.
	first.Request.Header["X-Sdk-Default"][0] = "changed"
	first.Request.QueryParams["sdk-default"][0] = "changed"
	second, err := tr.NewRequest(context.Background()).Get("/api/test")
	require.NoError(t, err)
	require.Equal(t, "original", second.Request.RawRequest.Header.Get("X-SDK-Default"))
	require.Equal(t, "original", second.Request.RawRequest.URL.Query().Get("sdk-default"))
}
