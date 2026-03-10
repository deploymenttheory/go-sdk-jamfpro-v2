package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type AdvancedUserSearchesMock struct {
	*mocks.GenericMock
}

func NewAdvancedUserSearchesMock() *AdvancedUserSearchesMock {
	return &AdvancedUserSearchesMock{
		GenericMock: mocks.NewXMLMock("AdvancedUserSearchesMock"),
	}
}

func (m *AdvancedUserSearchesMock) RegisterMocks() {
	m.RegisterListMock()
	m.RegisterGetByIDMock()
	m.RegisterGetByNameMock()
	m.RegisterCreateMock()
	m.RegisterUpdateByIDMock()
	m.RegisterUpdateByNameMock()
	m.RegisterDeleteByIDMock()
	m.RegisterDeleteByNameMock()
}

func (m *AdvancedUserSearchesMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *AdvancedUserSearchesMock) RegisterListMock() {
	m.Register("GET", "/JSSResource/advancedusersearches", 200, "validate_list_advanced_user_searches.xml")
}

func (m *AdvancedUserSearchesMock) RegisterGetByIDMock() {
	m.Register("GET", "/JSSResource/advancedusersearches/id/1", 200, "validate_get_advanced_user_search.xml")
}

func (m *AdvancedUserSearchesMock) RegisterGetByNameMock() {
	m.Register("GET", "/JSSResource/advancedusersearches/name/Test Search", 200, "validate_get_advanced_user_search.xml")
}

func (m *AdvancedUserSearchesMock) RegisterCreateMock() {
	m.Register("POST", "/JSSResource/advancedusersearches/id/0", 201, "validate_create_advanced_user_search.xml")
}

func (m *AdvancedUserSearchesMock) RegisterUpdateByIDMock() {
	m.Register("PUT", "/JSSResource/advancedusersearches/id/1", 200, "validate_update_advanced_user_search.xml")
}

func (m *AdvancedUserSearchesMock) RegisterUpdateByNameMock() {
	m.Register("PUT", "/JSSResource/advancedusersearches/name/Test Search", 200, "validate_update_advanced_user_search.xml")
}

func (m *AdvancedUserSearchesMock) RegisterDeleteByIDMock() {
	m.Register("DELETE", "/JSSResource/advancedusersearches/id/1", 200, "")
}

func (m *AdvancedUserSearchesMock) RegisterDeleteByNameMock() {
	m.Register("DELETE", "/JSSResource/advancedusersearches/name/Test Search", 200, "")
}

func (m *AdvancedUserSearchesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/advancedusersearches/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *AdvancedUserSearchesMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/advancedusersearches/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): An advanced user search with that name already exists")
}

