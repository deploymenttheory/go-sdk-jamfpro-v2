package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type NotificationsMock struct {
	*mocks.GenericMock
}

func NewNotificationsMock() *NotificationsMock {
	return &NotificationsMock{
		GenericMock: mocks.NewJSONMock("NotificationsMock"),
	}
}

func (m *NotificationsMock) RegisterMocks() {
	m.Register("GET", "/api/v1/notifications", 200, "validate_list.json")
	m.Register("DELETE", "/api/v1/notifications/APNS_CERT_REVOKED/1", 204, "")
}

func (m *NotificationsMock) RegisterListErrorMock() {
	m.RegisterError("GET", "/api/v1/notifications", 500, "", "Jamf Pro API error (500) [INTERNAL]: server error")
}

func (m *NotificationsMock) RegisterDeleteErrorMock() {
	m.RegisterError("DELETE", "/api/v1/notifications/APNS_CERT_REVOKED/1", 500, "", "Jamf Pro API error (500) [INTERNAL]: server error")
}
