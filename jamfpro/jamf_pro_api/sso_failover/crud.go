package sso_failover

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the SSO failover-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-sso-failover
	SsoFailover struct {
		client client.Client
	}
)

func NewSsoFailover(client client.Client) *SsoFailover {
	return &SsoFailover{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - SSO Failover Operations
// -----------------------------------------------------------------------------

// GetV1 retrieves the current SSO failover settings.
// URL: GET /api/v1/sso/failover
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-sso-failover
func (s *SsoFailover) GetV1(ctx context.Context) (*FailoverSettings, *resty.Response, error) {
	var result FailoverSettings

	endpoint := constants.EndpointJamfProSSOFailoverV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// RegenerateV1 generates a new SSO failover URL.
// URL: POST /api/v1/sso/failover/generate
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-sso-failover-generate
func (s *SsoFailover) RegenerateV1(ctx context.Context) (*FailoverSettings, *resty.Response, error) {
	var result FailoverSettings

	endpoint := constants.EndpointJamfProSSOFailoverGenerateV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
