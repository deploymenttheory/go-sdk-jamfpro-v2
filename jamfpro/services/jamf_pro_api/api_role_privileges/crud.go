package api_role_privileges

import (
	"context"
	"fmt"
	"net/url"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

type (
	// APIRolePrivilegesServiceInterface defines the interface for API role privilege operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/api-role-privileges
	APIRolePrivilegesServiceInterface interface {
		// ListPrivilegesV1 returns all API role privileges (Get API Role Privileges).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-api-role-privileges
		ListPrivilegesV1(ctx context.Context) (*ListResponse, *interfaces.Response, error)

		// SearchPrivilegesByNameV1 returns privileges matching the given name (Get API Role Privileges by name).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-api-role-privileges-search
		SearchPrivilegesByNameV1(ctx context.Context, name string, limit int) (*ListResponse, *interfaces.Response, error)
	}

	// Service handles communication with the API role privileges-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/api-role-privileges
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ APIRolePrivilegesServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// ListPrivilegesV1 returns all API role privileges.
// URL: GET /api/v1/api-role-privileges
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-api-role-privileges
func (s *Service) ListPrivilegesV1(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse
	resp, err := s.client.Get(ctx, EndpointAPIRolePrivilegesV1, nil, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// SearchPrivilegesByNameV1 returns privileges matching the given name.
// URL: GET /api/v1/api-role-privileges/search?name=...&limit=...
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-api-role-privileges-search
func (s *Service) SearchPrivilegesByNameV1(ctx context.Context, name string, limit int) (*ListResponse, *interfaces.Response, error) {
	if limit <= 0 {
		limit = 100
	}
	endpoint := fmt.Sprintf("%s/search?name=%s&limit=%d", EndpointAPIRolePrivilegesV1, url.QueryEscape(name), limit)
	var result ListResponse
	resp, err := s.client.Get(ctx, endpoint, nil, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
