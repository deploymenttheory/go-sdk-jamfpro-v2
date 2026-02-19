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

func TestUnitListPrinters_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListPrintersMock()

	result, resp, err := svc.ListPrinters(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
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

func TestUnitGetPrinterByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetPrinterByIDMock()

	result, resp, err := svc.GetPrinterByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Office Printer", result.Name)
	assert.Equal(t, "Printers", result.Category)
	assert.Equal(t, "HP LaserJet", result.Model)
}

func TestUnitGetPrinterByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetPrinterByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "printer ID must be a positive integer")
}

func TestUnitGetPrinterByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetPrinterByID(context.Background(), -1)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "printer ID must be a positive integer")
}

func TestUnitGetPrinterByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	result, resp, err := svc.GetPrinterByID(context.Background(), 999)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

// =============================================================================
// GetPrinterByName
// =============================================================================

func TestUnitGetPrinterByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetPrinterByNameMock()

	result, resp, err := svc.GetPrinterByName(context.Background(), "Office Printer")
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Office Printer", result.Name)
}

func TestUnitGetPrinterByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetPrinterByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "printer name is required")
}

// =============================================================================
// CreatePrinter
// =============================================================================

func TestUnitCreatePrinter_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreatePrinterMock()

	req := &RequestPrinter{
		Name:     "New Printer",
		CUPSName: "New_Printer",
		URI:      "ipp://newprinter.example.com/ipp",
	}
	result, resp, err := svc.CreatePrinter(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 3, result.ID)
}

func TestUnitCreatePrinter_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.CreatePrinter(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreatePrinter_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestPrinter{Name: "Office Printer"}
	result, resp, err := svc.CreatePrinter(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 409, resp.StatusCode)
}

// =============================================================================
// UpdatePrinterByID
// =============================================================================

func TestUnitUpdatePrinterByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdatePrinterByIDMock()

	req := &RequestPrinter{Name: "Office Printer Updated"}
	result, resp, err := svc.UpdatePrinterByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnitUpdatePrinterByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdatePrinterByID(context.Background(), 0, &RequestPrinter{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "printer ID must be a positive integer")
}

func TestUnitUpdatePrinterByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdatePrinterByID(context.Background(), 1, nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdatePrinterByName
// =============================================================================

func TestUnitUpdatePrinterByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdatePrinterByNameMock()

	req := &RequestPrinter{Name: "Office Printer Updated"}
	result, resp, err := svc.UpdatePrinterByName(context.Background(), "Office Printer", req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnitUpdatePrinterByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdatePrinterByName(context.Background(), "", &RequestPrinter{Name: "x"})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "printer name is required")
}

func TestUnitUpdatePrinterByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdatePrinterByName(context.Background(), "Office Printer", nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeletePrinterByID
// =============================================================================

func TestUnitDeletePrinterByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeletePrinterByIDMock()

	resp, err := svc.DeletePrinterByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeletePrinterByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeletePrinterByID(context.Background(), 0)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "printer ID must be a positive integer")
}

// =============================================================================
// DeletePrinterByName
// =============================================================================

func TestUnitDeletePrinterByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeletePrinterByNameMock()

	resp, err := svc.DeletePrinterByName(context.Background(), "Office Printer")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeletePrinterByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.DeletePrinterByName(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "printer name is required")
}
