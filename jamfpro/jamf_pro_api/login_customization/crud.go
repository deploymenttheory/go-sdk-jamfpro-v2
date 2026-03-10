package login_customization

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the login customization-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-login-customization
	LoginCustomization struct {
		client client.Client
	}
)

func NewLoginCustomization(client client.Client) *LoginCustomization {
	return &LoginCustomization{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Login Customization Operations
// -----------------------------------------------------------------------------

// GetV1 returns the current login customization settings.
// URL: GET /api/v1/login-customization
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-login-customization
func (s *LoginCustomization) GetV1(ctx context.Context) (*ResourceLoginCustomizationV1, *resty.Response, error) {
	var result ResourceLoginCustomizationV1

	endpoint := constants.EndpointJamfProLoginCustomizationV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateV1 updates login customization settings.
// URL: PUT /api/v1/login-customization
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-login-customization
func (s *LoginCustomization) UpdateV1(ctx context.Context, request *ResourceLoginCustomizationV1) (*ResourceLoginCustomizationV1, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceLoginCustomizationV1

	endpoint := constants.EndpointJamfProLoginCustomizationV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Put(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
