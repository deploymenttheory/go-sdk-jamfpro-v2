package locales

// ResourceLocale represents a locale resource.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-locales
type ResourceLocale struct {
	Description string `json:"description"`
	Identifier  string `json:"identifier"`
}
