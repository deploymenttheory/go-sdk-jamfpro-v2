package certificate_authority

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the certificate authority-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-active
	CertificateAuthority struct {
		client client.Client
	}
)

func NewCertificateAuthority(client client.Client) *CertificateAuthority {
	return &CertificateAuthority{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Certificate Authority Operations
// -----------------------------------------------------------------------------

// GetV1 returns the active certificate authority X.509 details.
// URL: GET /api/v1/pki/certificate-authority/active
// https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-active
func (s *CertificateAuthority) GetV1(ctx context.Context) (*ResourceActiveCertificateAuthorityV1, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/active", constants.EndpointJamfProCertificateAuthorityV1)
	var result ResourceActiveCertificateAuthorityV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetActiveCertificateAuthorityDERV1 returns the active certificate authority in DER format.
// URL: GET /api/v1/pki/certificate-authority/active/der
// https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-active-der
func (s *CertificateAuthority) GetActiveCertificateAuthorityDERV1(ctx context.Context) (*resty.Response, []byte, error) {
	endpoint := fmt.Sprintf("%s/active/der", constants.EndpointJamfProCertificateAuthorityV1)

	resp, body, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationPKIXCert).
		GetBytes(endpoint)

	return resp, body, err
}

// GetActiveCertificateAuthorityPEMV1 returns the active certificate authority in PEM format.
// URL: GET /api/v1/pki/certificate-authority/active/pem
// https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-active-pem
func (s *CertificateAuthority) GetActiveCertificateAuthorityPEMV1(ctx context.Context) (*resty.Response, []byte, error) {
	endpoint := fmt.Sprintf("%s/active/pem", constants.EndpointJamfProCertificateAuthorityV1)

	resp, body, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationPEMCertificateChain).
		GetBytes(endpoint)

	return resp, body, err
}

// GetCertificateAuthorityByIDV1 returns X.509 details of the certificate authority with the given ID.
// URL: GET /api/v1/pki/certificate-authority/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-id
func (s *CertificateAuthority) GetCertificateAuthorityByIDV1(ctx context.Context, id string) (*ResourceActiveCertificateAuthorityV1, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/%s", constants.EndpointJamfProCertificateAuthorityV1, id)
	var result ResourceActiveCertificateAuthorityV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)

	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetCertificateAuthorityByIDDERV1 returns the certificate authority with the given ID in DER format.
// URL: GET /api/v1/pki/certificate-authority/{id}/der
// https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-id-der
func (s *CertificateAuthority) GetCertificateAuthorityByIDDERV1(ctx context.Context, id string) (*resty.Response, []byte, error) {
	endpoint := fmt.Sprintf("%s/%s/der", constants.EndpointJamfProCertificateAuthorityV1, id)

	resp, body, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationPKIXCert).
		GetBytes(endpoint)

	return resp, body, err
}

// GetCertificateAuthorityByIDPEMV1 returns the certificate authority with the given ID in PEM format.
// URL: GET /api/v1/pki/certificate-authority/{id}/pem
// https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-id-pem
func (s *CertificateAuthority) GetCertificateAuthorityByIDPEMV1(ctx context.Context, id string) (*resty.Response, []byte, error) {
	endpoint := fmt.Sprintf("%s/%s/pem", constants.EndpointJamfProCertificateAuthorityV1, id)

	resp, body, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationPEMCertificateChain).
		GetBytes(endpoint)

	return resp, body, err
}
