package engage

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/models"
)

type ResourceEngageSettings struct {
	IsEnabled bool `json:"isEnabled"`
}

// ResourceHistoryEntry is an alias to the shared history item struct.
type ResourceHistoryEntry = models.SharedHistoryItem

// HistoryResponse is an alias to the shared history response struct.
type HistoryResponse = models.SharedHistoryResponse

// RequestAddHistoryNotes is an alias to the shared history note request struct.
type RequestAddHistoryNotes = models.SharedHistoryNoteRequest

type ResponseAddHistoryNotes struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Date     string  `json:"date"`
	Note     string  `json:"note"`
	Details  *string `json:"details"`
}
