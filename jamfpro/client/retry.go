package client

import "resty.dev/v3"

// retryCondition is the resty AddRetryConditions callback.
// It returns true when the request should be retried.
//
// Retry rules (aligned with Jamf Pro API scalability best practices):
//   - Only idempotent HTTP methods are retried (GET, HEAD, OPTIONS, PUT, DELETE).
//     POST and PATCH are never retried to prevent duplicate resource creation.
//   - Transient server errors (500, 502, 503, 504) are retried.
//   - 429 Too Many Requests is retried using resty's exponential backoff.
//     Jamf Pro does not send Retry-After headers; backoff timing is derived
//     from RetryWaitTime/RetryMaxWaitTime configured on the resty client.
//   - Definitive client errors (4xx excluding 429) are never retried.
//   - Network-level errors (resp == nil) are retried if method is idempotent.
//
// See: https://developer.jamf.com/jamf-pro/docs/jamf-pro-api-scalability-best-practices
func retryCondition(resp *resty.Response, err error) bool {
	method := ""
	if resp != nil && resp.Request != nil {
		method = resp.Request.Method
	}

	// Network / transport error with no response — retry if safe to do so.
	if err != nil {
		return isIdempotentMethod(method)
	}

	if resp == nil {
		return false
	}

	code := resp.StatusCode()

	// Never retry definitive client-side failures.
	if isNonRetryableStatusCode(code) {
		return false
	}

	// Only retry safe methods.
	if !isIdempotentMethod(method) {
		return false
	}

	return isTransientStatusCode(code) || code == StatusTooManyRequests
}

// isIdempotentMethod returns true for HTTP methods that are safe to retry:
// GET, HEAD, OPTIONS, PUT, DELETE.
// POST and PATCH are excluded because retrying them may produce duplicate
// side-effects (created resources, sent notifications, etc.).
func isIdempotentMethod(method string) bool {
	switch method {
	case "GET", "HEAD", "OPTIONS", "PUT", "DELETE":
		return true
	default:
		return false
	}
}

// isTransientStatusCode returns true for server-side errors that are likely
// temporary and worth retrying with exponential backoff.
func isTransientStatusCode(code int) bool {
	switch code {
	case 500, 502, 503, 504:
		return true
	default:
		return false
	}
}

// isNonRetryableStatusCode returns true for definitive client-side errors
// that will not succeed on retry regardless of timing.
func isNonRetryableStatusCode(code int) bool {
	switch code {
	case 400, 401, 402, 403, 404, 405, 406, 407, 409, 410,
		411, 412, 413, 414, 415, 416, 417, 422, 423, 424,
		426, 428, 431, 451:
		return true
	default:
		return false
	}
}
