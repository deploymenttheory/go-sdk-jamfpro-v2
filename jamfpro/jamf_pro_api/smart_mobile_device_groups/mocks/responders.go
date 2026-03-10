package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type SmartMobileDeviceGroupsMock struct {
	*mocks.GenericMock
}

func NewSmartMobileDeviceGroupsMock() *SmartMobileDeviceGroupsMock {
	return &SmartMobileDeviceGroupsMock{
		GenericMock: mocks.NewJSONMock("SmartMobileDeviceGroupsMock"),
	}
}

func (m *SmartMobileDeviceGroupsMock) RegisterListMock() {
	m.Register("GET", "/api/v2/mobile-device-groups/smart-groups", 200, "validate_list.json")
}

func (m *SmartMobileDeviceGroupsMock) RegisterListEmptyMock() {
	m.Register("GET", "/api/v2/mobile-device-groups/smart-groups", 200, "validate_list_empty.json")
}

func (m *SmartMobileDeviceGroupsMock) RegisterGetMock() {
	m.Register("GET", "/api/v2/mobile-device-groups/smart-groups/1", 200, "validate_get.json")
}

func (m *SmartMobileDeviceGroupsMock) RegisterGetMembershipMock() {
	m.Register("GET", "/api/v2/mobile-device-groups/smart-group-membership/1", 200, "validate_get_membership.json")
}

func (m *SmartMobileDeviceGroupsMock) RegisterCreateMock() {
	m.Register("POST", "/api/v2/mobile-device-groups/smart-groups", 201, "validate_create.json")
}

func (m *SmartMobileDeviceGroupsMock) RegisterUpdateMock() {
	m.Register("PUT", "/api/v2/mobile-device-groups/smart-groups/1", 200, "validate_update.json")
}

func (m *SmartMobileDeviceGroupsMock) RegisterDeleteMock() {
	m.Register("DELETE", "/api/v2/mobile-device-groups/smart-groups/1", 204, "")
}

func (m *SmartMobileDeviceGroupsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v2/mobile-device-groups/smart-groups/999", 404, "error_not_found.json", "")
}

func (m *SmartMobileDeviceGroupsMock) RegisterListNoResponseErrorMock() {
	m.RegisterError("GET", "/api/v2/mobile-device-groups/smart-groups", 500, "error_internal.json", "no response registered")
}

func (m *SmartMobileDeviceGroupsMock) RegisterGetMembershipNoResponseErrorMock() {
	m.RegisterError("GET", "/api/v2/mobile-device-groups/smart-group-membership/1", 500, "error_internal.json", "no response registered")
}

func (m *SmartMobileDeviceGroupsMock) RegisterCreateNoResponseErrorMock() {
	m.RegisterError("POST", "/api/v2/mobile-device-groups/smart-groups", 500, "error_internal.json", "no response registered")
}

func (m *SmartMobileDeviceGroupsMock) RegisterUpdateNoResponseErrorMock() {
	m.RegisterError("PUT", "/api/v2/mobile-device-groups/smart-groups/1", 500, "error_internal.json", "no response registered")
}

func (m *SmartMobileDeviceGroupsMock) RegisterDeleteNoResponseErrorMock() {
	m.RegisterError("DELETE", "/api/v2/mobile-device-groups/smart-groups/1", 500, "error_internal.json", "no response registered")
}
