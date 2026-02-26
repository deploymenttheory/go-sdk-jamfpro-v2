package self_service_settings

// ResourceSelfServiceSettings is the self-service settings resource (get/update).
type ResourceSelfServiceSettings struct {
	InstallSettings       InstallSettings       `json:"installSettings"`
	LoginSettings         LoginSettings         `json:"loginSettings"`
	ConfigurationSettings ConfigurationSettings `json:"configurationSettings"`
}

// InstallSettings holds install options.
type InstallSettings struct {
	InstallAutomatically bool   `json:"installAutomatically"`
	InstallLocation      string `json:"installLocation"`
}

// LoginSettings holds login-level options.
type LoginSettings struct {
	UserLoginLevel  string `json:"userLoginLevel"`
	AllowRememberMe bool   `json:"allowRememberMe"`
	UseFido2        bool   `json:"useFido2"`
	AuthType        string `json:"authType"`
}

// ConfigurationSettings holds notifications and landing page options.
type ConfigurationSettings struct {
	NotificationsEnabled  bool   `json:"notificationsEnabled"`
	AlertUserApprovedMdm  bool   `json:"alertUserApprovedMdm"`
	DefaultLandingPage    string `json:"defaultLandingPage"`
	DefaultHomeCategoryId int    `json:"defaultHomeCategoryId"`
	BookmarksName         string `json:"bookmarksName"`
}

// HistoryObject represents a single Self Service settings history entry.
type HistoryObject struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}

// HistoryResponse is the response for GetHistoryV1.
type HistoryResponse struct {
	TotalCount int            `json:"totalCount"`
	Results    []HistoryObject `json:"results"`
}

// AddHistoryNotesRequest is the request body for AddHistoryNotesV1.
type AddHistoryNotesRequest struct {
	Note string `json:"note"`
}

// AddHistoryNotesResponse is the response for AddHistoryNotesV1 (201 Created).
type AddHistoryNotesResponse struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
}
