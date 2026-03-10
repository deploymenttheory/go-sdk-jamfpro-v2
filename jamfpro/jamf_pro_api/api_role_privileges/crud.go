package api_role_privileges

import (
	"context"
	"fmt"
	"net/url"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the API role privileges-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/api-role-privileges
	ApiRolePrivileges struct {
		client client.Client
	}
)

func NewApiRolePrivileges(client client.Client) *ApiRolePrivileges {
	return &ApiRolePrivileges{client: client}
}

// ListV1 returns all API role privileges.
// URL: GET /api/v1/api-role-privileges
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-api-role-privileges
func (s *ApiRolePrivileges) ListV1(ctx context.Context) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointJamfProAPIRolePrivilegesV1

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
	endpoint := fmt.Sprintf("%s/search?name=%s&limit=%d", constants.EndpointJamfProAPIRolePrivilegesV1, url.QueryEscape(name), limit)
	var result ListResponse

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
