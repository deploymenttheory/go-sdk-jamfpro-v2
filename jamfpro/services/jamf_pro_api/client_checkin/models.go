package client_checkin

import "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"

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

// ResourceClientCheckinHistoryEntry is an alias to the shared history item struct with string IDs.
type ResourceClientCheckinHistoryEntry = shared.SharedHistoryItemString

// ResourceClientCheckinHistory is an alias to the shared history response struct with string IDs.
type ResourceClientCheckinHistory = shared.SharedHistoryResponseString

// RequestClientCheckinHistoryNote is an alias to the shared history note request struct.
type RequestClientCheckinHistoryNote = shared.SharedHistoryNoteRequest

// CreateHistoryResponse is the response for AddHistoryNoteV3 (POST history).
type CreateHistoryResponse struct {
	ID   string `json:"id"`
	HREF string `json:"href"`
}
