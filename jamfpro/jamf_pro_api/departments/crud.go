package departments

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the departments-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-departments
	Departments struct {
		client client.Client
	}
)

func NewDepartments(client client.Client) *Departments {
	return &Departments{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Departments CRUD Operations
// -----------------------------------------------------------------------------

// ListV1 returns all department objects (Get Department objects).
// URL: GET /api/v1/departments
// Query Params: page, pageSize, sort (optional)
// https://developer.jamf.com/jamf-pro/reference/get_v1-departments
func (s *Departments) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointJamfProDepartmentsV1

	mergePage := func(pageData []byte) error {
		var pageResults []ResourceDepartment
		if err := json.Unmarshal(pageData, &pageResults); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResults...)
		return nil
	}

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list departments: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByIDV1 returns the specified department by ID (Get specified Department object).
// URL: GET /api/v1/departments/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-departments-id
func (s *Departments) GetByIDV1(ctx context.Context, id string) (*ResourceDepartment, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("department ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProDepartmentsV1, id)

	var result ResourceDepartment

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateV1 creates a new department record (Create Department record).
// URL: POST /api/v1/departments
// Body: JSON with name
// https://developer.jamf.com/jamf-pro/reference/post_v1-departments
func (s *Departments) CreateV1(ctx context.Context, request *RequestDepartment) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := constants.EndpointJamfProDepartmentsV1

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByIDV1 updates the specified department by ID (Update specified Department object).
// URL: PUT /api/v1/departments/{id}
// Body: JSON with name
// https://developer.jamf.com/jamf-pro/reference/put_v1-departments-id
func (s *Departments) UpdateByIDV1(ctx context.Context, id string, request *RequestDepartment) (*ResourceDepartment, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProDepartmentsV1, id)

	var result ResourceDepartment

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByIDV1 removes the specified department by ID (Remove specified Department record).
// URL: DELETE /api/v1/departments/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-departments-id
func (s *Departments) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("department ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProDepartmentsV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetDepartmentHistoryV1 returns the history object for the specified department.
// URL: GET /api/v1/departments/{id}/history
// Query Params: filter, sort, page, page-size (optional)
// https://developer.jamf.com/jamf-pro/reference/get_v1-departments-id-history
func (s *Departments) GetDepartmentHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("department ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProDepartmentsV1, id)

	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var pageResults []HistoryObject
		if err := json.Unmarshal(pageData, &pageResults); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResults...)
		return nil
	}

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get department history: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// AddDepartmentHistoryNotesV1 adds notes to the specified department history.
// URL: POST /api/v1/departments/{id}/history
// Body: JSON with note
// https://developer.jamf.com/jamf-pro/reference/post_v1-departments-id-history
func (s *Departments) AddDepartmentHistoryNotesV1(ctx context.Context, id string, req *AddHistoryNotesRequest) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("department ID is required")
	}
	if req == nil {
		return nil, fmt.Errorf("request body is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProDepartmentsV1, id)

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// DeleteDepartmentsByIDV1 deletes multiple departments by their IDs.
// URL: POST /api/v1/departments/delete-multiple
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-departments-delete-multiple
func (s *Departments) DeleteDepartmentsByIDV1(ctx context.Context, req *DeleteDepartmentsByIDRequest) (*resty.Response, error) {
	if req == nil || len(req.IDs) == 0 {
		return nil, fmt.Errorf("department IDs are required")
	}

	endpoint := constants.EndpointJamfProDepartmentsV1 + "/delete-multiple"

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
