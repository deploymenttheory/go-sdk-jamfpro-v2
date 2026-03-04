package printers

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/printers/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh PrintersMock.
func setupMockService(t *testing.T) (*Service, *mocks.PrintersMock) {
	t.Helper()
	mock := mocks.NewPrintersMock()
	return NewService(mock), mock
}

// =============================================================================
// ListPrinters
// =============================================================================

func TestUnit_Printers_List_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListPrintersMock()

	result, resp, err := svc.List(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.Size)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "Office Printer", result.Results[0].Name)
	assert.Equal(t, 2, result.Results[1].ID)
	assert.Equal(t, "Lab Printer", result.Results[1].Name)
}

// =============================================================================
// GetPrinterByID
// =============================================================================

func TestUnit_Printers_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetPrinterByIDMock()

	result, resp, err := svc.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Office Printer", result.Name)
	assert.Equal(t, "Printers", result.Category)
	assert.Equal(t, "HP LaserJet", result.Model)
}

func TestUnit_Printers_GetByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "printer ID must be a positive integer")
}

func TestUnit_Printers_GetByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "printer ID must be a positive integer")
}

func TestUnit_Printers_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
}

// =============================================================================
// GetPrinterByName
// =============================================================================

func TestUnit_Printers_GetByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetPrinterByNameMock()

	result, resp, err := svc.GetByName(context.Background(), "Office Printer")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Office Printer", result.Name)
}

func TestUnit_Printers_GetByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "printer name is required")
}

// =============================================================================
// CreatePrinter
// =============================================================================

func TestUnit_Printers_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreatePrinterMock()

	req := &RequestPrinter{
		Name:     "New Printer",
		CUPSName: "New_Printer",
		URI:      "ipp://newprinter.example.com/ipp",
	}
	result, resp, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode())
	assert.Equal(t, 3, result.ID)
}

func TestUnit_Printers_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.Create(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Printers_Create_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestPrinter{Name: "Office Printer"}
	result, resp, err := svc.Create(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode())
}

// =============================================================================
// UpdatePrinterByID
// =============================================================================

func TestUnit_Printers_UpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdatePrinterByIDMock()

	req := &RequestPrinter{Name: "Office Printer Updated"}
	result, resp, err := svc.UpdateByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
}

func TestUnit_Printers_UpdateByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByID(context.Background(), 0, &RequestPrinter{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "printer ID must be a positive integer")
}

func TestUnit_Printers_UpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByID(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdatePrinterByName
// =============================================================================

func TestUnit_Printers_UpdateByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdatePrinterByNameMock()

	req := &RequestPrinter{Name: "Office Printer Updated"}
	result, resp, err := svc.UpdateByName(context.Background(), "Office Printer", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.ID)
}

func TestUnit_Printers_UpdateByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByName(context.Background(), "", &RequestPrinter{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "printer name is required")
}

func TestUnit_Printers_UpdateByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateByName(context.Background(), "Office Printer", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeletePrinterByID
// =============================================================================

func TestUnit_Printers_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeletePrinterByIDMock()

	resp, err := svc.DeleteByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_Printers_DeleteByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "printer ID must be a positive integer")
}

// =============================================================================
// DeletePrinterByName
// =============================================================================

func TestUnit_Printers_DeleteByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeletePrinterByNameMock()

	resp, err := svc.DeleteByName(context.Background(), "Office Printer")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_Printers_DeleteByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeleteByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "printer name is required")
}
