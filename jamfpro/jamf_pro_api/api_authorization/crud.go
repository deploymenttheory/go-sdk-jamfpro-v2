package api_authorization

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// ApiAuthorizationServiceInterface defines the interface for API authorization operations (read-only).
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-auth
	ApiAuthorizationServiceInterface interface {
		// GetV1 returns the current authorization details associated with the API token.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-auth
		GetV1(ctx context.Context) (*ResourceAuthV1, *resty.Response, error)
	}

	// Service handles communication with the API authorization-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-auth
	ApiAuthorization struct {
		client transport.HTTPClient
	}
)

var _ ApiAuthorizationServiceInterface = (*ApiAuthorization)(nil)

func NewApiAuthorization(client transport.HTTPClient) *ApiAuthorization {
	return &ApiAuthorization{client: client}
}

// GetV1 returns the current authorization details associated with the API token.
// URL: GET /api/v1/auth
// https://developer.jamf.com/jamf-pro/reference/get_v1-auth
func (s *ApiAuthorization) GetV1(ctx context.Context) (*ResourceAuthV1, *resty.Response, error) {
	var result ResourceAuthV1

	endpoint := EndpointAuthV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
