package device_communication_settings

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the device communication settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-communication-settings
	DeviceCommunicationSettings struct {
		client client.Client
	}
)

func NewDeviceCommunicationSettings(client client.Client) *DeviceCommunicationSettings {
	return &DeviceCommunicationSettings{client: client}
}

// GetV1 retrieves the current device communication settings.
// URL: GET /api/v1/device-communication-settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-device-communication-settings
func (s *DeviceCommunicationSettings) GetV1(ctx context.Context) (*ResourceDeviceCommunicationSettings, *resty.Response, error) {
	var result ResourceDeviceCommunicationSettings

	endpoint := constants.EndpointJamfProDeviceCommunicationSettingsV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).Get(endpoint)

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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Put(endpoint)

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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetQueryParams(rsqlQuery).
		GetPaginated(endpoint, mergePage)

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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(req).
		SetResult(&result).
		Post(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
