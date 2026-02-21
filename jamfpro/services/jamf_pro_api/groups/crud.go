package groups

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// GroupsServiceInterface defines the interface for Groups operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-groups
	GroupsServiceInterface interface {
		// ListV1 retrieves a paginated list of groups.
		//
		// Supports optional RSQL filtering and pagination via rsqlQuery
		// (keys: filter, sort, page, page-size).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-groups
		ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetByIDV1 retrieves a group by its platform ID (groupPlatformId).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-groups-id
		GetByIDV1(ctx context.Context, id string) (*ResourceGroup, *interfaces.Response, error)

		// GetComputerGroupByNameV1 retrieves a computer group by its name (groupName).
		//
		// This method filters by groupName and groupType=COMPUTER.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-groups
		GetComputerGroupByNameV1(ctx context.Context, name string) (*ResourceGroup, *interfaces.Response, error)

		// GetMobileGroupByNameV1 retrieves a mobile device group by its name (groupName).
		//
		// This method filters by groupName and groupType=MOBILE.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-groups
		GetMobileGroupByNameV1(ctx context.Context, name string) (*ResourceGroup, *interfaces.Response, error)

		// GetComputerGroupByIDV1 retrieves a computer group by its Jamf Pro ID (groupJamfProId).
		//
		// This method filters by groupJamfProId and groupType=COMPUTER.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-groups
		GetComputerGroupByIDV1(ctx context.Context, jamfProID string) (*ResourceGroup, *interfaces.Response, error)

		// GetMobileGroupByIDV1 retrieves a mobile device group by its Jamf Pro ID (groupJamfProId).
		//
		// This method filters by groupJamfProId and groupType=MOBILE.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-groups
		GetMobileGroupByIDV1(ctx context.Context, jamfProID string) (*ResourceGroup, *interfaces.Response, error)

		// UpdateByIDV1 updates a group by its platform ID.
		//
		// For both smart and static groups, groupName and groupDescription can be updated.
		// For smart groups, criteria can also be updated.
		// For static groups, assignments can also be updated.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-groups-id
		UpdateByIDV1(ctx context.Context, id string, req *RequestUpdateGroup) (*ResourceGroup, *interfaces.Response, error)

		// DeleteByIDV1 removes the specified group by its platform ID.
		//
		// Returns a 400 error if the group is being used as a dependency.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-groups-id
		DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error)
	}

	// Service handles communication with the Groups-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-groups
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ GroupsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// ListV1 retrieves a paginated list of groups.
// URL: GET /api/v1/groups
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v1-groups
func (s *Service) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error) {
	endpoint := EndpointGroupsV1

	var result ListResponse

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

// GetByIDV1 retrieves a group by its platform ID (groupPlatformId).
// URL: GET /api/v1/groups/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-groups-id
func (s *Service) GetByIDV1(ctx context.Context, id string) (*ResourceGroup, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointGroupsV1, id)

	var result ResourceGroup

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

// GetComputerGroupByNameV1 retrieves a computer group by its name (groupName).
// URL: GET /api/v1/groups?filter=groupName=="name"
// https://developer.jamf.com/jamf-pro/reference/get_v1-groups
func (s *Service) GetComputerGroupByNameV1(ctx context.Context, name string) (*ResourceGroup, *interfaces.Response, error) {
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
func (s *Service) GetMobileGroupByNameV1(ctx context.Context, name string) (*ResourceGroup, *interfaces.Response, error) {
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
func (s *Service) GetComputerGroupByIDV1(ctx context.Context, jamfProID string) (*ResourceGroup, *interfaces.Response, error) {
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
func (s *Service) GetMobileGroupByIDV1(ctx context.Context, jamfProID string) (*ResourceGroup, *interfaces.Response, error) {
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
func (s *Service) UpdateByIDV1(ctx context.Context, id string, req *RequestUpdateGroup) (*ResourceGroup, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("group ID is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request body is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointGroupsV1, id)

	var result ResourceGroup

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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
func (s *Service) DeleteByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointGroupsV1, id)

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
