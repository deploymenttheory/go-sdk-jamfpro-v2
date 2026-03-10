package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type OIDCMock struct {
	*mocks.GenericMock
}

func NewOIDCMock() *OIDCMock {
	return &OIDCMock{
		GenericMock: mocks.NewJSONMock("OIDCMock"),
	}
}

func (m *OIDCMock) RegisterMocks() {
	m.RegisterGetDirectIdPLoginURLMock()
	m.RegisterGetPublicKeyMock()
	m.RegisterGetPublicFeaturesMock()
	m.RegisterGenerateCertificateMock()
	m.RegisterGetRedirectURLMock()
}

func (m *OIDCMock) RegisterGetDirectIdPLoginURLMock() {
	m.Register("GET", "/api/v1/oidc/direct-idp-login-url", 200, "validate_direct_idp_login_url.json")
}

func (m *OIDCMock) RegisterGetPublicKeyMock() {
	m.Register("GET", "/api/v1/oidc/public-key", 200, "validate_public_key.json")
}

func (m *OIDCMock) RegisterGetPublicFeaturesMock() {
	m.Register("GET", "/api/v1/oidc/public-features", 200, "validate_public_features.json")
}

func (m *OIDCMock) RegisterGenerateCertificateMock() {
	m.Register("POST", "/api/v1/oidc/generate-certificate", 204, "")
}

func (m *OIDCMock) RegisterGetRedirectURLMock() {
	m.Register("POST", "/api/v2/oidc/dispatch", 200, "validate_redirect_url.json")
}

func (m *OIDCMock) RegisterGetDirectIdPLoginURLErrorMock() {
	m.RegisterError("GET", "/api/v1/oidc/direct-idp-login-url", 500, "error_internal.json", "no response registered")
}

func (m *OIDCMock) RegisterGetPublicKeyErrorMock() {
	m.RegisterError("GET", "/api/v1/oidc/public-key", 500, "error_internal.json", "no response registered")
}

func (m *OIDCMock) RegisterGetPublicFeaturesErrorMock() {
	m.RegisterError("GET", "/api/v1/oidc/public-features", 500, "error_internal.json", "no response registered")
}

func (m *OIDCMock) RegisterGenerateCertificateErrorMock() {
	m.RegisterError("POST", "/api/v1/oidc/generate-certificate", 500, "error_internal.json", "no response registered")
}

func (m *OIDCMock) RegisterGetRedirectURLErrorMock() {
	m.RegisterError("POST", "/api/v2/oidc/dispatch", 500, "error_internal.json", "no response registered")
}
