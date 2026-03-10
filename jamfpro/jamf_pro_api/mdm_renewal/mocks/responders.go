package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type MDMRenewalMock struct {
	*mocks.GenericMock
}

func NewMDMRenewalMock() *MDMRenewalMock {
	return &MDMRenewalMock{
		GenericMock: mocks.NewJSONMock("MDMRenewalMock"),
	}
}

func (m *MDMRenewalMock) RegisterUpdateDeviceCommonDetailsMock() {
	m.Register("PATCH", "/api/v1/mdm-renewal/device-common-details", 204, "")
}

func (m *MDMRenewalMock) RegisterGetDeviceCommonDetailsMock(clientManagementID string) {
	m.Register("GET", "/api/v1/mdm-renewal/device-common-details/"+clientManagementID, 200, "validate_get_device_common_details.json")
}

func (m *MDMRenewalMock) RegisterGetRenewalStrategiesMock(clientManagementID string) {
	m.Register("GET", "/api/v1/mdm-renewal/renewal-strategies/"+clientManagementID, 200, "validate_get_renewal_strategies.json")
}

func (m *MDMRenewalMock) RegisterDeleteRenewalStrategiesMock(clientManagementID string) {
	m.Register("DELETE", "/api/v1/mdm-renewal/renewal-strategies/"+clientManagementID, 204, "")
}

func (m *MDMRenewalMock) RegisterNotFoundErrorMock(clientManagementID string) {
	m.RegisterError("GET", "/api/v1/mdm-renewal/device-common-details/"+clientManagementID, 404, "error_not_found.json", "")
	m.RegisterError("GET", "/api/v1/mdm-renewal/renewal-strategies/"+clientManagementID, 404, "error_not_found.json", "")
	m.RegisterError("DELETE", "/api/v1/mdm-renewal/renewal-strategies/"+clientManagementID, 404, "error_not_found.json", "")
}

func (m *MDMRenewalMock) RegisterUpdateDeviceCommonDetailsErrorMock() {
	m.RegisterError("PATCH", "/api/v1/mdm-renewal/device-common-details", 500, "error_internal.json", "")
}

func (m *MDMRenewalMock) RegisterGetDeviceCommonDetailsErrorMock(clientManagementID string) {
	m.RegisterError("GET", "/api/v1/mdm-renewal/device-common-details/"+clientManagementID, 500, "error_internal.json", "")
}

func (m *MDMRenewalMock) RegisterGetRenewalStrategiesErrorMock(clientManagementID string) {
	m.RegisterError("GET", "/api/v1/mdm-renewal/renewal-strategies/"+clientManagementID, 500, "error_internal.json", "")
}

func (m *MDMRenewalMock) RegisterDeleteRenewalStrategiesErrorMock(clientManagementID string) {
	m.RegisterError("DELETE", "/api/v1/mdm-renewal/renewal-strategies/"+clientManagementID, 500, "error_internal.json", "")
}
