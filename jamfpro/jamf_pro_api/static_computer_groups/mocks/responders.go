package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type StaticComputerGroupsMock struct {
	*mocks.GenericMock
}

func NewStaticComputerGroupsMock() *StaticComputerGroupsMock {
	return &StaticComputerGroupsMock{
		GenericMock: mocks.NewJSONMock("StaticComputerGroupsMock"),
	}
}

func (m *StaticComputerGroupsMock) RegisterListMock() {
	m.Register("GET", "/api/v2/computer-groups/static-groups", 200, "validate_list.json")
}

func (m *StaticComputerGroupsMock) RegisterGetByIDMock() {
	m.Register("GET", "/api/v2/computer-groups/static-groups/1", 200, "validate_get.json")
}

func (m *StaticComputerGroupsMock) RegisterCreateMock() {
	m.Register("POST", "/api/v2/computer-groups/static-groups", 201, "validate_create.json")
}

func (m *StaticComputerGroupsMock) RegisterUpdateMock() {
	m.Register("PUT", "/api/v2/computer-groups/static-groups/1", 200, "validate_update.json")
}

func (m *StaticComputerGroupsMock) RegisterDeleteMock() {
	m.Register("DELETE", "/api/v2/computer-groups/static-groups/1", 204, "")
}

func (m *StaticComputerGroupsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v2/computer-groups/static-groups/999", 404, "error_not_found.json", "")
}

func (m *StaticComputerGroupsMock) RegisterUpdateNotFoundErrorMock() {
	m.RegisterError("PUT", "/api/v2/computer-groups/static-groups/999", 404, "error_not_found.json", "")
}

func (m *StaticComputerGroupsMock) RegisterDeleteNotFoundErrorMock() {
	m.RegisterError("DELETE", "/api/v2/computer-groups/static-groups/999", 404, "error_not_found.json", "")
}

func (m *StaticComputerGroupsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/api/v2/computer-groups/static-groups", 409, "error_conflict.json", "")
}

func (m *StaticComputerGroupsMock) RegisterListEmptyMock() {
	m.Register("GET", "/api/v2/computer-groups/static-groups", 200, "validate_list_empty.json")
}

func (m *StaticComputerGroupsMock) RegisterListNoResponseErrorMock() {
	m.RegisterError("GET", "/api/v2/computer-groups/static-groups", 500, "error_internal.json", "")
}
