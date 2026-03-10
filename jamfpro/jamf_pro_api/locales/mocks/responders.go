package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type LocalesMock struct {
	*mocks.GenericMock
}

func NewLocalesMock() *LocalesMock {
	return &LocalesMock{
		GenericMock: mocks.NewJSONMock("LocalesMock"),
	}
}

func (m *LocalesMock) RegisterMocks() {
	m.Register("GET", "/api/v1/locales", 200, "validate_list.json")
}

func (m *LocalesMock) RegisterListV1ErrorMock() {
	m.RegisterError("GET", "/api/v1/locales", 500, "", "mock client error")
}

func (m *LocalesMock) RegisterListV1InvalidJSONMock() {
	m.RegisterRawBody("GET", "/api/v1/locales", 200, []byte(`{invalid json`))
}

func (m *LocalesMock) RegisterListV1NoResponseErrorMock() {
	m.RegisterError("GET", "/api/v1/locales", 500, "error_internal.json", "no response")
}
