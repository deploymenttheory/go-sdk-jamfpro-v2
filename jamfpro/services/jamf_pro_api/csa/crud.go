package csa

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
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
		GetTokenExchangeDetailsV1(ctx context.Context) (*ResourceTokenExchangeDetails, *interfaces.Response, error)

		// GetTenantIDV1 returns the CSA tenant ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-csa-token-tenant-id
		GetTenantIDV1(ctx context.Context) (*ResourceTenantID, *interfaces.Response, error)

		// DeleteTokenExchangeV1 deletes the CSA token exchange, disabling Jamf Pro's ability to authenticate with cloud-hosted services.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-csa-token
		DeleteTokenExchangeV1(ctx context.Context) (*interfaces.Response, error)
	}

	// Service handles communication with the CSA Token-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-csa-token
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ CSAServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetTokenExchangeDetailsV1 returns details regarding the CSA token exchange.
// URL: GET /api/v1/csa/token
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-csa-token
func (s *Service) GetTokenExchangeDetailsV1(ctx context.Context) (*ResourceTokenExchangeDetails, *interfaces.Response, error) {
	var result ResourceTokenExchangeDetails

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, EndpointCSAV1, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetTenantIDV1 returns the CSA tenant ID.
// URL: GET /api/v1/csa/token/tenant-id
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-csa-token-tenant-id
func (s *Service) GetTenantIDV1(ctx context.Context) (*ResourceTenantID, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/tenant-id", EndpointCSAV1)

	var result ResourceTenantID

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

// DeleteTokenExchangeV1 deletes the CSA token exchange, disabling Jamf Pro's ability to authenticate with cloud-hosted services.
// URL: DELETE /api/v1/csa/token
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v1-csa-token
func (s *Service) DeleteTokenExchangeV1(ctx context.Context) (*interfaces.Response, error) {
	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Delete(ctx, EndpointCSAV1, nil, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
