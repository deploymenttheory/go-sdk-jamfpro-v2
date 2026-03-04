package jamf_pro_notifications

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

// ServiceInterface defines the interface for Jamf Pro notifications operations.
//
// Jamf Pro notifications provide alerts and messages for the currently authenticated
// user within the context of their assigned site.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-notifications
type ServiceInterface interface {
	// GetForUserAndSiteV1 retrieves all notifications for the current user and site.
	//
	// Returns a list of notifications filtered by the authenticated user's permissions
	// and site assignment. Notifications include system alerts, updates, and messages
	// relevant to the user's role.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-notifications
	GetForUserAndSiteV1(ctx context.Context) ([]ResourceNotification, *resty.Response, error)

	// DeleteByTypeAndIDV1 deletes a notification by type and ID.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-notifications-type-id
	DeleteByTypeAndIDV1(ctx context.Context, notificationType, id string) (*resty.Response, error)
}

type (
	// Service handles communication with the Jamf Pro notifications-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-notifications
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ServiceInterface = (*Service)(nil)

// NewService creates a new Jamf Pro notifications service.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetForUserAndSiteV1 retrieves all notifications for the current user and site.
// URL: GET /api/v1/notifications
// https://developer.jamf.com/jamf-pro/reference/get_v1-notifications
func (s *Service) GetForUserAndSiteV1(ctx context.Context) ([]ResourceNotification, *resty.Response, error) {
	endpoint := EndpointNotificationsV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
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
func (s *Service) DeleteByTypeAndIDV1(ctx context.Context, notificationType, id string) (*resty.Response, error) {
	if notificationType == "" {
		return nil, fmt.Errorf("notification type is required")
	}
	if id == "" {
		return nil, fmt.Errorf("notification id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/%s", EndpointNotificationsV1, notificationType, id)

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
