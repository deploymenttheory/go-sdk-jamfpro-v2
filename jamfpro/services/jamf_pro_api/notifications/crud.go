package notifications

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// NotificationsServiceInterface defines the interface for notifications (read-only).
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-notifications
	NotificationsServiceInterface interface {
		// ListV1 returns all notifications for the current user and site (Get Notifications).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-notifications
		ListV1(ctx context.Context) ([]ResourceNotification, *interfaces.Response, error)
	}

	// Service handles communication with the notifications methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-notifications
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ NotificationsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// ListV1 returns all notifications for the current user and site.
// URL: GET /api/v1/notifications
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-notifications
func (s *Service) ListV1(ctx context.Context) ([]ResourceNotification, *interfaces.Response, error) {
	var result []ResourceNotification

	endpoint := EndpointNotificationsV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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
