package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type SsoSettingsMock struct {
	*mocks.GenericMock
}

func NewSsoSettingsMock() *SsoSettingsMock {
	return &SsoSettingsMock{
		GenericMock: mocks.NewJSONMock("SsoSettingsMock"),
	}
}

func (m *SsoSettingsMock) RegisterGetMock() {
	m.Register("GET", "/api/v3/sso", 200, "validate_get.json")
}

func (m *SsoSettingsMock) RegisterUpdateMock() {
	m.Register("PUT", "/api/v3/sso", 200, "validate_get.json")
}

func (m *SsoSettingsMock) RegisterGetDependenciesMock() {
	m.Register("GET", "/api/v3/sso/dependencies", 200, "validate_dependencies.json")
}

func (m *SsoSettingsMock) RegisterDisableMock() {
	m.Register("POST", "/api/v3/sso/disable", 204, "")
}

func (m *SsoSettingsMock) RegisterGetHistoryMock() {
	m.Register("GET", "/api/v3/sso/history", 200, "validate_history.json")
}

func (m *SsoSettingsMock) RegisterAddHistoryNoteMock() {
	m.Register("POST", "/api/v3/sso/history", 201, "validate_add_history_note.json")
}

func (m *SsoSettingsMock) RegisterDownloadMetadataMock() {
	m.Register("GET", "/api/v3/sso/metadata/download", 200, "")
}

func (m *SsoSettingsMock) RegisterGetErrorMock() {
	m.RegisterError("GET", "/api/v3/sso", 500, "error_not_found.json", "")
}

func (m *SsoSettingsMock) RegisterUpdateErrorMock() {
	m.RegisterError("PUT", "/api/v3/sso", 500, "error_not_found.json", "")
}

func (m *SsoSettingsMock) RegisterGetDependenciesErrorMock() {
	m.RegisterError("GET", "/api/v3/sso/dependencies", 500, "error_not_found.json", "")
}

func (m *SsoSettingsMock) RegisterDisableErrorMock() {
	m.RegisterError("POST", "/api/v3/sso/disable", 500, "error_not_found.json", "")
}

func (m *SsoSettingsMock) RegisterGetHistoryErrorMock() {
	m.RegisterError("GET", "/api/v3/sso/history", 500, "error_not_found.json", "")
}

func (m *SsoSettingsMock) RegisterAddHistoryNoteErrorMock() {
	m.RegisterError("POST", "/api/v3/sso/history", 500, "error_not_found.json", "")
}

func (m *SsoSettingsMock) RegisterDownloadMetadataErrorMock() {
	m.RegisterError("GET", "/api/v3/sso/metadata/download", 500, "error_not_found.json", "")
}
