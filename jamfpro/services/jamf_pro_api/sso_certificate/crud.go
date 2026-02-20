package sso_certificate

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// SsoCertificateServiceInterface defines the interface for SSO certificate operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-sso-cert
	SsoCertificateServiceInterface interface {
		// GetV2 returns the certificate currently configured for SSO (Get SSO Certificate).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-sso-cert
		GetV2(ctx context.Context) (*ResourceSSOKeystoreResponse, *interfaces.Response, error)

		// CreateV2 generates a new certificate for signing SSO requests (Create SSO Certificate).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-sso-cert
		CreateV2(ctx context.Context) (*ResourceSSOKeystoreResponse, *interfaces.Response, error)

		// DeleteV2 removes the currently configured SSO certificate (Delete SSO Certificate).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-sso-cert
		DeleteV2(ctx context.Context) (*interfaces.Response, error)
	}

	// Service handles communication with the SSO certificate-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-sso-cert
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ SsoCertificateServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - SSO Certificate Operations
// -----------------------------------------------------------------------------

// GetV2 returns the certificate currently configured for SSO.
// URL: GET /api/v2/sso/cert
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-sso-cert
func (s *Service) GetV2(ctx context.Context) (*ResourceSSOKeystoreResponse, *interfaces.Response, error) {
	var result ResourceSSOKeystoreResponse
	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}
	resp, err := s.client.Get(ctx, EndpointSSOCertV2, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// CreateV2 generates a new certificate for signing SSO requests.
// URL: POST /api/v2/sso/cert
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-sso-cert
func (s *Service) CreateV2(ctx context.Context) (*ResourceSSOKeystoreResponse, *interfaces.Response, error) {
	var result ResourceSSOKeystoreResponse
	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}
	resp, err := s.client.Post(ctx, EndpointSSOCertV2, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// DeleteV2 removes the currently configured SSO certificate.
// URL: DELETE /api/v2/sso/cert
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-sso-cert
func (s *Service) DeleteV2(ctx context.Context) (*interfaces.Response, error) {
	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}
	resp, err := s.client.Delete(ctx, EndpointSSOCertV2, nil, headers, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
