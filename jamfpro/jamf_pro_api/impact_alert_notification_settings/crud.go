package impact_alert_notification_settings

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the impact alert notification settings-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-impact-alert-notification-settings
	ImpactAlertNotificationSettings struct {
		client client.Client
	}
)

func NewImpactAlertNotificationSettings(client client.Client) *ImpactAlertNotificationSettings {
	return &ImpactAlertNotificationSettings{client: client}
}

// GetV1 retrieves the impact alert notification settings.
// URL: GET /api/v1/impact-alert-notification-settings
// https://developer.jamf.com/jamf-pro/reference/get_v1-impact-alert-notification-settings
func (s *ImpactAlertNotificationSettings) GetV1(ctx context.Context) (*ResourceImpactAlertNotificationSettings, *resty.Response, error) {
	var result ResourceImpactAlertNotificationSettings

	endpoint := constants.EndpointJamfProImpactAlertNotificationSettingsV1
	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateV1 updates the impact alert notification settings via PUT.
// URL: PUT /api/v1/impact-alert-notification-settings
// https://developer.jamf.com/jamf-pro/reference/put_v1-impact-alert-notification-settings
func (s *ImpactAlertNotificationSettings) UpdateV1(ctx context.Context, request *ResourceImpactAlertNotificationSettings) (*resty.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := constants.EndpointJamfProImpactAlertNotificationSettingsV1
	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
