package buildings

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/normalization"
)

// ResourceBuilding represents a building resource.
type ResourceBuilding struct {
	ID               string `json:"id,omitempty"`
	Name             string `json:"name"` // required
	StreetAddress1   string `json:"streetAddress1,omitempty"`
	StreetAddress2   string `json:"streetAddress2,omitempty"`
	City             string `json:"city,omitempty"`
	StateProvince    string `json:"stateProvince,omitempty"`
	ZipPostalCode    string `json:"zipPostalCode,omitempty"`
	Country          string `json:"country,omitempty"`
}

// ListResponse is the response for ListBuildings.
type ListResponse struct {
	TotalCount int               `json:"totalCount"`
	Results    []ResourceBuilding `json:"results"`
}

// RequestBuilding is the body for creating and updating buildings.
type RequestBuilding struct {
	Name             string `json:"name"` // required
	StreetAddress1   string `json:"streetAddress1,omitempty"`
	StreetAddress2   string `json:"streetAddress2,omitempty"`
	City             string `json:"city,omitempty"`
	StateProvince    string `json:"stateProvince,omitempty"`
	ZipPostalCode    string `json:"zipPostalCode,omitempty"`
	Country          string `json:"country,omitempty"`
}

// CreateResponse is the response for CreateBuilding.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// HistoryObject represents a building history entry.
type HistoryObject struct {
	ID       normalization.IDAsString `json:"id"`
	Username string                   `json:"username"`
	Date     string                   `json:"date"`
	Note     string                   `json:"note"`
	Details  *string                  `json:"details"`
}

// HistoryResponse is the response for GetBuildingHistoryV1.
type HistoryResponse struct {
	TotalCount int             `json:"totalCount"`
	Results    []HistoryObject `json:"results"`
}

// AddHistoryNotesRequest is the body for AddBuildingHistoryNotesV1.
type AddHistoryNotesRequest struct {
	Note string `json:"note"`
}

// DeleteBuildingsByIDRequest is the body for DeleteBuildingsByIDV1 (delete multiple).
type DeleteBuildingsByIDRequest struct {
	IDs []string `json:"ids"`
}
