package user

import (
	"context"
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the user-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_api-user
	User struct {
		client client.Client
	}
)

func NewUser(client client.Client) *User {
	return &User{client: client}
}

// Get returns the current authenticated user information.
// URL: GET /api/user
func (s *User) Get(ctx context.Context) (*ResourceUser, *resty.Response, error) {
	endpoint := constants.EndpointJamfProUser

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	var result ResourceUser
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get current user: %w", err)
	}

	return &result, resp, nil
}

// ChangePassword changes the current user's password.
// URL: POST /api/v1/user/change-password
// Response: 204 No Content
func (s *User) ChangePassword(ctx context.Context, request *RequestChangePassword) (*resty.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}
	if strings.TrimSpace(request.CurrentPassword) == "" {
		return nil, fmt.Errorf("currentPassword is required and cannot be empty")
	}
	if strings.TrimSpace(request.NewPassword) == "" {
		return nil, fmt.Errorf("newPassword is required and cannot be empty")
	}

	endpoint := constants.EndpointJamfProChangePasswordV1

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to change password: %w", err)
	}

	return resp, nil
}

// UpdateSession updates the current user's session (change site).
// URL: POST /api/user/updateSession
// Response: 204 No Content
func (s *User) UpdateSession(ctx context.Context, request *RequestUpdateSession) (*resty.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := constants.EndpointJamfProUpdateSession

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to update session: %w", err)
	}

	return resp, nil
}
