package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ManagedSoftwareUpdatesMock struct {
	*mocks.GenericMock
}

func NewManagedSoftwareUpdatesMock() *ManagedSoftwareUpdatesMock {
	return &ManagedSoftwareUpdatesMock{
		GenericMock: mocks.NewJSONMock("ManagedSoftwareUpdatesMock"),
	}
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetAvailableUpdatesMock() {
	m.Register("GET", "/api/v1/managed-software-updates/available-updates", 200, "validate_available_updates.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetPlansMock() {
	m.Register("GET", "/api/v1/managed-software-updates/plans", 200, "validate_plans_list.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetPlanByUUIDMock(uuid string) {
	m.Register("GET", "/api/v1/managed-software-updates/plans/"+uuid, 200, "validate_plan_detail.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetDeclarationsByPlanUUIDMock(uuid string) {
	m.Register("GET", "/api/v1/managed-software-updates/plans/"+uuid+"/declarations", 200, "validate_declarations_list.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterCreatePlanByDeviceIDMock() {
	m.Register("POST", "/api/v1/managed-software-updates/plans", 201, "validate_plan_create.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterCreatePlanByGroupIDMock() {
	m.Register("POST", "/api/v1/managed-software-updates/plans/group", 201, "validate_plan_create.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetPlansByGroupIDMock(groupID string) {
	m.Register("GET", "/api/v1/managed-software-updates/plans/group/"+groupID+"?group-type=COMPUTER", 200, "validate_plans_list.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetFeatureToggleMock() {
	m.Register("GET", "/api/v1/managed-software-updates/plans/feature-toggle", 200, "validate_feature_toggle_get.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterUpdateFeatureToggleMock() {
	m.Register("PUT", "/api/v1/managed-software-updates/plans/feature-toggle", 200, "validate_feature_toggle_response.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetFeatureToggleStatusMock() {
	m.Register("GET", "/api/v1/managed-software-updates/plans/feature-toggle/status", 200, "validate_feature_toggle_status.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterForceStopFeatureToggleProcessMock() {
	m.Register("POST", "/api/v1/managed-software-updates/plans/feature-toggle/abandon", 200, "validate_error_response.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterEmptyPlansMock() {
	m.Register("GET", "/api/v1/managed-software-updates/plans", 200, "validate_empty_plans_list.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetUpdateStatusesMock() {
	m.Register("GET", "/api/v1/managed-software-updates/update-statuses", 200, "validate_update_statuses.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetUpdateStatusesArrayFormatMock() {
	m.Register("GET", "/api/v1/managed-software-updates/update-statuses", 200, "validate_update_statuses_array.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetPlansInvalidMock() {
	m.Register("GET", "/api/v1/managed-software-updates/plans", 200, "validate_plans_list_invalid.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetUpdateStatusesInvalidMock() {
	m.Register("GET", "/api/v1/managed-software-updates/update-statuses", 200, "validate_update_statuses_invalid.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetPlanEventsByUUIDMock(uuid string) {
	m.Register("GET", "/api/v1/managed-software-updates/plans/"+uuid+"/events", 200, "validate_plan_events.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetUpdateStatusesByComputerGroupMock(id string) {
	m.Register("GET", "/api/v1/managed-software-updates/update-statuses/computer-groups/"+id, 200, "validate_update_statuses.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetUpdateStatusesByComputerMock(id string) {
	m.Register("GET", "/api/v1/managed-software-updates/update-statuses/computers/"+id, 200, "validate_update_statuses.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetUpdateStatusesByMobileDeviceGroupMock(id string) {
	m.Register("GET", "/api/v1/managed-software-updates/update-statuses/mobile-device-groups/"+id, 200, "validate_update_statuses.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterGetUpdateStatusesByMobileDeviceMock(id string) {
	m.Register("GET", "/api/v1/managed-software-updates/update-statuses/mobile-devices/"+id, 200, "validate_update_statuses.json")
}

func (m *ManagedSoftwareUpdatesMock) RegisterErrorMock(method, path, errMsg string) {
	m.RegisterError(method, path, 500, "", errMsg)
}
