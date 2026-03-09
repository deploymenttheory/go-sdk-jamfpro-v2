package static_mobile_device_groups

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// ServiceInterface defines the interface for static mobile device group operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-device-groups-static-groups
	ServiceInterface interface {
		// List returns all static mobile device groups (Get Static Mobile Device Group objects).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-device-groups-static-groups
		List(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error)

		// GetByID returns the specified static mobile device group by ID (Get specified Static Mobile Device Group object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-device-groups-static-groups-id
		GetByID(ctx context.Context, id string) (*ResourceStaticMobileDeviceGroup, *resty.Response, error)

		// Create creates a new static mobile device group (Create Static Mobile Device Group record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-mobile-device-groups-static-groups
		Create(ctx context.Context, request *RequestStaticMobileDeviceGroup) (*CreateResponse, *resty.Response, error)

		// UpdateByID updates the specified static mobile device group by ID (Update specified Static Mobile Device Group object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v2-mobile-device-groups-static-groups-id
		UpdateByID(ctx context.Context, id string, request *RequestStaticMobileDeviceGroup) (*ResourceStaticMobileDeviceGroup, *resty.Response, error)

		// DeleteByID removes the specified static mobile device group by ID (Remove specified Static Mobile Device Group record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-mobile-device-groups-static-groups-id
		DeleteByID(ctx context.Context, id string) (*resty.Response, error)
	}

	// Service handles communication with the static mobile device groups-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-device-groups-static-groups
	StaticMobileDeviceGroups struct {
		client interfaces.HTTPClient
	}
)

var _ ServiceInterface = (*StaticMobileDeviceGroups)(nil)

func NewStaticMobileDeviceGroups(client interfaces.HTTPClient) *StaticMobileDeviceGroups {
	return &StaticMobileDeviceGroups{client: client}
}

// List returns all static mobile device groups.
// URL: GET /api/v2/mobile-device-groups/static-groups
func (s *StaticMobileDeviceGroups) List(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := EndpointStaticGroupsV2

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByID returns the specified static mobile device group by ID.
// URL: GET /api/v2/mobile-device-groups/static-groups/{id}
func (s *StaticMobileDeviceGroups) GetByID(ctx context.Context, id string) (*ResourceStaticMobileDeviceGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("static mobile device group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointStaticGroupsV2, id)

	var result ResourceStaticMobileDeviceGroup

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// Create creates a new static mobile device group.
// URL: POST /api/v2/mobile-device-groups/static-groups
func (s *StaticMobileDeviceGroups) Create(ctx context.Context, request *RequestStaticMobileDeviceGroup) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	if request.Assignments == nil {
		request.Assignments = []StaticMobileDeviceGroupAssignment{}
	}

	var result CreateResponse

	endpoint := EndpointStaticGroupsV2

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

// UpdateByID updates the specified static mobile device group by ID (PATCH).
// URL: PATCH /api/v2/mobile-device-groups/static-groups/{id}
func (s *StaticMobileDeviceGroups) UpdateByID(ctx context.Context, id string, request *RequestStaticMobileDeviceGroup) (*ResourceStaticMobileDeviceGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	if request.Assignments == nil {
		request.Assignments = []StaticMobileDeviceGroupAssignment{}
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointStaticGroupsV2, id)

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

// DeleteByID removes the specified static mobile device group by ID.
// URL: DELETE /api/v2/mobile-device-groups/static-groups/{id}
func (s *StaticMobileDeviceGroups) DeleteByID(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("static mobile device group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointStaticGroupsV2, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
