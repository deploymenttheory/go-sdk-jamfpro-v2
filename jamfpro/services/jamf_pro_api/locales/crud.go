package locales

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

type (
	// LocalesServiceInterface defines the interface for locale operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-locales
	LocalesServiceInterface interface {
		// ListLocalesV1 returns all available locales (Get Locales).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-locales
		ListLocalesV1(ctx context.Context) ([]ResourceLocale, *interfaces.Response, error)
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

// ListLocalesV1 returns all available locales.
// URL: GET /api/v1/locales
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-locales
func (s *Service) ListLocalesV1(ctx context.Context) ([]ResourceLocale, *interfaces.Response, error) {
	var result []ResourceLocale
	resp, err := s.client.Get(ctx, EndpointLocalesV1, nil, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}
