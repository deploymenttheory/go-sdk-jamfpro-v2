package client

import "time"

const (
	UserAgentBase = "go-sdk-jamfpro-v2"
)

// HTTP client defaults. Aligned with Jamf Pro API scalability guidance:
// max 5 concurrent connections, exponential backoff for transient errors.
const (
	DefaultTimeout   = 120 * time.Second
	MaxRetries       = 3
	RetryWaitTime    = 2 * time.Second
	RetryMaxWaitTime = 30 * time.Second

	// DefaultMaxConcurrentRequests is the Jamf-recommended maximum of 5
	// concurrent API connections. Set to 0 to use WithMaxConcurrentRequests.
	DefaultMaxConcurrentRequests = 5

	// DefaultPageSize is the number of results fetched per page in GetPaginated.
	DefaultPageSize = 200

	// adaptiveDelayMax is the ceiling applied to the adaptive inter-request
	// delay computed from response-time EMA tracking. Prevents unbounded
	// stalls when the server is under extreme load.
	adaptiveDelayMax = 5 * time.Second
)

// Jamf Pro API authentication and token management endpoints.
// See: https://developer.jamf.com/jamf-pro/docs/classic-api-authentication-changes
const (
	oAuthTokenEndpoint      = "/api/v1/oauth/token"
	bearerTokenEndpoint     = "/api/v1/auth/token"
	invalidateTokenEndpoint = "/api/v1/auth/invalidate-token"
	keepAliveTokenEndpoint  = "/api/v1/auth/keep-alive"
)
