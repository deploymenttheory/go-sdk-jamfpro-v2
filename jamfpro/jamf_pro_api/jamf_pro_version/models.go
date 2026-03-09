package jamf_pro_version

// ResourceJamfProVersion is the response for GetV1.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-version
type ResourceJamfProVersion struct {
	Version *string `json:"Version,omitempty"`
}
