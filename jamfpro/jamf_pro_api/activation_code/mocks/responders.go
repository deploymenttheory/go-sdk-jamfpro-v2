package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ActivationCodeMock struct {
	*mocks.GenericMock
}

func NewActivationCodeMock() *ActivationCodeMock {
	return &ActivationCodeMock{
		GenericMock: mocks.NewJSONMock("ActivationCodeMock"),
	}
}

func (m *ActivationCodeMock) RegisterMocks() {
	m.RegisterGetHistoryMock()
	m.RegisterUpdateActivationCodeMock()
	m.RegisterUpdateOrganizationNameMock()
	m.RegisterAddHistoryNoteMock()
	m.RegisterExportHistoryMock()
}

func (m *ActivationCodeMock) RegisterGetHistoryMock() {
	m.Register("GET", "/api/v1/activation-code/history", 200, "validate_history.json")
}

func (m *ActivationCodeMock) RegisterUpdateActivationCodeMock() {
	m.Register("PUT", "/api/v1/activation-code", 202, "")
}

func (m *ActivationCodeMock) RegisterUpdateOrganizationNameMock() {
	m.Register("PATCH", "/api/v1/activation-code/organization-name", 202, "")
}

func (m *ActivationCodeMock) RegisterAddHistoryNoteMock() {
	m.Register("POST", "/api/v1/activation-code/history", 201, "validate_add_history_note.json")
}

func (m *ActivationCodeMock) RegisterExportHistoryMock() {
	m.Register("POST", "/api/v1/activation-code/history/export", 200, "export_history.json")
}
