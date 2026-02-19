package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestAPIError_Error(t *testing.T) {
	e := &APIError{Code: "ERR", Message: "msg", StatusCode: 404, Status: "Not Found", Method: "GET", Endpoint: "/api"}
	s := e.Error()
	assert.NotEmpty(t, s)
	assert.GreaterOrEqual(t, len(s), 10)

	e2 := &APIError{Message: "m", StatusCode: 500, Status: "Internal", Method: "POST", Endpoint: "/x"}
	assert.NotEmpty(t, e2.Error())
}

func TestParseErrorResponse(t *testing.T) {
	logger := zap.NewNop()
	tests := []struct {
		name       string
		body       []byte
		statusCode int
		status     string
		wantMsg    string
	}{
		{"json body", []byte(`{"code":"ERR","message":"not found"}`), 404, "Not Found", "not found"},
		{"empty body", []byte(``), 404, "Not Found", "Resource not found"},
		{"plain body", []byte("plain"), 500, "Internal", "plain"},
		{"default 400", nil, 400, "Bad Request", "Bad request"},
		{"default 401", nil, 401, "Unauthorized", "Authentication required"},
		{"default 403", nil, 403, "Forbidden", "Forbidden"},
		{"default 409", nil, 409, "Conflict", "Conflict"},
		{"default 429", nil, 429, "Too Many", "Too many requests"},
		{"default 503", nil, 503, "Unavailable", "Service unavailable"},
		{"unknown status", nil, 418, "Teapot", "Unknown error"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ParseErrorResponse(tt.body, tt.statusCode, tt.status, "GET", "/api", logger)
			require.Error(t, err)
			apiErr, ok := err.(*APIError)
			require.True(t, ok, "expected *APIError, got %T", err)
			assert.Equal(t, tt.wantMsg, apiErr.Message)
		})
	}
}

func TestIsNotFound(t *testing.T) {
	assert.False(t, IsNotFound(nil))
	assert.False(t, IsNotFound(&APIError{StatusCode: 500}))
	assert.True(t, IsNotFound(&APIError{StatusCode: StatusNotFound}))
}

func TestIsUnauthorized(t *testing.T) {
	assert.True(t, IsUnauthorized(&APIError{StatusCode: StatusUnauthorized}))
	assert.False(t, IsUnauthorized(&APIError{StatusCode: 200}))
}

func TestIsBadRequest(t *testing.T) {
	assert.True(t, IsBadRequest(&APIError{StatusCode: StatusBadRequest}))
}

func TestIsServerError(t *testing.T) {
	assert.True(t, IsServerError(&APIError{StatusCode: 500}))
	assert.True(t, IsServerError(&APIError{StatusCode: 503}))
	assert.False(t, IsServerError(&APIError{StatusCode: 400}))
}

func TestIsNotFound_NonAPIError(t *testing.T) {
	assert.False(t, IsNotFound(assert.AnError))
}

func TestIsUnauthorized_NonAPIError(t *testing.T) {
	assert.False(t, IsUnauthorized(assert.AnError))
}

func TestIsBadRequest_NonAPIError(t *testing.T) {
	assert.False(t, IsBadRequest(assert.AnError))
}

func TestIsServerError_NonAPIError(t *testing.T) {
	assert.False(t, IsServerError(assert.AnError))
}

func TestParseErrorResponse_DefaultMessageForStatus(t *testing.T) {
	logger := zap.NewNop()
	for code, want := range map[int]string{
		StatusForbidden: "Forbidden",
		StatusConflict:  "Conflict",
	} {
		err := ParseErrorResponse(nil, code, "", "GET", "/", logger)
		require.Error(t, err)
		assert.Equal(t, want, err.(*APIError).Message)
	}
}
