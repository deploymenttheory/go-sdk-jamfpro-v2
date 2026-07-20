package version_locking

import (
	"context"
	"net/http"

	"resty.dev/v3"
)

// DefaultMaxAttempts bounds how many times a write is re-driven after a
// rejected attempt. Each attempt costs one GET and one write.
const DefaultMaxAttempts = 3

// Fetch retrieves the resource's current server-side state, including its
// version locks.
type Fetch[T any] func(ctx context.Context) (*T, *resty.Response, error)

// Submit issues the write (PUT/POST) with the supplied request body.
type Submit[T any] func(ctx context.Context, request *T) (*T, *resty.Response, error)

// Option configures a locked write.
type Option func(*config)

type config struct{ maxAttempts int }

// WithMaxAttempts overrides DefaultMaxAttempts.
func WithMaxAttempts(n int) Option {
	return func(c *config) {
		if n > 0 {
			c.maxAttempts = n
		}
	}
}

// Update performs a version-locked update, handling the locking protocol so
// callers never have to.
//
// For each attempt it fetches current server state, copies every version lock
// in the tree onto the request (see SyncAll), and submits once.
//
// Two Jamf behaviours make the naive "submit, retry on error" approach unsafe,
// and this function exists to absorb both:
//
// Writes that succeed but report failure. Prestage endpoints have been observed
// returning HTTP 500 with an empty error list while having fully applied the
// change. Retrying such a write resubmits a body whose locks the server has
// already consumed, which then fails with OPTIMISTIC_LOCK_FAILED — so the
// caller is told a write failed when it actually succeeded, twice over. After
// any failed attempt this function re-reads the resource: if the version lock
// has advanced beyond the value submitted, the write landed and the refreshed
// resource is returned as a success.
//
// Stale locks. A genuine conflict means the submitted locks were already
// consumed, and resubmitting the same body can never succeed. Retrying is only
// useful with locks re-read from the server, which is what the next iteration
// does.
//
// The landed-check runs only for ambiguous failures, never for conflicts.
// Inferring success from an advanced lock is sound only when the server has not
// said whether it applied the write; a conflict is an explicit rejection, and
// running the check there would misattribute a competing writer's commit to us
// and silently drop the caller's change.
func Update[T any](ctx context.Context, request *T, fetch Fetch[T], submit Submit[T], opts ...Option) (*T, *resty.Response, error) {
	cfg := &config{maxAttempts: DefaultMaxAttempts}
	for _, o := range opts {
		o(cfg)
	}

	var lastResp *resty.Response
	var lastErr error

	for attempt := 1; attempt <= cfg.maxAttempts; attempt++ {
		current, resp, err := fetch(ctx)
		if err != nil {
			return nil, resp, err
		}

		submitted, hasLock := TopLock(current)
		SyncAll(current, request)

		result, resp, err := submit(ctx, request)
		if err == nil {
			return result, resp, nil
		}
		lastResp, lastErr = resp, err

		// A conflict is the server explicitly reporting that it rejected the
		// write, so there is nothing to reconcile — go straight to a retry
		// carrying freshly read locks. Checking whether state advanced here
		// would misread a competing writer's commit as our own success.
		if !isConflict(resp) {
			// Any other failure is ambiguous: these endpoints return 5xx on
			// writes they have fully applied. A lock beyond what we submitted
			// means ours was the write that moved it.
			if hasLock {
				if after, _, ferr := fetch(ctx); ferr == nil {
					if now, ok := TopLock(after); ok && now > submitted {
						return after, resp, nil
					}
				}
			}
		}

		if !worthRetrying(resp) {
			return nil, resp, err
		}
		// Next iteration re-reads state, so the retry carries fresh locks.
	}

	return nil, lastResp, lastErr
}

// Create performs a create, zeroing every version lock in the tree first.
// Jamf requires a lock of 0 on POST; a create has no prior state to echo.
func Create[T any](ctx context.Context, request *T, submit Submit[T]) (*T, *resty.Response, error) {
	ZeroAll(request)
	return submit(ctx, request)
}

// worthRetrying reports whether resubmitting with freshly read locks could
// plausibly succeed. A conflict is the canonical stale-lock signal; 5xx covers
// the transient server faults these endpoints exhibit. Everything else (bad
// request, unauthorised, not found) fails identically no matter the lock.
func worthRetrying(resp *resty.Response) bool {
	if resp == nil {
		return false
	}
	return isConflict(resp) || resp.StatusCode() >= http.StatusInternalServerError
}

// isConflict reports whether the server rejected the write because the
// supplied locks were stale (HTTP 409, OPTIMISTIC_LOCK_FAILED).
func isConflict(resp *resty.Response) bool {
	return resp != nil && resp.StatusCode() == http.StatusConflict
}
