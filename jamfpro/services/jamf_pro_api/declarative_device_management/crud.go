package declarative_device_management

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// DeclarativeDeviceManagementServiceInterface defines the interface for Declarative Device Management operations.
	// Uses v1 API. Manages DDM synchronization and status reporting for devices.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-ddm-clientmanagementid-sync
	DeclarativeDeviceManagementServiceInterface interface {
		// ForceSyncV1 initiates a DDM synchronization for a specific client management ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-ddm-clientmanagementid-sync
		ForceSyncV1(ctx context.Context, clientManagementID string) (*interfaces.Response, error)

		// GetStatusItemsV1 retrieves the latest status report items for a specific device by its client management ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-ddm-clientmanagementid-status-items
		GetStatusItemsV1(ctx context.Context, clientManagementID string) (*ResourceStatusItems, *interfaces.Response, error)

		// GetStatusItemByKeyV1 retrieves a specific status report item by its client management ID and status item key.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-ddm-clientmanagementid-status-items-key
		GetStatusItemByKeyV1(ctx context.Context, clientManagementID string, key string) (*StatusItem, *interfaces.Response, error)
	}

	// Service handles communication with the Declarative Device Management-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-ddm-clientmanagementid-sync
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ DeclarativeDeviceManagementServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// ForceSyncV1 initiates a DDM synchronization for a specific client management ID.
// URL: POST /api/v1/ddm/{clientManagementId}/sync
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-ddm-clientmanagementid-sync
func (s *Service) ForceSyncV1(ctx context.Context, clientManagementID string) (*interfaces.Response, error) {
	if clientManagementID == "" {
		return nil, fmt.Errorf("clientManagementID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/sync", EndpointDeclarativeDeviceManagementV1, clientManagementID)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// GetStatusItemsV1 retrieves the latest status report items for a specific device by its client management ID.
// URL: GET /api/v1/ddm/{clientManagementId}/status-items
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-ddm-clientmanagementid-status-items
func (s *Service) GetStatusItemsV1(ctx context.Context, clientManagementID string) (*ResourceStatusItems, *interfaces.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/status-items", EndpointDeclarativeDeviceManagementV1, clientManagementID)

	var result ResourceStatusItems

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

// GetStatusItemByKeyV1 retrieves a specific status report item by its client management ID and status item key.
// URL: GET /api/v1/ddm/{clientManagementId}/status-items/{key}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-ddm-clientmanagementid-status-items-key
func (s *Service) GetStatusItemByKeyV1(ctx context.Context, clientManagementID string, key string) (*StatusItem, *interfaces.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}
	if key == "" {
		return nil, nil, fmt.Errorf("key is required")
	}

	endpoint := fmt.Sprintf("%s/%s/status-items/%s", EndpointDeclarativeDeviceManagementV1, clientManagementID, key)

	var result StatusItem

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
