package mobile_device_extension_attributes

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// MobileDeviceExtensionAttributesServiceInterface defines the interface for mobile device extension attribute operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-extension-attributes
	MobileDeviceExtensionAttributesServiceInterface interface {
		// ListV1 returns all mobile device extension attribute objects (Get Mobile Device Extension Attribute objects).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-extension-attributes
		ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error)

		// GetByIDV1 returns the specified mobile device extension attribute by ID (Get specified Mobile Device Extension Attribute object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-extension-attributes-id
		GetByIDV1(ctx context.Context, id string) (*ResourceMobileDeviceExtensionAttribute, *resty.Response, error)

		// CreateV1 creates a new mobile device extension attribute (Create Mobile Device Extension Attribute record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-extension-attributes
		CreateV1(ctx context.Context, request *RequestMobileDeviceExtensionAttribute) (*CreateResponse, *resty.Response, error)

		// UpdateByIDV1 updates the specified mobile device extension attribute by ID (Update specified Mobile Device Extension Attribute object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-mobile-device-extension-attributes-id
		UpdateByIDV1(ctx context.Context, id string, request *RequestMobileDeviceExtensionAttribute) (*ResourceMobileDeviceExtensionAttribute, *resty.Response, error)

		// DeleteByIDV1 removes the specified mobile device extension attribute by ID (Remove specified Mobile Device Extension Attribute record).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-mobile-device-extension-attributes-id
		DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error)

		// DeleteMobileDeviceExtensionAttributesByIDV1 deletes multiple mobile device extension attributes by their IDs (Delete multiple Mobile Device Extension Attributes by their IDs).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-extension-attributes-delete-multiple
		DeleteMobileDeviceExtensionAttributesByIDV1(ctx context.Context, req *DeleteMobileDeviceExtensionAttributesByIDRequest) (*resty.Response, error)

		// GetHistoryByIDV1 returns the history for the specified mobile device extension attribute by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-extension-attributes-id-history
		GetHistoryByIDV1(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error)

		// AddHistoryNoteByIDV1 adds a history note to the specified mobile device extension attribute.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-extension-attributes-id-history
		AddHistoryNoteByIDV1(ctx context.Context, id string, req *AddHistoryNoteRequest) (*resty.Response, error)

		// GetDataDependencyByIDV1 returns smart group dependent objects for the specified mobile device extension attribute.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-extension-attributes-id-data-dependency
		GetDataDependencyByIDV1(ctx context.Context, id string) (*DataDependencyResponse, *resty.Response, error)
	}

	// Service handles communication with the mobile device extension attributes-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-extension-attributes
	MobileDeviceExtensionAttributes struct {
		client transport.HTTPClient
	}
)

var _ MobileDeviceExtensionAttributesServiceInterface = (*MobileDeviceExtensionAttributes)(nil)

func NewMobileDeviceExtensionAttributes(client transport.HTTPClient) *MobileDeviceExtensionAttributes {
	return &MobileDeviceExtensionAttributes{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Mobile Device Extension Attributes CRUD Operations
// -----------------------------------------------------------------------------

// ListV1 returns all mobile device extension attribute objects.
// URL: GET /api/v1/mobile-device-extension-attributes
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-extension-attributes
func (s *MobileDeviceExtensionAttributes) ListV1(ctx context.Context, rsqlQuery map[string]string) (*ListResponse, *resty.Response, error) {
	var result ListResponse

	endpoint := constants.EndpointJamfProMobileDeviceExtensionAttributesV1

	mergePage := func(pageData []byte) error {
		var pageItems []ResourceMobileDeviceExtensionAttribute
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
		return nil, resp, fmt.Errorf("failed to list mobile device extension attributes: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// GetByIDV1 returns the specified mobile device extension attribute by ID.
// URL: GET /api/v1/mobile-device-extension-attributes/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-extension-attributes-id
func (s *MobileDeviceExtensionAttributes) GetByIDV1(ctx context.Context, id string) (*ResourceMobileDeviceExtensionAttribute, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("mobile device extension attribute ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProMobileDeviceExtensionAttributesV1, id)

	var result ResourceMobileDeviceExtensionAttribute

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// CreateV1 creates a new mobile device extension attribute.
// URL: POST /api/v1/mobile-device-extension-attributes
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-extension-attributes
func (s *MobileDeviceExtensionAttributes) CreateV1(ctx context.Context, request *RequestMobileDeviceExtensionAttribute) (*CreateResponse, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result CreateResponse

	endpoint := constants.EndpointJamfProMobileDeviceExtensionAttributesV1

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

// UpdateByIDV1 updates the specified mobile device extension attribute by ID.
// URL: PUT /api/v1/mobile-device-extension-attributes/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-mobile-device-extension-attributes-id
func (s *MobileDeviceExtensionAttributes) UpdateByIDV1(ctx context.Context, id string, request *RequestMobileDeviceExtensionAttribute) (*ResourceMobileDeviceExtensionAttribute, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProMobileDeviceExtensionAttributesV1, id)

	var result ResourceMobileDeviceExtensionAttribute

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

// DeleteByIDV1 removes the specified mobile device extension attribute by ID.
// URL: DELETE /api/v1/mobile-device-extension-attributes/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-mobile-device-extension-attributes-id
func (s *MobileDeviceExtensionAttributes) DeleteByIDV1(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("mobile device extension attribute ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProMobileDeviceExtensionAttributesV1, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *MobileDeviceExtensionAttributes) DeleteMobileDeviceExtensionAttributesByIDV1(ctx context.Context, req *DeleteMobileDeviceExtensionAttributesByIDRequest) (*resty.Response, error) {
	if req == nil || len(req.IDs) == 0 {
		return nil, fmt.Errorf("ids are required")
	}

	endpoint := constants.EndpointJamfProMobileDeviceExtensionAttributesV1 + "/delete-multiple"

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetHistoryByIDV1 returns the history for the specified mobile device extension attribute by ID.
// URL: GET /api/v1/mobile-device-extension-attributes/{id}/history
// Query params (optional): filter (RSQL), sort, page, page-size.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-extension-attributes-id-history
func (s *MobileDeviceExtensionAttributes) GetHistoryByIDV1(ctx context.Context, id string, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("mobile device extension attribute ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProMobileDeviceExtensionAttributesV1, id)

	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var pageItems []HistoryItem
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
		return nil, resp, fmt.Errorf("failed to get mobile device extension attribute history: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// AddHistoryNoteByIDV1 adds a history note to the specified mobile device extension attribute.
// URL: POST /api/v1/mobile-device-extension-attributes/{id}/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-extension-attributes-id-history
func (s *MobileDeviceExtensionAttributes) AddHistoryNoteByIDV1(ctx context.Context, id string, req *AddHistoryNoteRequest) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("mobile device extension attribute ID is required")
	}
	if req == nil || req.Note == "" {
		return nil, fmt.Errorf("note is required")
	}

	endpoint := fmt.Sprintf("%s/%s/history", constants.EndpointJamfProMobileDeviceExtensionAttributesV1, id)

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetDataDependencyByIDV1 returns smart group dependent objects for the specified mobile device extension attribute.
// URL: GET /api/v1/mobile-device-extension-attributes/{id}/data-dependency
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-extension-attributes-id-data-dependency
func (s *MobileDeviceExtensionAttributes) GetDataDependencyByIDV1(ctx context.Context, id string) (*DataDependencyResponse, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("mobile device extension attribute ID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/data-dependency", constants.EndpointJamfProMobileDeviceExtensionAttributesV1, id)

	var result DataDependencyResponse

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
