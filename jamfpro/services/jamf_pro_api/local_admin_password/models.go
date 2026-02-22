package local_admin_password

// SettingsResource represents the Jamf Pro LAPS settings.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-settings
type SettingsResource struct {
	AutoDeployEnabled        bool `json:"autoDeployEnabled"`
	PasswordRotationTime     int  `json:"passwordRotationTime"`
	AutoRotateEnabled        bool `json:"autoRotateEnabled"`
	AutoRotateExpirationTime int  `json:"autoRotateExpirationTime"`
}

// SetPasswordRequest is the body for setting LAPS passwords.
type SetPasswordRequest struct {
	LapsUserPasswordList []LapsUserPassword `json:"lapsUserPasswordList"`
}

// LapsUserPassword represents a username/password pair for LAPS.
type LapsUserPassword struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// SetPasswordResponse represents the response after setting LAPS passwords.
type SetPasswordResponse struct {
	LapsUserPasswordList []LapsUserPasswordResponse `json:"lapsUserPasswordList"`
}

// LapsUserPasswordResponse represents the response for a single user after setting password.
type LapsUserPasswordResponse struct {
	Username string `json:"username"`
}

// PendingRotationsResponse represents the response for pending LAPS rotations.
type PendingRotationsResponse struct {
	TotalCount int                  `json:"totalCount"`
	Results    []PendingRotationItem `json:"results"`
}

// PendingRotationItem represents a single pending rotation entry.
type PendingRotationItem struct {
	LapsUser    LapsUser `json:"lapsUser"`
	CreatedDate string   `json:"createdDate"`
}

// LapsUser represents a LAPS user account.
type LapsUser struct {
	ClientManagementID string `json:"clientManagementId"`
	GUID               string `json:"guid"`
	Username           string `json:"username"`
	UserSource         string `json:"userSource"`
}

// PasswordHistoryResponse represents the password view history for a specific username.
type PasswordHistoryResponse struct {
	TotalCount int                   `json:"totalCount"`
	Results    []PasswordHistoryItem `json:"results"`
}

// PasswordHistoryItem represents a single password history entry.
type PasswordHistoryItem struct {
	Password       string         `json:"password"`
	DateLastSeen   string         `json:"dateLastSeen"`
	ExpirationTime string         `json:"expirationTime"`
	Audits         []PasswordAudit `json:"audits"`
}

// PasswordAudit represents who viewed a password and when.
type PasswordAudit struct {
	ViewedBy string `json:"viewedBy"`
	DateSeen string `json:"dateSeen"`
}

// CurrentPasswordResponse represents the current LAPS password for a user.
type CurrentPasswordResponse struct {
	Password string `json:"password"`
}

// FullHistoryResponse represents the complete history of all local admin passwords.
type FullHistoryResponse struct {
	TotalCount int              `json:"totalCount"`
	Results    []FullHistoryEvent `json:"results"`
}

// FullHistoryEvent represents a single event in the full LAPS history.
type FullHistoryEvent struct {
	Username   string `json:"username"`
	EventType  string `json:"eventType"`
	EventTime  string `json:"eventTime"`
	ViewedBy   string `json:"viewedBy"`
	UserSource string `json:"userSource"`
}

// CapableAccountsResponse represents LAPS-capable accounts on a device.
type CapableAccountsResponse struct {
	TotalCount int              `json:"totalCount"`
	Results    []CapableAccount `json:"results"`
}

// CapableAccount represents a single LAPS-capable account.
type CapableAccount struct {
	ClientManagementID string `json:"clientManagementId"`
	GUID               string `json:"guid"`
	Username           string `json:"username"`
	UserSource         string `json:"userSource"`
}
