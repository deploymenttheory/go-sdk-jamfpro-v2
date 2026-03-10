package mocks

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"resty.dev/v3"
)

// NewMockResponse creates a resty.Response for testing purposes.
// This helper constructs a valid resty.Response with the provided status code,
// headers, and body content. It properly initializes all required fields to avoid
// nil pointer dereferences when calling methods like Bytes().
func NewMockResponse(statusCode int, headers http.Header, body []byte) *resty.Response {
	if headers == nil {
		headers = make(http.Header)
	}
	if body == nil {
		body = []byte{}
	}
	
	status := http.StatusText(statusCode)
	if status == "" {
		status = fmt.Sprintf("%d", statusCode)
	}
	
	bodyReader := io.NopCloser(bytes.NewReader(body))

	// Create a minimal Request to avoid nil pointer dereferences
	req := &resty.Request{
		URL:                "",
		DoNotParseResponse: false,
	}

	resp := &resty.Response{
		Request: req,
		// Body must be set on the resty.Response directly so that resp.Bytes()
		// (which reads r.Body via readAll) can access the body content.
		Body: io.NopCloser(bytes.NewReader(body)),
		RawResponse: &http.Response{
			StatusCode: statusCode,
			Status:     status,
			Header:     headers,
			Body:       bodyReader,
		},
	}

	return resp
}
