package jamf_account_preferences

// Endpoints for the Jamf Pro account preferences API (Jamf Pro API v3).
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-account-preferences
const (
	// EndpointAccountPreferencesV3 is the path for account preferences (GET/PATCH).
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-account-preferences
	EndpointAccountPreferencesV3 = "/api/v3/account-preferences"
)

// User interface display theme values.
const (
	ThemeMatchSystem = "MATCH_SYSTEM"
	ThemeLight       = "LIGHT"
	ThemeDark        = "DARK"
)

// Search method values for computer, mobile device, and user search preferences.
const (
	SearchExactMatch  = "EXACT_MATCH"
	SearchStartsWith  = "STARTS_WITH"
	SearchContains    = "CONTAINS"
)
