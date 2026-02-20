package mobile_device_groups

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
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
		CreateSmartV1(ctx context.Context, req *RequestSmartMobileDeviceGroup) (*CreateSmartResponse, *interfaces.Response, error)

		// UpdateSmartByIDV1 updates the specified smart mobile device group by ID (Update specified Smart Mobile Device Group object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-mobile-device-groups-smart-groups-id
		UpdateSmartByIDV1(ctx context.Context, id string, req *RequestSmartMobileDeviceGroup) (*ResourceSmartMobileDeviceGroup, *interfaces.Response, error)

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
		CreateStaticV1(ctx context.Context, req *RequestStaticMobileDeviceGroup) (*CreateStaticResponse, *interfaces.Response, error)

		// UpdateStaticByIDV1 updates the specified static mobile device group by ID (Update specified Static Mobile Device Group object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-mobile-device-groups-static-groups-id
		UpdateStaticByIDV1(ctx context.Context, id string, req *RequestStaticMobileDeviceGroup) (*ResourceStaticMobileDeviceGroup, *interfaces.Response, error)

		// DeleteStaticByIDV1 removes the specified static mobile device group by ID (Remove specified Static Mobile Device Group record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-mobile-device-groups-static-groups-id
		DeleteStaticByIDV1(ctx context.Context, id string) (*interfaces.Response, error)
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
func (s *Service) CreateSmartV1(ctx context.Context, req *RequestSmartMobileDeviceGroup) (*CreateSmartResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateSmartResponse

	endpoint := EndpointSmartGroupsV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateSmartByIDV1 updates the specified smart mobile device group by ID.
// URL: PUT /api/v1/mobile-device-groups/smart-groups/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-mobile-device-groups-smart-groups-id
func (s *Service) UpdateSmartByIDV1(ctx context.Context, id string, req *RequestSmartMobileDeviceGroup) (*ResourceSmartMobileDeviceGroup, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointSmartGroupsV1, id)

	var result ResourceSmartMobileDeviceGroup

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, req, headers, &result)
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
func (s *Service) CreateStaticV1(ctx context.Context, req *RequestStaticMobileDeviceGroup) (*CreateStaticResponse, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Assignments == nil {
		req.Assignments = []StaticMobileDeviceGroupAssignment{}
	}

	var result CreateStaticResponse

	endpoint := EndpointStaticGroupsV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateStaticByIDV1 updates the specified static mobile device group by ID (PATCH).
// URL: PATCH /api/v1/mobile-device-groups/static-groups/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v1-mobile-device-groups-static-groups-id
func (s *Service) UpdateStaticByIDV1(ctx context.Context, id string, req *RequestStaticMobileDeviceGroup) (*ResourceStaticMobileDeviceGroup, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Assignments == nil {
		req.Assignments = []StaticMobileDeviceGroupAssignment{}
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointStaticGroupsV1, id)

	var result ResourceStaticMobileDeviceGroup

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
