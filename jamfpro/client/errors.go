package client

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
)

const (
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusConflict            = 409
	StatusPreconditionFailed  = 412
	StatusUnprocessableEntity = 422
	StatusTooManyRequests     = 429
	StatusInternalServerError = 500
	StatusServiceUnavailable  = 503
)

// APIError represents an error response from the Jamf Pro API.
type APIError struct {
	Code       string
	Message    string
	StatusCode int
	Status     string
	Endpoint   string
	Method     string
}

// Error implements the error interface.
func (e *APIError) Error() string {
	if e.Code != "" {
		return fmt.Sprintf("Jamf Pro API error (%d %s) [%s] at %s %s: %s",
			e.StatusCode, e.Status, e.Code, e.Method, e.Endpoint, e.Message)
	}
	return fmt.Sprintf("Jamf Pro API error (%d %s) at %s %s: %s",
		e.StatusCode, e.Status, e.Method, e.Endpoint, e.Message)
}

// jamfErrorBody is a common shape for Jamf error responses.
type jamfErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ParseErrorResponse parses an error response from the API.
func ParseErrorResponse(body []byte, statusCode int, status, method, endpoint string, logger *zap.Logger) error {
	apiError := &APIError{
		StatusCode: statusCode,
		Status:     status,
		Endpoint:   endpoint,
		Method:     method,
	}
	var parsed jamfErrorBody
	if err := json.Unmarshal(body, &parsed); err == nil && (parsed.Code != "" || parsed.Message != "") {
		apiError.Code = parsed.Code
		apiError.Message = parsed.Message
	} else {
		apiError.Message = string(body)
		if apiError.Message == "" {
			apiError.Message = defaultMessageForStatus(statusCode)
		}
	}
	logger.Error("API error response",
		zap.Int("status_code", statusCode),
		zap.String("method", method),
		zap.String("endpoint", endpoint),
		zap.String("message", apiError.Message))
	return apiError
}

func defaultMessageForStatus(statusCode int) string {
	switch statusCode {
	case 400:
		return "The request could not be understood by the server due to malformed syntax."
	case 401:
		return "The request has not been applied because it lacks valid authentication credentials for the target resource."
	case 403:
		return "Authentication required or token invalid. The server understood the request but refuses to authorize it."
	case 404:
		return "The server has not found anything matching the Request-URI."
	case 409:
		return "The request could not be completed due to a conflict with the current state of the resource."
	case 412:
		return "One or more conditions given in the request header fields evaluated to false when tested on the server."
	case 422:
		return "The request has correct syntax, but has a field with a bad value, such as an ID which does not exist, an illegal enum value, or a field is missing entirely."
	case 429:
		return "The user has sent too many requests in a given amount of time (rate limiting)."
	case 500:
		return "The server encountered an unexpected condition which prevented it from fulfilling the request."
	case 503:
		return "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server."
	default:
		return "Unknown error"
	}
}

// IsNotFound checks if the error is a not found error (404).
func IsNotFound(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == StatusNotFound
	}
	return false
}

// IsUnauthorized checks if the error is an authentication error (401).
func IsUnauthorized(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == StatusUnauthorized
	}
	return false
}

// IsBadRequest checks if the error is a bad request error (400).
func IsBadRequest(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == StatusBadRequest
	}
	return false
}

// IsServerError checks if the error is a server error (5xx).
func IsServerError(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode >= 500 && apiErr.StatusCode < 600
	}
	return false
}
