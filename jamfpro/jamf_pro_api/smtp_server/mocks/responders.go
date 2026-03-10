package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type SMTPServerMock struct {
	*mocks.GenericMock
}

func NewSMTPServerMock() *SMTPServerMock {
	return &SMTPServerMock{
		GenericMock: mocks.NewJSONMock("SMTPServerMock"),
	}
}

func (m *SMTPServerMock) RegisterMocks() {
	m.Register("GET", "/api/v2/smtp-server", 200, "validate_get.json")
	m.Register("PUT", "/api/v2/smtp-server", 200, "validate_get.json")
	m.Register("GET", "/api/v1/smtp-server/history", 200, "validate_history.json")
	m.Register("POST", "/api/v1/smtp-server/history", 201, "validate_add_history_note.json")
	m.Register("POST", "/api/v1/smtp-server/test", 202, "")
}

func (m *SMTPServerMock) RegisterGetErrorMock() {
	m.RegisterError("GET", "/api/v2/smtp-server", 500, "validate_get.json", "Jamf Pro API error (500): server error")
}

func (m *SMTPServerMock) RegisterPutErrorMock() {
	m.RegisterError("PUT", "/api/v2/smtp-server", 500, "validate_get.json", "Jamf Pro API error (500): server error")
}

func (m *SMTPServerMock) RegisterGetHistoryErrorMock() {
	m.RegisterError("GET", "/api/v1/smtp-server/history", 500, "validate_history.json", "Jamf Pro API error (500): server error")
}

func (m *SMTPServerMock) RegisterGetHistoryInvalidJSONMock() {
	m.RegisterRawBody("GET", "/api/v1/smtp-server/history", 200, []byte(`{invalid json`))
}

func (m *SMTPServerMock) RegisterGetHistoryInvalidItemMock() {
	m.Register("GET", "/api/v1/smtp-server/history", 200, "validate_history_invalid_item.json")
}

func (m *SMTPServerMock) RegisterAddHistoryNoteErrorMock() {
	m.RegisterError("POST", "/api/v1/smtp-server/history", 500, "validate_add_history_note.json", "Jamf Pro API error (500): server error")
}

func (m *SMTPServerMock) RegisterTestErrorMock() {
	m.RegisterError("POST", "/api/v1/smtp-server/test", 500, "", "Jamf Pro API error (500): server error")
}
