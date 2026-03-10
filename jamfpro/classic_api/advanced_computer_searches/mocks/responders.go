package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type AdvancedComputerSearchesMock struct {
	*mocks.GenericMock
}

func NewAdvancedComputerSearchesMock() *AdvancedComputerSearchesMock {
	return &AdvancedComputerSearchesMock{
		GenericMock: mocks.NewXMLMock("AdvancedComputerSearchesMock"),
	}
}

func (m *AdvancedComputerSearchesMock) RegisterMocks() {
	m.RegisterListAdvancedComputerSearchesMock()
	m.RegisterGetAdvancedComputerSearchByIDMock()
	m.RegisterGetAdvancedComputerSearchByNameMock()
	m.RegisterCreateAdvancedComputerSearchMock()
	m.RegisterUpdateAdvancedComputerSearchByIDMock()
	m.RegisterUpdateAdvancedComputerSearchByNameMock()
	m.RegisterDeleteAdvancedComputerSearchByIDMock()
	m.RegisterDeleteAdvancedComputerSearchByNameMock()
}

func (m *AdvancedComputerSearchesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *AdvancedComputerSearchesMock) RegisterListAdvancedComputerSearchesMock() {
	m.Register("GET", "/JSSResource/advancedcomputersearches", 200, "validate_list_advanced_computer_searches.xml")
}

func (m *AdvancedComputerSearchesMock) RegisterGetAdvancedComputerSearchByIDMock() {
	m.Register("GET", "/JSSResource/advancedcomputersearches/id/1", 200, "validate_get_advanced_computer_search.xml")
}

func (m *AdvancedComputerSearchesMock) RegisterGetAdvancedComputerSearchByNameMock() {
	m.Register("GET", "/JSSResource/advancedcomputersearches/name/Test Search", 200, "validate_get_advanced_computer_search.xml")
}

func (m *AdvancedComputerSearchesMock) RegisterCreateAdvancedComputerSearchMock() {
	m.Register("POST", "/JSSResource/advancedcomputersearches/id/0", 201, "validate_create_advanced_computer_search.xml")
}

func (m *AdvancedComputerSearchesMock) RegisterUpdateAdvancedComputerSearchByIDMock() {
	m.Register("PUT", "/JSSResource/advancedcomputersearches/id/1", 200, "validate_update_advanced_computer_search.xml")
}

func (m *AdvancedComputerSearchesMock) RegisterUpdateAdvancedComputerSearchByNameMock() {
	m.Register("PUT", "/JSSResource/advancedcomputersearches/name/Test Search", 200, "validate_update_advanced_computer_search.xml")
}

func (m *AdvancedComputerSearchesMock) RegisterDeleteAdvancedComputerSearchByIDMock() {
	m.Register("DELETE", "/JSSResource/advancedcomputersearches/id/1", 200, "")
}

func (m *AdvancedComputerSearchesMock) RegisterDeleteAdvancedComputerSearchByNameMock() {
	m.Register("DELETE", "/JSSResource/advancedcomputersearches/name/Test Search", 200, "")
}

func (m *AdvancedComputerSearchesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/advancedcomputersearches/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *AdvancedComputerSearchesMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/advancedcomputersearches/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): An advanced computer search with that name already exists")
}

