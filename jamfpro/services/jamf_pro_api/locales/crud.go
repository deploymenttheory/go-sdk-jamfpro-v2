package locales

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// LocalesServiceInterface defines the interface for locale operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-locales
	LocalesServiceInterface interface {
		// ListV1 returns all available locales (Get Locales).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-locales
		ListV1(ctx context.Context) ([]ResourceLocale, *interfaces.Response, error)
	}

	// Service handles communication with the locales-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-locales
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ LocalesServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Locales Operations
// -----------------------------------------------------------------------------

// ListV1 returns all available locales.
// URL: GET /api/v1/locales
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-locales
func (s *Service) ListV1(ctx context.Context) ([]ResourceLocale, *interfaces.Response, error) {
	var result []ResourceLocale

	endpoint := EndpointLocalesV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
