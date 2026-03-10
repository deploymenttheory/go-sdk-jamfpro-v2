package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type SelfServiceSettingsMock struct {
	*mocks.GenericMock
}

func NewSelfServiceSettingsMock() *SelfServiceSettingsMock {
	return &SelfServiceSettingsMock{
		GenericMock: mocks.NewJSONMock("SelfServiceSettingsMock"),
	}
}

func (m *SelfServiceSettingsMock) RegisterGetMock() {
	m.Register("GET", "/api/v1/self-service/settings", 200, "validate_get.json")
}

func (m *SelfServiceSettingsMock) RegisterUpdateMock() {
	m.Register("PUT", "/api/v1/self-service/settings", 200, "validate_get.json")
}

func (m *SelfServiceSettingsMock) RegisterGetHistoryMock() {
	m.Register("GET", "/api/v1/self-service/settings/history", 200, "validate_history.json")
}

func (m *SelfServiceSettingsMock) RegisterGetHistoryInvalidMock() {
	m.Register("GET", "/api/v1/self-service/settings/history", 200, "validate_history_invalid.json")
}

func (m *SelfServiceSettingsMock) RegisterAddHistoryNotesMock() {
	m.Register("POST", "/api/v1/self-service/settings/history", 201, "validate_add_history_notes.json")
}

func (m *SelfServiceSettingsMock) RegisterGetErrorMock() {
	m.RegisterError("GET", "/api/v1/self-service/settings", 500, "error_internal.json", "")
}

func (m *SelfServiceSettingsMock) RegisterUpdateErrorMock() {
	m.RegisterError("PUT", "/api/v1/self-service/settings", 500, "error_internal.json", "")
}

func (m *SelfServiceSettingsMock) RegisterGetHistoryErrorMock() {
	m.RegisterError("GET", "/api/v1/self-service/settings/history", 500, "error_internal.json", "")
}

func (m *SelfServiceSettingsMock) RegisterAddHistoryNotesErrorMock() {
	m.RegisterError("POST", "/api/v1/self-service/settings/history", 500, "error_internal.json", "")
}
