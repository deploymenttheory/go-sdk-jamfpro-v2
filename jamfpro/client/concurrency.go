package client

import "context"

// semaphore is a buffered-channel based concurrency limiter.
// A nil semaphore means unlimited concurrent requests (default).
//
// Jamf Pro API scalability guidance recommends no more than 5 concurrent
// connections to avoid disrupting other Jamf Pro tasks and managed devices.
//
// See: https://developer.jamf.com/jamf-pro/docs/jamf-pro-api-scalability-best-practices
type semaphore struct {
	ch chan struct{}
}

// newSemaphore creates a semaphore that allows at most n concurrent holders.
func newSemaphore(n int) *semaphore {
	return &semaphore{ch: make(chan struct{}, n)}
}

// acquire blocks until a slot is available or ctx is cancelled.
// Returns ctx.Err() if the context is cancelled while waiting.
func (s *semaphore) acquire(ctx context.Context) error {
	select {
	case s.ch <- struct{}{}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// release returns a slot to the semaphore.
func (s *semaphore) release() {
	select {
	case <-s.ch:
	default:
		// Guard against releasing more than acquired; should not happen.
	}
}
