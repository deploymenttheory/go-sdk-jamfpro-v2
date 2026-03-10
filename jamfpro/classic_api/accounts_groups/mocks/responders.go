package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type AccountGroupsMock struct {
	*mocks.GenericMock
}

func NewAccountGroupsMock() *AccountGroupsMock {
	return &AccountGroupsMock{
		GenericMock: mocks.NewXMLMock("AccountGroupsMock"),
	}
}

func (m *AccountGroupsMock) RegisterMocks() {
	m.RegisterGetAccountGroupByIDMock()
	m.RegisterGetAccountGroupByNameMock()
	m.RegisterCreateAccountGroupMock()
	m.RegisterUpdateAccountGroupByIDMock()
	m.RegisterUpdateAccountGroupByNameMock()
	m.RegisterDeleteAccountGroupByIDMock()
	m.RegisterDeleteAccountGroupByNameMock()
}

func (m *AccountGroupsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *AccountGroupsMock) RegisterGetAccountGroupByIDMock() {
	m.Register("GET", "/JSSResource/accounts/groupid/1", 200, "validate_get_account_group.xml")
}

func (m *AccountGroupsMock) RegisterGetAccountGroupByNameMock() {
	m.Register("GET", "/JSSResource/accounts/groupname/testgroup1", 200, "validate_get_account_group.xml")
}

func (m *AccountGroupsMock) RegisterCreateAccountGroupMock() {
	m.Register("POST", "/JSSResource/accounts/groupid/0", 201, "validate_create_account_group.xml")
}

func (m *AccountGroupsMock) RegisterUpdateAccountGroupByIDMock() {
	m.Register("PUT", "/JSSResource/accounts/groupid/1", 200, "validate_update_account_group.xml")
}

func (m *AccountGroupsMock) RegisterUpdateAccountGroupByNameMock() {
	m.Register("PUT", "/JSSResource/accounts/groupname/testgroup1", 200, "validate_update_account_group.xml")
}

func (m *AccountGroupsMock) RegisterDeleteAccountGroupByIDMock() {
	m.Register("DELETE", "/JSSResource/accounts/groupid/1", 200, "")
}

func (m *AccountGroupsMock) RegisterDeleteAccountGroupByNameMock() {
	m.Register("DELETE", "/JSSResource/accounts/groupname/testgroup1", 200, "")
}

func (m *AccountGroupsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/accounts/groupid/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *AccountGroupsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/accounts/groupid/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): An account group with that name already exists")
}
