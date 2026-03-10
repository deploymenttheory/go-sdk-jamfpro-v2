package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type MobileDevicesMock struct {
	*mocks.GenericMock
}

func NewMobileDevicesMock() *MobileDevicesMock {
	return &MobileDevicesMock{
		GenericMock: mocks.NewXMLMock("MobileDevicesMock"),
	}
}

func (m *MobileDevicesMock) RegisterMocks() {
	m.RegisterListMobileDevicesMock()
	m.RegisterGetMobileDeviceByIDMock()
	m.RegisterGetMobileDeviceByNameMock()
	m.RegisterGetMobileDeviceByIDAndDataSubsetMock()
	m.RegisterGetMobileDeviceByNameAndDataSubsetMock()
	m.RegisterCreateMobileDeviceMock()
	m.RegisterUpdateMobileDeviceByIDMock()
	m.RegisterUpdateMobileDeviceByNameMock()
	m.RegisterDeleteMobileDeviceByIDMock()
	m.RegisterDeleteMobileDeviceByNameMock()
}

func (m *MobileDevicesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
}

func (m *MobileDevicesMock) RegisterListMobileDevicesMock() {
	m.Register("GET", "/JSSResource/mobiledevices", 200, "validate_list_mobile_devices.xml")
}

func (m *MobileDevicesMock) RegisterGetMobileDeviceByIDMock() {
	m.Register("GET", "/JSSResource/mobiledevices/id/1", 200, "validate_get_mobile_device.xml")
}

func (m *MobileDevicesMock) RegisterGetMobileDeviceByNameMock() {
	m.Register("GET", "/JSSResource/mobiledevices/name/iPhone-01", 200, "validate_get_mobile_device.xml")
}

func (m *MobileDevicesMock) RegisterGetMobileDeviceByIDAndDataSubsetMock() {
	m.Register("GET", "/JSSResource/mobiledevices/id/1/subset/General", 200, "validate_get_mobile_device.xml")
}

func (m *MobileDevicesMock) RegisterGetMobileDeviceByNameAndDataSubsetMock() {
	m.Register("GET", "/JSSResource/mobiledevices/name/iPhone-01/subset/General", 200, "validate_get_mobile_device.xml")
}

func (m *MobileDevicesMock) RegisterCreateMobileDeviceMock() {
	m.Register("POST", "/JSSResource/mobiledevices/id/0", 201, "validate_create_mobile_device.xml")
}

func (m *MobileDevicesMock) RegisterUpdateMobileDeviceByIDMock() {
	m.Register("PUT", "/JSSResource/mobiledevices/id/1", 200, "validate_update_mobile_device.xml")
}

func (m *MobileDevicesMock) RegisterUpdateMobileDeviceByNameMock() {
	m.Register("PUT", "/JSSResource/mobiledevices/name/iPhone-01", 200, "validate_update_mobile_device.xml")
}

func (m *MobileDevicesMock) RegisterDeleteMobileDeviceByIDMock() {
	m.Register("DELETE", "/JSSResource/mobiledevices/id/1", 200, "")
}

func (m *MobileDevicesMock) RegisterDeleteMobileDeviceByNameMock() {
	m.Register("DELETE", "/JSSResource/mobiledevices/name/iPhone-01", 200, "")
}

func (m *MobileDevicesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/mobiledevices/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

