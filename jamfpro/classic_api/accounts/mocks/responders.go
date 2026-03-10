package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type AccountsMock struct {
	*mocks.GenericMock
}

func NewAccountsMock() *AccountsMock {
	return &AccountsMock{
		GenericMock: mocks.NewXMLMock("AccountsMock"),
	}
}

func (m *AccountsMock) RegisterMocks() {
	m.RegisterListAccountsMock()
	m.RegisterGetAccountByIDMock()
	m.RegisterGetAccountByNameMock()
	m.RegisterCreateAccountMock()
	m.RegisterUpdateAccountByIDMock()
	m.RegisterUpdateAccountByNameMock()
	m.RegisterDeleteAccountByIDMock()
	m.RegisterDeleteAccountByNameMock()
}

func (m *AccountsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *AccountsMock) RegisterListAccountsMock() {
	m.Register("GET", "/JSSResource/accounts", 200, "validate_list_accounts.xml")
}

func (m *AccountsMock) RegisterGetAccountByIDMock() {
	m.Register("GET", "/JSSResource/accounts/userid/1", 200, "validate_get_account.xml")
}

func (m *AccountsMock) RegisterGetAccountByNameMock() {
	m.Register("GET", "/JSSResource/accounts/username/testuser1", 200, "validate_get_account.xml")
}

func (m *AccountsMock) RegisterCreateAccountMock() {
	m.Register("POST", "/JSSResource/accounts/userid/0", 201, "validate_create_account.xml")
}

func (m *AccountsMock) RegisterUpdateAccountByIDMock() {
	m.Register("PUT", "/JSSResource/accounts/userid/1", 200, "validate_update_account.xml")
}

func (m *AccountsMock) RegisterUpdateAccountByNameMock() {
	m.Register("PUT", "/JSSResource/accounts/username/testuser1", 200, "validate_update_account.xml")
}

func (m *AccountsMock) RegisterDeleteAccountByIDMock() {
	m.Register("DELETE", "/JSSResource/accounts/userid/1", 200, "")
}

func (m *AccountsMock) RegisterDeleteAccountByNameMock() {
	m.Register("DELETE", "/JSSResource/accounts/username/testuser1", 200, "")
}

func (m *AccountsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/JSSResource/accounts/userid/999", 404, "error_not_found.xml", "Jamf Pro Classic API error (404): Resource not found")
}

func (m *AccountsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/JSSResource/accounts/userid/0", 409, "error_conflict.xml", "Jamf Pro Classic API error (409): An account with that name already exists")
}
