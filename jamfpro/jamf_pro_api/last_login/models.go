package last_login

// ResourceLastLogin represents the last login event information.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-last-login
type ResourceLastLogin struct {
	Date string `json:"date"` // ISO 8601 timestamp of the last login event
}
