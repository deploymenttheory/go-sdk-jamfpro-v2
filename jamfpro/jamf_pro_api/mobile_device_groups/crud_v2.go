package mobile_device_groups

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/smartgroupvalidation"
	"resty.dev/v3"
)

// -----------------------------------------------------------------------------
// Mobile Device Groups CRUD (V2) — Jamf Pro 11.28, replaces the V1 surface.
//
// The JSON request/response shapes are identical to the V1 endpoints (verified
// against a live Jamf Pro 11.28.1 instance), so the V1 model types are reused.
// V2 adds stricter client-side validation: the 255-character group-name cap, an
// andOr enum ("and"/"or") for smart-group criteria, and set semantics
// (deduplication) for static-group assignments.
// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------
// Smart Groups CRUD (V2)
// -----------------------------------------------------------------------------

// ListSmartV2 returns all smart mobile device groups.
// URL: GET /api/v2/mobile-device-groups/smart-groups
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-device-groups-smart-groups
func (s *MobileDeviceGroups) ListSmartV2(ctx context.Context, rsqlQuery map[string]string) (*ListSmartResponse, *resty.Response, error) {
	var result ListSmartResponse

	endpoint := constants.EndpointJamfProSmartMobileDeviceGroupsV2

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

// GetSmartByIDV2 returns the specified smart mobile device group by ID.
// URL: GET /api/v2/mobile-device-groups/smart-groups/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-device-groups-smart-groups-id
func (s *MobileDeviceGroups) GetSmartByIDV2(ctx context.Context, id string) (*ResourceSmartMobileDeviceGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("smart mobile device group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSmartMobileDeviceGroupsV2, id)

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

// CreateSmartV2 creates a new smart mobile device group.
// URL: POST /api/v2/mobile-device-groups/smart-groups
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-mobile-device-groups-smart-groups
func (s *MobileDeviceGroups) CreateSmartV2(ctx context.Context, request *RequestSmartMobileDeviceGroup) (*CreateSmartResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if err := validateSmartMobileDeviceGroupV2(request); err != nil {
		return nil, nil, err
	}

	var result CreateSmartResponse

	endpoint := constants.EndpointJamfProSmartMobileDeviceGroupsV2

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

// UpdateSmartByIDV2 updates the specified smart mobile device group by ID.
// URL: PUT /api/v2/mobile-device-groups/smart-groups/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-mobile-device-groups-smart-groups-id
func (s *MobileDeviceGroups) UpdateSmartByIDV2(ctx context.Context, id string, request *RequestSmartMobileDeviceGroup) (*ResourceSmartMobileDeviceGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if err := validateSmartMobileDeviceGroupV2(request); err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSmartMobileDeviceGroupsV2, id)

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

// DeleteSmartByIDV2 removes the specified smart mobile device group by ID.
// URL: DELETE /api/v2/mobile-device-groups/smart-groups/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-mobile-device-groups-smart-groups-id
func (s *MobileDeviceGroups) DeleteSmartByIDV2(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("smart mobile device group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSmartMobileDeviceGroupsV2, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// -----------------------------------------------------------------------------
// Static Groups CRUD (V2)
// -----------------------------------------------------------------------------

// ListStaticV2 returns all static mobile device groups.
// URL: GET /api/v2/mobile-device-groups/static-groups
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-device-groups-static-groups
func (s *MobileDeviceGroups) ListStaticV2(ctx context.Context, rsqlQuery map[string]string) (*ListStaticResponse, *resty.Response, error) {
	var result ListStaticResponse

	endpoint := constants.EndpointJamfProStaticMobileDeviceGroupsV2

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

// GetStaticByIDV2 returns the specified static mobile device group by ID.
// URL: GET /api/v2/mobile-device-groups/static-groups/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-device-groups-static-groups-id
func (s *MobileDeviceGroups) GetStaticByIDV2(ctx context.Context, id string) (*ResourceStaticMobileDeviceGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("static mobile device group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProStaticMobileDeviceGroupsV2, id)

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

// CreateStaticV2 creates a new static mobile device group. Assignments are a set
// (uniqueItems) and are deduplicated before sending; a nil assignments slice is
// serialized as an empty array because the API returns HTTP 500 when the
// "assignments" key is missing.
// URL: POST /api/v2/mobile-device-groups/static-groups
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-mobile-device-groups-static-groups
func (s *MobileDeviceGroups) CreateStaticV2(ctx context.Context, request *RequestStaticMobileDeviceGroup) (*CreateStaticResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if err := validateStaticMobileDeviceGroupV2(request); err != nil {
		return nil, nil, err
	}
	request.Assignments = dedupeStaticAssignments(request.Assignments)

	var result CreateStaticResponse

	endpoint := constants.EndpointJamfProStaticMobileDeviceGroupsV2

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

// UpdateStaticByIDV2 updates the specified static mobile device group by ID (PATCH).
// Assignments are deduplicated and always serialized as a non-nil array.
// URL: PATCH /api/v2/mobile-device-groups/static-groups/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v2-mobile-device-groups-static-groups-id
func (s *MobileDeviceGroups) UpdateStaticByIDV2(ctx context.Context, id string, request *RequestStaticMobileDeviceGroup) (*ResourceStaticMobileDeviceGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if err := validateStaticMobileDeviceGroupV2(request); err != nil {
		return nil, nil, err
	}
	request.Assignments = dedupeStaticAssignments(request.Assignments)

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProStaticMobileDeviceGroupsV2, id)

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

// DeleteStaticByIDV2 removes the specified static mobile device group by ID.
// URL: DELETE /api/v2/mobile-device-groups/static-groups/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-mobile-device-groups-static-groups-id
func (s *MobileDeviceGroups) DeleteStaticByIDV2(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("static mobile device group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProStaticMobileDeviceGroupsV2, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// -----------------------------------------------------------------------------
// List All, Membership, and Erase (V2)
// -----------------------------------------------------------------------------

// ListAllV2 returns all mobile device groups (smart + static) as a simple list.
// URL: GET /api/v2/mobile-device-groups
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-device-groups
func (s *MobileDeviceGroups) ListAllV2(ctx context.Context) ([]ResourceMobileDeviceGroupSummary, *resty.Response, error) {
	var result []ResourceMobileDeviceGroupSummary

	endpoint := constants.EndpointJamfProMobileDeviceGroupsV2

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}

// GetStaticGroupMembershipV2 returns the mobile devices in the specified static group.
// URL: GET /api/v2/mobile-device-groups/static-group-membership/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-device-groups-static-group-membership-id
func (s *MobileDeviceGroups) GetStaticGroupMembershipV2(ctx context.Context, id string, rsqlQuery map[string]string) (*GroupMembershipResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("static group ID is required")
	}

	var result GroupMembershipResponse

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProStaticMobileDeviceGroupMembershipV2, id)

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

// GetSmartGroupMembershipV2 returns the mobile devices in the specified smart group.
// URL: GET /api/v2/mobile-device-groups/smart-group-membership/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-device-groups-smart-group-membership-id
func (s *MobileDeviceGroups) GetSmartGroupMembershipV2(ctx context.Context, id string, rsqlQuery map[string]string) (*GroupMembershipResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("smart group ID is required")
	}

	var result GroupMembershipResponse

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSmartMobileDeviceGroupMembershipV2, id)

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

// EraseDevicesByGroupIDV2 erases all devices in the specified mobile device group.
// URL: POST /api/v2/mobile-device-groups/{id}/erase
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-mobile-device-groups-id-erase
func (s *MobileDeviceGroups) EraseDevicesByGroupIDV2(ctx context.Context, id string, request *RequestEraseDevices) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/erase", constants.EndpointJamfProMobileDeviceGroupsV2, id)

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

// dedupeStaticAssignments removes duplicate device assignments (assignments are a
// set / uniqueItems) and guarantees a non-nil slice so JSON serialization always
// emits the "assignments" array (a missing key causes the API to return HTTP 500).
func dedupeStaticAssignments(in []StaticMobileDeviceGroupAssignment) []StaticMobileDeviceGroupAssignment {
	if len(in) == 0 {
		return []StaticMobileDeviceGroupAssignment{}
	}
	ids := make([]string, 0, len(in))
	byID := make(map[string]StaticMobileDeviceGroupAssignment, len(in))
	for _, a := range in {
		ids = append(ids, a.MobileDeviceID)
		byID[a.MobileDeviceID] = a
	}
	uniqueIDs := smartgroupvalidation.DedupeStrings(ids)
	out := make([]StaticMobileDeviceGroupAssignment, 0, len(uniqueIDs))
	for _, id := range uniqueIDs {
		out = append(out, byID[id])
	}
	return out
}
