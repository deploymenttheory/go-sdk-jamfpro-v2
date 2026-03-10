package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type StaticUserGroupsMock struct {
	*mocks.GenericMock
}

func NewStaticUserGroupsMock() *StaticUserGroupsMock {
	return &StaticUserGroupsMock{
		GenericMock: mocks.NewXMLMock("StaticUserGroupsMock"),
	}
}

func (m *StaticUserGroupsMock) RegisterMocks() {
	m.RegisterListStaticUserGroupsMock()
	m.RegisterGetStaticUserGroupByIDMock()
	m.RegisterGetStaticUserGroupByNameMock()
	m.RegisterCreateStaticUserGroupMock()
	m.RegisterUpdateStaticUserGroupByIDMock()
	m.RegisterUpdateStaticUserGroupByNameMock()
	m.RegisterDeleteStaticUserGroupByIDMock()
	m.RegisterDeleteStaticUserGroupByNameMock()
}

func (m *StaticUserGroupsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *StaticUserGroupsMock) RegisterListStaticUserGroupsMock() {
	m.Register("GET", "/JSSResource/usergroups", 200, "validate_list_user_groups.xml")
}

func (m *StaticUserGroupsMock) RegisterGetStaticUserGroupByIDMock() {
	m.Register("GET", "/JSSResource/usergroups/id/1", 200, "validate_get_user_group.xml")
}

func (m *StaticUserGroupsMock) RegisterGetStaticUserGroupByNameMock() {
	m.Register("GET", "/JSSResource/usergroups/name/Static Test Group", 200, "validate_get_user_group.xml")
}

func (m *StaticUserGroupsMock) RegisterCreateStaticUserGroupMock() {
	m.Register("POST", "/JSSResource/usergroups/id/0", 201, "validate_create_user_group.xml")
}

func (m *StaticUserGroupsMock) RegisterUpdateStaticUserGroupByIDMock() {
	m.Register("PUT", "/JSSResource/usergroups/id/1", 200, "validate_update_user_group.xml")
}

func (m *StaticUserGroupsMock) RegisterUpdateStaticUserGroupByNameMock() {
	m.Register("PUT", "/JSSResource/usergroups/name/Static Test Group", 200, "validate_update_user_group.xml")
}

func (m *StaticUserGroupsMock) RegisterDeleteStaticUserGroupByIDMock() {
	m.Register("DELETE", "/JSSResource/usergroups/id/1", 200, "")
}

func (m *StaticUserGroupsMock) RegisterDeleteStaticUserGroupByNameMock() {
	m.Register("DELETE", "/JSSResource/usergroups/name/Static Test Group", 200, "")
}

func (m *StaticUserGroupsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/usergroups/id/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *StaticUserGroupsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/usergroups/id/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): A user group with that name already exists")
}

