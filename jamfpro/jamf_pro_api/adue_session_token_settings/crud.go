package adue_session_token_settings

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// AdueSessionTokenSettingsServiceInterface defines the interface for ADUE session token settings operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-adue-session-token-settings
	AdueSessionTokenSettingsServiceInterface interface {
		// GetV1 retrieves ADUE session token settings (Get ADUE Session Token Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-adue-session-token-settings
		GetV1(ctx context.Context) (*ResourceADUETokenSettings, *resty.Response, error)

		// UpdateV1 updates ADUE session token settings (Update ADUE Session Token Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-adue-session-token-settings
		UpdateV1(ctx context.Context, request *ResourceADUETokenSettings) (*ResourceADUETokenSettings, *resty.Response, error)
	}

	// Service handles communication with the ADUE session token settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-adue-session-token-settings
	AdueSessionTokenSettings struct {
		client interfaces.HTTPClient
	}
)

var _ AdueSessionTokenSettingsServiceInterface = (*AdueSessionTokenSettings)(nil)

func NewAdueSessionTokenSettings(client interfaces.HTTPClient) *AdueSessionTokenSettings {
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

	endpoint := EndpointADUESessionTokenSettingsV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
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

	endpoint := EndpointADUESessionTokenSettingsV1
	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
