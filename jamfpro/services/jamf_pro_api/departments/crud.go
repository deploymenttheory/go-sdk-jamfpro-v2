package departments

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

type (
	// DepartmentsServiceInterface defines the interface for department operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-departments
	DepartmentsServiceInterface interface {
		// ListDepartmentsV1 returns all department objects (Get Department objects).
		//
		// Returns a paged list of department objects. Optional query parameters support
		// filtering and pagination (page, pageSize, sort).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-departments
		ListDepartmentsV1(ctx context.Context, queryParams map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetDepartmentByIDV1 returns the specified department by ID (Get specified Department object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-departments-id
		GetDepartmentByIDV1(ctx context.Context, id string) (*ResourceDepartment, *interfaces.Response, error)

		// CreateDepartmentV1 creates a new department record (Create Department record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-departments
		CreateDepartmentV1(ctx context.Context, req *RequestDepartment) (*CreateResponse, *interfaces.Response, error)

		// UpdateDepartmentByIDV1 updates the specified department by ID (Update specified Department object).
		//
		// Returns the full updated department resource.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-departments-id
		UpdateDepartmentByIDV1(ctx context.Context, id string, req *RequestDepartment) (*ResourceDepartment, *interfaces.Response, error)

		// DeleteDepartmentByIDV1 removes the specified department by ID (Remove specified Department record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-departments-id
		DeleteDepartmentByIDV1(ctx context.Context, id string) (*interfaces.Response, error)

		// GetDepartmentHistoryV1 returns the history object for the specified department.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-departments-id-history
		GetDepartmentHistoryV1(ctx context.Context, id string, queryParams map[string]string) (*HistoryResponse, *interfaces.Response, error)

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

// ListDepartmentsV1 returns all department objects (Get Department objects).
// URL: GET /api/v1/departments
// Query Params: page, pageSize, sort (optional)
// https://developer.jamf.com/jamf-pro/reference/get_v1-departments
func (s *Service) ListDepartmentsV1(ctx context.Context, queryParams map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	resp, err := s.client.Get(ctx, EndpointDepartmentsV1, queryParams, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetDepartmentByIDV1 returns the specified department by ID (Get specified Department object).
// URL: GET /api/v1/departments/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-departments-id
func (s *Service) GetDepartmentByIDV1(ctx context.Context, id string) (*ResourceDepartment, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("department ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointDepartmentsV1, id)

	var result ResourceDepartment

	resp, err := s.client.Get(ctx, endpoint, nil, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateDepartmentV1 creates a new department record (Create Department record).
// URL: POST /api/v1/departments
// Body: JSON with name
// https://developer.jamf.com/jamf-pro/reference/post_v1-departments
func (s *Service) CreateDepartmentV1(ctx context.Context, req *RequestDepartment) (*CreateResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	resp, err := s.client.Post(ctx, EndpointDepartmentsV1, req, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateDepartmentByIDV1 updates the specified department by ID (Update specified Department object).
// URL: PUT /api/v1/departments/{id}
// Body: JSON with name
// https://developer.jamf.com/jamf-pro/reference/put_v1-departments-id
func (s *Service) UpdateDepartmentByIDV1(ctx context.Context, id string, req *RequestDepartment) (*ResourceDepartment, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointDepartmentsV1, id)

	var result ResourceDepartment

	resp, err := s.client.Put(ctx, endpoint, req, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteDepartmentByIDV1 removes the specified department by ID (Remove specified Department record).
// URL: DELETE /api/v1/departments/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-departments-id
func (s *Service) DeleteDepartmentByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("department ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointDepartmentsV1, id)

	resp, err := s.client.Delete(ctx, endpoint, nil, shared.JSONHeaders(), nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetDepartmentHistoryV1 returns the history object for the specified department.
// URL: GET /api/v1/departments/{id}/history
// Query Params: filter, sort, page, page-size (optional)
// https://developer.jamf.com/jamf-pro/reference/get_v1-departments-id-history
func (s *Service) GetDepartmentHistoryV1(ctx context.Context, id string, queryParams map[string]string) (*HistoryResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("department ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", EndpointDepartmentsV1, id)

	var result HistoryResponse

	resp, err := s.client.Get(ctx, endpoint, queryParams, shared.JSONHeaders(), &result)
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

	resp, err := s.client.Post(ctx, endpoint, req, shared.JSONHeaders(), nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
