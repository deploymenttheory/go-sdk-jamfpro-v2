package mobile_device_groups

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the mobile device groups-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-smart-groups
	MobileDeviceGroups struct {
		client client.Client
	}
)

func NewMobileDeviceGroups(client client.Client) *MobileDeviceGroups {
	return &MobileDeviceGroups{client: client}
}

// -----------------------------------------------------------------------------
// Smart Groups CRUD (V1)
// -----------------------------------------------------------------------------

// ListSmartV1 returns all smart mobile device groups.
// URL: GET /api/v1/mobile-device-groups/smart-groups
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-smart-groups
func (s *MobileDeviceGroups) ListSmartV1(ctx context.Context, rsqlQuery map[string]string) (*ListSmartResponse, *resty.Response, error) {
	var result ListSmartResponse

	endpoint := constants.EndpointJamfProSmartMobileDeviceGroupsV1

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceSmartMobileDeviceGroup
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list smart mobile device groups: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetSmartByIDV1 returns the specified smart mobile device group by ID.
// URL: GET /api/v1/mobile-device-groups/smart-groups/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-smart-groups-id
func (s *MobileDeviceGroups) GetSmartByIDV1(ctx context.Context, id string) (*ResourceSmartMobileDeviceGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("smart mobile device group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSmartMobileDeviceGroupsV1, id)

	var result ResourceSmartMobileDeviceGroup

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateSmartV1 creates a new smart mobile device group.
// URL: POST /api/v1/mobile-device-groups/smart-groups
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-groups-smart-groups
func (s *MobileDeviceGroups) CreateSmartV1(ctx context.Context, request *RequestSmartMobileDeviceGroup) (*CreateSmartResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateSmartResponse

	endpoint := constants.EndpointJamfProSmartMobileDeviceGroupsV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateSmartByIDV1 updates the specified smart mobile device group by ID.
// URL: PUT /api/v1/mobile-device-groups/smart-groups/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-mobile-device-groups-smart-groups-id
func (s *MobileDeviceGroups) UpdateSmartByIDV1(ctx context.Context, id string, request *RequestSmartMobileDeviceGroup) (*ResourceSmartMobileDeviceGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSmartMobileDeviceGroupsV1, id)

	var result ResourceSmartMobileDeviceGroup

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Put(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteSmartByIDV1 removes the specified smart mobile device group by ID.
// URL: DELETE /api/v1/mobile-device-groups/smart-groups/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-mobile-device-groups-smart-groups-id
func (s *MobileDeviceGroups) DeleteSmartByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("smart mobile device group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSmartMobileDeviceGroupsV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// -----------------------------------------------------------------------------
// Static Groups CRUD (V1)
// -----------------------------------------------------------------------------

// ListStaticV1 returns all static mobile device groups.
// URL: GET /api/v1/mobile-device-groups/static-groups
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-static-groups
func (s *MobileDeviceGroups) ListStaticV1(ctx context.Context, rsqlQuery map[string]string) (*ListStaticResponse, *resty.Response, error) {
	var result ListStaticResponse

	endpoint := constants.EndpointJamfProStaticMobileDeviceGroupsV1

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceStaticMobileDeviceGroup
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list static mobile device groups: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetStaticByIDV1 returns the specified static mobile device group by ID.
// URL: GET /api/v1/mobile-device-groups/static-groups/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-static-groups-id
func (s *MobileDeviceGroups) GetStaticByIDV1(ctx context.Context, id string) (*ResourceStaticMobileDeviceGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("static mobile device group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProStaticMobileDeviceGroupsV1, id)

	var result ResourceStaticMobileDeviceGroup

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateStaticV1 creates a new static mobile device group.
// URL: POST /api/v1/mobile-device-groups/static-groups
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-groups-static-groups
func (s *MobileDeviceGroups) CreateStaticV1(ctx context.Context, request *RequestStaticMobileDeviceGroup) (*CreateStaticResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	if request.Assignments == nil {
		request.Assignments = []StaticMobileDeviceGroupAssignment{}
	}

	var result CreateStaticResponse

	endpoint := constants.EndpointJamfProStaticMobileDeviceGroupsV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateStaticByIDV1 updates the specified static mobile device group by ID (PATCH).
// URL: PATCH /api/v1/mobile-device-groups/static-groups/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-mobile-device-groups-static-groups-id
func (s *MobileDeviceGroups) UpdateStaticByIDV1(ctx context.Context, id string, request *RequestStaticMobileDeviceGroup) (*ResourceStaticMobileDeviceGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	if request.Assignments == nil {
		request.Assignments = []StaticMobileDeviceGroupAssignment{}
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProStaticMobileDeviceGroupsV1, id)

	var result ResourceStaticMobileDeviceGroup

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Patch(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteStaticByIDV1 removes the specified static mobile device group by ID.
// URL: DELETE /api/v1/mobile-device-groups/static-groups/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-mobile-device-groups-static-groups-id
func (s *MobileDeviceGroups) DeleteStaticByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("static mobile device group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProStaticMobileDeviceGroupsV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// -----------------------------------------------------------------------------
// List All, Membership, and Erase (V1)
// -----------------------------------------------------------------------------

// ListAllV1 returns all mobile device groups (smart + static) as a simple list.
// URL: GET /api/v1/mobile-device-groups
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups
func (s *MobileDeviceGroups) ListAllV1(ctx context.Context) ([]ResourceMobileDeviceGroupSummary, *resty.Response, error) {
	var result []ResourceMobileDeviceGroupSummary

	endpoint := constants.EndpointJamfProMobileDeviceGroupsV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// GetStaticGroupMembershipV1 returns the mobile devices in the specified static group.
// URL: GET /api/v1/mobile-device-groups/static-group-membership/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-static-group-membership-id
func (s *MobileDeviceGroups) GetStaticGroupMembershipV1(ctx context.Context, id string, rsqlQuery map[string]string) (*GroupMembershipResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("static group ID is required")
	}

	var result GroupMembershipResponse

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProStaticMobileDeviceGroupMembershipV1, id)

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceMobileDeviceMember
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get static group membership: %w", err)
	}

	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetSmartGroupMembershipV1 returns the mobile devices in the specified smart group.
// URL: GET /api/v1/mobile-device-groups/smart-group-membership/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-smart-group-membership-id
func (s *MobileDeviceGroups) GetSmartGroupMembershipV1(ctx context.Context, id string, rsqlQuery map[string]string) (*GroupMembershipResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("smart group ID is required")
	}

	var result GroupMembershipResponse

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSmartMobileDeviceGroupMembershipV1, id)

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceMobileDeviceMember
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get smart group membership: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// EraseDevicesByGroupIDV1 erases all devices in the specified mobile device group.
// URL: POST /api/v1/mobile-device-groups/{id}/erase
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-groups-id-erase
func (s *MobileDeviceGroups) EraseDevicesByGroupIDV1(ctx context.Context, id string, request *RequestEraseDevices) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/erase", constants.EndpointJamfProMobileDeviceGroupsV1, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		Post(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
