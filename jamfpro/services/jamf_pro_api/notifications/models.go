package notifications

// ResourceNotification represents a single notification.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-notifications
type ResourceNotification struct {
	Type   string         `json:"type"`
	ID     string         `json:"id"`
	Params map[string]any `json:"params"`
}
