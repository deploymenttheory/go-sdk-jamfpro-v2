package client

import (
	"resty.dev/v3"
)

// retryCondition is the resty AddRetryConditions callback.
// It returns true when the request should be retried.
//
// Resty handles the actual retry scheduling: it uses capped exponential backoff
// with full jitter (min * 2^attempt, capped at max, then halved and jittered).
// The min and max bounds are set by RetryWaitTime and RetryMaxWaitTime on the
// resty client. This function only decides whether a given response warrants
// another attempt.
//
// Retry rules (aligned with Jamf Pro API scalability best practices):
//   - Only idempotent HTTP methods are retried (GET, HEAD, OPTIONS, PUT, DELETE).
//     POST and PATCH are never retried to prevent duplicate resource creation.
//   - Transient server errors (500, 502, 503, 504) are retried.
//   - 408 Request Timeout is retried: the request did not reach the server and
//     is safe to resend on idempotent methods.
//   - Jamf Pro does not implement rate limiting and does not return 429.
//   - Definitive client errors (4xx excluding 408) are never retried.
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

	return isTransientStatusCode(code)
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

// isTransientStatusCode returns true for errors that are likely temporary and
// worth retrying with exponential backoff.
//
// 408 Request Timeout: the request did not complete before the server's
// timeout; the server did not process it so retrying is safe.
// 500, 502, 503, 504: standard transient server-side errors.
func isTransientStatusCode(code int) bool {
	switch code {
	case 408, 500, 502, 503, 504:
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
		426, 428, 429, 431, 451:
		return true
	default:
		return false
	}
}
