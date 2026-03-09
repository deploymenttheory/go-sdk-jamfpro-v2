package api_role_privileges

import (
	"context"
	"fmt"
	"net/url"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// APIRolePrivilegesServiceInterface defines the interface for API role privilege operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/api-role-privileges
	APIRolePrivilegesServiceInterface interface {
		// ListV1 returns all API role privileges (Get API Role Privileges).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-api-role-privileges
		ListV1(ctx context.Context) (*ListResponse, *resty.Response, error)

		// SearchPrivilegesByNameV1 returns privileges matching the given name (Get API Role Privileges by name).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-api-role-privileges-search
		SearchPrivilegesByNameV1(ctx context.Context, name string, limit int) (*ListResponse, *resty.Response, error)
	}

	// Service handles communication with the API role privileges-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/api-role-privileges
	ApiRolePrivileges struct {
		client interfaces.HTTPClient
	}
)

var _ APIRolePrivilegesServiceInterface = (*ApiRolePrivileges)(nil)

func NewApiRolePrivileges(client interfaces.HTTPClient) *ApiRolePrivileges {
	return &ApiRolePrivileges{client: client}
}

// ListV1 returns all API role privileges.
// URL: GET /api/v1/api-role-privileges
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-api-role-privileges
func (s *ApiRolePrivileges) ListV1(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := EndpointAPIRolePrivilegesV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// SearchPrivilegesByNameV1 returns privileges matching the given name.
// URL: GET /api/v1/api-role-privileges/search?name=...&limit=...
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-api-role-privileges-search
func (s *ApiRolePrivileges) SearchPrivilegesByNameV1(ctx context.Context, name string, limit int) (*ListResponse, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("name parameter is required")
	}
	if limit <= 0 {
		limit = 15
	}
	endpoint := fmt.Sprintf("%s/search?name=%s&limit=%d", EndpointAPIRolePrivilegesV1, url.QueryEscape(name), limit)
	var result ListResponse

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
