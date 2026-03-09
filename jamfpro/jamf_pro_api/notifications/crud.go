package notifications

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// NotificationsServiceInterface defines the interface for notifications (read-only).
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-notifications
	NotificationsServiceInterface interface {
		// ListV1 returns all notifications for the current user and site (Get Notifications).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-notifications
		ListV1(ctx context.Context) ([]ResourceNotification, *resty.Response, error)
	}

	// Service handles communication with the notifications methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-notifications
	Notifications struct {
		client interfaces.HTTPClient
	}
)

var _ NotificationsServiceInterface = (*Notifications)(nil)

func NewNotifications(client interfaces.HTTPClient) *Notifications {
	return &Notifications{client: client}
}

// ListV1 returns all notifications for the current user and site.
// URL: GET /api/v1/notifications
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-notifications
func (s *Notifications) ListV1(ctx context.Context) ([]ResourceNotification, *resty.Response, error) {
	var result []ResourceNotification

	endpoint := EndpointNotificationsV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	if result == nil {
		result = []ResourceNotification{}
	}

	return result, resp, nil
}
