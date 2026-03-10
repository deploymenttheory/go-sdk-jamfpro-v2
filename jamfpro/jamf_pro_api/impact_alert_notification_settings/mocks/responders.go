package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ImpactAlertNotificationSettingsMock struct {
	*mocks.GenericMock
}

func NewImpactAlertNotificationSettingsMock() *ImpactAlertNotificationSettingsMock {
	return &ImpactAlertNotificationSettingsMock{
		GenericMock: mocks.NewJSONMock("ImpactAlertNotificationSettingsMock"),
	}
}

func (m *ImpactAlertNotificationSettingsMock) RegisterMocks() {
	m.RegisterGetMock()
	m.RegisterUpdateMock()
}

func (m *ImpactAlertNotificationSettingsMock) RegisterGetMock() {
	m.Register("GET", "/api/v1/impact-alert-notification-settings", 200, "validate_get.json")
}

func (m *ImpactAlertNotificationSettingsMock) RegisterUpdateMock() {
	// Update returns 204 No Content
	m.Register("PUT", "/api/v1/impact-alert-notification-settings", 204, "")
}

func (m *ImpactAlertNotificationSettingsMock) RegisterGetErrorMock() {
	m.RegisterError("GET", "/api/v1/impact-alert-notification-settings", 404, "error_not_found.json", "")
}

func (m *ImpactAlertNotificationSettingsMock) RegisterUpdateErrorMock() {
	m.RegisterError("PUT", "/api/v1/impact-alert-notification-settings", 400, "error_update.json", "")
}

