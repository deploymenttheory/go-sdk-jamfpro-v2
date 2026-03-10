package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type ComputerPrestagesMock struct {
	*mocks.GenericMock
}

func NewComputerPrestagesMock() *ComputerPrestagesMock {
	return &ComputerPrestagesMock{
		GenericMock: mocks.NewJSONMock("ComputerPrestagesMock"),
	}
}

func (m *ComputerPrestagesMock) RegisterListMock() {
	m.Register("GET", "/api/v3/computer-prestages", 200, "validate_list.json")
}

func (m *ComputerPrestagesMock) RegisterGetByIDMock(id string) {
	m.Register("GET", "/api/v3/computer-prestages/"+id, 200, "validate_get.json")
}

func (m *ComputerPrestagesMock) RegisterCreateMock() {
	m.Register("POST", "/api/v3/computer-prestages", 200, "validate_create.json")
}

func (m *ComputerPrestagesMock) RegisterUpdateByIDMock(id string) {
	m.Register("PUT", "/api/v3/computer-prestages/"+id, 200, "validate_get.json")
}

func (m *ComputerPrestagesMock) RegisterDeleteByIDMock(id string) {
	m.Register("DELETE", "/api/v3/computer-prestages/"+id, 200, "")
}

func (m *ComputerPrestagesMock) RegisterGetDeviceScopeMock(id string) {
	m.Register("GET", "/api/v2/computer-prestages/"+id+"/scope", 200, "validate_scope.json")
}

func (m *ComputerPrestagesMock) RegisterReplaceDeviceScopeMock(id string) {
	m.Register("PUT", "/api/v2/computer-prestages/"+id+"/scope", 200, "validate_scope.json")
}

func (m *ComputerPrestagesMock) RegisterGetAllDeviceScopeMock() {
	m.Register("GET", "/api/v2/computer-prestages/scope", 200, "validate_all_scope.json")
}

func (m *ComputerPrestagesMock) RegisterAddDeviceScopeMock(id string) {
	m.Register("POST", "/api/v2/computer-prestages/"+id+"/scope", 200, "validate_scope.json")
}

func (m *ComputerPrestagesMock) RegisterRemoveDeviceScopeMock(id string) {
	m.Register("POST", "/api/v2/computer-prestages/"+id+"/scope/delete-multiple", 200, "validate_scope.json")
}
