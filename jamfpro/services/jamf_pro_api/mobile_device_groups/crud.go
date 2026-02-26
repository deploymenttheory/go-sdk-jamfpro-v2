package mobile_device_groups

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"github.com/mitchellh/mapstructure"
)

type (
	// MobileDeviceGroupsServiceInterface defines the interface for mobile device group operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-smart-groups
	MobileDeviceGroupsServiceInterface interface {
		// ListSmartV1 returns all smart mobile device groups (Get Smart Mobile Device Group objects).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-smart-groups
		ListSmartV1(ctx context.Context, rsqlQuery map[string]string) (*ListSmartResponse, *interfaces.Response, error)

		// GetSmartByIDV1 returns the specified smart mobile device group by ID (Get specified Smart Mobile Device Group object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-smart-groups-id
		GetSmartByIDV1(ctx context.Context, id string) (*ResourceSmartMobileDeviceGroup, *interfaces.Response, error)

		// CreateSmartV1 creates a new smart mobile device group (Create Smart Mobile Device Group record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-groups-smart-groups
		CreateSmartV1(ctx context.Context, request *RequestSmartMobileDeviceGroup) (*CreateSmartResponse, *interfaces.Response, error)

		// UpdateSmartByIDV1 updates the specified smart mobile device group by ID (Update specified Smart Mobile Device Group object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-mobile-device-groups-smart-groups-id
		UpdateSmartByIDV1(ctx context.Context, id string, request *RequestSmartMobileDeviceGroup) (*ResourceSmartMobileDeviceGroup, *interfaces.Response, error)

		// DeleteSmartByIDV1 removes the specified smart mobile device group by ID (Remove specified Smart Mobile Device Group record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-mobile-device-groups-smart-groups-id
		DeleteSmartByIDV1(ctx context.Context, id string) (*interfaces.Response, error)

		// ListStaticV1 returns all static mobile device groups (Get Static Mobile Device Group objects).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-static-groups
		ListStaticV1(ctx context.Context, rsqlQuery map[string]string) (*ListStaticResponse, *interfaces.Response, error)

		// GetStaticByIDV1 returns the specified static mobile device group by ID (Get specified Static Mobile Device Group object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-static-groups-id
		GetStaticByIDV1(ctx context.Context, id string) (*ResourceStaticMobileDeviceGroup, *interfaces.Response, error)

		// CreateStaticV1 creates a new static mobile device group (Create Static Mobile Device Group record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-groups-static-groups
		CreateStaticV1(ctx context.Context, request *RequestStaticMobileDeviceGroup) (*CreateStaticResponse, *interfaces.Response, error)

		// UpdateStaticByIDV1 updates the specified static mobile device group by ID (Update specified Static Mobile Device Group object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-mobile-device-groups-static-groups-id
		UpdateStaticByIDV1(ctx context.Context, id string, request *RequestStaticMobileDeviceGroup) (*ResourceStaticMobileDeviceGroup, *interfaces.Response, error)

		// DeleteStaticByIDV1 removes the specified static mobile device group by ID (Remove specified Static Mobile Device Group record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-mobile-device-groups-static-groups-id
		DeleteStaticByIDV1(ctx context.Context, id string) (*interfaces.Response, error)

		// ListAllV1 returns all mobile device groups (smart + static) as a simple list.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups
		ListAllV1(ctx context.Context) ([]ResourceMobileDeviceGroupSummary, *interfaces.Response, error)

		// GetStaticGroupMembershipV1 returns the mobile devices in the specified static group.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-static-group-membership-id
		GetStaticGroupMembershipV1(ctx context.Context, id string, rsqlQuery map[string]string) (*GroupMembershipResponse, *interfaces.Response, error)

		// GetSmartGroupMembershipV1 returns the mobile devices in the specified smart group.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-smart-group-membership-id
		GetSmartGroupMembershipV1(ctx context.Context, id string, rsqlQuery map[string]string) (*GroupMembershipResponse, *interfaces.Response, error)

		// EraseDevicesByGroupIDV1 erases all devices in the specified mobile device group.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-groups-id-erase
		EraseDevicesByGroupIDV1(ctx context.Context, id string, request *RequestEraseDevices) (*interfaces.Response, error)
	}

	// Service handles communication with the mobile device groups-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-smart-groups
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ MobileDeviceGroupsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Smart Groups CRUD (V1)
// -----------------------------------------------------------------------------

// ListSmartV1 returns all smart mobile device groups.
// URL: GET /api/v1/mobile-device-groups/smart-groups
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-smart-groups
func (s *Service) ListSmartV1(ctx context.Context, rsqlQuery map[string]string) (*ListSmartResponse, *interfaces.Response, error) {
	var result ListSmartResponse

	endpoint := EndpointSmartGroupsV1

	mergePage := func(pageData []byte) error {
		var rawData map[string]any
		if err := json.Unmarshal(pageData, &rawData); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}

		if totalCount, ok := rawData["totalCount"].(float64); ok {
			result.TotalCount = int(totalCount)
		}

		if results, ok := rawData["results"].([]any); ok {
			for _, item := range results {
				var group ResourceSmartMobileDeviceGroup
				if err := mapstructure.Decode(item, &group); err != nil {
					return fmt.Errorf("failed to decode smart mobile device group: %w", err)
				}
				result.Results = append(result.Results, group)
			}
		}

		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, nil, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list smart mobile device groups: %w", err)
	}

	return &result, resp, nil
}

