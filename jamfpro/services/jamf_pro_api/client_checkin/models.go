package client_checkin

// ResourceClientCheckinSettings represents client check-in settings (singleton).
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-check-in
type ResourceClientCheckinSettings struct {
	CheckInFrequency                 int  `json:"checkInFrequency"`
	CreateHooks                      bool `json:"createHooks"`
	HookLog                          bool `json:"hookLog"`
	HookPolicies                     bool `json:"hookPolicies"`
	CreateStartupScript              bool `json:"createStartupScript"`
	StartupLog                       bool `json:"startupLog"`
	StartupPolicies                  bool `json:"startupPolicies"`
	StartupSsh                       bool `json:"startupSsh"`
	EnableLocalConfigurationProfiles bool `json:"enableLocalConfigurationProfiles"`
}

// ResourceClientCheckinHistoryEntry represents a single client check-in history entry.
type ResourceClientCheckinHistoryEntry struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}

// ResourceClientCheckinHistory is the response for GetHistoryV3.
type ResourceClientCheckinHistory struct {
	TotalCount int                                `json:"totalCount"`
	Results    []ResourceClientCheckinHistoryEntry `json:"results"`
}

// RequestClientCheckinHistoryNote is the body for AddHistoryNoteV3.
type RequestClientCheckinHistoryNote struct {
	Note string `json:"note"`
}
