package computer_inventory_collection_settings

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// ComputerInventoryCollectionSettingsServiceInterface defines the interface for Computer Inventory Collection Settings operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-inventory-collection-settings
	ComputerInventoryCollectionSettingsServiceInterface interface {
		// GetV2 retrieves the computer inventory collection settings.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-inventory-collection-settings
		GetV2(ctx context.Context) (*ResourceComputerInventoryCollectionSettings, *resty.Response, error)

		// UpdateV2 updates the computer inventory collection settings.
		//
		// Returns 204 No Content on success with no response body.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v2-computer-inventory-collection-settings
		UpdateV2(ctx context.Context, settings *ResourceComputerInventoryCollectionSettings) (*resty.Response, error)

		// CreateCustomPathV2 creates a custom path for computer inventory collection settings.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-computer-inventory-collection-settings-custom-path
		CreateCustomPathV2(ctx context.Context, req *RequestCustomPath) (*SubsetPathItem, *resty.Response, error)

		// DeleteCustomPathByIDV2 deletes a custom path by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-computer-inventory-collection-settings-custom-path-id
		DeleteCustomPathByIDV2(ctx context.Context, id string) (*resty.Response, error)
	}

	// Service handles communication with the Computer Inventory Collection Settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-inventory-collection-settings
	ComputerInventoryCollectionSettings struct {
		client transport.HTTPClient
	}
)

var _ ComputerInventoryCollectionSettingsServiceInterface = (*ComputerInventoryCollectionSettings)(nil)

func NewComputerInventoryCollectionSettings(client transport.HTTPClient) *ComputerInventoryCollectionSettings {
	return &ComputerInventoryCollectionSettings{client: client}
}

// GetV2 retrieves the computer inventory collection settings.
// URL: GET /api/v2/computer-inventory-collection-settings
// https://developer.jamf.com/jamf-pro/reference/get_v2-computer-inventory-collection-settings
func (s *ComputerInventoryCollectionSettings) GetV2(ctx context.Context) (*ResourceComputerInventoryCollectionSettings, *resty.Response, error) {
	endpoint := constants.EndpointJamfProComputerInventoryCollectionSettingsV2

	var result ResourceComputerInventoryCollectionSettings

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateV2 updates the computer inventory collection settings.
// URL: PATCH /api/v2/computer-inventory-collection-settings
// Returns 204 No Content on success with no response body.
// https://developer.jamf.com/jamf-pro/reference/patch_v2-computer-inventory-collection-settings
func (s *ComputerInventoryCollectionSettings) UpdateV2(ctx context.Context, settings *ResourceComputerInventoryCollectionSettings) (*resty.Response, error) {
	if settings == nil {
		return nil, fmt.Errorf("settings is required")
	}

	endpoint := constants.EndpointJamfProComputerInventoryCollectionSettingsV2

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Patch(ctx, endpoint, settings, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// CreateCustomPathV2 creates a custom path for computer inventory collection settings.
// URL: POST /api/v2/computer-inventory-collection-settings/custom-path
// https://developer.jamf.com/jamf-pro/reference/post_v2-computer-inventory-collection-settings-custom-path
func (s *ComputerInventoryCollectionSettings) CreateCustomPathV2(ctx context.Context, req *RequestCustomPath) (*SubsetPathItem, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request body is required")
	}
	if req.Path == "" {
		return nil, nil, fmt.Errorf("path is required")
	}

	endpoint := fmt.Sprintf("%s/custom-path", constants.EndpointJamfProComputerInventoryCollectionSettingsV2)

	var result SubsetPathItem

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteCustomPathByIDV2 deletes a custom path by ID.
// URL: DELETE /api/v2/computer-inventory-collection-settings/custom-path/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v2-computer-inventory-collection-settings-custom-path-id
func (s *ComputerInventoryCollectionSettings) DeleteCustomPathByIDV2(ctx context.Context, id string) (*resty.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("custom path ID is required")
	}

	endpoint := fmt.Sprintf("%s/custom-path/%s", constants.EndpointJamfProComputerInventoryCollectionSettingsV2, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
