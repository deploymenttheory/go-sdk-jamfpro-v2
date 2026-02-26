package jamf_pro_server_url

// ResourceJamfProServerURL represents the Jamf Pro server URL settings.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-server-url
type ResourceJamfProServerURL struct {
	URL                    string `json:"url"`
	UnsecuredEnrollmentUrl string `json:"unsecuredEnrollmentUrl"`
}

// HistoryResponse is the paginated response for GetHistoryV1.
type HistoryResponse struct {
	TotalCount int             `json:"totalCount"`
	Results    []HistoryObject `json:"results"`
}

// HistoryObject represents a single history entry.
type HistoryObject struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Date     string  `json:"date"`
	Note     string  `json:"note"`
	Details  *string `json:"details"`
}

// CreateHistoryNoteRequest is the request body for CreateHistoryNoteV1.
type CreateHistoryNoteRequest struct {
	Note string `json:"note"`
}
