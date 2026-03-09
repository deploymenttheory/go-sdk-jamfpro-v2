package client

import (
	"net/http"
	"strings"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
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
		if contentType != "" &&
			!strings.HasPrefix(contentType, constants.ApplicationJSON) &&
			!strings.HasPrefix(contentType, constants.ApplicationXML) &&
			!strings.HasPrefix(contentType, constants.TextXML) {
			t.logger.Warn("Unexpected Content-Type in response",
				zap.String("method", method),
				zap.String("path", path),
				zap.String("content_type", contentType))
		}
	}
	return nil
}

// IsResponseSuccess returns true if the response status code is 2xx.
// Delegates to resty's native IsSuccess() method.
func IsResponseSuccess(resp *resty.Response) bool {
	if resp == nil {
		return false
	}
	return resp.IsSuccess()
}

// IsResponseError returns true if the response status code is 4xx or 5xx.
// Delegates to resty's native IsError() method.
func IsResponseError(resp *resty.Response) bool {
	if resp == nil {
		return false
	}
	return resp.IsError()
}

// GetResponseHeader returns a header value from the response by key.
func GetResponseHeader(resp *resty.Response, key string) string {
	if resp == nil {
		return ""
	}
	return resp.Header().Get(key)
}

// GetResponseHeaders returns all headers from the response.
func GetResponseHeaders(resp *resty.Response) http.Header {
	if resp == nil {
		return make(http.Header)
	}
	return resp.Header()
}
