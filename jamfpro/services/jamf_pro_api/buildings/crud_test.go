package buildings

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/buildings/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.BuildingsMock) {
	t.Helper()
	mock := mocks.NewBuildingsMock()
	return NewService(mock), mock
}

func TestUnitListBuildings_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListBuildingsMock()

	result, resp, err := svc.ListBuildingsV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "Main Office", result.Results[0].Name)
	assert.Equal(t, "2", result.Results[1].ID)
	assert.Equal(t, "Warehouse", result.Results[1].Name)
}

func TestUnitListBuildings_WithrsqlQuery(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListBuildingsMock()

	params := map[string]string{"page": "0", "page-size": "50", "sort": "name:asc"}
	result, resp, err := svc.ListBuildingsV1(context.Background(), params)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitListBuildings_WithRSQLFilter(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListBuildingsRSQLMock()

	rsqlQuery := map[string]string{"filter": `name=="Warehouse"`}
	result, resp, err := svc.ListBuildingsV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "2", result.Results[0].ID)
	assert.Equal(t, "Warehouse", result.Results[0].Name)
	assert.Equal(t, rsqlQuery, mock.LastRSQLQuery)
}

func TestUnitGetBuildingByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetBuildingMock()

	result, resp, err := svc.GetBuildingByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Main Office", result.Name)
	assert.Equal(t, "123 Main St", result.StreetAddress1)
	assert.Equal(t, "Austin", result.City)
}

func TestUnitGetBuildingByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetBuildingByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "building ID is required")
}

func TestUnitGetBuildingByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetBuildingByIDV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestUnitCreateBuilding_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateBuildingMock()

	req := &RequestBuilding{
		Name:           "Marketing",
		StreetAddress1: "100 Marketing Way",
		City:           "Austin",
		StateProvince:  "TX",
		Country:        "United States",
	}
	result, resp, err := svc.CreateBuildingV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "3", result.ID)
	assert.Contains(t, result.Href, "/api/v1/buildings/3")
}

func TestUnitCreateBuilding_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateBuildingV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateBuilding_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestBuilding{Name: "Duplicate"}
	result, resp, err := svc.CreateBuildingV1(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

func TestUnitUpdateBuildingByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateBuildingMock()

	req := &RequestBuilding{
		Name:           "Main Office Updated",
		StreetAddress1: "789 New St",
		City:           "Austin",
		StateProvince:  "TX",
		ZipPostalCode:  "78702",
		Country:        "United States",
	}
	result, resp, err := svc.UpdateBuildingByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Main Office Updated", result.Name)
}

func TestUnitUpdateBuildingByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateBuildingByIDV1(context.Background(), "", &RequestBuilding{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnitUpdateBuildingByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateBuildingByIDV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitDeleteBuildingByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteBuildingMock()

	resp, err := svc.DeleteBuildingByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitDeleteBuildingByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteBuildingByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "building ID is required")
}

func TestUnitDeleteBuildingsByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteBuildingsByIDMock()

	req := &DeleteBuildingsByIDRequest{IDs: []string{"1", "2"}}
	resp, err := svc.DeleteBuildingsByIDV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnitDeleteBuildingsByID_EmptyIDs(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteBuildingsByIDV1(context.Background(), &DeleteBuildingsByIDRequest{IDs: []string{}})
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "ids are required")
}

func TestUnitDeleteBuildingsByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteBuildingsByIDV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "ids are required")
}

func TestUnitGetBuildingHistoryV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetBuildingHistoryMock()

	result, resp, err := svc.GetBuildingHistoryV1(context.Background(), "1", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", string(result.Results[0].ID))
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "Building created", result.Results[0].Note)
}

func TestUnitGetBuildingHistoryV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetBuildingHistoryV1(context.Background(), "", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "building ID is required")
}

func TestUnitAddBuildingHistoryNotesV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddBuildingHistoryNotesMock()

	req := &AddHistoryNotesRequest{Note: "Added via SDK"}
	resp, err := svc.AddBuildingHistoryNotesV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
}

func TestUnitAddBuildingHistoryNotesV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.AddBuildingHistoryNotesV1(context.Background(), "", &AddHistoryNotesRequest{Note: "x"})
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "building ID is required")
}

func TestUnitAddBuildingHistoryNotesV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.AddBuildingHistoryNotesV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request body is required")
}
