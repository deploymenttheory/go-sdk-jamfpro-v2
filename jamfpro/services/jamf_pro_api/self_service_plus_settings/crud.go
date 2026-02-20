package self_service_plus_settings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// SelfServicePlusSettingsServiceInterface defines the interface for self-service plus settings operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getselfserviceplussettings
	SelfServicePlusSettingsServiceInterface interface {
		// GetFeatureToggleEnabledV1 returns whether the Self Service Plus feature toggle is enabled (Determines if Self Service Plus feature toggle is enabled).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getselfserviceplusfeaturetoggleenabled
		GetFeatureToggleEnabledV1(ctx context.Context) (bool, *interfaces.Response, error)

		// GetV1 retrieves the current Self Service Plus settings (Get Self Service Plus settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getselfserviceplussettings
		GetV1(ctx context.Context) (*ResourceSelfServicePlusSettings, *interfaces.Response, error)

		// UpdateV1 updates the Self Service Plus settings (Save Self Service Plus settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/putselfserviceplussettings
		UpdateV1(ctx context.Context, request *ResourceSelfServicePlusSettings) (*ResourceSelfServicePlusSettings, *interfaces.Response, error)
	}

	// Service handles communication with the self-service plus settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getselfserviceplussettings
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ SelfServicePlusSettingsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Self Service Plus Settings Operations
// -----------------------------------------------------------------------------

// GetFeatureToggleEnabledV1 returns whether the Self Service Plus feature toggle is enabled.
// URL: GET /api/v1/self-service-plus/feature-toggle/enabled
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getselfserviceplusfeaturetoggleenabled
func (s *Service) GetFeatureToggleEnabledV1(ctx context.Context) (bool, *interfaces.Response, error) {
	var result ResourceFeatureToggleEnabled
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Get(ctx, EndpointSelfServicePlusFeatureToggleEnabledV1, nil, headers, &result)
	if err != nil {
		return false, resp, err
	}
	return result.Enabled, resp, nil
}

// GetV1 retrieves the current Self Service Plus settings.
// URL: GET /api/v1/self-service-plus/settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/getselfserviceplussettings
func (s *Service) GetV1(ctx context.Context) (*ResourceSelfServicePlusSettings, *interfaces.Response, error) {
	var result ResourceSelfServicePlusSettings
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Get(ctx, EndpointSelfServicePlusSettingsV1, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// UpdateV1 updates the Self Service Plus settings.
// URL: PUT /api/v1/self-service-plus/settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/putselfserviceplussettings
func (s *Service) UpdateV1(ctx context.Context, request *ResourceSelfServicePlusSettings) (*ResourceSelfServicePlusSettings, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Put(ctx, EndpointSelfServicePlusSettingsV1, request, headers, nil)
	if err != nil {
		return nil, resp, err
	}
	if resp != nil && resp.StatusCode == http.StatusNoContent {
		return nil, resp, nil
	}
	return nil, resp, nil
}
