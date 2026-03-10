package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type SmartUserGroupsMock struct {
	*mocks.GenericMock
}

func NewSmartUserGroupsMock() *SmartUserGroupsMock {
	return &SmartUserGroupsMock{
		GenericMock: mocks.NewXMLMock("SmartUserGroupsMock"),
	}
}

func (m *SmartUserGroupsMock) RegisterMocks() {
	m.RegisterListSmartUserGroupsMock()
	m.RegisterGetSmartUserGroupByIDMock()
	m.RegisterGetSmartUserGroupByNameMock()
	m.RegisterCreateSmartUserGroupMock()
	m.RegisterUpdateSmartUserGroupByIDMock()
	m.RegisterUpdateSmartUserGroupByNameMock()
	m.RegisterDeleteSmartUserGroupByIDMock()
	m.RegisterDeleteSmartUserGroupByNameMock()
}

func (m *SmartUserGroupsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *SmartUserGroupsMock) RegisterListSmartUserGroupsMock() {
	m.Register("GET", "/JSSResource/usergroups", 200, "validate_list_user_groups.xml")
}

func (m *SmartUserGroupsMock) RegisterGetSmartUserGroupByIDMock() {
	m.Register("GET", "/JSSResource/usergroups/id/1", 200, "validate_get_user_group.xml")
}

func (m *SmartUserGroupsMock) RegisterGetSmartUserGroupByNameMock() {
	m.Register("GET", "/JSSResource/usergroups/name/All Users", 200, "validate_get_user_group.xml")
}

func (m *SmartUserGroupsMock) RegisterCreateSmartUserGroupMock() {
	m.Register("POST", "/JSSResource/usergroups/id/0", 201, "validate_create_user_group.xml")
}

func (m *SmartUserGroupsMock) RegisterUpdateSmartUserGroupByIDMock() {
	m.Register("PUT", "/JSSResource/usergroups/id/1", 200, "validate_update_user_group.xml")
}

func (m *SmartUserGroupsMock) RegisterUpdateSmartUserGroupByNameMock() {
	m.Register("PUT", "/JSSResource/usergroups/name/All Users", 200, "validate_update_user_group.xml")
}

func (m *SmartUserGroupsMock) RegisterDeleteSmartUserGroupByIDMock() {
	m.Register("DELETE", "/JSSResource/usergroups/id/1", 200, "")
}

func (m *SmartUserGroupsMock) RegisterDeleteSmartUserGroupByNameMock() {
	m.Register("DELETE", "/JSSResource/usergroups/name/All Users", 200, "")
}

func (m *SmartUserGroupsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/usergroups/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *SmartUserGroupsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/usergroups/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A user group with that name already exists")
}

