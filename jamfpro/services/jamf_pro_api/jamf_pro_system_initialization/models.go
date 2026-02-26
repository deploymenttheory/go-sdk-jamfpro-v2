package jamf_pro_system_initialization

// ResourceDatabasePassword is the request payload for database password initialization.
// Used when setting up the database password during Jamf Pro startup.
type ResourceDatabasePassword struct {
	Password string `json:"password"`
}

// ResourceSystemInitialize is the request payload for initializing a fresh Jamf Pro Server installation.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-system-initialize
type ResourceSystemInitialize struct {
	ActivationCode  string `json:"activationCode"`
	InstitutionName string `json:"institutionName"`
	EulaAccepted    bool   `json:"eulaAccepted"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Email           string `json:"email,omitempty"`
	JssUrl          string `json:"jssUrl"`
}

// ResourcePlatformInitialize is the request payload for initializing a Jamf Pro Server with OIDC SSO
// and a federated user (no password required).
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-system-platform-initialize
type ResourcePlatformInitialize struct {
	ActivationCode  string `json:"activationCode"`
	InstitutionName string `json:"institutionName"`
	EulaAccepted    bool   `json:"eulaAccepted"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	JssUrl          string `json:"jssUrl"`
}
