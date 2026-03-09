package user

// API endpoint constants for user operations.
const (
	// EndpointUser is the endpoint for getting current authenticated user information.
	EndpointUser = "/api/user"

	// EndpointChangePasswordV1 is the endpoint for changing the current user's password.
	EndpointChangePasswordV1 = "/api/v1/user/change-password"

	// EndpointUpdateSession is the endpoint for updating the current user's session (change site).
	EndpointUpdateSession = "/api/user/updateSession"
)
