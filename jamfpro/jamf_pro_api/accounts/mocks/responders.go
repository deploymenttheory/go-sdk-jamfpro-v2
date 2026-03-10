package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type AccountsMock struct {
	*mocks.GenericMock
}

func NewAccountsMock() *AccountsMock {
	return &AccountsMock{
		GenericMock: mocks.NewJSONMock("AccountsMock"),
	}
}

func (m *AccountsMock) RegisterMocks() {
	m.RegisterListAccountsMock()
	m.RegisterGetAccountMock()
	m.RegisterCreateAccountMock()
	m.RegisterDeleteAccountMock()
}

func (m *AccountsMock) RegisterListAccountsMock() {
	m.Register("GET", "/api/v1/accounts", 200, "validate_list.json")
}

func (m *AccountsMock) RegisterGetAccountMock() {
	m.Register("GET", "/api/v1/accounts/1", 200, "validate_get.json")
}

func (m *AccountsMock) RegisterCreateAccountMock() {
	m.Register("POST", "/api/v1/accounts", 201, "validate_create.json")
}

func (m *AccountsMock) RegisterDeleteAccountMock() {
	m.Register("DELETE", "/api/v1/accounts/1", 204, "")
}
