package user_sessions

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// UserSessions handles communication with the user sessions-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-user-sessions-active
	UserSessions struct {
		client client.Client
	}
)

// NewUserSessions returns a new UserSessions service.
func NewUserSessions(client client.Client) *UserSessions {
	return &UserSessions{client: client}
}

// GetActiveV1 returns detailed information about currently logged-in users.
// URL: GET /api/v1/user-sessions/active
// https://developer.jamf.com/jamf-pro/reference/get_v1-user-sessions-active
func (s *UserSessions) GetActiveV1(ctx context.Context) ([]ResourceActiveUserSession, *resty.Response, error) {
	endpoint := constants.EndpointJamfProUserSessionsActiveV1

	var result []ResourceActiveUserSession

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get active user sessions: %w", err)
	}

	return result, resp, nil
}

// GetCountV1 returns the number of currently logged-in users.
// URL: GET /api/v1/user-sessions/count
// https://developer.jamf.com/jamf-pro/reference/get_v1-user-sessions-count
func (s *UserSessions) GetCountV1(ctx context.Context) (*ResourceUserSessionCount, *resty.Response, error) {
	endpoint := constants.EndpointJamfProUserSessionsCountV1

	var result ResourceUserSessionCount

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get user session count: %w", err)
	}

	return &result, resp, nil
}

// GetActiveV2 returns detailed information about currently logged-in users.
// URL: GET /api/v2/user-sessions/active
// https://developer.jamf.com/jamf-pro/reference/get_v2-user-sessions-active
func (s *UserSessions) GetActiveV2(ctx context.Context) ([]ResourceActiveUserSessionV2, *resty.Response, error) {
	endpoint := constants.EndpointJamfProUserSessionsActiveV2

	var result []ResourceActiveUserSessionV2

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get active user sessions: %w", err)
	}

	return result, resp, nil
}

// GetCountV2 returns the number of currently logged-in users.
// URL: GET /api/v2/user-sessions/count
// https://developer.jamf.com/jamf-pro/reference/get_v2-user-sessions-count
func (s *UserSessions) GetCountV2(ctx context.Context) (*ResourceUserSessionCountV2, *resty.Response, error) {
	endpoint := constants.EndpointJamfProUserSessionsCountV2

	var result ResourceUserSessionCountV2

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get user session count: %w", err)
	}

	return &result, resp, nil
}
