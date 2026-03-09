package inventory_preload

import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/inventory_preload/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*InventoryPreload, *mocks.InventoryPreloadMock) {
	t.Helper()
	mock := mocks.NewInventoryPreloadMock()
	return NewInventoryPreload(mock), mock
}

func TestUnit_InventoryPreload_NewService(t *testing.T) {
	mock := mocks.NewInventoryPreloadMock()
	svc := NewInventoryPreload(mock)
	require.NotNil(t, svc)
}

// -----------------------------------------------------------------------------
// CSV operations
// -----------------------------------------------------------------------------

func TestUnit_InventoryPreload_CreateFromCSV_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateFromCSVMock()

	csvContent := []byte("Serial Number,Device Type\nSN001,Computer")
	result, resp, err := svc.CreateFromCSV(context.Background(), bytes.NewReader(csvContent), int64(len(csvContent)), "test.csv")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	require.Len(t, result, 2)
	assert.Equal(t, "1", result[0].ID)
	assert.Equal(t, "2", result[1].ID)
}

func TestUnit_InventoryPreload_CreateFromCSV_EmptyFileName(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateFromCSVMock()

	csvContent := []byte("Serial Number,Device Type\nSN001,Computer")
	result, resp, err := svc.CreateFromCSV(context.Background(), bytes.NewReader(csvContent), int64(len(csvContent)), "")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_InventoryPreload_CreateFromCSV_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	csvContent := []byte("Serial Number,Device Type\nSN001,Computer")
	_, resp, err := svc.CreateFromCSV(context.Background(), bytes.NewReader(csvContent), int64(len(csvContent)), "test.csv")
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_InventoryPreload_CreateFromCSVFile_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateFromCSVMock()

	tmpDir := t.TempDir()
	csvPath := filepath.Join(tmpDir, "preload.csv")
	err := os.WriteFile(csvPath, []byte("Serial Number,Device Type\nSN001,Computer"), 0644)
	require.NoError(t, err)

	result, resp, err := svc.CreateFromCSVFile(context.Background(), csvPath)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	require.Len(t, result, 2)
}

func TestUnit_InventoryPreload_CreateFromCSVFile_FileNotFound(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.CreateFromCSVFile(context.Background(), "/nonexistent/path/file.csv")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "open CSV file")
}

func TestUnit_InventoryPreload_GetCSVTemplate_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetCSVTemplateMock()

	body, resp, err := svc.GetCSVTemplate(context.Background())
	require.NoError(t, err)
	require.NotNil(t, body)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Contains(t, string(body), "Serial Number")
}

func TestUnit_InventoryPreload_GetCSVTemplate_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	_, resp, err := svc.GetCSVTemplate(context.Background())
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_InventoryPreload_ValidateCSV_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterValidateCSVMock()

	csvContent := []byte("Serial Number,Device Type\nSN001,Computer")
	result, resp, err := svc.ValidateCSV(context.Background(), bytes.NewReader(csvContent), int64(len(csvContent)), "test.csv")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 5, result.RecordCount)
}

func TestUnit_InventoryPreload_ValidateCSV_EmptyFileName(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterValidateCSVMock()

	csvContent := []byte("Serial Number,Device Type\nSN001,Computer")
	result, resp, err := svc.ValidateCSV(context.Background(), bytes.NewReader(csvContent), int64(len(csvContent)), "")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_InventoryPreload_ValidateCSV_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	csvContent := []byte("Serial Number,Device Type\nSN001,Computer")
	_, resp, err := svc.ValidateCSV(context.Background(), bytes.NewReader(csvContent), int64(len(csvContent)), "test.csv")
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_InventoryPreload_ValidateCSVFile_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterValidateCSVMock()

	tmpDir := t.TempDir()
	csvPath := filepath.Join(tmpDir, "validate.csv")
	err := os.WriteFile(csvPath, []byte("Serial Number,Device Type\nSN001,Computer"), 0644)
	require.NoError(t, err)

	result, resp, err := svc.ValidateCSVFile(context.Background(), csvPath)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_InventoryPreload_ValidateCSVFile_FileNotFound(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.ValidateCSVFile(context.Background(), "/nonexistent/path/file.csv")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "open CSV file")
}

// -----------------------------------------------------------------------------
// EA columns
// -----------------------------------------------------------------------------

func TestUnit_InventoryPreload_GetEAColumns_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetEAColumnsMock()

	result, resp, err := svc.GetEAColumns(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "Custom EA 1", result.Results[0].Name)
	assert.Equal(t, "Custom Extension Attribute 1", result.Results[0].FullName)
}

func TestUnit_InventoryPreload_GetEAColumns_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	_, resp, err := svc.GetEAColumns(context.Background())
	assert.Error(t, err)
	assert.Nil(t, resp)
}

// -----------------------------------------------------------------------------
// Export
// -----------------------------------------------------------------------------

func TestUnit_InventoryPreload_Export_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterExportMock()

	data, resp, err := svc.Export(context.Background(), nil, nil, constants.ApplicationJSON)
	require.NoError(t, err)
	require.NotNil(t, data)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Contains(t, string(data), "totalCount")
}

func TestUnit_InventoryPreload_Export_EmptyAcceptType(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterExportMock()

	data, resp, err := svc.Export(context.Background(), nil, nil, "")
	require.NoError(t, err)
	require.NotNil(t, data)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_InventoryPreload_Export_WithQueryAndBody(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterExportMock()

	query := map[string]string{"page": "0", "page-size": "100"}
	req := &ExportRequest{Page: intPtr(0), PageSize: intPtr(100)}
	data, resp, err := svc.Export(context.Background(), query, req, constants.ApplicationJSON)
	require.NoError(t, err)
	require.NotNil(t, data)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_InventoryPreload_Export_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	_, resp, err := svc.Export(context.Background(), nil, nil, constants.ApplicationJSON)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "export inventory preload")
}

