package reenrollment

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/models"
)

// ResourceReenrollmentSettings is the re-enrollment settings resource (get/update).
type ResourceReenrollmentSettings struct {
	FlushPolicyHistory              bool   `json:"isFlushPolicyHistoryEnabled"`
	FlushLocationInformation        bool   `json:"isFlushLocationInformationEnabled"`
	FlushLocationInformationHistory bool   `json:"isFlushLocationInformationHistoryEnabled"`
	FlushExtensionAttributes        bool   `json:"isFlushExtensionAttributesEnabled"`
	FlushSoftwareUpdatePlans        bool   `json:"isFlushSoftwareUpdatePlansEnabled"`
	FlushMdmQueue                   string `json:"flushMDMQueue"`
}

// ReenrollmentHistoryObject is an alias to the shared history item struct.
type ReenrollmentHistoryObject = models.SharedHistoryItem

// ReenrollmentHistoryResponse is an alias to the shared history response struct.
type ReenrollmentHistoryResponse = models.SharedHistoryResponse

// AddReenrollmentHistoryNotesRequest is an alias to the shared history note request struct.
type AddReenrollmentHistoryNotesRequest = models.SharedHistoryNoteRequest

// ExportReenrollmentHistoryRequest is the optional body for POST .../history/export (override query when URI exceeds ~2k chars).
type ExportReenrollmentHistoryRequest struct {
	Page     *int     `json:"page,omitempty"`
	PageSize *int     `json:"pageSize,omitempty"`
	Sort     []string `json:"sort,omitempty"`
	Filter   *string  `json:"filter,omitempty"`
	Fields   []any    `json:"fields,omitempty"`
}
