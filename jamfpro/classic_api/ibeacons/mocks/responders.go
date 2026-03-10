package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type IBeaconsMock struct {
	*mocks.GenericMock
}

func NewIBeaconsMock() *IBeaconsMock {
	return &IBeaconsMock{
		GenericMock: mocks.NewXMLMock("IBeaconsMock"),
	}
}

func (m *IBeaconsMock) RegisterMocks() {
	m.RegisterListIBeaconsMock()
	m.RegisterGetIBeaconByIDMock()
	m.RegisterGetIBeaconByNameMock()
	m.RegisterCreateIBeaconMock()
	m.RegisterUpdateIBeaconByIDMock()
	m.RegisterUpdateIBeaconByNameMock()
	m.RegisterDeleteIBeaconByIDMock()
	m.RegisterDeleteIBeaconByNameMock()
}

func (m *IBeaconsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *IBeaconsMock) RegisterListIBeaconsMock() {
	m.Register("GET", "/JSSResource/ibeacons", 200, "validate_list_ibeacons.xml")
}

func (m *IBeaconsMock) RegisterGetIBeaconByIDMock() {
	m.Register("GET", "/JSSResource/ibeacons/id/1", 200, "validate_get_ibeacon.xml")
}

func (m *IBeaconsMock) RegisterGetIBeaconByNameMock() {
	m.Register("GET", "/JSSResource/ibeacons/name/Lobby Beacon", 200, "validate_get_ibeacon.xml")
}

func (m *IBeaconsMock) RegisterCreateIBeaconMock() {
	m.Register("POST", "/JSSResource/ibeacons/id/0", 201, "validate_create_ibeacon.xml")
}

func (m *IBeaconsMock) RegisterUpdateIBeaconByIDMock() {
	m.Register("PUT", "/JSSResource/ibeacons/id/1", 200, "validate_update_ibeacon.xml")
}

func (m *IBeaconsMock) RegisterUpdateIBeaconByNameMock() {
	m.Register("PUT", "/JSSResource/ibeacons/name/Lobby Beacon", 200, "validate_update_ibeacon.xml")
}

func (m *IBeaconsMock) RegisterDeleteIBeaconByIDMock() {
	m.Register("DELETE", "/JSSResource/ibeacons/id/1", 200, "")
}

func (m *IBeaconsMock) RegisterDeleteIBeaconByNameMock() {
	m.Register("DELETE", "/JSSResource/ibeacons/name/Lobby Beacon", 200, "")
}

func (m *IBeaconsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/ibeacons/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *IBeaconsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/ibeacons/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): An iBeacon with that name already exists")
}

