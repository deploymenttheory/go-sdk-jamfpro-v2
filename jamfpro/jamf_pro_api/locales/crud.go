package locales

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// LocalesServiceInterface defines the interface for locale operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-locales
	LocalesServiceInterface interface {
		// ListV1 returns all available locales (Get Locales).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-locales
		ListV1(ctx context.Context) ([]ResourceLocale, *resty.Response, error)
	}

	// Service handles communication with the locales-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-locales
	Locales struct {
		client transport.HTTPClient
	}
)

var _ LocalesServiceInterface = (*Locales)(nil)

func NewLocales(client transport.HTTPClient) *Locales {
	return &Locales{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Locales Operations
// -----------------------------------------------------------------------------

// ListV1 returns all available locales.
// URL: GET /api/v1/locales
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-locales
func (s *Locales) ListV1(ctx context.Context) ([]ResourceLocale, *resty.Response, error) {
	var result []ResourceLocale

	endpoint := EndpointLocalesV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
