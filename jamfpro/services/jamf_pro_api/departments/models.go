package departments

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/normalization"
)

// ResourceDepartment represents a department resource.
type ResourceDepartment struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ListResponse is the response for ListDepartments.
type ListResponse struct {
	TotalCount int                  `json:"totalCount"`
	Results    []ResourceDepartment `json:"results"`
}

// RequestDepartment is the body for creating and updating departments.
type RequestDepartment struct {
	Name string `json:"name"`
}

// CreateResponse is the response for CreateDepartment.
type CreateResponse struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// HistoryObject represents a department history entry.
type HistoryObject struct {
	ID       normalization.IDAsString `json:"id"`
	Username string                   `json:"username"`
	Date     string                   `json:"date"`
	Note     string                   `json:"note"`
	Details  string                   `json:"details"`
}

// HistoryResponse is the response for GetDepartmentHistoryV1.
type HistoryResponse struct {
	TotalCount int             `json:"totalCount"`
	Results    []HistoryObject `json:"results"`
}

// AddHistoryNotesRequest is the body for AddDepartmentHistoryNotesV1.
type AddHistoryNotesRequest struct {
	Note string `json:"note"`
}
