package api_roles

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the API roles-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getallapiroles
	ApiRoles struct {
		client transport.HTTPClient
	}
)

func NewApiRoles(client transport.HTTPClient) *ApiRoles {
	return &ApiRoles{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - API Roles CRUD Operations
// -----------------------------------------------------------------------------

// ListV1 returns all API role objects with automatic pagination.
// URL: GET /api/v1/api-roles
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// Note: page and page-size are managed internally by GetPaginated.
// https://developer.jamf.com/jamf-pro/reference/getallapiroles
func (s *ApiRoles) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointJamfProAPIRolesV1

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	mergePage := func(pageData []byte) error {
		var pageResults []ResourceAPIRole
		if err := json.Unmarshal(pageData, &pageResults); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageResults...)
		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, err
	}

	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByIDV1 returns the specified API role by ID.
// URL: GET /api/v1/api-roles/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getoneapirole
func (s *ApiRoles) GetByIDV1(ctx context.Context, id string) (*ResourceAPIRole, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("API role ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProAPIRolesV1, id)

	var result ResourceAPIRole

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateV1 creates a new API role.
// URL: POST /api/v1/api-roles
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/postcreateapirole
func (s *ApiRoles) CreateV1(ctx context.Context, request *RequestAPIRole) (*ResourceAPIRole, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceAPIRole

	endpoint := constants.EndpointJamfProAPIRolesV1

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

// UpdateByIDV1 updates the specified API role by ID.
// URL: PUT /api/v1/api-roles/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/putupdateapirole
func (s *ApiRoles) UpdateByIDV1(ctx context.Context, id string, request *RequestAPIRole) (*ResourceAPIRole, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProAPIRolesV1, id)

	var result ResourceAPIRole

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

// DeleteByIDV1 removes the specified API role by ID.
// URL: DELETE /api/v1/api-roles/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/deleteapirole
func (s *ApiRoles) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("API role ID is required")
	}
	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProAPIRolesV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
