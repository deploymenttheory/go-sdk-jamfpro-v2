package self_service_settings

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// SelfServiceSettingsServiceInterface defines the interface for self-service settings operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-settings
	SelfServiceSettingsServiceInterface interface {
		// Get retrieves self-service settings (Get Self Service Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-settings
		Get(ctx context.Context) (*ResourceSelfServiceSettings, *interfaces.Response, error)

		// Update updates self-service settings (Update Self Service Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-self-service-settings
		Update(ctx context.Context, request *ResourceSelfServiceSettings) (*ResourceSelfServiceSettings, *interfaces.Response, error)
	}

	// Service handles communication with the self-service settings methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-settings
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ SelfServiceSettingsServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Self Service Settings Operations
// -----------------------------------------------------------------------------

// Get retrieves self-service settings.
// URL: GET /api/v1/self-service/settings
func (s *Service) Get(ctx context.Context) (*ResourceSelfServiceSettings, *interfaces.Response, error) {
	var result ResourceSelfServiceSettings

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, EndpointSelfServiceSettingsV1, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// Update updates self-service settings.
// URL: PUT /api/v1/self-service/settings
func (s *Service) Update(ctx context.Context, request *ResourceSelfServiceSettings) (*ResourceSelfServiceSettings, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceSelfServiceSettings

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, EndpointSelfServiceSettingsV1, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
