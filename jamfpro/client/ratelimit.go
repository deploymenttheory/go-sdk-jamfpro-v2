package client

import (
	"sync"
	"time"
)

// responseTimeTracker computes an exponential moving average (EMA) of API
// response durations and signals when the caller should pause before the next
// request.
//
// Jamf Pro does not emit rate-limit HTTP headers (no Retry-After,
// X-RateLimit-Remaining, or X-RateLimit-Reset). Throttling must be inferred
// entirely from observed response times, as recommended by Jamf:
//
//	"Measure response times and dynamically adjust time between requests
//	 accordingly."
//
// When the server begins responding more slowly than its own EMA baseline,
// the excess latency is returned as a suggested pause. This gives the server
// time to recover before the next request arrives without imposing a fixed
// static delay.
//
// See: https://developer.jamf.com/jamf-pro/docs/jamf-pro-api-scalability-best-practices
type responseTimeTracker struct {
	mu          sync.Mutex
	ema         time.Duration
	alpha       float64 // EMA smoothing factor: 0 < alpha < 1
	initialized bool
}

// newResponseTimeTracker returns a tracker using alpha=0.2.
// alpha=0.2 weights the most recent sample at 20%, providing a stable
// baseline that still reacts to sustained slowdowns within a few samples.
func newResponseTimeTracker() *responseTimeTracker {
	return &responseTimeTracker{alpha: 0.2}
}

// record adds a response duration sample and returns the adaptive delay the
// caller should sleep before issuing the next request.
//
// When the observed duration exceeds 2× the current EMA, the server is under
// measurable pressure. The excess (d − ema) is returned as the suggested
// pause, capped at adaptiveDelayMax to prevent unbounded stalls.
// No delay is returned while the server responds at or below its baseline.
func (r *responseTimeTracker) record(d time.Duration) time.Duration {
	r.mu.Lock()
	defer r.mu.Unlock()

	if !r.initialized {
		r.ema = d
		r.initialized = true
		return 0
	}

	// EMA update: ema = alpha*current + (1-alpha)*previous
	r.ema = time.Duration(float64(d)*r.alpha + float64(r.ema)*(1-r.alpha))

	// Server is responding at or below its own baseline — no pause needed.
	if d <= 2*r.ema {
		return 0
	}

	// Return the excess above the EMA, capped to prevent excessive stalls.
	excess := d - r.ema
	if excess > adaptiveDelayMax {
		return adaptiveDelayMax
	}
	return excess
}
