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
