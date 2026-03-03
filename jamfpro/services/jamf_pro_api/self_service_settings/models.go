package self_service_settings

import "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"

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

// HistoryObject is an alias to the shared history item struct with string IDs.
type HistoryObject = shared.SharedHistoryItemString

// HistoryResponse is an alias to the shared history response struct with string IDs.
type HistoryResponse = shared.SharedHistoryResponseString

// AddHistoryNotesRequest is an alias to the shared history note request struct.
type AddHistoryNotesRequest = shared.SharedHistoryNoteRequest

// AddHistoryNotesResponse is the response for AddHistoryNotesV1 (201 Created).
type AddHistoryNotesResponse struct {
	ID   int    `json:"id"`
	Href string `json:"href"`
}
