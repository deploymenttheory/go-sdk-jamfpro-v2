package user_sessions

// ResourceActiveUserSession represents a currently logged-in user session.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-user-sessions-active
type ResourceActiveUserSession struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	LastAccess string `json:"lastAccess,omitempty"` // ISO 8601 timestamp
}

// ListActiveUserSessionsResponse is the response for GetActiveV1.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-user-sessions-active
type ListActiveUserSessionsResponse struct {
	TotalCount int                          `json:"totalCount"`
	Results    []ResourceActiveUserSession  `json:"results"`
}

// ResourceUserSessionCount represents the count of currently logged-in users.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-user-sessions-count
type ResourceUserSessionCount struct {
	Count int `json:"count"`
}
