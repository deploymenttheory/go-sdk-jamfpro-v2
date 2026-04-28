package last_login

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// LastLogin handles communication with the last login-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-last-login
	LastLogin struct {
		client client.Client
	}
)

// NewLastLogin returns a new LastLogin service.
func NewLastLogin(client client.Client) *LastLogin {
	return &LastLogin{client: client}
}

// GetV1 returns the date of the last login event.
// URL: GET /api/v1/last-login
// https://developer.jamf.com/jamf-pro/reference/get_v1-last-login
func (s *LastLogin) GetV1(ctx context.Context) (*ResourceLastLogin, *resty.Response, error) {
	endpoint := constants.EndpointJamfProLastLoginV1

	var result ResourceLastLogin

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get last login: %w", err)
	}

	return &result, resp, nil
}
