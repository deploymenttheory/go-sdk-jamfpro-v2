package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type JCDSMock struct {
	*mocks.GenericMock
}

func NewJCDSMock() *JCDSMock {
	return &JCDSMock{
		GenericMock: mocks.NewJSONMock("JCDSMock"),
	}
}

func (m *JCDSMock) RegisterGetPackagesMock() {
	m.Register("GET", "/api/v1/jcds/files", 200, "validate_get_packages.json")
}

func (m *JCDSMock) RegisterGetPackageURIByNameMock() {
	m.Register("GET", "/api/v1/jcds/files/test-package.pkg", 200, "validate_get_package_uri.json")
}

func (m *JCDSMock) RegisterRenewCredentialsMock() {
	m.Register("POST", "/api/v1/jcds/renew-credentials", 200, "validate_renew_credentials.json")
}

func (m *JCDSMock) RegisterRefreshInventoryMock() {
	m.Register("POST", "/api/v1/jcds/refresh-inventory", 204, "")
}

func (m *JCDSMock) RegisterUploadCredentialsMock() {
	m.Register("POST", "/api/v1/jcds/files", 200, "validate_upload_credentials.json")
}

func (m *JCDSMock) RegisterIncompleteCredentialsMock() {
	m.Register("POST", "/api/v1/jcds/files", 200, "validate_upload_credentials_incomplete.json")
}

func (m *JCDSMock) RegisterErrorMock(method, path string, errMsg string) {
	m.RegisterError(method, path, 500, "", errMsg)
}
