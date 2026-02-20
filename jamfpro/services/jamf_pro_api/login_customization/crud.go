package login_customization

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// LoginCustomizationServiceInterface defines the interface for login customization operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-login-customization
	LoginCustomizationServiceInterface interface {
		// GetLoginCustomizationV1 returns the current login customization settings (Get Login Customization).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-login-customization
		GetLoginCustomizationV1(ctx context.Context) (*ResourceLoginCustomizationV1, *interfaces.Response, error)

		// UpdateLoginCustomizationV1 updates login customization settings (Update Login Customization / PUT).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-login-customization
		UpdateLoginCustomizationV1(ctx context.Context, request *ResourceLoginCustomizationV1) (*ResourceLoginCustomizationV1, *interfaces.Response, error)
	}

	// Service handles communication with the login customization-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-login-customization
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ LoginCustomizationServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Login Customization Operations
// -----------------------------------------------------------------------------

// GetLoginCustomizationV1 returns the current login customization settings.
// URL: GET /api/v1/login-customization
// https://developer.jamf.com/jamf-pro/reference/get_v1-login-customization
func (s *Service) GetLoginCustomizationV1(ctx context.Context) (*ResourceLoginCustomizationV1, *interfaces.Response, error) {
	var result ResourceLoginCustomizationV1

	endpoint := EndpointLoginCustomizationV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateLoginCustomizationV1 updates login customization settings.
// URL: PUT /api/v1/login-customization
// https://developer.jamf.com/jamf-pro/reference/put_v1-login-customization
func (s *Service) UpdateLoginCustomizationV1(ctx context.Context, request *ResourceLoginCustomizationV1) (*ResourceLoginCustomizationV1, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceLoginCustomizationV1

	endpoint := EndpointLoginCustomizationV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
