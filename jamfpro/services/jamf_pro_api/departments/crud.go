package departments

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// DepartmentsServiceInterface defines the interface for department operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-departments
	DepartmentsServiceInterface interface {
		// ListV1 returns all department objects (Get Department objects).
		//
		// Returns a paged list of department objects. Optional query parameters support
		// filtering and pagination (page, pageSize, sort).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-departments
		ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetByIDV1 returns the specified department by ID (Get specified Department object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-departments-id
		GetByIDV1(ctx context.Context, id string) (*ResourceDepartment, *interfaces.Response, error)

		// CreateV1 creates a new department record (Create Department record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-departments
		CreateV1(ctx context.Context, request *RequestDepartment) (*CreateResponse, *interfaces.Response, error)

		// UpdateByIDV1 updates the specified department by ID (Update specified Department object).
		//
		// Returns the full updated department resource.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-departments-id
		UpdateByIDV1(ctx context.Context, id string, request *RequestDepartment) (*ResourceDepartment, *interfaces.Response, error)

		// DeleteByIDV1 removes the specified department by ID (Remove specified Department record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-departments-id
		DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error)

		// GetDepartmentHistoryV1 returns the history object for the specified department.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-departments-id-history
		GetDepartmentHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error)

		// AddDepartmentHistoryNotesV1 adds notes to the specified department history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-departments-id-history
		AddDepartmentHistoryNotesV1(ctx context.Context, id string, req *AddHistoryNotesRequest) (*interfaces.Response, error)
	}

	// Service handles communication with the departments-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-departments
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ DepartmentsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Departments CRUD Operations
// -----------------------------------------------------------------------------

// ListV1 returns all department objects (Get Department objects).
// URL: GET /api/v1/departments
// Query Params: page, pageSize, sort (optional)
// https://developer.jamf.com/jamf-pro/reference/get_v1-departments
func (s *Service) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointDepartmentsV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByIDV1 returns the specified department by ID (Get specified Department object).
// URL: GET /api/v1/departments/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-departments-id
func (s *Service) GetByIDV1(ctx context.Context, id string) (*ResourceDepartment, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("department ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointDepartmentsV1, id)

	var result ResourceDepartment

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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
func (s *Service) CreateV1(ctx context.Context, request *RequestDepartment) (*CreateResponse, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := EndpointDepartmentsV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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
func (s *Service) UpdateByIDV1(ctx context.Context, id string, request *RequestDepartment) (*ResourceDepartment, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointDepartmentsV1, id)

	var result ResourceDepartment

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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
func (s *Service) DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("department ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointDepartmentsV1, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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
func (s *Service) GetDepartmentHistoryV1(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("department ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointDepartmentsV1, id)

	var result HistoryResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// AddDepartmentHistoryNotesV1 adds notes to the specified department history.
// URL: POST /api/v1/departments/{id}/history
// Body: JSON with note
// https://developer.jamf.com/jamf-pro/reference/post_v1-departments-id-history
func (s *Service) AddDepartmentHistoryNotesV1(ctx context.Context, id string, req *AddHistoryNotesRequest) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("department ID is required")
	}
	if req == nil {
		return nil, fmt.Errorf("request body is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointDepartmentsV1, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
