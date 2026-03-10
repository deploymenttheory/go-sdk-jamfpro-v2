package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type TimeZonesMock struct {
	*mocks.GenericMock
}

func NewTimeZonesMock() *TimeZonesMock {
	return &TimeZonesMock{
		GenericMock: mocks.NewJSONMock("TimeZonesMock"),
	}
}

func (m *TimeZonesMock) RegisterMocks() {
	m.Register("GET", "/api/v1/time-zones", 200, "validate_list.json")
}

func (m *TimeZonesMock) RegisterListV1ErrorMock() {
	m.RegisterError("GET", "/api/v1/time-zones", 500, "", "mock client error")
}

func (m *TimeZonesMock) RegisterListV1InvalidJSONMock() {
	m.RegisterRawBody("GET", "/api/v1/time-zones", 200, []byte(`{invalid json`))
}

func (m *TimeZonesMock) RegisterListV1EmptyMock() {
	m.Register("GET", "/api/v1/time-zones", 200, "validate_list_empty.json")
}

func (m *TimeZonesMock) RegisterListV1NoResponseErrorMock() {
	m.RegisterError("GET", "/api/v1/time-zones", 500, "error_internal.json", "no response registered")
}
