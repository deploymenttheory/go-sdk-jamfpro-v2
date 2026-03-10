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
}

func (m *NotificationsMock) RegisterListErrorMock() {
	m.RegisterError("GET", "/api/v1/notifications", 500, "", "Jamf Pro API error (500) [INTERNAL]: server error")
}
