package departments

import (
	"context"
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

func TestUnitListDepartments_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListDepartmentsMock()

	result, resp, err := svc.ListDepartmentsV1(context.Background(), nil)
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

func TestUnitListDepartments_WithQueryParams(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListDepartmentsMock()

	params := map[string]string{"page": "0", "page-size": "50", "sort": "name:asc"}
	result, resp, err := svc.ListDepartmentsV1(context.Background(), params)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitListDepartments_WithRSQLFilter(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListDepartmentsRSQLMock()

	rsqlQuery := map[string]string{"filter": `name=="Sales"`}
	result, resp, err := svc.ListDepartmentsV1(context.Background(), rsqlQuery)
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

func TestUnitGetDepartmentByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDepartmentMock()

	result, resp, err := svc.GetDepartmentByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Engineering", result.Name)
}

func TestUnitGetDepartmentByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetDepartmentByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "department ID is required")
}

func TestUnitGetDepartmentByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetDepartmentByIDV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestUnitCreateDepartment_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateDepartmentMock()

	req := &RequestDepartment{Name: "Marketing"}
	result, resp, err := svc.CreateDepartmentV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "3", result.ID)
	assert.Contains(t, result.Href, "/api/v1/departments/3")
}

func TestUnitCreateDepartment_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateDepartmentV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateDepartment_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestDepartment{Name: "Duplicate"}
	result, resp, err := svc.CreateDepartmentV1(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

func TestUnitUpdateDepartmentByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateDepartmentMock()

	req := &RequestDepartment{Name: "Engineering Updated"}
	result, resp, err := svc.UpdateDepartmentByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Engineering Updated", result.Name)
}

func TestUnitUpdateDepartmentByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateDepartmentByIDV1(context.Background(), "", &RequestDepartment{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnitUpdateDepartmentByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateDepartmentByIDV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitDeleteDepartmentByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteDepartmentMock()

	resp, err := svc.DeleteDepartmentByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitDeleteDepartmentByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteDepartmentByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "department ID is required")
}

func TestUnitGetDepartmentHistoryV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDepartmentHistoryMock()

	result, resp, err := svc.GetDepartmentHistoryV1(context.Background(), "1", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", string(result.Results[0].ID))
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "Department created", result.Results[0].Note)
}

func TestUnitGetDepartmentHistoryV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetDepartmentHistoryV1(context.Background(), "", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "department ID is required")
}

func TestUnitAddDepartmentHistoryNotesV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddDepartmentHistoryNotesMock()

	req := &AddHistoryNotesRequest{Note: "Added via SDK"}
	resp, err := svc.AddDepartmentHistoryNotesV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
}

func TestUnitAddDepartmentHistoryNotesV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.AddDepartmentHistoryNotesV1(context.Background(), "", &AddHistoryNotesRequest{Note: "x"})
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "department ID is required")
}

func TestUnitAddDepartmentHistoryNotesV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.AddDepartmentHistoryNotesV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request body is required")
}
