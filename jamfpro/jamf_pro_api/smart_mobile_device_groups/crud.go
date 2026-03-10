package smart_mobile_device_groups

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the smart mobile device groups-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-device-groups-smart-groups
	SmartMobileDeviceGroups struct {
		client client.Client
	}
)

// NewService returns a new smart mobile device groups service.
func NewSmartMobileDeviceGroups(client client.Client) *SmartMobileDeviceGroups {
	return &SmartMobileDeviceGroups{client: client}
}

// List returns all smart mobile device groups.
// URL: GET /api/v2/mobile-device-groups/smart-groups
func (s *SmartMobileDeviceGroups) List(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointJamfProSmartMobileDeviceGroups2V2

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByID returns the specified smart mobile device group by ID.
// URL: GET /api/v2/mobile-device-groups/smart-groups/{id}
func (s *SmartMobileDeviceGroups) GetByID(ctx context.Context, id string) (*ResourceSmartMobileDeviceGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("smart mobile device group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSmartMobileDeviceGroups2V2, id)

	var result ResourceSmartMobileDeviceGroup

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetByName returns a smart mobile device group by name.
// Uses List with a filter and returns the first match.
func (s *SmartMobileDeviceGroups) GetByName(ctx context.Context, name string) (*ListItem, *resty.Response, error) {
	if name == "" {
		return nil, nil, fmt.Errorf("smart mobile device group name is required")
	}

	query := map[string]string{
		"filter": fmt.Sprintf("groupName==\"%s\"", name),
	}

	list, resp, err := s.List(ctx, query)
	if err != nil {
		return nil, resp, err
	}

	if len(list.Results) == 0 {
		return nil, resp, fmt.Errorf("smart mobile device group with name %q not found", name)
	}

	return &list.Results[0], resp, nil
}

// GetMembership returns the membership of a smart mobile device group by ID.
// URL: GET /api/v2/mobile-device-groups/smart-group-membership/{id}
func (s *SmartMobileDeviceGroups) GetMembership(ctx context.Context, id string, rsqlQuery map[string]string) (*MembershipResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("smart mobile device group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSmartMobileDeviceGroupMembership, id)

	var result MembershipResponse

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// Create creates a new smart mobile device group.
// URL: POST /api/v2/mobile-device-groups/smart-groups
func (s *SmartMobileDeviceGroups) Create(ctx context.Context, request *RequestSmartMobileDeviceGroup) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := constants.EndpointJamfProSmartMobileDeviceGroups2V2

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateByID updates the specified smart mobile device group by ID.
// URL: PUT /api/v2/mobile-device-groups/smart-groups/{id}
func (s *SmartMobileDeviceGroups) UpdateByID(ctx context.Context, id string, request *RequestSmartMobileDeviceGroup) (*ResourceSmartMobileDeviceGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSmartMobileDeviceGroups2V2, id)

	var result ResourceSmartMobileDeviceGroup

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteByID removes the specified smart mobile device group by ID.
// URL: DELETE /api/v2/mobile-device-groups/smart-groups/{id}
func (s *SmartMobileDeviceGroups) DeleteByID(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("smart mobile device group ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProSmartMobileDeviceGroups2V2, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
