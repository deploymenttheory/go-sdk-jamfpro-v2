package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Device Enrollments
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ListV1(ctx, rsqlQuery) - Lists device enrollments with optional pagination/sorting
//   • GetByIDV1(ctx, id) - Retrieves a device enrollment by ID
//   • GetByNameV1(ctx, name) - Retrieves a device enrollment by name (client-side search)
//   • GetHistoryV1(ctx, id, rsqlQuery) - Retrieves device enrollment history with optional RSQL filtering
//   • GetSyncStatesV1(ctx, id) - Retrieves all sync states for a device enrollment instance
//   • GetLatestSyncStateV1(ctx, id) - Retrieves the latest sync state for a device enrollment instance
//   • GetAllSyncStatesV1(ctx) - Retrieves all sync states for all device enrollment instances
//   • GetPublicKeyV1(ctx) - Retrieves the public key for device enrollments as PEM file
//   • CreateWithTokenV1(ctx, request) - Creates a new device enrollment instance with token
//   • UpdateByIDV1(ctx, id, request) - Updates device enrollment metadata
//   • UpdateTokenByIDV1(ctx, id, request) - Updates device enrollment token
//   • DeleteByIDV1(ctx, id) - Deletes a device enrollment by ID
//   • DisownDevicesByIDV1(ctx, id, request) - Disowns devices from a device enrollment instance
//   • AddHistoryNotesV1(ctx, id, request) - Adds notes to device enrollment history
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 4: Read-Only with Existing Data
//     -- Reason: Device enrollments typically exist and cannot be easily created in tests
//     -- Tests: TestAcceptance_DeviceEnrollments_ListAndGet
//     -- Flow: List → If exists, test GetByID, GetByName, GetHistory, GetSyncStates
//
//   ✓ Pattern 3: Read-Only Information
//     -- Reason: Public key and sync states are read-only system information
//     -- Tests: TestAcceptance_DeviceEnrollments_PublicKey, TestAcceptance_DeviceEnrollments_AllSyncStates
//     -- Flow: Get public key → Verify PEM format; Get all sync states → Verify structure
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ List operations (with pagination)
//   ✓ Read operations (GetByID, GetByName)
//   ✓ History operations (GetHistory with RSQL filtering)
//   ✓ Sync state operations (GetSyncStates, GetLatestSyncState, GetAllSyncStates)
//   ✓ Public key retrieval (PEM file format)
//   ✗ Create operations (requires valid MDM token - not tested)
//   ✗ Update operations (requires existing enrollment - not tested)
//   ✗ Delete operations (requires existing enrollment - not tested)
//   ✗ Disown operations (requires existing enrollment with devices - not tested)
//
// Notes
// -----------------------------------------------------------------------------
//   • Device enrollments are typically pre-configured and managed through Apple Business Manager
//   • Tests skip gracefully if no device enrollments exist on the tenant
//   • Public key is returned as PEM file (application/x-pem-file)
//   • Sync states track connection status with Apple's servers
//   • Create/Update/Delete operations require valid MDM tokens and are not tested in acceptance suite
//
// =============================================================================

