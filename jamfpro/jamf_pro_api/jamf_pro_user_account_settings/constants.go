package jamf_pro_user_account_settings

// Endpoints for the Jamf Pro user account settings API (Jamf Pro API v1).
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-user-preferences-settings-keyid
const (
	// EndpointUserPreferencesSettingsV1 is the path prefix for user preferences settings.
	// Use with fmt.Sprintf("%s/%s", EndpointUserPreferencesSettingsV1, keyId).
	// Returns username, key, values (array of strings).
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-user-preferences-settings-keyid
	EndpointUserPreferencesSettingsV1 = "/api/v1/user/preferences/settings"

	// EndpointUserPreferencesV1 is the path prefix for user preferences (get/put/delete single setting).
	// Use with fmt.Sprintf("%s/%s", EndpointUserPreferencesV1, keyId).
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-user-preferences-keyid
	EndpointUserPreferencesV1 = "/api/v1/user/preferences"
)
