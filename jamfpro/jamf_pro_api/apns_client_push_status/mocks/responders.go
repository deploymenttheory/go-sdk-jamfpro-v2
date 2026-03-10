package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type APNSClientPushStatusMock struct {
	*mocks.GenericMock
}

func NewAPNSClientPushStatusMock() *APNSClientPushStatusMock {
	return &APNSClientPushStatusMock{
		GenericMock: mocks.NewJSONMock("APNSClientPushStatusMock"),
	}
}

func (m *APNSClientPushStatusMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterEnableAllClientsMock()
	m.RegisterGetEnableAllClientsStatusMock()
	m.RegisterEnableClientMock()
}

func (m *APNSClientPushStatusMock) RegisterListMock() {
	m.Register("GET", "/api/v1/apns-client-push-status", 200, "validate_list.json")
}

func (m *APNSClientPushStatusMock) RegisterEnableAllClientsMock() {
	m.Register("POST", "/api/v1/apns-client-push-status/enable-all-clients", 202, "")
}

func (m *APNSClientPushStatusMock) RegisterGetEnableAllClientsStatusMock() {
	m.Register("GET", "/api/v1/apns-client-push-status/enable-all-clients/status", 200, "validate_get_enable_all_clients_status.json")
}

func (m *APNSClientPushStatusMock) RegisterEnableClientMock() {
	m.Register("POST", "/api/v1/apns-client-push-status/enable-client", 204, "")
}
