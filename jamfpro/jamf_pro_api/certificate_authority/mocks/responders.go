package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type CertificateAuthorityMock struct {
	*mocks.GenericMock
}

func NewCertificateAuthorityMock() *CertificateAuthorityMock {
	return &CertificateAuthorityMock{
		GenericMock: mocks.NewJSONMock("CertificateAuthorityMock"),
	}
}

func (m *CertificateAuthorityMock) RegisterGetActiveCertificateAuthorityMock() {
	m.Register("GET", "/api/v1/pki/certificate-authority/active", 200, "validate_get.json")
}

func (m *CertificateAuthorityMock) RegisterGetActiveCertificateAuthorityDERMock() {
	m.Register("GET", "/api/v1/pki/certificate-authority/active/der", 200, "validate_active_der.der")
}

func (m *CertificateAuthorityMock) RegisterGetActiveCertificateAuthorityPEMMock() {
	m.Register("GET", "/api/v1/pki/certificate-authority/active/pem", 200, "validate_active_pem.pem")
}

func (m *CertificateAuthorityMock) RegisterGetCertificateAuthorityByIDMock(id string) {
	base := "/api/v1/pki/certificate-authority/" + id
	m.Register("GET", base, 200, "validate_get.json")
	m.Register("GET", base+"/der", 200, "validate_active_der.der")
	m.Register("GET", base+"/pem", 200, "validate_active_pem.pem")
}
