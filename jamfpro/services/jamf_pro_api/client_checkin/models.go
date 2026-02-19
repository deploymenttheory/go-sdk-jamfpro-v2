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
