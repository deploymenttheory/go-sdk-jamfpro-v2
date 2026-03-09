package devices

// ResourceGroup represents a group that a device belongs to.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-devices-id-groups
type ResourceGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
