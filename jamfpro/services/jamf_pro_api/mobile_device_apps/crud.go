package mobile_device_apps

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// MobileDeviceAppsServiceInterface defines the interface for mobile device apps operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-apps-reinstall-app-config
	MobileDeviceAppsServiceInterface interface {
		// ReinstallAppConfigV1 redeploys the managed app configuration for a specific app on a specific device
		// using the $APP_CONFIG_REINSTALL_CODE generated during deployment.
		// This endpoint does not require authorization, only the re-install code.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-apps-reinstall-app-config
		ReinstallAppConfigV1(ctx context.Context, request *RequestReinstallAppConfig) (*interfaces.Response, error)
	}

	// Service handles communication with the mobile device apps-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-apps-reinstall-app-config
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ MobileDeviceAppsServiceInterface = (*Service)(nil)

// NewService returns a new mobile device apps Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Mobile Device Apps Operations
// -----------------------------------------------------------------------------

// ReinstallAppConfigV1 redeploys the managed app configuration for a specific app on a specific device.
// URL: POST /api/v1/mobile-device-apps/reinstall-app-config
// Response: 204 No Content on success.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-mobile-device-apps-reinstall-app-config
func (s *Service) ReinstallAppConfigV1(ctx context.Context, request *RequestReinstallAppConfig) (*interfaces.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}
	if request.ReinstallCode == "" {
		return nil, fmt.Errorf("reinstallCode is required")
	}

	endpoint := fmt.Sprintf("%s/reinstall-app-config", EndpointMobileDeviceAppsV1)

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to reinstall app config: %w", err)
	}

	return resp, nil
}