// -----------------------------------------------------------------------------
// History
// -----------------------------------------------------------------------------

func TestUnit_InventoryPreload_ListHistory_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListHistoryMock()

	result, resp, err := svc.ListHistory(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "admin", result.Results[0].Username)
	assert.Equal(t, "CSV import completed", result.Results[0].Note)
}

func TestUnit_InventoryPreload_ListHistory_WithRSQL(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListHistoryMock()

	rsqlQuery := map[string]string{"page": "0", "page-size": "50"}
	result, resp, err := svc.ListHistory(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, rsqlQuery, mock.LastRSQLQuery)
}

func TestUnit_InventoryPreload_ListHistory_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	_, resp, err := svc.ListHistory(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_InventoryPreload_AddHistoryNote_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterAddHistoryNoteMock()

	req := &AddHistoryNoteRequest{Note: "Test note"}
	result, resp, err := svc.AddHistoryNote(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "3", result.ID)
}

func TestUnit_InventoryPreload_AddHistoryNote_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.AddHistoryNote(context.Background(), nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "request body is required")
}

func TestUnit_InventoryPreload_AddHistoryNote_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.AddHistoryNote(context.Background(), &AddHistoryNoteRequest{Note: "x"})
	assert.Error(t, err)
}

// -----------------------------------------------------------------------------
// Records CRUD
// -----------------------------------------------------------------------------

func TestUnit_InventoryPreload_ListRecords_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListRecordsMock()

	result, resp, err := svc.ListRecords(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "SN001", result.Results[0].SerialNumber)
	assert.Equal(t, "Computer", result.Results[0].DeviceType)
}

func TestUnit_InventoryPreload_ListRecords_WithRSQL(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListRecordsMock()

	rsqlQuery := map[string]string{"filter": `serialNumber=="SN001"`, "sort": "id:asc"}
	result, resp, err := svc.ListRecords(context.Background(), rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, rsqlQuery, mock.LastRSQLQuery)
}

func TestUnit_InventoryPreload_ListRecords_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	_, resp, err := svc.ListRecords(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "list inventory preload records")
}

func TestUnit_InventoryPreload_ListRecords_InvalidJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListRecordsInvalidJSONMock()

	_, resp, err := svc.ListRecords(context.Background(), nil)
	assert.Error(t, err)
	require.NotNil(t, resp)
	assert.Contains(t, err.Error(), "mergePage failed")
}

func TestUnit_InventoryPreload_CreateRecord_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateRecordMock()

	record := &InventoryPreloadRecord{
		SerialNumber: "SN003",
		DeviceType:   "Computer",
	}
	result, resp, err := svc.CreateRecord(context.Background(), record)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, "3", result.ID)
	assert.Contains(t, result.Href, "/api/v2/inventory-preload/records/3")
}

func TestUnit_InventoryPreload_CreateRecord_NilRecord(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.CreateRecord(context.Background(), nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "record is required")
}

func TestUnit_InventoryPreload_CreateRecord_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.CreateRecord(context.Background(), &InventoryPreloadRecord{SerialNumber: "SN", DeviceType: "Computer"})
	assert.Error(t, err)
}

func TestUnit_InventoryPreload_DeleteAllRecords_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteAllRecordsMock()

	resp, err := svc.DeleteAllRecords(context.Background())
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_InventoryPreload_DeleteAllRecords_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteAllRecords(context.Background())
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_InventoryPreload_GetRecordByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetRecordByIDMock("1")

	result, resp, err := svc.GetRecordByID(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "SN001", result.SerialNumber)
	assert.Equal(t, "Computer", result.DeviceType)
}

func TestUnit_InventoryPreload_GetRecordByID_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.GetRecordByID(context.Background(), "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_InventoryPreload_GetRecordByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock("999")

	_, resp, err := svc.GetRecordByID(context.Background(), "999")
	assert.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_InventoryPreload_GetRecordByID_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	_, resp, err := svc.GetRecordByID(context.Background(), "1")
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestUnit_InventoryPreload_UpdateRecord_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateRecordMock("1")

	record := &InventoryPreloadRecord{
		SerialNumber: "SN001",
		DeviceType:   "Computer",
	}
	result, resp, err := svc.UpdateRecord(context.Background(), "1", record)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
}

func TestUnit_InventoryPreload_UpdateRecord_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.UpdateRecord(context.Background(), "", &InventoryPreloadRecord{SerialNumber: "SN", DeviceType: "Computer"})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_InventoryPreload_UpdateRecord_NilRecord(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.UpdateRecord(context.Background(), "1", nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "record is required")
}

func TestUnit_InventoryPreload_UpdateRecord_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock("999")

	record := &InventoryPreloadRecord{SerialNumber: "SN", DeviceType: "Computer"}
	_, resp, err := svc.UpdateRecord(context.Background(), "999", record)
	assert.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_InventoryPreload_UpdateRecord_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	_, _, err := svc.UpdateRecord(context.Background(), "1", &InventoryPreloadRecord{SerialNumber: "SN", DeviceType: "Computer"})
	assert.Error(t, err)
}

func TestUnit_InventoryPreload_DeleteRecord_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteRecordMock("1")

	resp, err := svc.DeleteRecord(context.Background(), "1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_InventoryPreload_DeleteRecord_EmptyID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteRecord(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_InventoryPreload_DeleteRecord_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock("999")

	resp, err := svc.DeleteRecord(context.Background(), "999")
	assert.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_InventoryPreload_DeleteRecord_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteRecord(context.Background(), "1")
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func intPtr(i int) *int { return &i }
