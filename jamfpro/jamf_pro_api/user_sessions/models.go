package user_sessions

// ResourceActiveUserSession represents a currently logged-in user session.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-user-sessions-active
type ResourceActiveUserSession struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	LastAccess string `json:"lastAccess,omitempty"` // ISO 8601 timestamp
}

// ResourceUserSessionCount represents the count of currently logged-in users.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-user-sessions-count
type ResourceUserSessionCount struct {
	Count int `json:"count"`
}

type ResourceActiveUserSessionV2 struct {
	Username         string  `json:"username"`
	SessionID        string  `json:"sessionId"`
	LastAccessedTime string  `json:"lastAccessedTime"`
	CreationTime     string  `json:"creationTime"`
	UserAgent        *string `json:"userAgent"`
	IPAddress        *string `json:"ipAddress"`
}

type ResourceUserSessionCountV2 struct {
	ActiveUserCount int `json:"activeUserCount"`
}
