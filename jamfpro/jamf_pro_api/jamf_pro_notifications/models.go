package jamf_pro_notifications

// ResourceNotification represents a single notification in the Jamf Pro API.
type ResourceNotification struct {
	Type   string         `json:"type"`   // The type of notification
	ID     string         `json:"id"`     // The unique identifier for the notification
	Params map[string]any `json:"params"` // Additional parameters for the notification
}
