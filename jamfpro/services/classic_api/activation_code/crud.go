package activation_code

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ActivationCodeServiceInterface defines the interface for Classic API activation code operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/activationcode
	ActivationCodeServiceInterface interface {
		// GetActivationCode retrieves the activation code information.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findactivationcode
		GetActivationCode(ctx context.Context) (*ResourceActivationCode, *interfaces.Response, error)

		// UpdateActivationCode updates the activation code information.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateactivationcode
		UpdateActivationCode(ctx context.Context, request *RequestActivationCode) (*interfaces.Response, error)
	}

	// Service handles communication with the activation code-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/activationcode
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ActivationCodeServiceInterface = (*Service)(nil)

// NewService returns a new activation code Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Activation Code Operations
// -----------------------------------------------------------------------------

// GetActivationCode retrieves the activation code information.
// URL: GET /JSSResource/activationcode
// https://developer.jamf.com/jamf-pro/reference/findactivationcode
func (s *Service) GetActivationCode(ctx context.Context) (*ResourceActivationCode, *interfaces.Response, error) {
	var result ResourceActivationCode

	endpoint := EndpointClassicActivationCode

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateActivationCode updates the activation code information.
// URL: PUT /JSSResource/activationcode
// https://developer.jamf.com/jamf-pro/reference/updateactivationcode
func (s *Service) UpdateActivationCode(ctx context.Context, request *RequestActivationCode) (*interfaces.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := EndpointClassicActivationCode

	headers := map[string]string{
		"Accept":       mime.ApplicationXML,
		"Content-Type": mime.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
