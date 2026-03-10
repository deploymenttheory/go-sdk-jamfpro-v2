package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type MobileDeviceAppsMock struct {
	*mocks.GenericMock
}

func NewMobileDeviceAppsMock() *MobileDeviceAppsMock {
	return &MobileDeviceAppsMock{
		GenericMock: mocks.NewJSONMock("MobileDeviceAppsMock"),
	}
}

func (m *MobileDeviceAppsMock) RegisterReinstallAppConfigMock() {
	m.Register("POST", "/api/v1/mobile-device-apps/reinstall-app-config", 204, "")
}

func (m *MobileDeviceAppsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("POST", "/api/v1/mobile-device-apps/reinstall-app-config", 404, "error_not_found.json", "")
}

func (m *MobileDeviceAppsMock) RegisterReinstallAppConfigErrorMock() {
	m.RegisterError("POST", "/api/v1/mobile-device-apps/reinstall-app-config", 500, "error_internal.json", "")
}
