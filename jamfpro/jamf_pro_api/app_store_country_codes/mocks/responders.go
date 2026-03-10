package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type AppStoreCountryCodesMock struct {
	*mocks.GenericMock
}

func NewAppStoreCountryCodesMock() *AppStoreCountryCodesMock {
	return &AppStoreCountryCodesMock{
		GenericMock: mocks.NewJSONMock("AppStoreCountryCodesMock"),
	}
}

func (m *AppStoreCountryCodesMock) RegisterMocks() {
	m.Register("GET", "/api/v1/app-store-country-codes", 200, "validate_list.json")
}

func (m *AppStoreCountryCodesMock) RegisterEmptyListMock() {
	m.Register("GET", "/api/v1/app-store-country-codes", 200, "empty_list.json")
}

func (m *AppStoreCountryCodesMock) RegisterListErrorMock() {
	m.RegisterError("GET", "/api/v1/app-store-country-codes", 500, "error_internal.json", "")
}
