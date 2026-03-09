package declarative_device_management

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
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
		ForceSyncV1(ctx context.Context, clientManagementID string) (*resty.Response, error)

		// GetStatusItemsV1 retrieves the latest status report items for a specific device by its client management ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-ddm-clientmanagementid-status-items
		GetStatusItemsV1(ctx context.Context, clientManagementID string) (*ResourceStatusItems, *resty.Response, error)

		// GetStatusItemByKeyV1 retrieves a specific status report item by its client management ID and status item key.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-ddm-clientmanagementid-status-items-key
		GetStatusItemByKeyV1(ctx context.Context, clientManagementID string, key string) (*StatusItem, *resty.Response, error)
	}

	// Service handles communication with the Declarative Device Management-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-ddm-clientmanagementid-sync
	DeclarativeDeviceManagement struct {
		client transport.HTTPClient
	}
)

var _ DeclarativeDeviceManagementServiceInterface = (*DeclarativeDeviceManagement)(nil)

func NewDeclarativeDeviceManagement(client transport.HTTPClient) *DeclarativeDeviceManagement {
	return &DeclarativeDeviceManagement{client: client}
}

// ForceSyncV1 initiates a DDM synchronization for a specific client management ID.
// URL: POST /api/v1/ddm/{clientManagementId}/sync
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-ddm-clientmanagementid-sync
func (s *DeclarativeDeviceManagement) ForceSyncV1(ctx context.Context, clientManagementID string) (*resty.Response, error) {
	if clientManagementID == "" {
		return nil, fmt.Errorf("clientManagementID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/sync", constants.EndpointJamfProDeclarativeDeviceManagementV1, clientManagementID)

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
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
func (s *DeclarativeDeviceManagement) GetStatusItemsV1(ctx context.Context, clientManagementID string) (*ResourceStatusItems, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}

	endpoint := fmt.Sprintf("%s/%s/status-items", constants.EndpointJamfProDeclarativeDeviceManagementV1, clientManagementID)

	var result ResourceStatusItems

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *DeclarativeDeviceManagement) GetStatusItemByKeyV1(ctx context.Context, clientManagementID string, key string) (*StatusItem, *resty.Response, error) {
	if clientManagementID == "" {
		return nil, nil, fmt.Errorf("clientManagementID is required")
	}
	if key == "" {
		return nil, nil, fmt.Errorf("key is required")
	}

	endpoint := fmt.Sprintf("%s/%s/status-items/%s", constants.EndpointJamfProDeclarativeDeviceManagementV1, clientManagementID, key)

	var result StatusItem

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
