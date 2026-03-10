package mocks

import (
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mocks"
)

type DepartmentsMock struct {
	*mocks.GenericMock
}

func NewDepartmentsMock() *DepartmentsMock {
	return &DepartmentsMock{
		GenericMock: mocks.NewJSONMock("DepartmentsMock"),
	}
}

func (m *DepartmentsMock) RegisterMocks() {
	m.RegisterListDepartmentsMock()
	m.RegisterGetDepartmentMock()
	m.RegisterCreateDepartmentMock()
	m.RegisterUpdateDepartmentMock()
	m.RegisterDeleteDepartmentMock()
	m.RegisterGetDepartmentHistoryMock()
	m.RegisterAddDepartmentHistoryNotesMock()
}

func (m *DepartmentsMock) RegisterErrorMocks() {
	m.RegisterNotFoundErrorMock()
	m.RegisterConflictErrorMock()
}

func (m *DepartmentsMock) RegisterListDepartmentsMock() {
	m.Register("GET", "/api/v1/departments", 200, "validate_list_departments.json")
}

func (m *DepartmentsMock) RegisterListDepartmentsRSQLMock() {
	m.Register("GET", "/api/v1/departments", 200, "validate_list_departments_rsql.json")
}

func (m *DepartmentsMock) RegisterGetDepartmentMock() {
	m.Register("GET", "/api/v1/departments/1", 200, "validate_get_department.json")
}

func (m *DepartmentsMock) RegisterCreateDepartmentMock() {
	m.Register("POST", "/api/v1/departments", 201, "validate_create_department.json")
}

func (m *DepartmentsMock) RegisterUpdateDepartmentMock() {
	m.Register("PUT", "/api/v1/departments/1", 200, "validate_update_department.json")
}

func (m *DepartmentsMock) RegisterDeleteDepartmentMock() {
	m.Register("DELETE", "/api/v1/departments/1", 204, "")
}

func (m *DepartmentsMock) RegisterNotFoundErrorMock() {
	m.RegisterError("GET", "/api/v1/departments/999", 404, "error_not_found.json", "")
}

func (m *DepartmentsMock) RegisterConflictErrorMock() {
	m.RegisterError("POST", "/api/v1/departments", 409, "error_conflict.json", "")
}

func (m *DepartmentsMock) RegisterGetDepartmentHistoryMock() {
	m.Register("GET", "/api/v1/departments/1/history", 200, "validate_get_history.json")
}

func (m *DepartmentsMock) RegisterAddDepartmentHistoryNotesMock() {
	m.Register("POST", "/api/v1/departments/1/history", 201, "")
}

func (m *DepartmentsMock) RegisterDeleteDepartmentsByIDMock() {
	m.Register("POST", "/api/v1/departments/delete-multiple", 204, "")
}

func (m *DepartmentsMock) RegisterListInvalidJSONMock() {
	m.Register("GET", "/api/v1/departments", 200, "validate_list_invalid.json")
}

func (m *DepartmentsMock) RegisterHistoryInvalidJSONMock() {
	m.Register("GET", "/api/v1/departments/1/history", 200, "validate_history_invalid.json")
}

func (m *DepartmentsMock) RegisterListAPIErrorMock() {
	m.RegisterError("GET", "/api/v1/departments", 500, "error_not_found.json", "")
}

func (m *DepartmentsMock) RegisterHistoryAPIErrorMock() {
	m.RegisterError("GET", "/api/v1/departments/1/history", 500, "error_not_found.json", "")
}

func (m *DepartmentsMock) RegisterDeleteDepartmentsByIDErrorMock() {
	m.RegisterError("POST", "/api/v1/departments/delete-multiple", 500, "error_not_found.json", "")
}
