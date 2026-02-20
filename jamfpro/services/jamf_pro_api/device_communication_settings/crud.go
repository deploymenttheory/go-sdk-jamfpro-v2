package device_communication_settings

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// DeviceCommunicationSettingsServiceInterface defines the interface for device communication settings operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-communication-settings
	DeviceCommunicationSettingsServiceInterface interface {
		// GetV1 retrieves the current device communication settings (Get Device Communication Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-communication-settings
		GetV1(ctx context.Context) (*ResourceDeviceCommunicationSettings, *interfaces.Response, error)

		// UpdateV1 updates the device communication settings (Update Device Communication Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-device-communication-settings
		UpdateV1(ctx context.Context, request *ResourceDeviceCommunicationSettings) (*ResourceDeviceCommunicationSettings, *interfaces.Response, error)
	}

	// Service handles communication with the device communication settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-communication-settings
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ DeviceCommunicationSettingsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetV1 retrieves the current device communication settings.
// URL: GET /api/v1/device-communication-settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-communication-settings
func (s *Service) GetV1(ctx context.Context) (*ResourceDeviceCommunicationSettings, *interfaces.Response, error) {
	var result ResourceDeviceCommunicationSettings
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Get(ctx, EndpointDeviceCommunicationSettingsV1, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// UpdateV1 updates the device communication settings.
// URL: PUT /api/v1/device-communication-settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-device-communication-settings
func (s *Service) UpdateV1(ctx context.Context, request *ResourceDeviceCommunicationSettings) (*ResourceDeviceCommunicationSettings, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	var result ResourceDeviceCommunicationSettings
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Put(ctx, EndpointDeviceCommunicationSettingsV1, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
