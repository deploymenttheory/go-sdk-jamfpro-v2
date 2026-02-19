package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/printers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_Printers_Lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → GetByName → UpdateByID →
// UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_Printers_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Printers
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test printer")

	printerName := uniqueName("acc-test-printer")
	createReq := &printers.RequestPrinter{
		Name:     printerName,
		CUPSName: "acc_test_printer",
		URI:      "ipp://printer.example.com/ipp",
		Location: "Test Lab",
		Model:    "Test Printer Model",
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.CreatePrinter(ctx1, createReq)
	require.NoError(t, err, "CreatePrinter should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created printer ID should be a positive integer")

	printerID := created.ID
	acc.LogTestSuccess(t, "Printer created with ID=%d name=%q", printerID, printerName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeletePrinterByID(cleanupCtx, printerID)
		acc.LogCleanupDeleteError(t, "printer", fmt.Sprintf("%d", printerID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new printer appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing printers to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListPrinters(ctx2)
	require.NoError(t, err, "ListPrinters should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, p := range list.Results {
		if p.ID == printerID {
			found = true
			assert.Equal(t, printerName, p.Name)
			break
		}
	}
	assert.True(t, found, "newly created printer should appear in list")
	acc.LogTestSuccess(t, "Printer ID=%d found in list (%d total)", printerID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching printer by ID=%d", printerID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetPrinterByID(ctx3, printerID)
	require.NoError(t, err, "GetPrinterByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, printerID, fetched.ID)
	assert.Equal(t, printerName, fetched.Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.ID, fetched.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching printer by name=%q", printerName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetPrinterByName(ctx4, printerName)
	require.NoError(t, err, "GetPrinterByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, printerID, fetchedByName.ID)
	assert.Equal(t, printerName, fetchedByName.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.ID, fetchedByName.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := uniqueName("acc-test-printer-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating printer ID=%d to name=%q", printerID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &printers.RequestPrinter{
		Name:     updatedName,
		CUPSName: "acc_test_printer_updated",
		URI:      "ipp://printer.example.com/ipp",
		Location: "Test Lab",
		Model:    "Test Printer Model",
	}
	updated, updateResp, err := svc.UpdatePrinterByID(ctx5, printerID, updateReq)
	require.NoError(t, err, "UpdatePrinterByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating printer name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &printers.RequestPrinter{
		Name:     printerName,
		CUPSName: "acc_test_printer",
		URI:      "ipp://printer.example.com/ipp",
		Location: "Test Lab",
		Model:    "Test Printer Model",
	}
	reverted, revertResp, err := svc.UpdatePrinterByName(ctx6, updatedName, revertReq)
	require.NoError(t, err, "UpdatePrinterByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode)

	// ------------------------------------------------------------------
	// 7. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	verified, verifyResp, err := svc.GetPrinterByID(ctx7, printerID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, printerName, verified.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.Name)

	// ------------------------------------------------------------------
	// 8. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting printer ID=%d", printerID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeletePrinterByID(ctx8, printerID)
	require.NoError(t, err, "DeletePrinterByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Printer ID=%d deleted", printerID)
}

// =============================================================================
// TestAcceptance_Printers_DeleteByName creates a printer then deletes by name.
// =============================================================================

func TestAcceptance_Printers_DeleteByName(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Printers
	ctx := context.Background()

	printerName := uniqueName("acc-test-printer-dbn")
	createReq := &printers.RequestPrinter{
		Name:     printerName,
		CUPSName: "acc_test_printer_dbn",
		URI:      "ipp://printer.example.com/ipp",
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.CreatePrinter(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	printerID := created.ID
	acc.LogTestSuccess(t, "Created printer ID=%d name=%q for delete-by-name test", printerID, printerName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeletePrinterByID(cleanupCtx, printerID)
		acc.LogCleanupDeleteError(t, "printer", fmt.Sprintf("%d", printerID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeletePrinterByName(ctx2, printerName)
	require.NoError(t, err, "DeletePrinterByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Printer %q deleted by name", printerName)
}

// =============================================================================
// TestAcceptance_Printers_ValidationErrors tests client-side validation
// without making any network calls.
// =============================================================================

func TestAcceptance_Printers_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Printers

	t.Run("GetPrinterByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetPrinterByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "printer ID must be a positive integer")
	})

	t.Run("GetPrinterByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetPrinterByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "printer name is required")
	})

	t.Run("CreatePrinter_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreatePrinter(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdatePrinterByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdatePrinterByID(context.Background(), 0, &printers.RequestPrinter{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "printer ID must be a positive integer")
	})

	t.Run("UpdatePrinterByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdatePrinterByName(context.Background(), "", &printers.RequestPrinter{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "printer name is required")
	})

	t.Run("DeletePrinterByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeletePrinterByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "printer ID must be a positive integer")
	})

	t.Run("DeletePrinterByName_EmptyName", func(t *testing.T) {
		_, err := svc.DeletePrinterByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "printer name is required")
	})
}
