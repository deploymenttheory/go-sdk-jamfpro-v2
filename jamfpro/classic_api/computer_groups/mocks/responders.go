package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ComputerGroupsMock struct {
	*mocks.GenericMock
}

func NewComputerGroupsMock() *ComputerGroupsMock {
	return &ComputerGroupsMock{
		GenericMock: mocks.NewXMLMock("ComputerGroupsMock"),
	}
}

func (m *ComputerGroupsMock) RegisterMocks() {
	m.RegisterListComputerGroupsMock()
	m.RegisterGetComputerGroupByIDMock()
	m.RegisterGetComputerGroupByNameMock()
	m.RegisterCreateComputerGroupMock()
	m.RegisterUpdateComputerGroupByIDMock()
	m.RegisterUpdateComputerGroupByNameMock()
	m.RegisterDeleteComputerGroupByIDMock()
	m.RegisterDeleteComputerGroupByNameMock()
}

func (m *ComputerGroupsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *ComputerGroupsMock) RegisterListComputerGroupsMock() {
	m.Register("GET", "/JSSResource/computergroups", 200, "validate_list_computer_groups.xml")
}

func (m *ComputerGroupsMock) RegisterGetComputerGroupByIDMock() {
	m.Register("GET", "/JSSResource/computergroups/id/1", 200, "validate_get_computer_group.xml")
}

func (m *ComputerGroupsMock) RegisterGetComputerGroupByNameMock() {
	m.Register("GET", "/JSSResource/computergroups/name/All Managed Clients", 200, "validate_get_computer_group.xml")
}

func (m *ComputerGroupsMock) RegisterCreateComputerGroupMock() {
	m.Register("POST", "/JSSResource/computergroups/id/0", 201, "validate_create_computer_group.xml")
}

func (m *ComputerGroupsMock) RegisterUpdateComputerGroupByIDMock() {
	m.Register("PUT", "/JSSResource/computergroups/id/1", 200, "validate_update_computer_group.xml")
}

func (m *ComputerGroupsMock) RegisterUpdateComputerGroupByNameMock() {
	m.Register("PUT", "/JSSResource/computergroups/name/All Managed Clients", 200, "validate_update_computer_group.xml")
}

func (m *ComputerGroupsMock) RegisterDeleteComputerGroupByIDMock() {
	m.Register("DELETE", "/JSSResource/computergroups/id/1", 200, "")
}

func (m *ComputerGroupsMock) RegisterDeleteComputerGroupByNameMock() {
	m.Register("DELETE", "/JSSResource/computergroups/name/All Managed Clients", 200, "")
}

func (m *ComputerGroupsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/computergroups/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *ComputerGroupsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/computergroups/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A computer group with that name already exists")
}

