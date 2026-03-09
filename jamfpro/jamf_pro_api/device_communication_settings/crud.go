package device_communication_settings

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// DeviceCommunicationSettingsServiceInterface defines the interface for device communication settings operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-communication-settings
	DeviceCommunicationSettingsServiceInterface interface {
		// GetV1 retrieves the current device communication settings (Get Device Communication Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-communication-settings
		GetV1(ctx context.Context) (*ResourceDeviceCommunicationSettings, *resty.Response, error)

		// UpdateV1 updates the device communication settings (Update Device Communication Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-device-communication-settings
		UpdateV1(ctx context.Context, request *ResourceDeviceCommunicationSettings) (*ResourceDeviceCommunicationSettings, *resty.Response, error)

		// GetHistoryV1 returns the history for the device communication settings.
		//
		// Query params (optional, pass via rsqlQuery): page, page-size, sort, filter (RSQL).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-communication-settings-history
		GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error)

		// AddHistoryNotesV1 adds a note to the device communication settings history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-device-communication-settings-history
		AddHistoryNotesV1(ctx context.Context, req *RequestAddHistoryNotes) (*ResponseAddHistoryNotes, *resty.Response, error)
	}

	// Service handles communication with the device communication settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-communication-settings
	DeviceCommunicationSettings struct {
		client transport.HTTPClient
	}
)

var _ DeviceCommunicationSettingsServiceInterface = (*DeviceCommunicationSettings)(nil)

func NewDeviceCommunicationSettings(client transport.HTTPClient) *DeviceCommunicationSettings {
	return &DeviceCommunicationSettings{client: client}
}

// GetV1 retrieves the current device communication settings.
// URL: GET /api/v1/device-communication-settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-communication-settings
func (s *DeviceCommunicationSettings) GetV1(ctx context.Context) (*ResourceDeviceCommunicationSettings, *resty.Response, error) {
	var result ResourceDeviceCommunicationSettings
	endpoint := constants.EndpointJamfProDeviceCommunicationSettingsV1
	headers := map[string]string{"Accept": constants.ApplicationJSON}
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// UpdateV1 updates the device communication settings.
// URL: PUT /api/v1/device-communication-settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-device-communication-settings
func (s *DeviceCommunicationSettings) UpdateV1(ctx context.Context, request *ResourceDeviceCommunicationSettings) (*ResourceDeviceCommunicationSettings, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	var result ResourceDeviceCommunicationSettings
	endpoint := constants.EndpointJamfProDeviceCommunicationSettingsV1
	headers := map[string]string{"Accept": constants.ApplicationJSON, "Content-Type": constants.ApplicationJSON}
	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// GetHistoryV1 returns the history for the device communication settings.
// URL: GET /api/v1/device-communication-settings/history
// Query params (optional): page, page-size, sort, filter (RSQL).
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-communication-settings-history
func (s *DeviceCommunicationSettings) GetHistoryV1(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error) {
	var result HistoryResponse

	mergePage := func(pageData []byte) error {
		var pageItems []HistoryItem
		if err := json.Unmarshal(pageData, &pageItems); err != nil {
			return fmt.Errorf("failed to unmarshal page: %w", err)
		}
		result.Results = append(result.Results, pageItems...)
		return nil
	}

	endpoint := constants.EndpointJamfProDeviceCommunicationSettingsHistoryV1
	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}
	resp, err := s.client.GetPaginated(ctx, endpoint, rsqlQuery, headers, mergePage)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get device communication settings history: %w", err)
	}
	result.TotalCount = len(result.Results)
	return &result, resp, nil
}

// AddHistoryNotesV1 adds a note to the device communication settings history.
// URL: POST /api/v1/device-communication-settings/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-device-communication-settings-history
func (s *DeviceCommunicationSettings) AddHistoryNotesV1(ctx context.Context, req *RequestAddHistoryNotes) (*ResponseAddHistoryNotes, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if req.Note == "" {
		return nil, nil, fmt.Errorf("note is required")
	}

	var result ResponseAddHistoryNotes
	endpoint := constants.EndpointJamfProDeviceCommunicationSettingsHistoryV1
	headers := map[string]string{"Accept": constants.ApplicationJSON, "Content-Type": constants.ApplicationJSON}
	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
