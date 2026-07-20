package client

import (
	"context"

	"resty.dev/v3"
)

// noRetryKey marks a request that must never be retried by the transport.
type noRetryContextKey struct{}

// DisableRetry opts this request out of the transport's automatic retries.
//
// The retry layer assumes PUT and DELETE are idempotent, which is true for most
// of the API but false for endpoints using optimistic locking. Those requests
// carry a versionLock that the server consumes on the first successful write,
// so a replayed body is rejected with OPTIMISTIC_LOCK_FAILED — and because
// these endpoints can report a fault on a write they actually applied, the
// replay happens after the change already landed. The caller is then told the
// write failed when it succeeded.
//
// Retrying a version-locked write is only meaningful with a lock re-read from
// the server, which is a decision for the version_locking layer, not a blind
// transport replay. Services performing such writes should call this.
func (b *RequestBuilder) DisableRetry() *RequestBuilder {
	b.req.SetContext(context.WithValue(b.req.Context(), noRetryContextKey{}, true))
	return b
}

// retryDisabled reports whether the request opted out of transport retries.
func retryDisabled(resp *resty.Response) bool {
	if resp == nil || resp.Request == nil {
		return false
	}
	ctx := resp.Request.Context()
	if ctx == nil {
		return false
	}
	disabled, _ := ctx.Value(noRetryContextKey{}).(bool)
	return disabled
}
