package icons

// ResourceIcon represents icon metadata (get or upload response).
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-icon-id
type ResourceIcon struct {
	ID   int    `json:"id"`
	URL  string `json:"url"`
	Name string `json:"name"`
}
