package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type SmartComputerGroupsMock struct {
	*mocks.GenericMock
}

func NewSmartComputerGroupsMock() *SmartComputerGroupsMock {
	return &SmartComputerGroupsMock{
		GenericMock: mocks.NewJSONMock("SmartComputerGroupsMock"),
	}
}

func (m *SmartComputerGroupsMock) RegisterListMock() {
	m.Register("GET", "/api/v2/computer-groups/smart-groups", 200, "validate_list.json")
}

func (m *SmartComputerGroupsMock) RegisterListEmptyMock() {
	m.Register("GET", "/api/v2/computer-groups/smart-groups", 200, "validate_list_empty.json")
}

func (m *SmartComputerGroupsMock) RegisterGetByIDMock() {
	m.Register("GET", "/api/v2/computer-groups/smart-groups/1", 200, "validate_get_by_id.json")
}

func (m *SmartComputerGroupsMock) RegisterGetMembershipMock() {
	m.Register("GET", "/api/v2/computer-groups/smart-group-membership/1", 200, "validate_membership.json")
}

func (m *SmartComputerGroupsMock) RegisterCreateMock() {
	m.Register("POST", "/api/v2/computer-groups/smart-groups", 201, "validate_create.json")
}

func (m *SmartComputerGroupsMock) RegisterUpdateMock() {
	m.Register("PUT", "/api/v2/computer-groups/smart-groups/1", 200, "validate_update.json")
}

func (m *SmartComputerGroupsMock) RegisterDeleteMock() {
	m.Register("DELETE", "/api/v2/computer-groups/smart-groups/1", 204, "")
}

func (m *SmartComputerGroupsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v2/computer-groups/smart-groups/999", 404, "error_not_found.json", "")
	m.RegisterError("GET", "/api/v2/computer-groups/smart-group-membership/999", 404, "error_not_found.json", "")
}

func (m *SmartComputerGroupsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/api/v2/computer-groups/smart-groups", 409, "error_conflict.json", "")
}
