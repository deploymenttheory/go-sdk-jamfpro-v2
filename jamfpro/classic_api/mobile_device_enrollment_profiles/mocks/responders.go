package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type MobileDeviceEnrollmentProfilesMock struct {
	*mocks.GenericMock
}

func NewMobileDeviceEnrollmentProfilesMock() *MobileDeviceEnrollmentProfilesMock {
	return &MobileDeviceEnrollmentProfilesMock{
		GenericMock: mocks.NewXMLMock("MobileDeviceEnrollmentProfilesMock"),
	}
}

func (m *MobileDeviceEnrollmentProfilesMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetByIDMock()
	m.RegisterGetByNameMock()
	m.RegisterGetByInvitationMock()
	m.RegisterGetByIDWithSubsetMock()
	m.RegisterGetByNameWithSubsetMock()
	m.RegisterCreateMock()
	m.RegisterUpdateByIDMock()
	m.RegisterUpdateByNameMock()
	m.RegisterUpdateByInvitationMock()
	m.RegisterDeleteByIDMock()
	m.RegisterDeleteByNameMock()
	m.RegisterDeleteByInvitationMock()
}

func (m *MobileDeviceEnrollmentProfilesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
}

func (m *MobileDeviceEnrollmentProfilesMock) RegisterListMock() {
	m.Register("GET", "/JSSResource/mobiledeviceenrollmentprofiles", 200, "validate_list_mobile_device_enrollment_profiles.xml")
}

func (m *MobileDeviceEnrollmentProfilesMock) RegisterGetByIDMock() {
	m.Register("GET", "/JSSResource/mobiledeviceenrollmentprofiles/id/1", 200, "validate_get_mobile_device_enrollment_profile.xml")
}

func (m *MobileDeviceEnrollmentProfilesMock) RegisterGetByNameMock() {
	m.Register("GET", "/JSSResource/mobiledeviceenrollmentprofiles/name/Test Enrollment Profile", 200, "validate_get_mobile_device_enrollment_profile.xml")
}

func (m *MobileDeviceEnrollmentProfilesMock) RegisterGetByInvitationMock() {
	m.Register("GET", "/JSSResource/mobiledeviceenrollmentprofiles/invitation/1234567890.123456", 200, "validate_get_mobile_device_enrollment_profile.xml")
}

func (m *MobileDeviceEnrollmentProfilesMock) RegisterGetByIDWithSubsetMock() {
	m.Register("GET", "/JSSResource/mobiledeviceenrollmentprofiles/id/1/subset/General", 200, "validate_get_mobile_device_enrollment_profile.xml")
}

func (m *MobileDeviceEnrollmentProfilesMock) RegisterGetByNameWithSubsetMock() {
	m.Register("GET", "/JSSResource/mobiledeviceenrollmentprofiles/name/Test Enrollment Profile/subset/General", 200, "validate_get_mobile_device_enrollment_profile.xml")
}

func (m *MobileDeviceEnrollmentProfilesMock) RegisterCreateMock() {
	m.Register("POST", "/JSSResource/mobiledeviceenrollmentprofiles/id/0", 201, "validate_create_mobile_device_enrollment_profile.xml")
}

func (m *MobileDeviceEnrollmentProfilesMock) RegisterUpdateByIDMock() {
	m.Register("PUT", "/JSSResource/mobiledeviceenrollmentprofiles/id/1", 200, "validate_update_mobile_device_enrollment_profile.xml")
}

func (m *MobileDeviceEnrollmentProfilesMock) RegisterUpdateByNameMock() {
	m.Register("PUT", "/JSSResource/mobiledeviceenrollmentprofiles/name/Test Enrollment Profile", 200, "validate_update_mobile_device_enrollment_profile.xml")
}

func (m *MobileDeviceEnrollmentProfilesMock) RegisterUpdateByInvitationMock() {
	m.Register("PUT", "/JSSResource/mobiledeviceenrollmentprofiles/invitation/1234567890.123456", 200, "validate_update_mobile_device_enrollment_profile.xml")
}

func (m *MobileDeviceEnrollmentProfilesMock) RegisterDeleteByIDMock() {
	m.Register("DELETE", "/JSSResource/mobiledeviceenrollmentprofiles/id/1", 200, "")
}

func (m *MobileDeviceEnrollmentProfilesMock) RegisterDeleteByNameMock() {
	m.Register("DELETE", "/JSSResource/mobiledeviceenrollmentprofiles/name/Test Enrollment Profile", 200, "")
}

func (m *MobileDeviceEnrollmentProfilesMock) RegisterDeleteByInvitationMock() {
	m.Register("DELETE", "/JSSResource/mobiledeviceenrollmentprofiles/invitation/1234567890.123456", 200, "")
}

func (m *MobileDeviceEnrollmentProfilesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/mobiledeviceenrollmentprofiles/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

