package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type SoftwareUpdateServersMock struct {
	*mocks.GenericMock
}

func NewSoftwareUpdateServersMock() *SoftwareUpdateServersMock {
	return &SoftwareUpdateServersMock{
		GenericMock: mocks.NewXMLMock("SoftwareUpdateServersMock"),
	}
}

func (m *SoftwareUpdateServersMock) RegisterMocks() {
	m.RegisterListSoftwareUpdateServersMock()
	m.RegisterGetSoftwareUpdateServerByIDMock()
	m.RegisterGetSoftwareUpdateServerByNameMock()
	m.RegisterCreateSoftwareUpdateServerMock()
	m.RegisterUpdateSoftwareUpdateServerByIDMock()
	m.RegisterUpdateSoftwareUpdateServerByNameMock()
	m.RegisterDeleteSoftwareUpdateServerByIDMock()
	m.RegisterDeleteSoftwareUpdateServerByNameMock()
}

func (m *SoftwareUpdateServersMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *SoftwareUpdateServersMock) RegisterListSoftwareUpdateServersMock() {
	m.Register("GET", "/JSSResource/softwareupdateservers", 200, "validate_list_software_update_servers.xml")
}

func (m *SoftwareUpdateServersMock) RegisterGetSoftwareUpdateServerByIDMock() {
	m.Register("GET", "/JSSResource/softwareupdateservers/id/1", 200, "validate_get_software_update_server.xml")
}

func (m *SoftwareUpdateServersMock) RegisterGetSoftwareUpdateServerByNameMock() {
	m.Register("GET", "/JSSResource/softwareupdateservers/name/Primary SUS", 200, "validate_get_software_update_server.xml")
}

func (m *SoftwareUpdateServersMock) RegisterCreateSoftwareUpdateServerMock() {
	m.Register("POST", "/JSSResource/softwareupdateservers/id/0", 201, "validate_create_software_update_server.xml")
}

func (m *SoftwareUpdateServersMock) RegisterUpdateSoftwareUpdateServerByIDMock() {
	m.Register("PUT", "/JSSResource/softwareupdateservers/id/1", 200, "validate_update_software_update_server.xml")
}

func (m *SoftwareUpdateServersMock) RegisterUpdateSoftwareUpdateServerByNameMock() {
	m.Register("PUT", "/JSSResource/softwareupdateservers/name/Primary SUS", 200, "validate_update_software_update_server.xml")
}

func (m *SoftwareUpdateServersMock) RegisterDeleteSoftwareUpdateServerByIDMock() {
	m.Register("DELETE", "/JSSResource/softwareupdateservers/id/1", 200, "")
}

func (m *SoftwareUpdateServersMock) RegisterDeleteSoftwareUpdateServerByNameMock() {
	m.Register("DELETE", "/JSSResource/softwareupdateservers/name/Primary SUS", 200, "")
}

func (m *SoftwareUpdateServersMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/softwareupdateservers/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *SoftwareUpdateServersMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/softwareupdateservers/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A software update server with that name already exists")
}

