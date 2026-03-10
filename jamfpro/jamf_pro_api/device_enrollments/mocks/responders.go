package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type DeviceEnrollmentsMock struct {
	*mocks.GenericMock
}

func NewDeviceEnrollmentsMock() *DeviceEnrollmentsMock {
	return &DeviceEnrollmentsMock{
		GenericMock: mocks.NewJSONMock("DeviceEnrollmentsMock"),
	}
}

func (m *DeviceEnrollmentsMock) RegisterListMock() {
	m.Register("GET", "/api/v1/device-enrollments", 200, "validate_list.json")
}

func (m *DeviceEnrollmentsMock) RegisterGetByIDMock(id string) {
	m.Register("GET", "/api/v1/device-enrollments/"+id, 200, "validate_get.json")
}

func (m *DeviceEnrollmentsMock) RegisterGetHistoryMock(id string) {
	m.Register("GET", "/api/v1/device-enrollments/"+id+"/history", 200, "validate_history.json")
}

func (m *DeviceEnrollmentsMock) RegisterGetSyncStatesMock(id string) {
	m.Register("GET", "/api/v1/device-enrollments/"+id+"/syncs", 200, "validate_sync_states.json")
}

func (m *DeviceEnrollmentsMock) RegisterGetPublicKeyMock() {
	m.Register("GET", "/api/v1/device-enrollments/public-key", 200, "validate_public_key.txt")
}

func (m *DeviceEnrollmentsMock) RegisterGetLatestSyncStateMock(id string) {
	m.Register("GET", "/api/v1/device-enrollments/"+id+"/syncs/latest", 200, "validate_latest_sync_state.json")
}

func (m *DeviceEnrollmentsMock) RegisterGetAllSyncStatesMock() {
	m.Register("GET", "/api/v1/device-enrollments/syncs", 200, "validate_sync_states.json")
}

func (m *DeviceEnrollmentsMock) RegisterCreateWithTokenMock() {
	m.Register("POST", "/api/v1/device-enrollments/upload-token", 201, "validate_create.json")
}

func (m *DeviceEnrollmentsMock) RegisterUpdateByIDMock(id string) {
	m.Register("PUT", "/api/v1/device-enrollments/"+id, 200, "validate_get.json")
}

func (m *DeviceEnrollmentsMock) RegisterUpdateTokenByIDMock(id string) {
	m.Register("PUT", "/api/v1/device-enrollments/"+id+"/upload-token", 200, "validate_get.json")
}

func (m *DeviceEnrollmentsMock) RegisterDeleteByIDMock(id string) {
	m.Register("DELETE", "/api/v1/device-enrollments/"+id, 204, "")
}

func (m *DeviceEnrollmentsMock) RegisterDisownDevicesMock(id string) {
	m.Register("POST", "/api/v1/device-enrollments/"+id+"/disown", 200, "validate_disown.json")
}

func (m *DeviceEnrollmentsMock) RegisterAddHistoryNotesMock(id string) {
	m.Register("POST", "/api/v1/device-enrollments/"+id+"/history", 201, "validate_add_history_notes.json")
}

func (m *DeviceEnrollmentsMock) RegisterGetDevicesByIDMock(id string) {
	m.Register("GET", "/api/v1/device-enrollments/"+id+"/devices", 200, "validate_devices.json")
}
