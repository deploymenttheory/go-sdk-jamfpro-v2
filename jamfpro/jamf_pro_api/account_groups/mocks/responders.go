package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type AccountGroupsMock struct {
	*mocks.GenericMock
}

func NewAccountGroupsMock() *AccountGroupsMock {
	return &AccountGroupsMock{
		GenericMock: mocks.NewJSONMock("AccountGroupsMock"),
	}
}

func (m *AccountGroupsMock) RegisterMocks() {
	m.RegisterListAccountGroupsMock()
	m.RegisterGetAccountGroupMock()
}

func (m *AccountGroupsMock) RegisterListAccountGroupsMock() {
	m.Register("GET", "/api/v1/account-groups", 200, "validate_list.json")
}

func (m *AccountGroupsMock) RegisterGetAccountGroupMock() {
	m.Register("GET", "/api/v1/account-groups/1", 200, "validate_get.json")
}
