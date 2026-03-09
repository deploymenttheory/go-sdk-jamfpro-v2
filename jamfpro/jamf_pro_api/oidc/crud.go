package oidc

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the OIDC-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-oidc-public-key
	Oidc struct {
		client transport.HTTPClient
	}
)

// NewService creates a new OIDC service.
func NewOidc(client transport.HTTPClient) *Oidc {
	return &Oidc{client: client}
}

// GetDirectIdPLoginURLV1 retrieves the direct IdP login URL for OIDC.
// URL: GET /api/v1/oidc/direct-idp-login-url
// https://developer.jamf.com/jamf-pro/reference/get_v1-oidc-direct-idp-login-url
func (s *Oidc) GetDirectIdPLoginURLV1(ctx context.Context) (*ResourceOIDCDirectIdPLoginURL, *resty.Response, error) {
	endpoint := constants.EndpointJamfProOIDCV1 + "/direct-idp-login-url"

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	var result ResourceOIDCDirectIdPLoginURL
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get OIDC direct IdP login URL: %w", err)
	}

	return &result, resp, nil
}

// GetPublicKeyV1 retrieves the public key used for signing OIDC messages.
// URL: GET /api/v1/oidc/public-key
// https://developer.jamf.com/jamf-pro/reference/get_v1-oidc-public-key
func (s *Oidc) GetPublicKeyV1(ctx context.Context) (*ResourceOIDCPublicKey, *resty.Response, error) {
	endpoint := constants.EndpointJamfProOIDCV1 + "/public-key"

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	var result ResourceOIDCPublicKey
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get OIDC public key: %w", err)
	}

	return &result, resp, nil
}

// GetPublicFeaturesV1 retrieves the public OIDC configuration features.
// URL: GET /api/v1/oidc/public-features
// https://developer.jamf.com/jamf-pro/reference/get_v1-oidc-public-features
func (s *Oidc) GetPublicFeaturesV1(ctx context.Context) (*ResourcePublicFeatures, *resty.Response, error) {
	endpoint := constants.EndpointJamfProOIDCV1 + "/public-features"

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	var result ResourcePublicFeatures
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get OIDC public features: %w", err)
	}

	return &result, resp, nil
}

// GenerateCertificateV1 generates a new certificate for signing OIDC messages.
// URL: POST /api/v1/oidc/generate-certificate
// https://developer.jamf.com/jamf-pro/reference/post_v1-oidc-generate-certificate
func (s *Oidc) GenerateCertificateV1(ctx context.Context) (*resty.Response, error) {
	endpoint := constants.EndpointJamfProOIDCV1 + "/generate-certificate"

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to generate OIDC certificate: %w", err)
	}

	return resp, nil
}

// GetRedirectURLV1 provides the redirect URL for OIDC authentication.
// URL: POST /api/v2/oidc/dispatch
// https://developer.jamf.com/jamf-pro/reference/post_v2-oidc-dispatch
func (s *Oidc) GetRedirectURLV1(ctx context.Context, request *RequestOIDCRedirectURL) (*ResourceOIDCRedirectURL, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("OIDC redirect URL request cannot be nil")
	}

	endpoint := constants.EndpointJamfProOIDCV2 + "/dispatch"

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	var result ResourceOIDCRedirectURL
	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get OIDC redirect URL: %w", err)
	}

	return &result, resp, nil
}
