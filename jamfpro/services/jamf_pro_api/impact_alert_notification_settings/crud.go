package impact_alert_notification_settings

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ImpactAlertNotificationSettingsServiceInterface defines the interface for impact alert notification settings operations.
	//
	// Manages impact alert notification settings for scopeable and deployable objects.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-impact-alert-notification-settings
	ImpactAlertNotificationSettingsServiceInterface interface {
		// GetV1 retrieves the impact alert notification settings.
		//
		// Returns current configuration for alert notifications and confirmation codes.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-impact-alert-notification-settings
		GetV1(ctx context.Context) (*ResourceImpactAlertNotificationSettings, *interfaces.Response, error)

		// UpdateV1 updates the impact alert notification settings via PUT.
		//
		// Returns 204 No Content on success.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-impact-alert-notification-settings
		UpdateV1(ctx context.Context, request *ResourceImpactAlertNotificationSettings) (*interfaces.Response, error)
	}

	// Service handles communication with the impact alert notification settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-impact-alert-notification-settings
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ImpactAlertNotificationSettingsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetV1 retrieves the impact alert notification settings.
// URL: GET /api/v1/impact-alert-notification-settings
// https://developer.jamf.com/jamf-pro/reference/get_v1-impact-alert-notification-settings
func (s *Service) GetV1(ctx context.Context) (*ResourceImpactAlertNotificationSettings, *interfaces.Response, error) {
	var result ResourceImpactAlertNotificationSettings

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, EndpointImpactAlertNotificationSettingsV1, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateV1 updates the impact alert notification settings via PUT.
// URL: PUT /api/v1/impact-alert-notification-settings
// https://developer.jamf.com/jamf-pro/reference/put_v1-impact-alert-notification-settings
func (s *Service) UpdateV1(ctx context.Context, request *ResourceImpactAlertNotificationSettings) (*interfaces.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, EndpointImpactAlertNotificationSettingsV1, request, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
