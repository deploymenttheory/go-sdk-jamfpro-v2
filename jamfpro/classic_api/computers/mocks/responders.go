package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ComputersMock struct {
	*mocks.GenericMock
}

func NewComputersMock() *ComputersMock {
	return &ComputersMock{
		GenericMock: mocks.NewXMLMock("ComputersMock"),
	}
}

func (m *ComputersMock) RegisterMocks() {
	m.RegisterListComputersMock()
	m.RegisterGetComputerByIDMock()
	m.RegisterGetComputerByNameMock()
	m.RegisterCreateComputerMock()
	m.RegisterUpdateComputerByIDMock()
	m.RegisterUpdateComputerByNameMock()
	m.RegisterDeleteComputerByIDMock()
	m.RegisterDeleteComputerByNameMock()
}

func (m *ComputersMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
}

func (m *ComputersMock) RegisterListComputersMock() {
	m.Register("GET", "/JSSResource/computers", 200, "validate_list_computers.xml")
}

func (m *ComputersMock) RegisterGetComputerByIDMock() {
	m.Register("GET", "/JSSResource/computers/id/1", 200, "validate_get_computer.xml")
}

func (m *ComputersMock) RegisterGetComputerByNameMock() {
	m.Register("GET", "/JSSResource/computers/name/MacBook-Pro-01", 200, "validate_get_computer.xml")
}

func (m *ComputersMock) RegisterCreateComputerMock() {
	m.Register("POST", "/JSSResource/computers", 201, "validate_create_computer.xml")
}

func (m *ComputersMock) RegisterUpdateComputerByIDMock() {
	m.Register("PUT", "/JSSResource/computers/id/1", 200, "validate_update_computer.xml")
}

func (m *ComputersMock) RegisterUpdateComputerByNameMock() {
	m.Register("PUT", "/JSSResource/computers/name/MacBook-Pro-01", 200, "validate_update_computer.xml")
}

func (m *ComputersMock) RegisterDeleteComputerByIDMock() {
	m.Register("DELETE", "/JSSResource/computers/id/1", 200, "")
}

func (m *ComputersMock) RegisterDeleteComputerByNameMock() {
	m.Register("DELETE", "/JSSResource/computers/name/MacBook-Pro-01", 200, "")
}

func (m *ComputersMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/computers/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

