package adue_session_token_settings

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the ADUE session token settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-adue-session-token-settings
	AdueSessionTokenSettings struct {
		client client.Client
	}
)

func NewAdueSessionTokenSettings(client client.Client) *AdueSessionTokenSettings {
	return &AdueSessionTokenSettings{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - ADUE Session Token Settings Operations
// -----------------------------------------------------------------------------

// GetV1 retrieves ADUE session token settings.
// URL: GET /api/v1/adue-session-token-settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-adue-session-token-settings
func (s *AdueSessionTokenSettings) GetV1(ctx context.Context) (*ResourceADUETokenSettings, *resty.Response, error) {
	var result ResourceADUETokenSettings

	endpoint := constants.EndpointJamfProADUESessionTokenSettingsV1

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateV1 updates ADUE session token settings.
// URL: PUT /api/v1/adue-session-token-settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-adue-session-token-settings
func (s *AdueSessionTokenSettings) UpdateV1(ctx context.Context, request *ResourceADUETokenSettings) (*ResourceADUETokenSettings, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceADUETokenSettings

	endpoint := constants.EndpointJamfProADUESessionTokenSettingsV1
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
