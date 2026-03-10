package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type DeclarativeDeviceManagementMock struct {
	*mocks.GenericMock
}

func NewDeclarativeDeviceManagementMock() *DeclarativeDeviceManagementMock {
	return &DeclarativeDeviceManagementMock{
		GenericMock: mocks.NewJSONMock("DeclarativeDeviceManagementMock"),
	}
}

func (m *DeclarativeDeviceManagementMock) RegisterForceSyncMock(clientManagementID string) {
	m.Register("POST", "/api/v1/ddm/"+clientManagementID+"/sync", 204, "")
}

func (m *DeclarativeDeviceManagementMock) RegisterGetStatusItemsMock(clientManagementID string) {
	m.Register("GET", "/api/v1/ddm/"+clientManagementID+"/status-items", 200, "validate_status_items.json")
}

func (m *DeclarativeDeviceManagementMock) RegisterGetStatusItemByKeyMock(clientManagementID, key string) {
	m.Register("GET", "/api/v1/ddm/"+clientManagementID+"/status-items/"+key, 200, "validate_status_item.json")
}
