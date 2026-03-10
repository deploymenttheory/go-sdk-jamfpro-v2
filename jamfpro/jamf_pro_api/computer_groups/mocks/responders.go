package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ComputerGroupsMock struct {
	*mocks.GenericMock
}

func NewComputerGroupsMock() *ComputerGroupsMock {
	return &ComputerGroupsMock{
		GenericMock: mocks.NewJSONMock("ComputerGroupsMock"),
	}
}

func (m *ComputerGroupsMock) RegisterMocks() {
	m.RegisterListSmartGroupsMock()
	m.RegisterGetSmartGroupMock()
	m.RegisterCreateSmartGroupMock()
	m.RegisterUpdateSmartGroupMock()
	m.RegisterDeleteSmartGroupMock()
	m.RegisterListStaticGroupsMock()
	m.RegisterGetStaticGroupMock()
	m.RegisterCreateStaticGroupMock()
	m.RegisterUpdateStaticGroupMock()
	m.RegisterDeleteStaticGroupMock()
}

func (m *ComputerGroupsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *ComputerGroupsMock) RegisterListSmartGroupsMock() {
	m.Register("GET", "/api/v2/computer-groups/smart-groups", 200, "validate_list_smart_groups.json")
}

func (m *ComputerGroupsMock) RegisterGetSmartGroupMock() {
	m.Register("GET", "/api/v2/computer-groups/smart-groups/1", 200, "validate_get_smart_group.json")
}

func (m *ComputerGroupsMock) RegisterCreateSmartGroupMock() {
	m.Register("POST", "/api/v2/computer-groups/smart-groups", 201, "validate_create_smart_group.json")
}

func (m *ComputerGroupsMock) RegisterUpdateSmartGroupMock() {
	m.Register("PUT", "/api/v2/computer-groups/smart-groups/1", 200, "validate_update_smart_group.json")
}

func (m *ComputerGroupsMock) RegisterDeleteSmartGroupMock() {
	m.Register("DELETE", "/api/v2/computer-groups/smart-groups/1", 204, "")
}

func (m *ComputerGroupsMock) RegisterListStaticGroupsMock() {
	m.Register("GET", "/api/v2/computer-groups/static-groups", 200, "validate_list_static_groups.json")
}

func (m *ComputerGroupsMock) RegisterGetStaticGroupMock() {
	m.Register("GET", "/api/v2/computer-groups/static-groups/10", 200, "validate_get_static_group.json")
}

func (m *ComputerGroupsMock) RegisterCreateStaticGroupMock() {
	m.Register("POST", "/api/v2/computer-groups/static-groups", 201, "validate_create_static_group.json")
}

func (m *ComputerGroupsMock) RegisterUpdateStaticGroupMock() {
	m.Register("PUT", "/api/v2/computer-groups/static-groups/10", 200, "validate_update_static_group.json")
}

func (m *ComputerGroupsMock) RegisterDeleteStaticGroupMock() {
	m.Register("DELETE", "/api/v2/computer-groups/static-groups/10", 204, "")
}

func (m *ComputerGroupsMock) RegisterListAllV1Mock() {
	m.Register("GET", "/api/v1/computer-groups", 200, "validate_list_all_v1.json")
}

func (m *ComputerGroupsMock) RegisterGetSmartGroupMembershipMock() {
	m.Register("GET", "/api/v2/computer-groups/smart-group-membership/1", 200, "validate_get_smart_group_membership.json")
}

func (m *ComputerGroupsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v2/computer-groups/smart-groups/999", 404, "error_not_found.json", "")
}

func (m *ComputerGroupsMock) RegisterStaticNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v2/computer-groups/static-groups/999", 404, "error_not_found.json", "")
}

func (m *ComputerGroupsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/api/v2/computer-groups/smart-groups", 409, "error_conflict.json", "")
}

func (m *ComputerGroupsMock) RegisterStaticConflictErrorMock() {
	m.RegisterError("POST", "/api/v2/computer-groups/static-groups", 409, "error_conflict.json", "")
}

func (m *ComputerGroupsMock) RegisterSmartUpdateNotFoundErrorMock() {
	m.RegisterError("PUT", "/api/v2/computer-groups/smart-groups/999", 404, "error_not_found.json", "")
}

func (m *ComputerGroupsMock) RegisterStaticUpdateNotFoundErrorMock() {
	m.RegisterError("PUT", "/api/v2/computer-groups/static-groups/999", 404, "error_not_found.json", "")
}

func (m *ComputerGroupsMock) RegisterSmartDeleteNotFoundErrorMock() {
	m.RegisterError("DELETE", "/api/v2/computer-groups/smart-groups/999", 404, "error_not_found.json", "")
}

func (m *ComputerGroupsMock) RegisterStaticDeleteNotFoundErrorMock() {
	m.RegisterError("DELETE", "/api/v2/computer-groups/static-groups/999", 404, "error_not_found.json", "")
}
