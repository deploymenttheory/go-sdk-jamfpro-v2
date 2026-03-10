package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type MacApplicationsMock struct {
	*mocks.GenericMock
}

func NewMacApplicationsMock() *MacApplicationsMock {
	return &MacApplicationsMock{
		GenericMock: mocks.NewXMLMock("MacApplicationsMock"),
	}
}

func (m *MacApplicationsMock) RegisterMocks() {
	m.RegisterListMacApplicationsMock()
	m.RegisterGetMacApplicationByIDMock()
	m.RegisterGetMacApplicationByNameMock()
	m.RegisterGetMacApplicationByIDAndSubsetMock()
	m.RegisterGetMacApplicationByNameAndSubsetMock()
	m.RegisterCreateMacApplicationMock()
	m.RegisterUpdateMacApplicationByIDMock()
	m.RegisterUpdateMacApplicationByNameMock()
	m.RegisterDeleteMacApplicationByIDMock()
	m.RegisterDeleteMacApplicationByNameMock()
}

func (m *MacApplicationsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *MacApplicationsMock) RegisterListMacApplicationsMock() {
	m.Register("GET", "/JSSResource/macapplications", 200, "validate_list_mac_applications.xml")
}

func (m *MacApplicationsMock) RegisterGetMacApplicationByIDMock() {
	m.Register("GET", "/JSSResource/macapplications/id/1", 200, "validate_get_mac_application.xml")
}

func (m *MacApplicationsMock) RegisterGetMacApplicationByNameMock() {
	m.Register("GET", "/JSSResource/macapplications/name/Sample Mac App", 200, "validate_get_mac_application.xml")
}

func (m *MacApplicationsMock) RegisterGetMacApplicationByIDAndSubsetMock() {
	m.Register("GET", "/JSSResource/macapplications/id/1/subset/General", 200, "validate_get_mac_application.xml")
}

func (m *MacApplicationsMock) RegisterGetMacApplicationByNameAndSubsetMock() {
	m.Register("GET", "/JSSResource/macapplications/name/Sample Mac App/subset/General", 200, "validate_get_mac_application.xml")
}

func (m *MacApplicationsMock) RegisterCreateMacApplicationMock() {
	m.Register("POST", "/JSSResource/macapplications/id/0", 201, "validate_create_mac_application.xml")
}

func (m *MacApplicationsMock) RegisterUpdateMacApplicationByIDMock() {
	m.Register("PUT", "/JSSResource/macapplications/id/1", 200, "validate_update_mac_application.xml")
}

func (m *MacApplicationsMock) RegisterUpdateMacApplicationByNameMock() {
	m.Register("PUT", "/JSSResource/macapplications/name/Sample Mac App", 200, "validate_update_mac_application.xml")
}

func (m *MacApplicationsMock) RegisterDeleteMacApplicationByIDMock() {
	m.Register("DELETE", "/JSSResource/macapplications/id/1", 200, "")
}

func (m *MacApplicationsMock) RegisterDeleteMacApplicationByNameMock() {
	m.Register("DELETE", "/JSSResource/macapplications/name/Sample Mac App", 200, "")
}

func (m *MacApplicationsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/macapplications/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *MacApplicationsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/macapplications/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A Mac application with that name already exists")
}

