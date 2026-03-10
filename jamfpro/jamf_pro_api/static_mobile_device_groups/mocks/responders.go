package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type StaticMobileDeviceGroupsMock struct {
	*mocks.GenericMock
}

func NewStaticMobileDeviceGroupsMock() *StaticMobileDeviceGroupsMock {
	return &StaticMobileDeviceGroupsMock{
		GenericMock: mocks.NewJSONMock("StaticMobileDeviceGroupsMock"),
	}
}

func (m *StaticMobileDeviceGroupsMock) RegisterListMock() {
	m.Register("GET", "/api/v2/mobile-device-groups/static-groups", 200, "validate_list.json")
}

func (m *StaticMobileDeviceGroupsMock) RegisterGetMock() {
	m.Register("GET", "/api/v2/mobile-device-groups/static-groups/10", 200, "validate_get.json")
}

func (m *StaticMobileDeviceGroupsMock) RegisterCreateMock() {
	m.Register("POST", "/api/v2/mobile-device-groups/static-groups", 201, "validate_create.json")
}

func (m *StaticMobileDeviceGroupsMock) RegisterUpdateMock() {
	m.Register("PATCH", "/api/v2/mobile-device-groups/static-groups/10", 200, "validate_update.json")
}

func (m *StaticMobileDeviceGroupsMock) RegisterDeleteMock() {
	m.Register("DELETE", "/api/v2/mobile-device-groups/static-groups/10", 204, "")
}

func (m *StaticMobileDeviceGroupsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v2/mobile-device-groups/static-groups/999", 404, "error_not_found.json", "")
}
