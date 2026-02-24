package activation_code

// ActivationCodeRequest is the body for updating the activation code.
type ActivationCodeRequest struct {
	ActivationCode string `json:"activationCode"` // Length between 32 and 39 characters. Hyphens are optional.
}

// OrganizationNameRequest is the body for updating the organization name.
type OrganizationNameRequest struct {
	OrganizationName string `json:"organizationName"`
}

// HistoryEntry represents a single activation code history record.
type HistoryEntry struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`    // ISO 8601 timestamp
	Note     string `json:"note"`
	Details  string `json:"details"`
}

// HistoryResponse is the response for listing activation code history.
type HistoryResponse struct {
	TotalCount int            `json:"totalCount"`
	Results    []HistoryEntry `json:"results"`
}

// HistoryNoteRequest is the body for adding a note to activation code history.
type HistoryNoteRequest struct {
	Note string `json:"note"`
}
