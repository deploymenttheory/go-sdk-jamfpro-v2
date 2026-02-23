package jamf_pro_server_url

// ResourceJamfProServerURL represents the Jamf Pro server URL settings.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-server-url
type ResourceJamfProServerURL struct {
	URL                    string `json:"url"`
	UnsecuredEnrollmentUrl string `json:"unsecuredEnrollmentUrl"`
}
