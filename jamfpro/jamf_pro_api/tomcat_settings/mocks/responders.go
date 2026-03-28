package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type TomcatSettingsMock struct {
	*mocks.GenericMock
}

func NewTomcatSettingsMock() *TomcatSettingsMock {
	return &TomcatSettingsMock{
		GenericMock: mocks.NewJSONMock("TomcatSettingsMock"),
	}
}

func (m *TomcatSettingsMock) RegisterIssueTomcatSslCertificateMock() {
	m.Register("POST", "/api/settings/issueTomcatSslCertificate", 200, "")
}

func (m *TomcatSettingsMock) RegisterIssueTomcatSslCertificateErrorMock() {
	m.RegisterError("POST", "/api/settings/issueTomcatSslCertificate", 500, "error_not_found.json", "Jamf Pro API error (500) [CERTIFICATE_ISSUE_FAILED]: SSL certificate issue failed")
}

func (m *TomcatSettingsMock) RegisterIssueTomcatSslCertificateNoResponseErrorMock() {
	m.RegisterError("POST", "/api/settings/issueTomcatSslCertificate", 500, "error_internal.json", "no response registered")
}
