package sso_failover

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ServiceInterface defines the interface for SSO failover operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-sso-failover
	ServiceInterface interface {
		// GetV1 retrieves the current SSO failover settings (Get SSO Failover Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-sso-failover
		GetV1(ctx context.Context) (*FailoverSettings, *interfaces.Response, error)

		// RegenerateV1 generates a new SSO failover URL (Regenerate SSO Failover URL).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-sso-failover-generate
		RegenerateV1(ctx context.Context) (*FailoverSettings, *interfaces.Response, error)
	}

	// Service handles communication with the SSO failover-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-sso-failover
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - SSO Failover Operations
// -----------------------------------------------------------------------------

// GetV1 retrieves the current SSO failover settings.
// URL: GET /api/v1/sso/failover
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-sso-failover
func (s *Service) GetV1(ctx context.Context) (*FailoverSettings, *interfaces.Response, error) {
	var result FailoverSettings

	endpoint := EndpointSSOFailoverV1
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

// RegenerateV1 generates a new SSO failover URL.
// URL: POST /api/v1/sso/failover/generate
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-sso-failover-generate
func (s *Service) RegenerateV1(ctx context.Context) (*FailoverSettings, *interfaces.Response, error) {
	var result FailoverSettings

	endpoint := EndpointSSOFailoverGenerateV1
	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
