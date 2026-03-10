package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type MobileDeviceExtensionAttributesMock struct {
	*mocks.GenericMock
}

func NewMobileDeviceExtensionAttributesMock() *MobileDeviceExtensionAttributesMock {
	return &MobileDeviceExtensionAttributesMock{
		GenericMock: mocks.NewJSONMock("MobileDeviceExtensionAttributesMock"),
	}
}

func (m *MobileDeviceExtensionAttributesMock) RegisterListMock() {
	m.Register("GET", "/api/v1/mobile-device-extension-attributes", 200, "validate_list.json")
}

func (m *MobileDeviceExtensionAttributesMock) RegisterGetMock() {
	m.Register("GET", "/api/v1/mobile-device-extension-attributes/1", 200, "validate_get.json")
}

func (m *MobileDeviceExtensionAttributesMock) RegisterCreateMock() {
	m.Register("POST", "/api/v1/mobile-device-extension-attributes", 201, "validate_create.json")
}

func (m *MobileDeviceExtensionAttributesMock) RegisterUpdateMock() {
	m.Register("PUT", "/api/v1/mobile-device-extension-attributes/1", 200, "validate_update.json")
}

func (m *MobileDeviceExtensionAttributesMock) RegisterDeleteMock() {
	m.Register("DELETE", "/api/v1/mobile-device-extension-attributes/1", 204, "")
}

func (m *MobileDeviceExtensionAttributesMock) RegisterDeleteMultipleMock() {
	m.Register("POST", "/api/v1/mobile-device-extension-attributes/delete-multiple", 204, "")
}

func (m *MobileDeviceExtensionAttributesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/mobile-device-extension-attributes/999", 404, "error_not_found.json", "")
}

func (m *MobileDeviceExtensionAttributesMock) RegisterGetHistoryMock() {
	m.Register("GET", "/api/v1/mobile-device-extension-attributes/1/history", 200, "validate_history.json")
}

func (m *MobileDeviceExtensionAttributesMock) RegisterAddHistoryNoteMock() {
	m.Register("POST", "/api/v1/mobile-device-extension-attributes/1/history", 201, "")
}
