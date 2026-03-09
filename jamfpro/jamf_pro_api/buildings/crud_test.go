package buildings

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/buildings/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Buildings, *mocks.BuildingsMock) {
	t.Helper()
	mock := mocks.NewBuildingsMock()
	return NewBuildings(mock), mock
}

func TestUnit_Buildings_List_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListBuildingsMock()

	result, resp, err := svc.ListV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "Main Office", result.Results[0].Name)
	assert.Equal(t, "2", result.Results[1].ID)
	assert.Equal(t, "Warehouse", result.Results[1].Name)
}

func TestUnit_Buildings_List_WithrsqlQuery(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListBuildingsMock()

	params := map[string]string{"page": "0", "page-size": "50", "sort": "name:asc"}
	result, resp, err := svc.ListV1(context.Background(), params)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_Buildings_List_WithRSQLFilter(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListBuildingsRSQLMock()

	rsqlQuery := map[string]string{"filter": `name=="Warehouse"`}
	result, resp, err := svc.ListV1(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "2", result.Results[0].ID)
	assert.Equal(t, "Warehouse", result.Results[0].Name)
	assert.Equal(t, rsqlQuery, mock.LastRSQLQuery)
}

func TestUnit_Buildings_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetBuildingMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Main Office", result.Name)
	assert.Equal(t, "123 Main St", result.StreetAddress1)
	assert.Equal(t, "Austin", result.City)
}

func TestUnit_Buildings_GetByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "building ID is required")
}

func TestUnit_Buildings_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByIDV1(context.Background(), "999")
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_Buildings_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateBuildingMock()

	req := &RequestBuilding{
		Name:           "Marketing",
		StreetAddress1: "100 Marketing Way",
		City:           "Austin",
		StateProvince:  "TX",
		Country:        "United States",
	}
	result, resp, err := svc.CreateV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "3", result.ID)
	assert.Contains(t, result.Href, "/api/v1/buildings/3")
}

func TestUnit_Buildings_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreateV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Buildings_Create_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestBuilding{Name: "Duplicate"}
	result, resp, err := svc.CreateV1(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode())
}

func TestUnit_Buildings_UpdateByID_Success(t *testing.T) {
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
	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Main Office Updated", result.Name)
}

func TestUnit_Buildings_UpdateByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByIDV1(context.Background(), "", &RequestBuilding{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_Buildings_UpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByIDV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Buildings_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteBuildingMock()

	resp, err := svc.DeleteByIDV1(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_Buildings_DeleteByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByIDV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "building ID is required")
}

func TestUnit_Buildings_DeleteMultipleByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteBuildingsByIDMock()

	req := &DeleteBuildingsByIDRequest{IDs: []string{"1", "2"}}
	resp, err := svc.DeleteBuildingsByIDV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_Buildings_DeleteMultipleByID_EmptyIDs(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteBuildingsByIDV1(context.Background(), &DeleteBuildingsByIDRequest{IDs: []string{}})
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "ids are required")
}

func TestUnit_Buildings_DeleteMultipleByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteBuildingsByIDV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "ids are required")
}

func TestUnit_Buildings_GetHistoryV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetBuildingHistoryMock()

	result, resp, err := svc.GetBuildingHistoryV1(context.Background(), "1", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "Building created", result.Results[0].Note)
	require.NotNil(t, result.Results[0].Details)
	assert.Equal(t, "Initial creation", *result.Results[0].Details)
}

func TestUnit_Buildings_GetHistoryV1_NullDetails(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetBuildingHistoryNullDetailsMock()

	result, resp, err := svc.GetBuildingHistoryV1(context.Background(), "1", nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Nil(t, result.Results[0].Details)
}

func TestUnit_Buildings_GetHistoryV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetBuildingHistoryV1(context.Background(), "", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "building ID is required")
}

func TestUnit_Buildings_AddHistoryNotesV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddBuildingHistoryNotesMock()

	req := &AddHistoryNotesRequest{Note: "Added via SDK"}
	resp, err := svc.AddBuildingHistoryNotesV1(context.Background(), "1", req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
}

func TestUnit_Buildings_AddHistoryNotesV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.AddBuildingHistoryNotesV1(context.Background(), "", &AddHistoryNotesRequest{Note: "x"})
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "building ID is required")
}

