package jamf_pro_server_url

import "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"

// ResourceJamfProServerURL represents the Jamf Pro server URL settings.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-server-url
type ResourceJamfProServerURL struct {
	URL                    string `json:"url"`
	UnsecuredEnrollmentUrl string `json:"unsecuredEnrollmentUrl"`
}

// HistoryObject is an alias to the shared history item struct.
type HistoryObject = shared.SharedHistoryItem

// HistoryResponse is an alias to the shared history response struct.
type HistoryResponse = shared.SharedHistoryResponse

// CreateHistoryNoteRequest is an alias to the shared history note request struct.
type CreateHistoryNoteRequest = shared.SharedHistoryNoteRequest
