package self_service_plus_settings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the self-service plus settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getselfserviceplussettings
	SelfServicePlusSettings struct {
		client client.Client
	}
)

func NewSelfServicePlusSettings(client client.Client) *SelfServicePlusSettings {
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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		Put(endpoint)
	if err != nil {
		return nil, resp, err
	}

	if resp != nil && resp.StatusCode() == http.StatusNoContent {
		return nil, resp, nil
	}

	return nil, resp, nil
}
