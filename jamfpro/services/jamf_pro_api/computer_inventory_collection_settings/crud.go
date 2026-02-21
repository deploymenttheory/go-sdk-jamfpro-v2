package computer_inventory_collection_settings

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ComputerInventoryCollectionSettingsServiceInterface defines the interface for Computer Inventory Collection Settings operations.
	// Uses v2 API for all operations. Manages settings for computer inventory collection and custom application paths.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-inventory-collection-settings
	ComputerInventoryCollectionSettingsServiceInterface interface {
		// GetV2 returns the computer inventory collection settings (Get Computer Inventory Collection Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-inventory-collection-settings
		GetV2(ctx context.Context) (*ResourceComputerInventoryCollectionSettings, *interfaces.Response, error)

		// UpdateV2 updates the computer inventory collection settings using merge-patch semantics (Update Computer Inventory Collection Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v2-computer-inventory-collection-settings
		UpdateV2(ctx context.Context, request *ResourceComputerInventoryCollectionSettings) (*interfaces.Response, error)

		// CreateCustomPathV2 creates a custom application path for inventory collection (Create Custom Path).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-computer-inventory-collection-settings-custom-path
		CreateCustomPathV2(ctx context.Context, request *CustomPathRequest) (*CustomPathResponse, *interfaces.Response, error)

		// DeleteCustomPathByIDV2 deletes a custom application path by ID (Delete Custom Path).
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

// GetV2 returns the computer inventory collection settings.
// URL: GET /api/v2/computer-inventory-collection-settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-inventory-collection-settings
func (s *Service) GetV2(ctx context.Context) (*ResourceComputerInventoryCollectionSettings, *interfaces.Response, error) {
	var result ResourceComputerInventoryCollectionSettings

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, EndpointComputerInventoryCollectionSettingsV2, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateV2 updates the computer inventory collection settings using merge-patch semantics.
// URL: PATCH /api/v2/computer-inventory-collection-settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/patch_v2-computer-inventory-collection-settings
func (s *Service) UpdateV2(ctx context.Context, request *ResourceComputerInventoryCollectionSettings) (*interfaces.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Patch(ctx, EndpointComputerInventoryCollectionSettingsV2, request, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// CreateCustomPathV2 creates a custom application path for inventory collection.
// URL: POST /api/v2/computer-inventory-collection-settings/custom-path
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-computer-inventory-collection-settings-custom-path
func (s *Service) CreateCustomPathV2(ctx context.Context, request *CustomPathRequest) (*CustomPathResponse, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := fmt.Sprintf("%s/custom-path", EndpointComputerInventoryCollectionSettingsV2)

	var result CustomPathResponse

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

// DeleteCustomPathByIDV2 deletes a custom application path by ID.
// URL: DELETE /api/v2/computer-inventory-collection-settings/custom-path/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-computer-inventory-collection-settings-custom-path-id
func (s *Service) DeleteCustomPathByIDV2(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("id is required")
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