func TestUnit_Buildings_AddHistoryNotesV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.AddBuildingHistoryNotesV1(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request body is required")
}

// TestUnit_Buildings_List_NoMockRegistered verifies error handling when no mock is registered.
func TestUnit_Buildings_List_NoMockRegistered(t *testing.T) {
	svc, _ := setupMockService(t)
	// Do not register any mock

	result, resp, err := svc.ListV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to list buildings")
}

// TestUnit_Buildings_DeleteMultipleByID_NoMockRegistered verifies error when no mock is registered.
func TestUnit_Buildings_DeleteMultipleByID_NoMockRegistered(t *testing.T) {
	svc, _ := setupMockService(t)
	// Do not register DeleteBuildingsByIDMock

	req := &DeleteBuildingsByIDRequest{IDs: []string{"1", "2"}}
	resp, err := svc.DeleteBuildingsByIDV1(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
}

// TestUnit_Buildings_ExportV1_Success tests successful buildings export.
func TestUnit_Buildings_ExportV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterExportBuildingsMock()

	data, resp, err := svc.ExportV1(context.Background(), nil, nil, mime.ApplicationJSON)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, data)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Contains(t, string(data), "totalCount")
}

// TestUnit_Buildings_ExportV1_EmptyAcceptType tests default to application/json.
func TestUnit_Buildings_ExportV1_EmptyAcceptType(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterExportBuildingsMock()

	data, resp, err := svc.ExportV1(context.Background(), nil, nil, "")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, data)
	assert.Equal(t, 200, resp.StatusCode())
}

// TestUnit_Buildings_ExportV1_WithQueryAndBody tests export with query params and body.
func TestUnit_Buildings_ExportV1_WithQueryAndBody(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterExportBuildingsMock()

	query := map[string]string{"page": "0", "page-size": "100"}
	req := &ExportRequest{Page: intPtr(0), PageSize: intPtr(100)}
	data, resp, err := svc.ExportV1(context.Background(), query, req, mime.ApplicationJSON)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, data)
	assert.Equal(t, 200, resp.StatusCode())
}

// TestUnit_Buildings_ExportV1_NoMockRegistered verifies error when no mock is registered.
func TestUnit_Buildings_ExportV1_NoMockRegistered(t *testing.T) {
	svc, _ := setupMockService(t)

	data, resp, err := svc.ExportV1(context.Background(), nil, nil, mime.ApplicationJSON)
	assert.Error(t, err)
	assert.Nil(t, data)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to export buildings")
}

// TestUnit_Buildings_ExportHistoryV1_Success tests successful building history export.
func TestUnit_Buildings_ExportHistoryV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterExportBuildingHistoryMock()

	data, resp, err := svc.ExportHistoryV1(context.Background(), "1", nil, nil, mime.ApplicationJSON)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, data)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Contains(t, string(data), "totalCount")
}

// TestUnit_Buildings_ExportHistoryV1_EmptyAcceptType tests default to application/json.
func TestUnit_Buildings_ExportHistoryV1_EmptyAcceptType(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterExportBuildingHistoryMock()

	data, resp, err := svc.ExportHistoryV1(context.Background(), "1", nil, nil, "")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.NotNil(t, data)
	assert.Equal(t, 200, resp.StatusCode())
}

// TestUnit_Buildings_ExportHistoryV1_EmptyID tests validation error for empty ID.
func TestUnit_Buildings_ExportHistoryV1_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	data, resp, err := svc.ExportHistoryV1(context.Background(), "", nil, nil, mime.ApplicationJSON)
	assert.Error(t, err)
	assert.Nil(t, data)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "building ID is required")
}

// TestUnit_Buildings_ExportHistoryV1_NoMockRegistered verifies error when no mock is registered.
func TestUnit_Buildings_ExportHistoryV1_NoMockRegistered(t *testing.T) {
	svc, _ := setupMockService(t)

	data, resp, err := svc.ExportHistoryV1(context.Background(), "1", nil, nil, mime.ApplicationJSON)
	assert.Error(t, err)
	assert.Nil(t, data)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to export building history")
}

func intPtr(i int) *int { return &i }
