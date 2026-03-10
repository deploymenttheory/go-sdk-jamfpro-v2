package buildings

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/models"
)

// ResourceBuilding represents a building resource.
type ResourceBuilding struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name"` // required
	StreetAddress1 string `json:"streetAddress1,omitempty"`
	StreetAddress2 string `json:"streetAddress2,omitempty"`
	City           string `json:"city,omitempty"`
	StateProvince  string `json:"stateProvince,omitempty"`
	ZipPostalCode  string `json:"zipPostalCode,omitempty"`
	Country        string `json:"country,omitempty"`
}

// ListResponse is the response for ListBuildings.
type ListResponse struct {
	TotalCount int                `json:"totalCount"`
	Results    []ResourceBuilding `json:"results"`
}

// RequestBuilding is the body for creating and updating buildings.
type RequestBuilding struct {
	Name           string `json:"name"` // required
	StreetAddress1 string `json:"streetAddress1,omitempty"`
	StreetAddress2 string `json:"streetAddress2,omitempty"`
	City           string `json:"city,omitempty"`
	StateProvince  string `json:"stateProvince,omitempty"`
	ZipPostalCode  string `json:"zipPostalCode,omitempty"`
	Country        string `json:"country,omitempty"`
}

// CreateResponse is the response for CreateBuilding.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// HistoryObject represents a building history entry.
type HistoryObject = models.SharedHistoryItem

// HistoryResponse is the response for GetBuildingHistoryV1.
type HistoryResponse = models.SharedHistoryResponse

// AddHistoryNotesRequest is the body for AddBuildingHistoryNotesV1.
type AddHistoryNotesRequest = models.SharedHistoryNoteRequest

// DeleteBuildingsByIDRequest is the body for DeleteBuildingsByIDV1 (delete multiple).
type DeleteBuildingsByIDRequest struct {
	IDs []string `json:"ids"`
}

// ExportRequest is the optional request body for ExportV1 and ExportHistoryV1.
// Overrides query parameters when URI would exceed ~2k characters.
type ExportRequest struct {
	Page     *int     `json:"page,omitempty"`
	PageSize *int     `json:"pageSize,omitempty"`
	Sort     []string `json:"sort,omitempty"`
	Filter   *string  `json:"filter,omitempty"`
	Fields   []string `json:"fields,omitempty"`
}
