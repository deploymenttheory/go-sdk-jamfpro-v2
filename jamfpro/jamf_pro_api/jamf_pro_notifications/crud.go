package jamf_pro_notifications

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the Jamf Pro notifications-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-notifications
	JamfProNotifications struct {
		client transport.HTTPClient
	}
)

// NewService creates a new Jamf Pro notifications service.
func NewJamfProNotifications(client transport.HTTPClient) *JamfProNotifications {
	return &JamfProNotifications{client: client}
}

// GetForUserAndSiteV1 retrieves all notifications for the current user and site.
// URL: GET /api/v1/notifications
// https://developer.jamf.com/jamf-pro/reference/get_v1-notifications
func (s *JamfProNotifications) GetForUserAndSiteV1(ctx context.Context) ([]ResourceNotification, *resty.Response, error) {
	endpoint := constants.EndpointJamfProNotificationsV1

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	var result []ResourceNotification
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get notifications: %w", err)
	}

	return result, resp, nil
}

// DeleteByTypeAndIDV1 deletes a notification by type and ID.
// URL: DELETE /api/v1/notifications/{type}/{id}
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-notifications-type-id
func (s *JamfProNotifications) DeleteByTypeAndIDV1(ctx context.Context, notificationType, id string) (*resty.Response, error) {
	if notificationType == "" {
		return nil, fmt.Errorf("notification type is required")
	}
	if id == "" {
		return nil, fmt.Errorf("notification id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/%s", constants.EndpointJamfProNotificationsV1, notificationType, id)

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
