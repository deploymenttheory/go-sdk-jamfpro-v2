package api_authorization

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ApiAuthorizationServiceInterface defines the interface for API authorization operations (read-only).
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-auth
	ApiAuthorizationServiceInterface interface {
		// GetV1 returns the current authorization details associated with the API token.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-auth
		GetV1(ctx context.Context) (*ResourceAuthV1, *interfaces.Response, error)
	}

	// Service handles communication with the API authorization-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-auth
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ApiAuthorizationServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetV1 returns the current authorization details associated with the API token.
// URL: GET /api/v1/auth
// https://developer.jamf.com/jamf-pro/reference/get_v1-auth
func (s *Service) GetV1(ctx context.Context) (*ResourceAuthV1, *interfaces.Response, error) {
	var result ResourceAuthV1

	endpoint := EndpointAuthV1

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
