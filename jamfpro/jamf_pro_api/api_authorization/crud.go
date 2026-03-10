package api_authorization

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the API authorization-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-auth
	ApiAuthorization struct {
		client client.Client
	}
)

func NewApiAuthorization(client client.Client) *ApiAuthorization {
	return &ApiAuthorization{client: client}
}

// GetV1 returns the current authorization details associated with the API token.
// URL: GET /api/v1/auth
// https://developer.jamf.com/jamf-pro/reference/get_v1-auth
func (s *ApiAuthorization) GetV1(ctx context.Context) (*ResourceAuthV1, *resty.Response, error) {
	var result ResourceAuthV1

	endpoint := constants.EndpointJamfProAuthV1

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
