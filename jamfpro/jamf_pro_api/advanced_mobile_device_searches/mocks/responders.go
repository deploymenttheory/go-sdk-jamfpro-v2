package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type AdvancedMobileDeviceSearchesMock struct {
	*mocks.GenericMock
}

func NewAdvancedMobileDeviceSearchesMock() *AdvancedMobileDeviceSearchesMock {
	return &AdvancedMobileDeviceSearchesMock{
		GenericMock: mocks.NewJSONMock("AdvancedMobileDeviceSearchesMock"),
	}
}

func (m *AdvancedMobileDeviceSearchesMock) RegisterMocks() {
	m.Register("GET", "/api/v1/advanced-mobile-device-searches", 200, "validate_list.json")
	m.Register("GET", "/api/v1/advanced-mobile-device-searches/1", 200, "validate_get.json")
	m.Register("POST", "/api/v1/advanced-mobile-device-searches", 201, "validate_create.json")
	m.Register("PUT", "/api/v1/advanced-mobile-device-searches/1", 200, "validate_get.json")
	m.Register("DELETE", "/api/v1/advanced-mobile-device-searches/1", 204, "")
	m.Register("POST", "/api/v1/advanced-mobile-device-searches/delete-multiple", 204, "")
	m.Register("GET", "/api/v1/advanced-mobile-device-searches/choices?criteria=Device%20Name&site=-1&contains=", 200, "validate_choices.json")
	m.Register("GET", "/api/v1/advanced-mobile-device-searches/choices?criteria=Device+Name&site=-1&contains=", 200, "validate_choices.json")
}

func (m *AdvancedMobileDeviceSearchesMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/advanced-mobile-device-searches/999", 404, "error_not_found.json", "")
}
