package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type DevicesMock struct {
	*mocks.GenericMock
}

func NewDevicesMock() *DevicesMock {
	return &DevicesMock{
		GenericMock: mocks.NewJSONMock("DevicesMock"),
	}
}

func (m *DevicesMock) RegisterGetGroupsMock() {
	m.Register("GET", "/api/v1/devices/1/groups", 200, "validate_get_groups.json")
}

func (m *DevicesMock) RegisterGetGroupsEmptyMock() {
	m.Register("GET", "/api/v1/devices/2/groups", 200, "validate_get_groups_empty.json")
}

func (m *DevicesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/devices/999/groups", 404, "error_not_found.json", "")
}

