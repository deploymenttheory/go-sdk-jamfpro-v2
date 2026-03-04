package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/inventory_preload"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"resty.dev/v3"
)

// =============================================================================
// Acceptance Tests: Inventory Preload
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ListRecords(ctx, rsqlQuery) - Lists paginated inventory preload records
//   • CreateRecord(ctx, record) - Creates a single inventory preload record
//   • GetRecordByID(ctx, id) - Gets a record by ID
//   • UpdateRecord(ctx, id, record) - Updates a record by ID
//   • DeleteRecord(ctx, id) - Deletes a record by ID
//   • DeleteAllRecords(ctx) - Deletes all records
//   • GetCSVTemplate(ctx) - Downloads the CSV template
//   • CreateFromCSV/CreateFromCSVFile - CSV bulk upload
//   • ValidateCSV/ValidateCSVFile - CSV validation
//   • GetEAColumns(ctx) - Gets EA columns
//   • Export(ctx, ...) - Exports records
//   • ListHistory/AddHistoryNote - History management
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 1: Full Record Lifecycle (Create → Get → Update → List → Delete)
//     -- Tests: TestAcceptance_InventoryPreload_record_lifecycle
//
//   ✓ Pattern 5: RSQL Filter
//     -- Included within lifecycle test
//
//   ✓ Pattern 3: Read-Only (CSV template, EA columns)
//     -- Tests: TestAcceptance_InventoryPreload_csv_template
//
//   ✓ Pattern 7: Validation Errors
//     -- Tests: TestAcceptance_InventoryPreload_validation_errors
//
// =============================================================================

// TestAcceptance_InventoryPreload_record_lifecycle exercises the full record CRUD.
func TestAcceptance_InventoryPreload_record_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.InventoryPreload
	ctx := context.Background()

	// 1. Create record
	acc.LogTestStage(t, "CreateRecord", "Creating inventory preload record")

	serial := fmt.Sprintf("SDKV2-%d", time.Now().UnixNano()%10000000)
	createReq := &inventory_preload.InventoryPreloadRecord{
		SerialNumber: serial,
		DeviceType:   "Computer",
	}

	created, createResp, err := svc.CreateRecord(ctx, createReq)
	require.NoError(t, err, "CreateRecord should not return an error")
	require.NotNil(t, created)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	recordID := created.ID
	acc.LogTestSuccess(t, "Created inventory preload record ID=%s serialNumber=%s", recordID, serial)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteRecord(cleanupCtx, recordID)
		acc.LogCleanupDeleteError(t, "inventory preload record", recordID, delErr)
	})

	// 2. GetRecordByID (with retry for eventual consistency)
	acc.LogTestStage(t, "GetRecordByID", "Getting record ID=%s", recordID)

	var fetched *inventory_preload.InventoryPreloadRecord
	var fetchResp *resty.Response
	err = acc.RetryOnNotFound(t, 3, 500*time.Millisecond, func() error {
		var getErr error
		fetched, fetchResp, getErr = svc.GetRecordByID(ctx, recordID)
		return getErr
	})
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, serial, fetched.SerialNumber)
	assert.Equal(t, "Computer", fetched.DeviceType)
	acc.LogTestSuccess(t, "GetRecordByID: ID=%s serialNumber=%s", recordID, fetched.SerialNumber)

	// 3. UpdateRecord
	acc.LogTestStage(t, "UpdateRecord", "Updating record ID=%s", recordID)

	updatedSerial := fmt.Sprintf("SDKV2U-%d", time.Now().UnixNano()%10000000)
	updateReq := &inventory_preload.InventoryPreloadRecord{
		SerialNumber: updatedSerial,
		DeviceType:   "Computer",
	}

	updated, updateResp, err := svc.UpdateRecord(ctx, recordID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode)
	assert.Equal(t, updatedSerial, updated.SerialNumber)
	acc.LogTestSuccess(t, "UpdateRecord: ID=%s serialNumber=%s", recordID, updated.SerialNumber)

	// 4. ListRecords with RSQL filter to verify the updated record
	acc.LogTestStage(t, "ListRecords", "Filtering records by serialNumber=%s", updatedSerial)

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`serialNumber=="%s"`, updatedSerial),
	}

	list, listResp, err := svc.ListRecords(ctx, rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, r := range list.Results {
		if r.ID == recordID {
			found = true
			break
		}
	}
	assert.True(t, found, "updated record should appear in RSQL-filtered results")
	acc.LogTestSuccess(t, "ListRecords: totalCount=%d found=%v", list.TotalCount, found)

	// 5. DeleteRecord
	acc.LogTestStage(t, "DeleteRecord", "Deleting record ID=%s", recordID)

	deleteResp, err := svc.DeleteRecord(ctx, recordID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "DeleteRecord: ID=%s deleted", recordID)
}

// TestAcceptance_InventoryPreload_csv_template verifies the CSV template download.
func TestAcceptance_InventoryPreload_csv_template(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.InventoryPreload
	ctx := context.Background()

	acc.LogTestStage(t, "GetCSVTemplate", "Downloading inventory preload CSV template")

	template, resp, err := svc.GetCSVTemplate(ctx)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotEmpty(t, template, "CSV template should not be empty")

	acc.LogTestSuccess(t, "GetCSVTemplate: %d bytes downloaded", len(template))
}

// TestAcceptance_InventoryPreload_validation_errors verifies input validation.
func TestAcceptance_InventoryPreload_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.InventoryPreload

	t.Run("CreateRecord_NilRecord", func(t *testing.T) {
		_, _, err := svc.CreateRecord(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "record is required")
	})

	t.Run("GetRecordByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetRecordByID(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("UpdateRecord_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateRecord(context.Background(), "", &inventory_preload.InventoryPreloadRecord{
			SerialNumber: "SN001",
			DeviceType:   "Computer",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("UpdateRecord_NilRecord", func(t *testing.T) {
		_, _, err := svc.UpdateRecord(context.Background(), "1", nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "record is required")
	})

	t.Run("DeleteRecord_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteRecord(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})
}
