package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type DeviceCommunicationSettingsMock struct {
	*mocks.GenericMock
}

func NewDeviceCommunicationSettingsMock() *DeviceCommunicationSettingsMock {
	return &DeviceCommunicationSettingsMock{
		GenericMock: mocks.NewJSONMock("DeviceCommunicationSettingsMock"),
	}
}

func (m *DeviceCommunicationSettingsMock) RegisterGetMock() {
	m.Register("GET", "/api/v1/device-communication-settings", 200, "validate_get.json")
}

func (m *DeviceCommunicationSettingsMock) RegisterPutMock() {
	m.Register("PUT", "/api/v1/device-communication-settings", 200, "validate_get.json")
}

func (m *DeviceCommunicationSettingsMock) RegisterGetHistoryMock() {
	m.Register("GET", "/api/v1/device-communication-settings/history", 200, "validate_history.json")
}

func (m *DeviceCommunicationSettingsMock) RegisterAddHistoryNotesMock() {
	m.Register("POST", "/api/v1/device-communication-settings/history", 201, "validate_add_history_note.json")
}

func (m *DeviceCommunicationSettingsMock) RegisterGetErrorMock() {
	m.RegisterError("GET", "/api/v1/device-communication-settings", 500, "", "Jamf Pro API error (500): server error")
}

func (m *DeviceCommunicationSettingsMock) RegisterPutErrorMock() {
	m.RegisterError("PUT", "/api/v1/device-communication-settings", 500, "", "Jamf Pro API error (500): server error")
}

func (m *DeviceCommunicationSettingsMock) RegisterGetHistoryErrorMock() {
	m.RegisterError("GET", "/api/v1/device-communication-settings/history", 500, "", "Jamf Pro API error (500): server error")
}

func (m *DeviceCommunicationSettingsMock) RegisterGetHistoryInvalidJSONMock() {
	m.Register("GET", "/api/v1/device-communication-settings/history", 200, "validate_history_invalid_json.json")
}

func (m *DeviceCommunicationSettingsMock) RegisterGetHistoryInvalidItemMock() {
	m.Register("GET", "/api/v1/device-communication-settings/history", 200, "validate_history_invalid_item.json")
}

func (m *DeviceCommunicationSettingsMock) RegisterAddHistoryNotesErrorMock() {
	m.RegisterError("POST", "/api/v1/device-communication-settings/history", 500, "", "Jamf Pro API error (500): server error")
}
