package locales

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the locales-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-locales
	Locales struct {
		client client.Client
	}
)

func NewLocales(client client.Client) *Locales {
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

	endpoint := constants.EndpointJamfProLocalesV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
