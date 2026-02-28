package departments

import (
	"context"
	"strings"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/departments/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.DepartmentsMock) {
	t.Helper()
	mock := mocks.NewDepartmentsMock()
	return NewService(mock), mock
}

func TestUnit_Departments_List_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListDepartmentsMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "Engineering", result.Results[0].Name)
	assert.Equal(t, "2", result.Results[1].ID)
	assert.Equal(t, "Sales", result.Results[1].Name)
}

func TestUnit_Departments_List_WithrsqlQuery(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListDepartmentsMock()

	params := map[string]string{"page": "0", "page-size": "50", "sort": "name:asc"}
	result, resp, err := svc.ListV1(context.Background(), params)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnit_Departments_List_WithRSQLFilter(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListDepartmentsRSQLMock()

	rsqlQuery := map[string]string{"filter": `name=="Sales"`}
	result, resp, err := svc.ListV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "2", result.Results[0].ID)
	assert.Equal(t, "Sales", result.Results[0].Name)
	assert.Equal(t, rsqlQuery, mock.LastRSQLQuery)
}

func TestUnit_Departments_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDepartmentMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Engineering", result.Name)
}

func TestUnit_Departments_GetByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "department ID is required")
}

func TestUnit_Departments_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestUnit_Departments_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateDepartmentMock()

	req := &RequestDepartment{Name: "Marketing"}
	result, resp, err := svc.CreateV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "3", result.ID)
	assert.Contains(t, result.Href, "/api/v1/departments/3")
}

func TestUnit_Departments_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Departments_Create_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestDepartment{Name: "Duplicate"}
	result, resp, err := svc.CreateV1(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

func TestUnit_Departments_UpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateDepartmentMock()

	req := &RequestDepartment{Name: "Engineering Updated"}
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Engineering Updated", result.Name)
}

func TestUnit_Departments_UpdateByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByIDV1(context.Background(), "", &RequestDepartment{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_Departments_UpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Departments_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteDepartmentMock()

	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnit_Departments_DeleteByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "department ID is required")
}

func TestUnit_Departments_GetHistoryV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDepartmentHistoryMock()

	result, resp, err := svc.GetDepartmentHistoryV1(context.Background(), "1", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "Department created", result.Results[0].Note)
}

func TestUnit_Departments_GetHistoryV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetDepartmentHistoryV1(context.Background(), "", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "department ID is required")
}

func TestUnit_Departments_AddHistoryNotesV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddDepartmentHistoryNotesMock()

	req := &AddHistoryNotesRequest{Note: "Added via SDK"}
	resp, err := svc.AddDepartmentHistoryNotesV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
}

func TestUnit_Departments_AddHistoryNotesV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.AddDepartmentHistoryNotesV1(context.Background(), "", &AddHistoryNotesRequest{Note: "x"})
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "department ID is required")
}

func TestUnit_Departments_AddHistoryNotesV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.AddDepartmentHistoryNotesV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request body is required")
}

func TestUnit_Departments_DeleteDepartmentsByIDV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteDepartmentsByIDMock()

	req := &DeleteDepartmentsByIDRequest{IDs: []string{"1", "2", "3"}}
	resp, err := svc.DeleteDepartmentsByIDV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnit_Departments_DeleteDepartmentsByIDV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteDepartmentsByIDV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "department IDs are required")
}

func TestUnit_Departments_DeleteDepartmentsByIDV1_EmptyIDs(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteDepartmentsByIDV1(context.Background(), &DeleteDepartmentsByIDRequest{IDs: []string{}})
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "department IDs are required")
}

func TestUnit_Departments_ListV1_InvalidJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListInvalidJSONMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	errMsg := err.Error()
	assert.True(t, strings.Contains(errMsg, "failed to unmarshal page") || strings.Contains(errMsg, "extract results field"), "error should contain JSON parsing error")
}

func TestUnit_Departments_ListV1_APIError(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListAPIErrorMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode)
}

func TestUnit_Departments_GetDepartmentHistoryV1_InvalidJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterHistoryInvalidJSONMock()

	result, resp, err := svc.GetDepartmentHistoryV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	errMsg := err.Error()
	assert.True(t, strings.Contains(errMsg, "failed to unmarshal page") || strings.Contains(errMsg, "extract results field"), "error should contain JSON parsing error")
}

func TestUnit_Departments_GetDepartmentHistoryV1_APIError(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterHistoryAPIErrorMock()

	result, resp, err := svc.GetDepartmentHistoryV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode)
}

func TestUnit_Departments_DeleteDepartmentsByIDV1_APIError(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteDepartmentsByIDErrorMock()

	req := &DeleteDepartmentsByIDRequest{IDs: []string{"1", "2"}}
	resp, err := svc.DeleteDepartmentsByIDV1(context.Background(), req)
	assert.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode)
}
