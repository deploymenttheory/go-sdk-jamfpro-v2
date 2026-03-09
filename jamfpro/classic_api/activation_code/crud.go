package activation_code

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// ActivationCodeServiceInterface defines the interface for Classic API activation code operations.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/activationcode
	ActivationCodeServiceInterface interface {
		// GetActivationCode retrieves the activation code information.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/findactivationcode
		GetActivationCode(ctx context.Context) (*ResourceActivationCode, *resty.Response, error)

		// UpdateActivationCode updates the activation code information.
		//
		// Classic API docs: https://developer.jamf.com/jamf-pro/reference/updateactivationcode
		UpdateActivationCode(ctx context.Context, request *RequestActivationCode) (*resty.Response, error)
	}

	// Service handles communication with the activation code-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/activationcode
	ActivationCode struct {
		client transport.HTTPClient
	}
)

var _ ActivationCodeServiceInterface = (*ActivationCode)(nil)

// NewService returns a new activation code Service backed by the provided HTTP client.
func NewActivationCode(client transport.HTTPClient) *ActivationCode {
	return &ActivationCode{client: client}
}

// -----------------------------------------------------------------------------
// Classic API - Activation Code Operations
// -----------------------------------------------------------------------------

// GetActivationCode retrieves the activation code information.
// URL: GET /JSSResource/activationcode
// https://developer.jamf.com/jamf-pro/reference/findactivationcode
func (s *ActivationCode) GetActivationCode(ctx context.Context) (*ResourceActivationCode, *resty.Response, error) {
	var result ResourceActivationCode

	endpoint := constants.EndpointClassicActivationCode

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
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
func (s *ActivationCode) UpdateActivationCode(ctx context.Context, request *RequestActivationCode) (*resty.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}

	endpoint := constants.EndpointClassicActivationCode

	headers := map[string]string{
		"Accept":       constants.ApplicationXML,
		"Content-Type": constants.ApplicationXML,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
