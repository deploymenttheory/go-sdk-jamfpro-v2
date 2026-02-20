package reenrollment

// ResourceReenrollmentSettings is the re-enrollment settings resource (get/update).
type ResourceReenrollmentSettings struct {
	FlushPolicyHistory              bool   `json:"isFlushPolicyHistoryEnabled"`
	FlushLocationInformation        bool   `json:"isFlushLocationInformationEnabled"`
	FlushLocationInformationHistory bool   `json:"isFlushLocationInformationHistoryEnabled"`
	FlushExtensionAttributes        bool   `json:"isFlushExtensionAttributesEnabled"`
	FlushSoftwareUpdatePlans        bool   `json:"isFlushSoftwareUpdatePlansEnabled"`
	FlushMdmQueue                   string `json:"flushMDMQueue"`
}

// ReenrollmentHistoryObject is a single re-enrollment history entry.
type ReenrollmentHistoryObject struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details,omitempty"`
}

// ReenrollmentHistoryResponse is the response for GET /api/v1/reenrollment/history.
type ReenrollmentHistoryResponse struct {
	TotalCount int                        `json:"totalCount"`
	Results    []ReenrollmentHistoryObject `json:"results"`
}

// AddReenrollmentHistoryNotesRequest is the body for POST /api/v1/reenrollment/history.
type AddReenrollmentHistoryNotesRequest struct {
	Note string `json:"note"`
}

// ExportReenrollmentHistoryRequest is the optional body for POST .../history/export (override query when URI exceeds ~2k chars).
type ExportReenrollmentHistoryRequest struct {
	Page     *int     `json:"page,omitempty"`
	PageSize *int     `json:"pageSize,omitempty"`
	Sort     []string `json:"sort,omitempty"`
	Filter   *string  `json:"filter,omitempty"`
	Fields   []any    `json:"fields,omitempty"`
}