func TestAcceptance_DeviceEnrollments_ListAndGet(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.DeviceEnrollments
	ctx := context.Background()

	acc.LogTestStage(t, "List", "Listing device enrollments")
	list, listResp, err := svc.ListV1(ctx, map[string]string{
		"page":      "0",
		"page-size": "100",
		"sort":      "id:asc",
	})

	if err != nil {
		t.Skipf("Failed to list device enrollments (may not be supported on this tenant): %v", err)
		return
	}

	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.GreaterOrEqual(t, list.TotalCount, 0)
	acc.LogTestSuccess(t, "Listed %d device enrollment(s)", list.TotalCount)

	if len(list.Results) == 0 {
		t.Skip("No device enrollments available for testing")
		return
	}

	firstEnrollment := list.Results[0]
	enrollmentID := firstEnrollment.ID
	enrollmentName := firstEnrollment.Name

	acc.LogTestStage(t, "GetByID", "Fetching device enrollment by ID=%s", enrollmentID)
	enrollment, getResp, err := svc.GetByIDV1(ctx, enrollmentID)
	require.NoError(t, err)
	require.NotNil(t, enrollment)
	assert.Equal(t, 200, getResp.StatusCode)
	assert.Equal(t, enrollmentID, enrollment.ID)
	assert.Equal(t, enrollmentName, enrollment.Name)
	acc.LogTestSuccess(t, "GetByID: name=%q", enrollment.Name)

	acc.LogTestStage(t, "GetByName", "Fetching device enrollment by name=%q", enrollmentName)
	enrollmentByName, getNameResp, err := svc.GetByNameV1(ctx, enrollmentName)
	require.NoError(t, err)
	require.NotNil(t, enrollmentByName)
	assert.Equal(t, 200, getNameResp.StatusCode)
	assert.Equal(t, enrollmentID, enrollmentByName.ID)
	assert.Equal(t, enrollmentName, enrollmentByName.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%s", enrollmentByName.ID)
}

func TestAcceptance_DeviceEnrollments_History(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.DeviceEnrollments
	ctx := context.Background()

	list, _, err := svc.ListV1(ctx, nil)
	if err != nil || len(list.Results) == 0 {
		t.Skip("No device enrollments available for history testing")
		return
	}

	enrollmentID := list.Results[0].ID

	acc.LogTestStage(t, "GetHistory", "Fetching history for device enrollment ID=%s", enrollmentID)
	history, histResp, err := svc.GetHistoryV1(ctx, enrollmentID, map[string]string{
		"page":      "0",
		"page-size": "100",
		"sort":      "date:desc",
	})

	require.NoError(t, err)
	require.NotNil(t, history)
	assert.Equal(t, 200, histResp.StatusCode)
	assert.GreaterOrEqual(t, history.TotalCount, 0)
	acc.LogTestSuccess(t, "History entries: %d", history.TotalCount)

	if history.TotalCount > 0 {
		acc.LogTestStage(t, "GetHistory", "Testing RSQL filter on history")
		filteredHistory, filtResp, err := svc.GetHistoryV1(ctx, enrollmentID, map[string]string{
			"filter": "username!=nonexistent",
		})
		require.NoError(t, err)
		require.NotNil(t, filteredHistory)
		assert.Equal(t, 200, filtResp.StatusCode)
		acc.LogTestSuccess(t, "RSQL filter returned %d result(s)", filteredHistory.TotalCount)
	}
}

func TestAcceptance_DeviceEnrollments_SyncStates(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.DeviceEnrollments
	ctx := context.Background()

	list, _, err := svc.ListV1(ctx, nil)
	if err != nil || len(list.Results) == 0 {
		t.Skip("No device enrollments available for sync state testing")
		return
	}

	enrollmentID := list.Results[0].ID

	acc.LogTestStage(t, "GetSyncStates", "Fetching sync states for device enrollment ID=%s", enrollmentID)
	syncStates, syncResp, err := svc.GetSyncStatesV1(ctx, enrollmentID)
	require.NoError(t, err)
	require.NotNil(t, syncStates)
	assert.Equal(t, 200, syncResp.StatusCode)
	acc.LogTestSuccess(t, "Retrieved %d sync state(s)", len(syncStates))

	acc.LogTestStage(t, "GetLatestSyncState", "Fetching latest sync state for device enrollment ID=%s", enrollmentID)
	latestSync, latestResp, err := svc.GetLatestSyncStateV1(ctx, enrollmentID)
	require.NoError(t, err)
	require.NotNil(t, latestSync)
	assert.Equal(t, 200, latestResp.StatusCode)
	assert.NotEmpty(t, latestSync.SyncState)
	assert.Equal(t, enrollmentID, latestSync.InstanceID)
	acc.LogTestSuccess(t, "Latest sync state: %s", latestSync.SyncState)
}

func TestAcceptance_DeviceEnrollments_AllSyncStates(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.DeviceEnrollments
	ctx := context.Background()

	acc.LogTestStage(t, "GetAllSyncStates", "Fetching all sync states for all device enrollments")
	allSyncStates, allSyncResp, err := svc.GetAllSyncStatesV1(ctx)

	if err != nil {
		t.Skipf("Failed to get all sync states (may not be supported on this tenant): %v", err)
		return
	}

	require.NotNil(t, allSyncStates)
	assert.Equal(t, 200, allSyncResp.StatusCode)
	acc.LogTestSuccess(t, "Retrieved %d total sync state(s) across all instances", len(allSyncStates))
}

func TestAcceptance_DeviceEnrollments_PublicKey(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.DeviceEnrollments
	ctx := context.Background()

	acc.LogTestStage(t, "GetPublicKey", "Fetching device enrollments public key")
	publicKey, keyResp, err := svc.GetPublicKeyV1(ctx)

	if err != nil {
		t.Skipf("Failed to get public key (may not be supported on this tenant): %v", err)
		return
	}

	require.NotNil(t, publicKey)
	assert.Equal(t, 200, keyResp.StatusCode)
	assert.NotEmpty(t, publicKey)

	keyStr := string(publicKey)
	assert.Contains(t, keyStr, "BEGIN", "Public key should be in PEM format")
	acc.LogTestSuccess(t, "Public key retrieved (%d bytes)", len(publicKey))
}
