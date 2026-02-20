package api_roles

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// APIRolesServiceInterface defines the interface for API role operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-api-roles
	APIRolesServiceInterface interface {
		// ListAPIRolesV1 returns all API role objects (Get API Role objects).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-api-roles
		ListAPIRolesV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetAPIRoleByIDV1 returns the specified API role by ID (Get specified API Role object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-api-roles-id
		GetAPIRoleByIDV1(ctx context.Context, id string) (*ResourceAPIRole, *interfaces.Response, error)

		// CreateAPIRoleV1 creates a new API role (Create API Role record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-api-roles
		CreateAPIRoleV1(ctx context.Context, request *RequestAPIRole) (*ResourceAPIRole, *interfaces.Response, error)

		// UpdateAPIRoleByIDV1 updates the specified API role by ID (Update specified API Role object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-api-roles-id
		UpdateAPIRoleByIDV1(ctx context.Context, id string, request *RequestAPIRole) (*ResourceAPIRole, *interfaces.Response, error)

		// DeleteAPIRoleByIDV1 removes the specified API role by ID (Remove specified API Role record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-api-roles-id
		DeleteAPIRoleByIDV1(ctx context.Context, id string) (*interfaces.Response, error)
	}

	// Service handles communication with the API roles-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-api-roles
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ APIRolesServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - API Roles CRUD Operations
// -----------------------------------------------------------------------------

// ListAPIRolesV1 returns all API role objects.
// URL: GET /api/v1/api-roles
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-api-roles
func (s *Service) ListAPIRolesV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointAPIRolesV1

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

// GetAPIRoleByIDV1 returns the specified API role by ID.
// URL: GET /api/v1/api-roles/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-api-roles-id
func (s *Service) GetAPIRoleByIDV1(ctx context.Context, id string) (*ResourceAPIRole, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("API role ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointAPIRolesV1, id)

	var result ResourceAPIRole

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

// CreateAPIRoleV1 creates a new API role.
// URL: POST /api/v1/api-roles
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-api-roles
func (s *Service) CreateAPIRoleV1(ctx context.Context, request *RequestAPIRole) (*ResourceAPIRole, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceAPIRole

	endpoint := EndpointAPIRolesV1

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

// UpdateAPIRoleByIDV1 updates the specified API role by ID.
// URL: PUT /api/v1/api-roles/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-api-roles-id
func (s *Service) UpdateAPIRoleByIDV1(ctx context.Context, id string, request *RequestAPIRole) (*ResourceAPIRole, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointAPIRolesV1, id)

	var result ResourceAPIRole

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

// DeleteAPIRoleByIDV1 removes the specified API role by ID.
// URL: DELETE /api/v1/api-roles/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-api-roles-id
func (s *Service) DeleteAPIRoleByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("API role ID is required")
	}
	endpoint := fmt.Sprintf("%s/%s", EndpointAPIRolesV1, id)

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
