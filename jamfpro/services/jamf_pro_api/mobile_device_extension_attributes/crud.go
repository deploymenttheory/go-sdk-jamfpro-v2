package mobile_device_extension_attributes

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// MobileDeviceExtensionAttributesServiceInterface defines the interface for mobile device extension attribute operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-extension-attributes
	MobileDeviceExtensionAttributesServiceInterface interface {
		// ListMobileDeviceExtensionAttributesV1 returns all mobile device extension attribute objects (Get Mobile Device Extension Attribute objects).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-extension-attributes
		ListMobileDeviceExtensionAttributesV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error)

		// GetMobileDeviceExtensionAttributeByIDV1 returns the specified mobile device extension attribute by ID (Get specified Mobile Device Extension Attribute object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-extension-attributes-id
		GetMobileDeviceExtensionAttributeByIDV1(ctx context.Context, id string) (*ResourceMobileDeviceExtensionAttribute, *interfaces.Response, error)

		// CreateMobileDeviceExtensionAttributeV1 creates a new mobile device extension attribute (Create Mobile Device Extension Attribute record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-extension-attributes
		CreateMobileDeviceExtensionAttributeV1(ctx context.Context, request *RequestMobileDeviceExtensionAttribute) (*CreateResponse, *interfaces.Response, error)

		// UpdateMobileDeviceExtensionAttributeByIDV1 updates the specified mobile device extension attribute by ID (Update specified Mobile Device Extension Attribute object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-mobile-device-extension-attributes-id
		UpdateMobileDeviceExtensionAttributeByIDV1(ctx context.Context, id string, request *RequestMobileDeviceExtensionAttribute) (*ResourceMobileDeviceExtensionAttribute, *interfaces.Response, error)

		// DeleteMobileDeviceExtensionAttributeByIDV1 removes the specified mobile device extension attribute by ID (Remove specified Mobile Device Extension Attribute record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-mobile-device-extension-attributes-id
		DeleteMobileDeviceExtensionAttributeByIDV1(ctx context.Context, id string) (*interfaces.Response, error)

		// DeleteMobileDeviceExtensionAttributesByIDV1 deletes multiple mobile device extension attributes by their IDs (Delete multiple Mobile Device Extension Attributes by their IDs).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-extension-attributes-delete-multiple
		DeleteMobileDeviceExtensionAttributesByIDV1(ctx context.Context, req *DeleteMobileDeviceExtensionAttributesByIDRequest) (*interfaces.Response, error)
	}

	// Service handles communication with the mobile device extension attributes-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-extension-attributes
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ MobileDeviceExtensionAttributesServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Mobile Device Extension Attributes CRUD Operations
// -----------------------------------------------------------------------------

// ListMobileDeviceExtensionAttributesV1 returns all mobile device extension attribute objects.
// URL: GET /api/v1/mobile-device-extension-attributes
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-extension-attributes
func (s *Service) ListMobileDeviceExtensionAttributesV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointMobileDeviceExtensionAttributesV1

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

// GetMobileDeviceExtensionAttributeByIDV1 returns the specified mobile device extension attribute by ID.
// URL: GET /api/v1/mobile-device-extension-attributes/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-extension-attributes-id
func (s *Service) GetMobileDeviceExtensionAttributeByIDV1(ctx context.Context, id string) (*ResourceMobileDeviceExtensionAttribute, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("mobile device extension attribute ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointMobileDeviceExtensionAttributesV1, id)

	var result ResourceMobileDeviceExtensionAttribute

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

// CreateMobileDeviceExtensionAttributeV1 creates a new mobile device extension attribute.
// URL: POST /api/v1/mobile-device-extension-attributes
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-extension-attributes
func (s *Service) CreateMobileDeviceExtensionAttributeV1(ctx context.Context, request *RequestMobileDeviceExtensionAttribute) (*CreateResponse, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := EndpointMobileDeviceExtensionAttributesV1

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

// UpdateMobileDeviceExtensionAttributeByIDV1 updates the specified mobile device extension attribute by ID.
// URL: PUT /api/v1/mobile-device-extension-attributes/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-mobile-device-extension-attributes-id
func (s *Service) UpdateMobileDeviceExtensionAttributeByIDV1(ctx context.Context, id string, request *RequestMobileDeviceExtensionAttribute) (*ResourceMobileDeviceExtensionAttribute, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointMobileDeviceExtensionAttributesV1, id)

	var result ResourceMobileDeviceExtensionAttribute

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

// DeleteMobileDeviceExtensionAttributeByIDV1 removes the specified mobile device extension attribute by ID.
// URL: DELETE /api/v1/mobile-device-extension-attributes/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-mobile-device-extension-attributes-id
func (s *Service) DeleteMobileDeviceExtensionAttributeByIDV1(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("mobile device extension attribute ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", EndpointMobileDeviceExtensionAttributesV1, id)

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

// DeleteMobileDeviceExtensionAttributesByIDV1 deletes multiple mobile device extension attributes by their IDs.
// URL: POST /api/v1/mobile-device-extension-attributes/delete-multiple
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-extension-attributes-delete-multiple
func (s *Service) DeleteMobileDeviceExtensionAttributesByIDV1(ctx context.Context, req *DeleteMobileDeviceExtensionAttributesByIDRequest) (*interfaces.Response, error) {
	if req == nil || len(req.IDs) == 0 {
		return nil, fmt.Errorf("ids are required")
	}

	endpoint := EndpointMobileDeviceExtensionAttributesV1 + "/delete-multiple"

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
