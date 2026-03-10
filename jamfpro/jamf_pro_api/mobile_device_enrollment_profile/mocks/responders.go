package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type MobileDeviceEnrollmentProfileMock struct {
	*mocks.GenericMock
}

func NewMobileDeviceEnrollmentProfileMock() *MobileDeviceEnrollmentProfileMock {
	return &MobileDeviceEnrollmentProfileMock{
		GenericMock: mocks.NewJSONMock("MobileDeviceEnrollmentProfileMock"),
	}
}

func (m *MobileDeviceEnrollmentProfileMock) RegisterGetDownloadProfileMock(id string) {
	path := "/api/v1/mobile-device-enrollment-profile/" + id + "/download-profile"
	m.Register("GET", path, 200, "validate_download_profile.bin")
}

func (m *MobileDeviceEnrollmentProfileMock) RegisterNotFoundErrorMock(id string) {
	path := "/api/v1/mobile-device-enrollment-profile/" + id + "/download-profile"
	m.RegisterError("GET", path, 404, "error_not_found.json", "")
}

func (m *MobileDeviceEnrollmentProfileMock) RegisterGetDownloadProfileErrorMock(id string) {
	path := "/api/v1/mobile-device-enrollment-profile/" + id + "/download-profile"
	m.RegisterError("GET", path, 500, "error_internal.json", "")
}
