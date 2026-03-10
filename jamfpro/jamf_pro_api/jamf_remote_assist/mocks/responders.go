package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type JamfRemoteAssistMock struct {
	*mocks.GenericMock
}

func NewJamfRemoteAssistMock() *JamfRemoteAssistMock {
	return &JamfRemoteAssistMock{
		GenericMock: mocks.NewJSONMock("JamfRemoteAssistMock"),
	}
}

func (m *JamfRemoteAssistMock) RegisterMocks() {
	m.Register("GET", "/api/v1/jamf-remote-assist/session", 200, "validate_list.json")
	m.Register("GET", "/api/v1/jamf-remote-assist/session/session-abc", 200, "validate_session.json")
	m.Register("GET", "/api/v2/jamf-remote-assist/session", 200, "validate_list_v2.json")
	m.Register("GET", "/api/v2/jamf-remote-assist/session/session-abc", 200, "validate_session.json")
	m.Register("POST", "/api/v2/jamf-remote-assist/session/export", 200, "")
}

func (m *JamfRemoteAssistMock) RegisterListSessionsV1ErrorMock() {
	m.RegisterError("GET", "/api/v1/jamf-remote-assist/session", 500, "error_not_found.json", "")
}

func (m *JamfRemoteAssistMock) RegisterListSessionsV2ErrorMock() {
	m.RegisterError("GET", "/api/v2/jamf-remote-assist/session", 500, "error_not_found.json", "")
}

func (m *JamfRemoteAssistMock) RegisterListSessionsV2InvalidMock() {
	m.Register("GET", "/api/v2/jamf-remote-assist/session", 200, "validate_list_v2_invalid.json")
}

func (m *JamfRemoteAssistMock) RegisterGetSessionByIDV2ErrorMock() {
	m.RegisterError("GET", "/api/v2/jamf-remote-assist/session/nonexistent", 404, "error_not_found.json", "")
}

func (m *JamfRemoteAssistMock) RegisterExportSessionsV2ErrorMock() {
	m.RegisterError("POST", "/api/v2/jamf-remote-assist/session/export", 500, "error_not_found.json", "")
}

