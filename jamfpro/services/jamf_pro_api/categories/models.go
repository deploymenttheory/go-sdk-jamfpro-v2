package categories

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/normalization"
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
type CreateUpdateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// DeleteCategoriesByIDRequest is the body for DeleteCategoriesByID (delete multiple).
type DeleteCategoriesByIDRequest struct {
	IDs []string `json:"ids"`
}

// HistoryObject represents a category history entry.
type HistoryObject struct {
	ID       normalization.IDAsString `json:"id"`
	Username string `json:"username"`
	Date     string `json:"date"`
	Note     string `json:"note"`
	Details  string `json:"details"`
}

// CategoryHistoryResponse is the response for GetCategoryHistory.
type CategoryHistoryResponse struct {
	TotalCount int             `json:"totalCount"`
	Results    []HistoryObject `json:"results"`
}

// AddCategoryHistoryNotesRequest is the body for AddCategoryHistoryNotes.
type AddCategoryHistoryNotesRequest struct {
	Note string `json:"note"`
}
