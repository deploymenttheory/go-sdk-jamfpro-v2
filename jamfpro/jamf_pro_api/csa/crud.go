package csa

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// CSAServiceInterface defines the interface for CSA Token operations.
	// Uses v1 API. Manages CSA token exchange for cloud-hosted services authentication.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-csa-token
	CSAServiceInterface interface {
		// GetTokenExchangeDetailsV1 returns details regarding the CSA token exchange.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-csa-token
		GetTokenExchangeDetailsV1(ctx context.Context) (*ResourceTokenExchangeDetails, *resty.Response, error)

		// GetTenantIDV1 returns the CSA tenant ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-csa-tenant-id
		GetTenantIDV1(ctx context.Context) (*ResourceTenantID, *resty.Response, error)

		// DeleteTokenExchangeV1 deletes the CSA token exchange, disabling Jamf Pro's ability to authenticate with cloud-hosted services.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-csa-token
		DeleteTokenExchangeV1(ctx context.Context) (*resty.Response, error)
	}

	// Service handles communication with the CSA Token-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-csa-token
	Csa struct {
		client transport.HTTPClient
	}
)

var _ CSAServiceInterface = (*Csa)(nil)

func NewCsa(client transport.HTTPClient) *Csa {
	return &Csa{client: client}
}

// GetTokenExchangeDetailsV1 returns details regarding the CSA token exchange.
// URL: GET /api/v1/csa/token
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-csa-token
func (s *Csa) GetTokenExchangeDetailsV1(ctx context.Context) (*ResourceTokenExchangeDetails, *resty.Response, error) {
	endpoint := constants.EndpointJamfProCSAV1

	var result ResourceTokenExchangeDetails

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
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

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
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

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
