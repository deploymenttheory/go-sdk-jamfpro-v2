package activation_code

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/models"
)

// ActivationCodeRequest is the body for updating the activation code.
type ActivationCodeRequest struct {
	ActivationCode string `json:"activationCode"` // Length between 32 and 39 characters. Hyphens are optional.
}

// OrganizationNameRequest is the body for updating the organization name.
type OrganizationNameRequest struct {
	OrganizationName string `json:"organizationName"`
}

// HistoryEntry represents a single activation code history record.
type HistoryEntry = models.SharedHistoryItem

// HistoryResponse is the response for listing activation code history.
type HistoryResponse = models.SharedHistoryResponse

// HistoryNoteRequest is the body for adding a note to activation code history.
type HistoryNoteRequest = models.SharedHistoryNoteRequest

// ExportField represents a field configuration for export.
type ExportField struct {
	Name  string `json:"name"`
	Label string `json:"label,omitempty"`
}

// HistoryExportRequest is the body for exporting activation code history.
type HistoryExportRequest struct {
	Page     *int          `json:"page,omitempty"`
	PageSize *int          `json:"pageSize,omitempty"`
	Sort     []string      `json:"sort,omitempty"`
	Filter   *string       `json:"filter,omitempty"`
	Fields   []ExportField `json:"fields,omitempty"`
}

// HistoryExportResponse represents the export response (could be JSON or CSV data).
type HistoryExportResponse struct {
	Data any `json:"-"`
}
