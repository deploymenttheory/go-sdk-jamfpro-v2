package self_service_plus_settings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// SelfServicePlusSettingsServiceInterface defines the interface for self-service plus settings operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getselfserviceplussettings
	SelfServicePlusSettingsServiceInterface interface {
		// GetFeatureToggleEnabledV1 returns whether the Self Service Plus feature toggle is enabled (Determines if Self Service Plus feature toggle is enabled).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getselfserviceplusfeaturetoggleenabled
		GetFeatureToggleEnabledV1(ctx context.Context) (bool, *resty.Response, error)

		// GetV1 retrieves the current Self Service Plus settings (Get Self Service Plus settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getselfserviceplussettings
		GetV1(ctx context.Context) (*ResourceSelfServicePlusSettings, *resty.Response, error)

		// UpdateV1 updates the Self Service Plus settings (Save Self Service Plus settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/putselfserviceplussettings
		UpdateV1(ctx context.Context, request *ResourceSelfServicePlusSettings) (*ResourceSelfServicePlusSettings, *resty.Response, error)
	}

	// Service handles communication with the self-service plus settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getselfserviceplussettings
	SelfServicePlusSettings struct {
		client transport.HTTPClient
	}
)

var _ SelfServicePlusSettingsServiceInterface = (*SelfServicePlusSettings)(nil)

func NewSelfServicePlusSettings(client transport.HTTPClient) *SelfServicePlusSettings {
	return &SelfServicePlusSettings{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Self Service Plus Settings Operations
// -----------------------------------------------------------------------------

// GetFeatureToggleEnabledV1 returns whether the Self Service Plus feature toggle is enabled.
// URL: GET /api/v1/self-service-plus/feature-toggle/enabled
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getselfserviceplusfeaturetoggleenabled
func (s *SelfServicePlusSettings) GetFeatureToggleEnabledV1(ctx context.Context) (bool, *resty.Response, error) {
	var result ResourceFeatureToggleEnabled
	endpoint := constants.EndpointJamfProSelfServicePlusFeatureToggleEnabledV1

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return false, resp, err
	}
	return result.Enabled, resp, nil
}

// GetV1 retrieves the current Self Service Plus settings.
// URL: GET /api/v1/self-service-plus/settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getselfserviceplussettings
func (s *SelfServicePlusSettings) GetV1(ctx context.Context) (*ResourceSelfServicePlusSettings, *resty.Response, error) {
	var result ResourceSelfServicePlusSettings
	endpoint := constants.EndpointJamfProSelfServicePlusSettingsV1
	headers := map[string]string{"Accept": constants.ApplicationJSON}
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// UpdateV1 updates the Self Service Plus settings.
// URL: PUT /api/v1/self-service-plus/settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/putselfserviceplussettings
func (s *SelfServicePlusSettings) UpdateV1(ctx context.Context, request *ResourceSelfServicePlusSettings) (*ResourceSelfServicePlusSettings, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	endpoint := constants.EndpointJamfProSelfServicePlusSettingsV1

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, nil)
	if err != nil {
		return nil, resp, err
	}

	if resp != nil && resp.StatusCode() == http.StatusNoContent {
		return nil, resp, nil
	}

	return nil, resp, nil
}
