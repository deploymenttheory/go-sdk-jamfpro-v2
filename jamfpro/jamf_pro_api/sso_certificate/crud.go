package sso_certificate

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// SsoCertificateServiceInterface defines the interface for SSO certificate operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-sso-cert
	SsoCertificateServiceInterface interface {
		// GetV2 returns the certificate currently configured for SSO (Get SSO Certificate).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-sso-cert
		GetV2(ctx context.Context) (*ResourceSSOKeystoreResponse, *resty.Response, error)

		// CreateV2 generates a new certificate for signing SSO requests (Create SSO Certificate).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-sso-cert
		CreateV2(ctx context.Context) (*ResourceSSOKeystoreResponse, *resty.Response, error)

		// UpdateV2 updates the certificate used for signing SSO requests (Update SSO Certificate).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-sso-cert
		UpdateV2(ctx context.Context, request *UpdateKeystoreRequest) (*ResourceSSOKeystoreResponse, *resty.Response, error)

		// DownloadV2 downloads the certificate used for signing SSO requests (Download SSO Certificate).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-sso-cert-download
		DownloadV2(ctx context.Context) ([]byte, *resty.Response, error)

		// ParseV2 parses the provided keystore file and returns keystore information (Parse SSO Certificate).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-sso-cert-parse
		ParseV2(ctx context.Context, request *ParseKeystoreRequest) (*ParseKeystoreResponse, *resty.Response, error)

		// DeleteV2 removes the currently configured SSO certificate (Delete SSO Certificate).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-sso-cert
		DeleteV2(ctx context.Context) (*resty.Response, error)
	}

	// Service handles communication with the SSO certificate-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-sso-cert
	SsoCertificate struct {
		client transport.HTTPClient
	}
)

var _ SsoCertificateServiceInterface = (*SsoCertificate)(nil)

func NewSsoCertificate(client transport.HTTPClient) *SsoCertificate {
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
	endpoint := EndpointSSOCertV2
	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
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
	endpoint := EndpointSSOCertV2
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

// UpdateV2 updates the certificate used for signing SSO requests.
// URL: PUT /api/v2/sso/cert
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-sso-cert
func (s *SsoCertificate) UpdateV2(ctx context.Context, request *UpdateKeystoreRequest) (*ResourceSSOKeystoreResponse, *resty.Response, error) {
	var result ResourceSSOKeystoreResponse
	endpoint := EndpointSSOCertV2
	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}
	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// DownloadV2 downloads the certificate used for signing SSO requests.
// URL: GET /api/v2/sso/cert/download
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-sso-cert-download
func (s *SsoCertificate) DownloadV2(ctx context.Context) ([]byte, *resty.Response, error) {
	var result []byte
	endpoint := EndpointSSOCertDownloadV2
	headers := map[string]string{
		"Accept": mime.ApplicationOctetStream,
	}
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return result, resp, nil
}

// ParseV2 parses the provided keystore file and returns keystore information.
// URL: POST /api/v2/sso/cert/parse
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-sso-cert-parse
func (s *SsoCertificate) ParseV2(ctx context.Context, request *ParseKeystoreRequest) (*ParseKeystoreResponse, *resty.Response, error) {
	var result ParseKeystoreResponse
	endpoint := EndpointSSOCertParseV2
	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}
	resp, err := s.client.Post(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// DeleteV2 removes the currently configured SSO certificate.
// URL: DELETE /api/v2/sso/cert
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/delete_v2-sso-cert
func (s *SsoCertificate) DeleteV2(ctx context.Context) (*resty.Response, error) {
	endpoint := EndpointSSOCertV2
	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}
	resp, err := s.client.Delete(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
