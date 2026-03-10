package static_mobile_device_groups

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the static mobile device groups-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-device-groups-static-groups
	StaticMobileDeviceGroups struct {
		client client.Client
	}
)

func NewStaticMobileDeviceGroups(client client.Client) *StaticMobileDeviceGroups {
	return &StaticMobileDeviceGroups{client: client}
}

// List returns all static mobile device groups.
// URL: GET /api/v2/mobile-device-groups/static-groups
func (s *StaticMobileDeviceGroups) List(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointJamfProStaticMobileDeviceGroups2V2

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		SetResult(&result).
		Get(endpoint)
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

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProStaticMobileDeviceGroups2V2, id)

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

	endpoint := constants.EndpointJamfProStaticMobileDeviceGroups2V2

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

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProStaticMobileDeviceGroups2V2, id)

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

// DeleteByID removes the specified static mobile device group by ID.
// URL: DELETE /api/v2/mobile-device-groups/static-groups/{id}
func (s *StaticMobileDeviceGroups) DeleteByID(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("static mobile device group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProStaticMobileDeviceGroups2V2, id)

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
