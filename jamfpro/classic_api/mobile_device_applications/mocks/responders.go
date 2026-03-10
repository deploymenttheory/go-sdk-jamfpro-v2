package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type MobileDeviceApplicationsMock struct {
	*mocks.GenericMock
}

func NewMobileDeviceApplicationsMock() *MobileDeviceApplicationsMock {
	return &MobileDeviceApplicationsMock{
		GenericMock: mocks.NewXMLMock("MobileDeviceApplicationsMock"),
	}
}

func (m *MobileDeviceApplicationsMock) RegisterMocks() {
	m.RegisterListMobileDeviceApplicationsMock()
	m.RegisterGetMobileDeviceApplicationByIDMock()
	m.RegisterGetMobileDeviceApplicationByNameMock()
	m.RegisterGetMobileDeviceApplicationByBundleIDMock()
	m.RegisterGetMobileDeviceApplicationByBundleIDAndVersionMock()
	m.RegisterGetMobileDeviceApplicationByIDAndSubsetMock()
	m.RegisterGetMobileDeviceApplicationByNameAndSubsetMock()
	m.RegisterCreateMobileDeviceApplicationMock()
	m.RegisterUpdateMobileDeviceApplicationByIDMock()
	m.RegisterUpdateMobileDeviceApplicationByNameMock()
	m.RegisterUpdateMobileDeviceApplicationByBundleIDMock()
	m.RegisterUpdateMobileDeviceApplicationByIDAndVersionMock()
	m.RegisterDeleteMobileDeviceApplicationByIDMock()
	m.RegisterDeleteMobileDeviceApplicationByNameMock()
	m.RegisterDeleteMobileDeviceApplicationByBundleIDMock()
	m.RegisterDeleteMobileDeviceApplicationByBundleIDAndVersionMock()
}

func (m *MobileDeviceApplicationsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *MobileDeviceApplicationsMock) RegisterListMobileDeviceApplicationsMock() {
	m.Register("GET", "/JSSResource/mobiledeviceapplications", 200, "validate_list_mobile_device_applications.xml")
}

func (m *MobileDeviceApplicationsMock) RegisterGetMobileDeviceApplicationByIDMock() {
	m.Register("GET", "/JSSResource/mobiledeviceapplications/id/1", 200, "validate_get_mobile_device_application.xml")
}

func (m *MobileDeviceApplicationsMock) RegisterGetMobileDeviceApplicationByNameMock() {
	m.Register("GET", "/JSSResource/mobiledeviceapplications/name/Sample iOS App", 200, "validate_get_mobile_device_application.xml")
}

func (m *MobileDeviceApplicationsMock) RegisterGetMobileDeviceApplicationByBundleIDMock() {
	m.Register("GET", "/JSSResource/mobiledeviceapplications/bundleid/com.example.app", 200, "validate_get_mobile_device_application.xml")
}

func (m *MobileDeviceApplicationsMock) RegisterGetMobileDeviceApplicationByBundleIDAndVersionMock() {
	m.Register("GET", "/JSSResource/mobiledeviceapplications/bundleid/com.example.app/version/1.0", 200, "validate_get_mobile_device_application.xml")
}

func (m *MobileDeviceApplicationsMock) RegisterGetMobileDeviceApplicationByIDAndSubsetMock() {
	m.Register("GET", "/JSSResource/mobiledeviceapplications/id/1/subset/General", 200, "validate_get_mobile_device_application.xml")
}

func (m *MobileDeviceApplicationsMock) RegisterGetMobileDeviceApplicationByNameAndSubsetMock() {
	m.Register("GET", "/JSSResource/mobiledeviceapplications/name/Sample iOS App/subset/General", 200, "validate_get_mobile_device_application.xml")
}

func (m *MobileDeviceApplicationsMock) RegisterCreateMobileDeviceApplicationMock() {
	m.Register("POST", "/JSSResource/mobiledeviceapplications/id/0", 201, "validate_create_mobile_device_application.xml")
}

func (m *MobileDeviceApplicationsMock) RegisterUpdateMobileDeviceApplicationByIDMock() {
	m.Register("PUT", "/JSSResource/mobiledeviceapplications/id/1", 200, "validate_update_mobile_device_application.xml")
}

func (m *MobileDeviceApplicationsMock) RegisterUpdateMobileDeviceApplicationByNameMock() {
	m.Register("PUT", "/JSSResource/mobiledeviceapplications/name/Sample iOS App", 200, "validate_update_mobile_device_application.xml")
}

func (m *MobileDeviceApplicationsMock) RegisterUpdateMobileDeviceApplicationByBundleIDMock() {
	m.Register("PUT", "/JSSResource/mobiledeviceapplications/bundleid/com.example.app", 200, "validate_update_mobile_device_application.xml")
}

func (m *MobileDeviceApplicationsMock) RegisterUpdateMobileDeviceApplicationByIDAndVersionMock() {
	m.Register("PUT", "/JSSResource/mobiledeviceapplications/id/1/version/1.0", 200, "validate_update_mobile_device_application.xml")
}

func (m *MobileDeviceApplicationsMock) RegisterDeleteMobileDeviceApplicationByIDMock() {
	m.Register("DELETE", "/JSSResource/mobiledeviceapplications/id/1", 200, "")
}

func (m *MobileDeviceApplicationsMock) RegisterDeleteMobileDeviceApplicationByNameMock() {
	m.Register("DELETE", "/JSSResource/mobiledeviceapplications/name/Sample iOS App", 200, "")
}

func (m *MobileDeviceApplicationsMock) RegisterDeleteMobileDeviceApplicationByBundleIDMock() {
	m.Register("DELETE", "/JSSResource/mobiledeviceapplications/bundleid/com.example.app", 200, "")
}

func (m *MobileDeviceApplicationsMock) RegisterDeleteMobileDeviceApplicationByBundleIDAndVersionMock() {
	m.Register("DELETE", "/JSSResource/mobiledeviceapplications/bundleid/com.example.app/version/1.0", 200, "")
}

func (m *MobileDeviceApplicationsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/mobiledeviceapplications/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *MobileDeviceApplicationsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/mobiledeviceapplications/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A mobile device application with that name already exists")
}

