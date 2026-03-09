package constants

// ============================================================================
// Authentication Methods
// ============================================================================

// AuthMethod constants for the Jamf Pro authentication methods.
const (
	AuthMethodOAuth2 = "oauth2"
	AuthMethodBasic  = "basic"
)

// ============================================================================
// Authentication Endpoints
// ============================================================================

// Jamf Pro API authentication and token management endpoints.
// See: https://developer.jamf.com/jamf-pro/docs/classic-api-authentication-changes
const (
	EndpointOAuthToken      = "/api/v1/oauth/token"
	EndpointBearerToken     = "/api/v1/auth/token"
	EndpointInvalidateToken = "/api/v1/auth/invalidate-token"
	EndpointKeepAliveToken  = "/api/v1/auth/keep-alive"
)
