package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ReenrollmentMock struct {
	*mocks.GenericMock
}

func NewReenrollmentMock() *ReenrollmentMock {
	return &ReenrollmentMock{
		GenericMock: mocks.NewJSONMock("ReenrollmentMock"),
	}
}

func (m *ReenrollmentMock) RegisterGetMock() {
	m.Register("GET", "/api/v1/reenrollment", 200, "validate_get.json")
}

func (m *ReenrollmentMock) RegisterUpdateMock() {
	m.Register("PUT", "/api/v1/reenrollment", 200, "validate_get.json")
}

func (m *ReenrollmentMock) RegisterGetHistoryMock() {
	m.Register("GET", "/api/v1/reenrollment/history", 200, "validate_history.json")
}

func (m *ReenrollmentMock) RegisterAddHistoryNotesMock() {
	m.Register("POST", "/api/v1/reenrollment/history", 201, "validate_add_history_note.json")
}

func (m *ReenrollmentMock) RegisterExportHistoryMock() {
	m.Register("POST", "/api/v1/reenrollment/history/export", 200, "validate_export.json")
}

func (m *ReenrollmentMock) RegisterGetErrorMock() {
	m.RegisterError("GET", "/api/v1/reenrollment", 500, "error_internal.json", "")
}

func (m *ReenrollmentMock) RegisterGetHistoryErrorMock() {
	m.RegisterError("GET", "/api/v1/reenrollment/history", 500, "error_internal.json", "")
}