// GetSmartByIDV1 returns the specified smart mobile device group by ID.
// URL: GET /api/v1/mobile-device-groups/smart-groups/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-smart-groups-id
func (s *Service) GetSmartByIDV1(ctx context.Context, id string) (*ResourceSmartMobileDeviceGroup, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("smart mobile device group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointSmartGroupsV1, id)

	var result ResourceSmartMobileDeviceGroup

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

// CreateSmartV1 creates a new smart mobile device group.
// URL: POST /api/v1/mobile-device-groups/smart-groups
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-groups-smart-groups
func (s *Service) CreateSmartV1(ctx context.Context, request *RequestSmartMobileDeviceGroup) (*CreateSmartResponse, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateSmartResponse

	endpoint := EndpointSmartGroupsV1

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

// UpdateSmartByIDV1 updates the specified smart mobile device group by ID.
// URL: PUT /api/v1/mobile-device-groups/smart-groups/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-mobile-device-groups-smart-groups-id
func (s *Service) UpdateSmartByIDV1(ctx context.Context, id string, request *RequestSmartMobileDeviceGroup) (*ResourceSmartMobileDeviceGroup, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointSmartGroupsV1, id)

	var result ResourceSmartMobileDeviceGroup

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

// DeleteSmartByIDV1 removes the specified smart mobile device group by ID.
// URL: DELETE /api/v1/mobile-device-groups/smart-groups/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-mobile-device-groups-smart-groups-id
func (s *Service) DeleteSmartByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("smart mobile device group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointSmartGroupsV1, id)

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

// -----------------------------------------------------------------------------
// Static Groups CRUD (V1)
// -----------------------------------------------------------------------------

// ListStaticV1 returns all static mobile device groups.
// URL: GET /api/v1/mobile-device-groups/static-groups
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-static-groups
func (s *Service) ListStaticV1(ctx context.Context, rsqlQuery map[string]string) (*ListStaticResponse, *interfaces.Response, error) {
	var result ListStaticResponse

	endpoint := EndpointStaticGroupsV1

	mergePage := func(pageData []byte) error {
		var rawData map[string]any
		if err := json.Unmarshal(pageData, &rawData); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}

		if totalCount, ok := rawData["totalCount"].(float64); ok {
			result.TotalCount = int(totalCount)
		}

		if results, ok := rawData["results"].([]any); ok {
			for _, item := range results {
				var group ResourceStaticMobileDeviceGroup
				if err := mapstructure.Decode(item, &group); err != nil {
					return fmt.Errorf("failed to decode static mobile device group: %w", err)
				}
				result.Results = append(result.Results, group)
			}
		}

		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, nil, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list static mobile device groups: %w", err)
	}

	return &result, resp, nil
}

// GetStaticByIDV1 returns the specified static mobile device group by ID.
// URL: GET /api/v1/mobile-device-groups/static-groups/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-static-groups-id
func (s *Service) GetStaticByIDV1(ctx context.Context, id string) (*ResourceStaticMobileDeviceGroup, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("static mobile device group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointStaticGroupsV1, id)

	var result ResourceStaticMobileDeviceGroup

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

// CreateStaticV1 creates a new static mobile device group.
// URL: POST /api/v1/mobile-device-groups/static-groups
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-groups-static-groups
func (s *Service) CreateStaticV1(ctx context.Context, request *RequestStaticMobileDeviceGroup) (*CreateStaticResponse, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	if request.Assignments == nil {
		request.Assignments = []StaticMobileDeviceGroupAssignment{}
	}

	var result CreateStaticResponse

	endpoint := EndpointStaticGroupsV1

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

// UpdateStaticByIDV1 updates the specified static mobile device group by ID (PATCH).
// URL: PATCH /api/v1/mobile-device-groups/static-groups/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-mobile-device-groups-static-groups-id
func (s *Service) UpdateStaticByIDV1(ctx context.Context, id string, request *RequestStaticMobileDeviceGroup) (*ResourceStaticMobileDeviceGroup, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	if request.Assignments == nil {
		request.Assignments = []StaticMobileDeviceGroupAssignment{}
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointStaticGroupsV1, id)

	var result ResourceStaticMobileDeviceGroup

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Patch(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteStaticByIDV1 removes the specified static mobile device group by ID.
// URL: DELETE /api/v1/mobile-device-groups/static-groups/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-mobile-device-groups-static-groups-id
func (s *Service) DeleteStaticByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("static mobile device group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointStaticGroupsV1, id)

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

// -----------------------------------------------------------------------------
// List All, Membership, and Erase (V1)
// -----------------------------------------------------------------------------

// ListAllV1 returns all mobile device groups (smart + static) as a simple list.
// URL: GET /api/v1/mobile-device-groups
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups
func (s *Service) ListAllV1(ctx context.Context) ([]ResourceMobileDeviceGroupSummary, *interfaces.Response, error) {
	var result []ResourceMobileDeviceGroupSummary

	endpoint := EndpointMobileDeviceGroupsV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// GetStaticGroupMembershipV1 returns the mobile devices in the specified static group.
// URL: GET /api/v1/mobile-device-groups/static-group-membership/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-static-group-membership-id
func (s *Service) GetStaticGroupMembershipV1(ctx context.Context, id string, rsqlQuery map[string]string) (*GroupMembershipResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("static group ID is required")
	}

	var result GroupMembershipResponse

	endpoint := fmt.Sprintf("%s/%s", EndpointStaticGroupMembershipV1, id)

	mergePage := func(pageData []byte) error {
		var rawData map[string]any
		if err := json.Unmarshal(pageData, &rawData); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}

		if totalCount, ok := rawData["totalCount"].(float64); ok {
			result.TotalCount = int(totalCount)
		}

		if results, ok := rawData["results"].([]any); ok {
			for _, item := range results {
				var member ResourceMobileDeviceMember
				if err := mapstructure.Decode(item, &member); err != nil {
					return fmt.Errorf("failed to decode mobile device member: %w", err)
				}
				result.Results = append(result.Results, member)
			}
		}

		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, nil, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get static group membership: %w", err)
	}

	return &result, resp, nil
}

// GetSmartGroupMembershipV1 returns the mobile devices in the specified smart group.
// URL: GET /api/v1/mobile-device-groups/smart-group-membership/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-smart-group-membership-id
func (s *Service) GetSmartGroupMembershipV1(ctx context.Context, id string, rsqlQuery map[string]string) (*GroupMembershipResponse, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("smart group ID is required")
	}

	var result GroupMembershipResponse

	endpoint := fmt.Sprintf("%s/%s", EndpointSmartGroupMembershipV1, id)

	mergePage := func(pageData []byte) error {
		var rawData map[string]any
		if err := json.Unmarshal(pageData, &rawData); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}

		if totalCount, ok := rawData["totalCount"].(float64); ok {
			result.TotalCount = int(totalCount)
		}

		if results, ok := rawData["results"].([]any); ok {
			for _, item := range results {
				var member ResourceMobileDeviceMember
				if err := mapstructure.Decode(item, &member); err != nil {
					return fmt.Errorf("failed to decode mobile device member: %w", err)
				}
				result.Results = append(result.Results, member)
			}
		}

		return nil
	}

	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, nil, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get smart group membership: %w", err)
	}

	return &result, resp, nil
}

// EraseDevicesByGroupIDV1 erases all devices in the specified mobile device group.
// URL: POST /api/v1/mobile-device-groups/{id}/erase
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-groups-id-erase
func (s *Service) EraseDevicesByGroupIDV1(ctx context.Context, id string, request *RequestEraseDevices) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/erase", EndpointMobileDeviceGroupsV1, id)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
