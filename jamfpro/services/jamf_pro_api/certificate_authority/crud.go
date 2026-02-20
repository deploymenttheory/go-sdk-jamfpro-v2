package certificate_authority

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// CertificateAuthorityServiceInterface defines the interface for certificate authority operations (read-only).
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-active
	CertificateAuthorityServiceInterface interface {
		// GetV1 returns the active certificate authority X.509 details.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-active
		GetV1(ctx context.Context) (*ResourceActiveCertificateAuthorityV1, *interfaces.Response, error)
		// GetActiveCertificateAuthorityDERV1 returns the active certificate authority in DER format.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-active-der
		GetActiveCertificateAuthorityDERV1(ctx context.Context) (*interfaces.Response, []byte, error)
		// GetActiveCertificateAuthorityPEMV1 returns the active certificate authority in PEM format.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-active-pem
		GetActiveCertificateAuthorityPEMV1(ctx context.Context) (*interfaces.Response, []byte, error)
		// GetCertificateAuthorityByIDV1 returns X.509 details of the certificate authority with the given ID.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-id
		GetCertificateAuthorityByIDV1(ctx context.Context, id string) (*ResourceActiveCertificateAuthorityV1, *interfaces.Response, error)
		// GetCertificateAuthorityByIDDERV1 returns the certificate authority with the given ID in DER format.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-id-der
		GetCertificateAuthorityByIDDERV1(ctx context.Context, id string) (*interfaces.Response, []byte, error)
		// GetCertificateAuthorityByIDPEMV1 returns the certificate authority with the given ID in PEM format.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-id-pem
		GetCertificateAuthorityByIDPEMV1(ctx context.Context, id string) (*interfaces.Response, []byte, error)
	}

	// Service handles communication with the certificate authority-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-active
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ CertificateAuthorityServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Certificate Authority Operations
// -----------------------------------------------------------------------------

// GetV1 returns the active certificate authority X.509 details.
// URL: GET /api/v1/pki/certificate-authority/active
// https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-active
func (s *Service) GetV1(ctx context.Context) (*ResourceActiveCertificateAuthorityV1, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/active", EndpointCertificateAuthorityV1)
	var result ResourceActiveCertificateAuthorityV1

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

// GetActiveCertificateAuthorityDERV1 returns the active certificate authority in DER format.
// URL: GET /api/v1/pki/certificate-authority/active/der
// https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-active-der
func (s *Service) GetActiveCertificateAuthorityDERV1(ctx context.Context) (*interfaces.Response, []byte, error) {
	endpoint := fmt.Sprintf("%s/active/der", EndpointCertificateAuthorityV1)

	headers := map[string]string{"Accept": mime.ApplicationPKIXCert}

	return s.client.GetBytes(ctx, endpoint, nil, headers)
}

// GetActiveCertificateAuthorityPEMV1 returns the active certificate authority in PEM format.
// URL: GET /api/v1/pki/certificate-authority/active/pem
// https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-active-pem
func (s *Service) GetActiveCertificateAuthorityPEMV1(ctx context.Context) (*interfaces.Response, []byte, error) {
	endpoint := fmt.Sprintf("%s/active/pem", EndpointCertificateAuthorityV1)

	headers := map[string]string{
		"Accept": mime.ApplicationPEMCertificateChain,
	}

	return s.client.GetBytes(ctx, endpoint, nil, headers)
}

// GetCertificateAuthorityByIDV1 returns X.509 details of the certificate authority with the given ID.
// URL: GET /api/v1/pki/certificate-authority/{id}
// https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-id
func (s *Service) GetCertificateAuthorityByIDV1(ctx context.Context, id string) (*ResourceActiveCertificateAuthorityV1, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/%s", EndpointCertificateAuthorityV1, id)
	var result ResourceActiveCertificateAuthorityV1

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

// GetCertificateAuthorityByIDDERV1 returns the certificate authority with the given ID in DER format.
// URL: GET /api/v1/pki/certificate-authority/{id}/der
// https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-id-der
func (s *Service) GetCertificateAuthorityByIDDERV1(ctx context.Context, id string) (*interfaces.Response, []byte, error) {
	endpoint := fmt.Sprintf("%s/%s/der", EndpointCertificateAuthorityV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationPKIXCert,
	}

	return s.client.GetBytes(ctx, endpoint, nil, headers)
}

// GetCertificateAuthorityByIDPEMV1 returns the certificate authority with the given ID in PEM format.
// URL: GET /api/v1/pki/certificate-authority/{id}/pem
// https://developer.jamf.com/jamf-pro/reference/get_v1-pki-certificate-authority-id-pem
func (s *Service) GetCertificateAuthorityByIDPEMV1(ctx context.Context, id string) (*interfaces.Response, []byte, error) {
	endpoint := fmt.Sprintf("%s/%s/pem", EndpointCertificateAuthorityV1, id)

	headers := map[string]string{
		"Accept": mime.ApplicationPEMCertificateChain,
	}

	return s.client.GetBytes(ctx, endpoint, nil, headers)
}
