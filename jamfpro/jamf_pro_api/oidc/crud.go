package oidc

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the OIDC-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-oidc-public-key
	Oidc struct {
		client client.Client
	}
)

// NewService creates a new OIDC service.
func NewOidc(client client.Client) *Oidc {
	return &Oidc{client: client}
}

// GetDirectIdPLoginURLV1 retrieves the direct IdP login URL for OIDC.
// URL: GET /api/v1/oidc/direct-idp-login-url
// https://developer.jamf.com/jamf-pro/reference/get_v1-oidc-direct-idp-login-url
func (s *Oidc) GetDirectIdPLoginURLV1(ctx context.Context) (*ResourceOIDCDirectIdPLoginURL, *resty.Response, error) {
	endpoint := constants.EndpointJamfProOIDCV1 + "/direct-idp-login-url"

	var result ResourceOIDCDirectIdPLoginURL

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
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

	var result ResourceOIDCPublicKey

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
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

	var result ResourcePublicFeatures

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
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

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		Post(endpoint)
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

	var result ResourceOIDCRedirectURL

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get OIDC redirect URL: %w", err)
	}

	return &result, resp, nil
}
