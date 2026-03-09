package mobile_device_apps

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the mobile device apps-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-apps-reinstall-app-config
	MobileDeviceApps struct {
		client transport.HTTPClient
	}
)

// NewService returns a new mobile device apps Service backed by the provided HTTP client.
func NewMobileDeviceApps(client transport.HTTPClient) *MobileDeviceApps {
	return &MobileDeviceApps{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Mobile Device Apps Operations
// -----------------------------------------------------------------------------

// ReinstallAppConfigV1 redeploys the managed app configuration for a specific app on a specific device.
// URL: POST /api/v1/mobile-device-apps/reinstall-app-config
// Response: 204 No Content on success.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-apps-reinstall-app-config
func (s *MobileDeviceApps) ReinstallAppConfigV1(ctx context.Context, request *RequestReinstallAppConfig) (*resty.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}
	if request.ReinstallCode == "" {
		return nil, fmt.Errorf("reinstallCode is required")
	}

	endpoint := fmt.Sprintf("%s/reinstall-app-config", constants.EndpointJamfProMobileDeviceAppsV1)

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to reinstall app config: %w", err)
	}

	return resp, nil
}
