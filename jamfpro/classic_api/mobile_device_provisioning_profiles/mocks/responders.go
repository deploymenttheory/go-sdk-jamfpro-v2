package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type MobileDeviceProvisioningProfilesMock struct {
	*mocks.GenericMock
}

func NewMobileDeviceProvisioningProfilesMock() *MobileDeviceProvisioningProfilesMock {
	return &MobileDeviceProvisioningProfilesMock{
		GenericMock: mocks.NewXMLMock("MobileDeviceProvisioningProfilesMock"),
	}
}

func (m *MobileDeviceProvisioningProfilesMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetByIDMock()
	m.RegisterGetByNameMock()
	m.RegisterGetByUUIDMock()
	m.RegisterCreateByIDMock()
	m.RegisterCreateByNameMock()
	m.RegisterCreateByUUIDMock()
	m.RegisterUpdateByIDMock()
	m.RegisterUpdateByNameMock()
	m.RegisterUpdateByUUIDMock()
	m.RegisterDeleteByIDMock()
	m.RegisterDeleteByNameMock()
	m.RegisterDeleteByUUIDMock()
}

func (m *MobileDeviceProvisioningProfilesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *MobileDeviceProvisioningProfilesMock) RegisterListMock() {
	m.Register("GET", "/JSSResource/mobiledeviceprovisioningprofiles", 200, "validate_list_mobile_device_provisioning_profiles.xml")
}

func (m *MobileDeviceProvisioningProfilesMock) RegisterGetByIDMock() {
	m.Register("GET", "/JSSResource/mobiledeviceprovisioningprofiles/id/1", 200, "validate_get_mobile_device_provisioning_profile.xml")
}

func (m *MobileDeviceProvisioningProfilesMock) RegisterGetByNameMock() {
	m.Register("GET", "/JSSResource/mobiledeviceprovisioningprofiles/name/Test Provisioning Profile", 200, "validate_get_mobile_device_provisioning_profile.xml")
}

func (m *MobileDeviceProvisioningProfilesMock) RegisterGetByUUIDMock() {
	m.Register("GET", "/JSSResource/mobiledeviceprovisioningprofiles/uuid/550e8400-e29b-41d4-a716-446655440000", 200, "validate_get_mobile_device_provisioning_profile.xml")
}

func (m *MobileDeviceProvisioningProfilesMock) RegisterCreateByIDMock() {
	m.Register("POST", "/JSSResource/mobiledeviceprovisioningprofiles/id/0", 201, "validate_create_mobile_device_provisioning_profile.xml")
}

func (m *MobileDeviceProvisioningProfilesMock) RegisterCreateByNameMock() {
	m.Register("POST", "/JSSResource/mobiledeviceprovisioningprofiles/name/New Profile", 201, "validate_create_mobile_device_provisioning_profile.xml")
}

func (m *MobileDeviceProvisioningProfilesMock) RegisterCreateByUUIDMock() {
	m.Register("POST", "/JSSResource/mobiledeviceprovisioningprofiles/uuid/550e8400-e29b-41d4-a716-446655440001", 201, "validate_create_mobile_device_provisioning_profile.xml")
}

func (m *MobileDeviceProvisioningProfilesMock) RegisterUpdateByIDMock() {
	m.Register("PUT", "/JSSResource/mobiledeviceprovisioningprofiles/id/1", 200, "validate_update_mobile_device_provisioning_profile.xml")
}

func (m *MobileDeviceProvisioningProfilesMock) RegisterUpdateByNameMock() {
	m.Register("PUT", "/JSSResource/mobiledeviceprovisioningprofiles/name/Test Provisioning Profile", 200, "validate_update_mobile_device_provisioning_profile.xml")
}

func (m *MobileDeviceProvisioningProfilesMock) RegisterUpdateByUUIDMock() {
	m.Register("PUT", "/JSSResource/mobiledeviceprovisioningprofiles/uuid/550e8400-e29b-41d4-a716-446655440000", 200, "validate_update_mobile_device_provisioning_profile.xml")
}

func (m *MobileDeviceProvisioningProfilesMock) RegisterDeleteByIDMock() {
	m.Register("DELETE", "/JSSResource/mobiledeviceprovisioningprofiles/id/1", 200, "")
}

func (m *MobileDeviceProvisioningProfilesMock) RegisterDeleteByNameMock() {
	m.Register("DELETE", "/JSSResource/mobiledeviceprovisioningprofiles/name/Test Provisioning Profile", 200, "")
}

func (m *MobileDeviceProvisioningProfilesMock) RegisterDeleteByUUIDMock() {
	m.Register("DELETE", "/JSSResource/mobiledeviceprovisioningprofiles/uuid/550e8400-e29b-41d4-a716-446655440000", 200, "")
}

func (m *MobileDeviceProvisioningProfilesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/mobiledeviceprovisioningprofiles/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *MobileDeviceProvisioningProfilesMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/mobiledeviceprovisioningprofiles/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A mobile device provisioning profile with that name already exists")
}

