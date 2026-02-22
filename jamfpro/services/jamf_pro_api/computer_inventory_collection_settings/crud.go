package computer_inventory_collection_settings

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ComputerInventoryCollectionSettingsServiceInterface defines the interface for Computer Inventory Collection Settings operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-inventory-collection-settings
	ComputerInventoryCollectionSettingsServiceInterface interface {
		// GetV2 retrieves the computer inventory collection settings.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-inventory-collection-settings
		GetV2(ctx context.Context) (*ResourceComputerInventoryCollectionSettings, *interfaces.Response, error)

		// UpdateV2 updates the computer inventory collection settings.
		//
		// Returns 204 No Content on success with no response body.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v2-computer-inventory-collection-settings
		UpdateV2(ctx context.Context, settings *ResourceComputerInventoryCollectionSettings) (*interfaces.Response, error)

		// CreateCustomPathV2 creates a custom path for computer inventory collection settings.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-computer-inventory-collection-settings-custom-path
		CreateCustomPathV2(ctx context.Context, req *RequestCustomPath) (*SubsetPathItem, *interfaces.Response, error)

		// DeleteCustomPathByIDV2 deletes a custom path by ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-computer-inventory-collection-settings-custom-path-id
		DeleteCustomPathByIDV2(ctx context.Context, id string) (*interfaces.Response, error)
	}

	// Service handles communication with the Computer Inventory Collection Settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-inventory-collection-settings
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ComputerInventoryCollectionSettingsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetV2 retrieves the computer inventory collection settings.
// URL: GET /api/v2/computer-inventory-collection-settings
// https://developer.jamf.com/jamf-pro/reference/get_v2-computer-inventory-collection-settings
func (s *Service) GetV2(ctx context.Context) (*ResourceComputerInventoryCollectionSettings, *interfaces.Response, error) {
	endpoint := EndpointComputerInventoryCollectionSettingsV2

	var result ResourceComputerInventoryCollectionSettings

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

// UpdateV2 updates the computer inventory collection settings.
// URL: PATCH /api/v2/computer-inventory-collection-settings
// Returns 204 No Content on success with no response body.
// https://developer.jamf.com/jamf-pro/reference/patch_v2-computer-inventory-collection-settings
func (s *Service) UpdateV2(ctx context.Context, settings *ResourceComputerInventoryCollectionSettings) (*interfaces.Response, error) {
	if settings == nil {
		return nil, fmt.Errorf("settings is required")
	}

	endpoint := EndpointComputerInventoryCollectionSettingsV2

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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
func (s *Service) CreateCustomPathV2(ctx context.Context, req *RequestCustomPath) (*SubsetPathItem, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request body is required")
	}
	if req.Path == "" {
		return nil, nil, fmt.Errorf("path is required")
	}

	endpoint := fmt.Sprintf("%s/custom-path", EndpointComputerInventoryCollectionSettingsV2)

	var result SubsetPathItem

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

// DeleteCustomPathByIDV2 deletes a custom path by ID.
// URL: DELETE /api/v2/computer-inventory-collection-settings/custom-path/{id}
// https://developer.jamf.com/jamf-pro/reference/delete_v2-computer-inventory-collection-settings-custom-path-id
func (s *Service) DeleteCustomPathByIDV2(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("custom path ID is required")
	}

	endpoint := fmt.Sprintf("%s/custom-path/%s", EndpointComputerInventoryCollectionSettingsV2, id)

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
