package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type MobileDeviceGroupsMock struct {
	*mocks.GenericMock
}

func NewMobileDeviceGroupsMock() *MobileDeviceGroupsMock {
	return &MobileDeviceGroupsMock{
		GenericMock: mocks.NewXMLMock("MobileDeviceGroupsMock"),
	}
}

func (m *MobileDeviceGroupsMock) RegisterMocks() {
	m.RegisterListMobileDeviceGroupsMock()
	m.RegisterGetMobileDeviceGroupByIDMock()
	m.RegisterGetMobileDeviceGroupByNameMock()
	m.RegisterCreateMobileDeviceGroupMock()
	m.RegisterUpdateMobileDeviceGroupByIDMock()
	m.RegisterUpdateMobileDeviceGroupByNameMock()
	m.RegisterDeleteMobileDeviceGroupByIDMock()
	m.RegisterDeleteMobileDeviceGroupByNameMock()
}

func (m *MobileDeviceGroupsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *MobileDeviceGroupsMock) RegisterListMobileDeviceGroupsMock() {
	m.Register("GET", "/JSSResource/mobiledevicegroups", 200, "validate_list_mobile_device_groups.xml")
}

func (m *MobileDeviceGroupsMock) RegisterGetMobileDeviceGroupByIDMock() {
	m.Register("GET", "/JSSResource/mobiledevicegroups/id/1", 200, "validate_get_mobile_device_group.xml")
}

func (m *MobileDeviceGroupsMock) RegisterGetMobileDeviceGroupByNameMock() {
	m.Register("GET", "/JSSResource/mobiledevicegroups/name/All Mobile Devices", 200, "validate_get_mobile_device_group.xml")
}

func (m *MobileDeviceGroupsMock) RegisterCreateMobileDeviceGroupMock() {
	m.Register("POST", "/JSSResource/mobiledevicegroups/id/0", 201, "validate_create_mobile_device_group.xml")
}

func (m *MobileDeviceGroupsMock) RegisterUpdateMobileDeviceGroupByIDMock() {
	m.Register("PUT", "/JSSResource/mobiledevicegroups/id/1", 200, "validate_update_mobile_device_group.xml")
}

func (m *MobileDeviceGroupsMock) RegisterUpdateMobileDeviceGroupByNameMock() {
	m.Register("PUT", "/JSSResource/mobiledevicegroups/name/All Mobile Devices", 200, "validate_update_mobile_device_group.xml")
}

func (m *MobileDeviceGroupsMock) RegisterDeleteMobileDeviceGroupByIDMock() {
	m.Register("DELETE", "/JSSResource/mobiledevicegroups/id/1", 200, "")
}

func (m *MobileDeviceGroupsMock) RegisterDeleteMobileDeviceGroupByNameMock() {
	m.Register("DELETE", "/JSSResource/mobiledevicegroups/name/All Mobile Devices", 200, "")
}

func (m *MobileDeviceGroupsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/mobiledevicegroups/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Mobile device group not found")
}

func (m *MobileDeviceGroupsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/mobiledevicegroups/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A mobile device group with that name already exists")
}

