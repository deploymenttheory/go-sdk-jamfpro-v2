package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type AdcsSettingsMock struct {
	*mocks.GenericMock
}

func NewAdcsSettingsMock() *AdcsSettingsMock {
	return &AdcsSettingsMock{
		GenericMock: mocks.NewJSONMock("AdcsSettingsMock"),
	}
}

func (m *AdcsSettingsMock) RegisterCreateMock() {
	m.Register("POST", "/api/v1/pki/adcs-settings", 201, "validate_create.json")
}

func (m *AdcsSettingsMock) RegisterGetByIDMock(id string) {
	m.Register("GET", "/api/v1/pki/adcs-settings/"+id, 200, "validate_get.json")
}

func (m *AdcsSettingsMock) RegisterUpdateByIDMock(id string) {
	m.Register("PATCH", "/api/v1/pki/adcs-settings/"+id, 204, "")
}

func (m *AdcsSettingsMock) RegisterDeleteByIDMock(id string) {
	m.Register("DELETE", "/api/v1/pki/adcs-settings/"+id, 204, "")
}

func (m *AdcsSettingsMock) RegisterValidateServerCertificateMock() {
	m.Register("POST", "/api/v1/pki/adcs-settings/validate-certificate", 204, "")
}

func (m *AdcsSettingsMock) RegisterValidateClientCertificateMock() {
	m.Register("POST", "/api/v1/pki/adcs-settings/validate-client-certificate", 204, "")
}

func (m *AdcsSettingsMock) RegisterGetDependenciesByIDMock(id string) {
	m.Register("GET", "/api/v1/pki/adcs-settings/"+id+"/dependencies", 200, "validate_dependencies.json")
}

func (m *AdcsSettingsMock) RegisterGetHistoryByIDMock(id string) {
	m.Register("GET", "/api/v1/pki/adcs-settings/"+id+"/history", 200, "validate_history.json")
}

func (m *AdcsSettingsMock) RegisterAddHistoryNoteMock(id string) {
	m.Register("POST", "/api/v1/pki/adcs-settings/"+id+"/history", 201, "validate_add_history_note.json")
}
