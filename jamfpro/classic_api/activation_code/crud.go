package activation_code

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the activation code-related Classic API methods.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/activationcode
	ActivationCode struct {
		client client.Client
	}
)

// NewService returns a new activation code Service backed by the provided HTTP client.
func NewActivationCode(client client.Client) *ActivationCode {
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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetResult(&result).
		Get(endpoint)

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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		SetHeader("Content-Type", constants.ApplicationXML).
		SetBody(request).
		Put(endpoint)

	if err != nil {
		return resp, err
	}

	return resp, nil
}
