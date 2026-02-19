package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"resty.dev/v3"
)

func TestIsIdempotentMethod(t *testing.T) {
	assert.True(t, isIdempotentMethod("GET"))
	assert.True(t, isIdempotentMethod("HEAD"))
	assert.True(t, isIdempotentMethod("OPTIONS"))
	assert.True(t, isIdempotentMethod("PUT"))
	assert.True(t, isIdempotentMethod("DELETE"))
	assert.False(t, isIdempotentMethod("POST"))
	assert.False(t, isIdempotentMethod("PATCH"))
	assert.False(t, isIdempotentMethod(""))
}

func TestIsTransientStatusCode(t *testing.T) {
	assert.True(t, isTransientStatusCode(500))
	assert.True(t, isTransientStatusCode(502))
	assert.True(t, isTransientStatusCode(503))
	assert.True(t, isTransientStatusCode(504))
	assert.False(t, isTransientStatusCode(200))
	assert.False(t, isTransientStatusCode(404))
}

func TestIsNonRetryableStatusCode(t *testing.T) {
	assert.True(t, isNonRetryableStatusCode(400))
	assert.True(t, isNonRetryableStatusCode(401))
	assert.True(t, isNonRetryableStatusCode(404))
	assert.True(t, isNonRetryableStatusCode(428))
	assert.False(t, isNonRetryableStatusCode(429))
	assert.False(t, isNonRetryableStatusCode(500))
	assert.False(t, isNonRetryableStatusCode(200))
}

func TestRetryCondition_WithRealResponse(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/503" {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		if r.URL.Path == "/404" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	client := resty.New().SetBaseURL(srv.URL)
	resp, err := client.R().Get("/503")
	requireNoError(t, err)
	assert.True(t, retryCondition(resp, nil))

	resp404, err := client.R().Get("/404")
	requireNoError(t, err)
	assert.False(t, retryCondition(resp404, nil))
}

func TestRetryCondition_WithError(t *testing.T) {
	// When resp is nil and err != nil, method is "" so not idempotent => no retry
	assert.False(t, retryCondition(nil, assert.AnError))
}

func TestRetryCondition_NilRespNoErr(t *testing.T) {
	assert.False(t, retryCondition(nil, nil))
}

func requireNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
