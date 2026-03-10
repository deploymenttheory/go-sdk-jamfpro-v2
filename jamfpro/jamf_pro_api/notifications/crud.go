package notifications

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the notifications methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-notifications
	Notifications struct {
		client client.Client
	}
)

func NewNotifications(client client.Client) *Notifications {
	return &Notifications{client: client}
}

// ListV1 returns all notifications for the current user and site.
// URL: GET /api/v1/notifications
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-notifications
func (s *Notifications) ListV1(ctx context.Context) ([]ResourceNotification, *resty.Response, error) {
	var result []ResourceNotification

	endpoint := constants.EndpointJamfProNotificationsV1

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
