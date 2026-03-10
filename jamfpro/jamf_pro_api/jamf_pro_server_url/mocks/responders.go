package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type JamfProServerURLMock struct {
	*mocks.GenericMock
}

func NewJamfProServerURLMock() *JamfProServerURLMock {
	return &JamfProServerURLMock{
		GenericMock: mocks.NewJSONMock("JamfProServerURLMock"),
	}
}

func (m *JamfProServerURLMock) RegisterGetMock() {
	m.Register("GET", "/api/v1/jamf-pro-server-url", 200, "validate_get.json")
}

func (m *JamfProServerURLMock) RegisterUpdateMock() {
	m.Register("PUT", "/api/v1/jamf-pro-server-url", 200, "validate_update.json")
}

func (m *JamfProServerURLMock) RegisterGetHistoryMock() {
	m.Register("GET", "/api/v1/jamf-pro-server-url/history", 200, "validate_history.json")
}

func (m *JamfProServerURLMock) RegisterCreateHistoryNoteMock() {
	m.Register("POST", "/api/v1/jamf-pro-server-url/history", 201, "validate_history_note.json")
}

