package jamf_pro_user_account_settings

// ResourceUserPreferencesSettings represents the response from GET /api/v1/user/preferences/settings/{keyId}.
// Contains the user preferences for the authenticated user and key.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-user-preferences-settings-keyid
type ResourceUserPreferencesSettings struct {
	Username string   `json:"username"`
	Key      string   `json:"key"`
	Values   []string `json:"values"`
}

// RequestUserPreferences is the body for PutV1 (PUT).
// Key-value pairs to persist; keys are setting names, values are strings.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-user-preferences-keyid
type RequestUserPreferences map[string]string
