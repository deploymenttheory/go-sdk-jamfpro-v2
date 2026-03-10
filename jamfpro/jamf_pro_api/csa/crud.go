package csa

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the CSA Token-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-csa-token
	Csa struct {
		client client.Client
	}
)

func NewCsa(client client.Client) *Csa {
	return &Csa{client: client}
}

// GetTokenExchangeDetailsV1 returns details regarding the CSA token exchange.
// URL: GET /api/v1/csa/token
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-csa-token
func (s *Csa) GetTokenExchangeDetailsV1(ctx context.Context) (*ResourceTokenExchangeDetails, *resty.Response, error) {
	endpoint := constants.EndpointJamfProCSAV1

	var result ResourceTokenExchangeDetails

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetTenantIDV1 returns the CSA tenant ID.
// URL: GET /api/v1/csa/token/tenant-id
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-csa-tenant-id
func (s *Csa) GetTenantIDV1(ctx context.Context) (*ResourceTenantID, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/tenant-id", constants.EndpointJamfProCSAV1)

	var result ResourceTenantID

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// DeleteTokenExchangeV1 deletes the CSA token exchange, disabling Jamf Pro's ability to authenticate with cloud-hosted services.
// URL: DELETE /api/v1/csa/token
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-csa-token
func (s *Csa) DeleteTokenExchangeV1(ctx context.Context) (*resty.Response, error) {
	endpoint := constants.EndpointJamfProCSAV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
