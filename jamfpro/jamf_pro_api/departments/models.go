package departments

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
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
type HistoryObject = shared.SharedHistoryItem

// HistoryResponse is the response for GetDepartmentHistoryV1.
type HistoryResponse = shared.SharedHistoryResponse

// AddHistoryNotesRequest is the body for AddDepartmentHistoryNotesV1.
type AddHistoryNotesRequest = shared.SharedHistoryNoteRequest

// DeleteDepartmentsByIDRequest is the body for DeleteDepartmentsByIDV1 (delete multiple).
type DeleteDepartmentsByIDRequest struct {
	IDs []string `json:"ids"`
}
