package groups

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the Groups-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-groups
	Groups struct {
		client client.Client
	}
)

func NewGroups(client client.Client) *Groups {
	return &Groups{client: client}
}

// ListV1 retrieves a paginated list of groups.
// URL: GET /api/v1/groups
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v1-groups
func (s *Groups) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	endpoint := constants.EndpointJamfProGroupsV1

	var result ListResponse

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceGroup
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list groups: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByIDV1 retrieves a group by its platform ID (groupPlatformId).
// URL: GET /api/v1/groups/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-groups-id
func (s *Groups) GetByIDV1(ctx context.Context, id string) (*ResourceGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProGroupsV1, id)

	var result ResourceGroup

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetComputerGroupByNameV1 retrieves a computer group by its name (groupName).
// URL: GET /api/v1/groups?filter=groupName=="name"
// https://developer.jamf.com/jamf-pro/reference/get_v1-groups
func (s *Groups) GetComputerGroupByNameV1(ctx context.Context, name string) (*ResourceGroup, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("group name is required")
	}

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`groupName=="%s"`, name),
	}

	listResult, resp, err := s.ListV1(ctx, rsqlQuery)
	if err != nil {
		return nil, resp, err
	}

	for _, group := range listResult.Results {
		if group.GroupName == name && group.GroupType == "COMPUTER" {
			return &group, resp, nil
		}
	}

	return nil, resp, fmt.Errorf("computer group with name %q not found", name)
}

// GetMobileGroupByNameV1 retrieves a mobile device group by its name (groupName).
// URL: GET /api/v1/groups?filter=groupName=="name"
// https://developer.jamf.com/jamf-pro/reference/get_v1-groups
func (s *Groups) GetMobileGroupByNameV1(ctx context.Context, name string) (*ResourceGroup, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("group name is required")
	}

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`groupName=="%s"`, name),
	}

	listResult, resp, err := s.ListV1(ctx, rsqlQuery)
	if err != nil {
		return nil, resp, err
	}

	for _, group := range listResult.Results {
		if group.GroupName == name && group.GroupType == "MOBILE" {
			return &group, resp, nil
		}
	}

	return nil, resp, fmt.Errorf("mobile group with name %q not found", name)
}

// GetComputerGroupByIDV1 retrieves a computer group by its Jamf Pro ID (groupJamfProId).
// URL: GET /api/v1/groups
// https://developer.jamf.com/jamf-pro/reference/get_v1-groups
func (s *Groups) GetComputerGroupByIDV1(ctx context.Context, jamfProID string) (*ResourceGroup, *resty.Response, error) {
	if jamfProID == "" {
		return nil, nil, fmt.Errorf("group Jamf Pro ID is required")
	}

	listResult, resp, err := s.ListV1(ctx, nil)
	if err != nil {
		return nil, resp, err
	}

	for _, group := range listResult.Results {
		if group.GroupJamfProId == jamfProID && group.GroupType == "COMPUTER" {
			return &group, resp, nil
		}
	}

	return nil, resp, fmt.Errorf("computer group with Jamf Pro ID %q not found", jamfProID)
}

// GetMobileGroupByIDV1 retrieves a mobile device group by its Jamf Pro ID (groupJamfProId).
// URL: GET /api/v1/groups
// https://developer.jamf.com/jamf-pro/reference/get_v1-groups
func (s *Groups) GetMobileGroupByIDV1(ctx context.Context, jamfProID string) (*ResourceGroup, *resty.Response, error) {
	if jamfProID == "" {
		return nil, nil, fmt.Errorf("group Jamf Pro ID is required")
	}

	listResult, resp, err := s.ListV1(ctx, nil)
	if err != nil {
		return nil, resp, err
	}

	for _, group := range listResult.Results {
		if group.GroupJamfProId == jamfProID && group.GroupType == "MOBILE" {
			return &group, resp, nil
		}
	}

	return nil, resp, fmt.Errorf("mobile group with Jamf Pro ID %q not found", jamfProID)
}

// UpdateByIDV1 updates a group by its platform ID.
// URL: PATCH /api/v1/groups/{id}
// https://developer.jamf.com/jamf-pro/reference/patch_v1-groups-id
func (s *Groups) UpdateByIDV1(ctx context.Context, id string, req *RequestUpdateGroup) (*ResourceGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("group ID is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request body is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProGroupsV1, id)

	var result ResourceGroup

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Patch(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByIDV1 removes the specified group by its platform ID.
// URL: DELETE /api/v1/groups/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v1-groups-id
func (s *Groups) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProGroupsV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
