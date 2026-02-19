package client

import (
	"net/http"
	"strings"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"go.uber.org/zap"
	"resty.dev/v3"
)

// validateResponse validates the HTTP response before processing.
func (t *Transport) validateResponse(resp *resty.Response, method, path string) error {
	bodyLen := len(resp.String())
	if resp.Header().Get("Content-Length") == "0" || bodyLen == 0 {
		t.logger.Debug("Empty response received",
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status_code", resp.StatusCode()))
		return nil
	}
	if !resp.IsError() && bodyLen > 0 {
		contentType := resp.Header().Get("Content-Type")
		if contentType != "" && !strings.HasPrefix(contentType, mime.ApplicationJSON) && !strings.HasPrefix(contentType, mime.ApplicationXML) {
			t.logger.Warn("Unexpected Content-Type in response",
				zap.String("method", method),
				zap.String("path", path),
				zap.String("content_type", contentType))
		}
	}
	return nil
}

// IsResponseSuccess returns true if the response status code is 2xx.
func IsResponseSuccess(resp *interfaces.Response) bool {
	if resp == nil {
		return false
	}
	return resp.StatusCode >= 200 && resp.StatusCode < 300
}

// IsResponseError returns true if the response status code is 4xx or 5xx.
func IsResponseError(resp *interfaces.Response) bool {
	if resp == nil {
		return false
	}
	return resp.StatusCode >= 400
}

// GetResponseHeader returns a header value from the response by key.
func GetResponseHeader(resp *interfaces.Response, key string) string {
	if resp == nil || resp.Headers == nil {
		return ""
	}
	return resp.Headers.Get(key)
}

// GetResponseHeaders returns all headers from the response.
func GetResponseHeaders(resp *interfaces.Response) http.Header {
	if resp == nil {
		return make(http.Header)
	}
	return resp.Headers
}
