package app_store_country_codes

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// AppStoreCountryCodesServiceInterface defines the interface for App Store country code operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/app-store-country-codes
	AppStoreCountryCodesServiceInterface interface {
		// ListV1 returns all App Store country codes (Get App Store Country Codes).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-app-store-country-codes
		ListV1(ctx context.Context) (*ListResponse, *interfaces.Response, error)
	}

	// Service handles communication with the App Store country codes-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/app-store-country-codes
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ AppStoreCountryCodesServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// ListV1 returns all App Store country codes.
// URL: GET /api/v1/app-store-country-codes
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-app-store-country-codes
func (s *Service) ListV1(ctx context.Context) (*ListResponse, *interfaces.Response, error) {
	var result ListResponse

	endpoint := EndpointAppStoreCountryCodesV1

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
