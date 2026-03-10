package sso_certificate

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the SSO certificate-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-sso-cert
	SsoCertificate struct {
		client client.Client
	}
)

func NewSsoCertificate(client client.Client) *SsoCertificate {
	return &SsoCertificate{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - SSO Certificate Operations
// -----------------------------------------------------------------------------

// GetV2 returns the certificate currently configured for SSO.
// URL: GET /api/v2/sso/cert
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-sso-cert
func (s *SsoCertificate) GetV2(ctx context.Context) (*ResourceSSOKeystoreResponse, *resty.Response, error) {
	var result ResourceSSOKeystoreResponse
	endpoint := constants.EndpointJamfProSSOCertV2

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// CreateV2 generates a new certificate for signing SSO requests.
// URL: POST /api/v2/sso/cert
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-sso-cert
func (s *SsoCertificate) CreateV2(ctx context.Context) (*ResourceSSOKeystoreResponse, *resty.Response, error) {
	var result ResourceSSOKeystoreResponse
	endpoint := constants.EndpointJamfProSSOCertV2

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

// UpdateV2 updates the certificate used for signing SSO requests.
// URL: PUT /api/v2/sso/cert
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-sso-cert
func (s *SsoCertificate) UpdateV2(ctx context.Context, request *UpdateKeystoreRequest) (*ResourceSSOKeystoreResponse, *resty.Response, error) {
	var result ResourceSSOKeystoreResponse
	endpoint := constants.EndpointJamfProSSOCertV2

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Put(endpoint)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// DownloadV2 downloads the certificate used for signing SSO requests.
// URL: GET /api/v2/sso/cert/download
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-sso-cert-download
func (s *SsoCertificate) DownloadV2(ctx context.Context) ([]byte, *resty.Response, error) {
	endpoint := constants.EndpointJamfProSSOCertDownloadV2

	resp, data, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationOctetStream).
		GetBytes(endpoint)
	if err != nil {
		return nil, resp, err
	}
	return data, resp, nil
}

// ParseV2 parses the provided keystore file and returns keystore information.
// URL: POST /api/v2/sso/cert/parse
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-sso-cert-parse
func (s *SsoCertificate) ParseV2(ctx context.Context, request *ParseKeystoreRequest) (*ParseKeystoreResponse, *resty.Response, error) {
	var result ParseKeystoreResponse
	endpoint := constants.EndpointJamfProSSOCertParseV2

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		SetBody(request).
		SetResult(&result).
		Post(endpoint)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// DeleteV2 removes the currently configured SSO certificate.
// URL: DELETE /api/v2/sso/cert
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-sso-cert
func (s *SsoCertificate) DeleteV2(ctx context.Context) (*resty.Response, error) {
	endpoint := constants.EndpointJamfProSSOCertV2

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Delete(endpoint)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
