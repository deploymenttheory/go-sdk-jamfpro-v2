package categories

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

// ResourceCategory represents a category resource.
type ResourceCategory struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}

// ListResponse is the response for ListCategories.
type ListResponse struct {
	TotalCount int                `json:"totalCount"`
	Results    []ResourceCategory `json:"results"`
}

// RequestCategory is the body for creating and updating categories.
type RequestCategory struct {
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}

// CreateUpdateResponse is the response for CreateCategory and UpdateCategoryByID.
// PUT update returns {id, name, priority}; POST create may return {id, href} or {id, name, priority}.
type CreateUpdateResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name,omitempty"`
	Priority int    `json:"priority,omitempty"`
	Href     string `json:"href,omitempty"`
}

// DeleteCategoriesByIDRequest is the body for DeleteCategoriesByID (delete multiple).
type DeleteCategoriesByIDRequest struct {
	IDs []string `json:"ids"`
}

// HistoryObject represents a category history entry.
type HistoryObject = shared.SharedHistoryItem

// CategoryHistoryResponse is the response for GetCategoryHistory.
type CategoryHistoryResponse = shared.SharedHistoryResponse

// AddCategoryHistoryNotesRequest is the body for AddCategoryHistoryNotes.
type AddCategoryHistoryNotesRequest = shared.SharedHistoryNoteRequest
