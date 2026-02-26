package user

import (
	"context"
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// UserServiceInterface defines the interface for user operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_api-user
	UserServiceInterface interface {
		// Get returns the current authenticated user information.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_api-user
		Get(ctx context.Context) (*ResourceUser, *interfaces.Response, error)

		// ChangePassword changes the current user's password.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-user-change-password
		ChangePassword(ctx context.Context, request *RequestChangePassword) (*interfaces.Response, error)

		// UpdateSession updates the current user's session (change site).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_api-user-updateSession
		UpdateSession(ctx context.Context, request *RequestUpdateSession) (*interfaces.Response, error)
	}

	// Service handles communication with the user-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_api-user
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ UserServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// Get returns the current authenticated user information.
// URL: GET /api/user
func (s *Service) Get(ctx context.Context) (*ResourceUser, *interfaces.Response, error) {
	endpoint := EndpointUser

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
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
func (s *Service) ChangePassword(ctx context.Context, request *RequestChangePassword) (*interfaces.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}
	if strings.TrimSpace(request.CurrentPassword) == "" {
		return nil, fmt.Errorf("currentPassword is required and cannot be empty")
	}
	if strings.TrimSpace(request.NewPassword) == "" {
		return nil, fmt.Errorf("newPassword is required and cannot be empty")
	}

	endpoint := EndpointChangePasswordV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
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
func (s *Service) UpdateSession(ctx context.Context, request *RequestUpdateSession) (*interfaces.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := EndpointUpdateSession

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to update session: %w", err)
	}

	return resp, nil
}
