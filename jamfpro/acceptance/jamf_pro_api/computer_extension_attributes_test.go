package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_extension_attributes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Computer Extension Attributes
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ListV1(ctx, rsqlQuery) - Lists computer extension attributes with optional RSQL filtering
//   • GetByIDV1(ctx, id) - Retrieves a computer extension attribute by ID
//   • CreateV1(ctx, request) - Creates a new computer extension attribute
//   • UpdateByIDV1(ctx, id, request) - Updates an existing computer extension attribute
//   • DeleteByIDV1(ctx, id) - Deletes a computer extension attribute by ID
//   • DeleteComputerExtensionAttributesByIDV1(ctx, request) - Deletes multiple computer extension attributes by IDs
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 1: Full CRUD Lifecycle
//     -- Reason: Service supports complete Create, Read, Update, Delete operations
//     -- Tests: TestAcceptance_ComputerExtensionAttributes_Lifecycle
//     -- Flow: Create → List → GetByID → Update → Verify → Delete
//
//   ✗ Pattern 5: RSQL Filter Testing [MANDATORY - MISSING]
//     -- Reason: ListV1 accepts rsqlQuery parameter for filtering
//     -- Tests: MISSING - Should be added as TestAcceptance_ComputerExtensionAttributes_ListWithRSQLFilter
//     -- Flow: Create unique EA → Filter with RSQL → Verify filtered results
//     -- Status: MANDATORY test not implemented
//
//   ✓ Pattern 6: Bulk Operations
//     -- Reason: Service provides DeleteComputerExtensionAttributesByIDV1 for bulk deletion
//     -- Tests: TestAcceptance_ComputerExtensionAttributes_DeleteMultiple
//     -- Flow: Create multiple → Bulk delete → Verify deletion
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Create operations (single EA creation)
//   ✓ Read operations (GetByID, List with pagination)
//   ✗ List with RSQL filtering [MANDATORY - MISSING]
//   ✓ Update operations (full resource update)
//   ✓ Delete operations (single delete)
//   ✓ Bulk delete operations (multiple EAs)
//   ✗ Input validation and error handling (not yet tested)
//   ✓ Cleanup and resource management
//
// Notes
// -----------------------------------------------------------------------------
//   • RSQL testing is MANDATORY because ListV1 supports filtering - currently missing
//   • All tests register cleanup handlers to remove test EAs
//   • Tests use acc.UniqueName() to avoid conflicts in shared test environments
//   • Computer Extension Attributes extend inventory data collection for devices
//   • DataType options: String, Integer, Date
//   • InputType options: TEXT, SCRIPT, POPUP_MENU, LDAP_ATTRIBUTE
//   • TODO: Add TestAcceptance_ComputerExtensionAttributes_ListWithRSQLFilter (MANDATORY)
//   • TODO: Add validation error tests for empty IDs, nil requests, etc.
//
// =============================================================================

func TestAcceptance_ComputerExtensionAttributes_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ComputerExtensionAttributes
	ctx := context.Background()

	acc.LogTestStage(t, "Create", "Creating test computer extension attribute")

	enabled := true
	createReq := &computer_extension_attributes.RequestComputerExtensionAttribute{
		Name:                 acc.UniqueName("acc-test-ea"),
		Description:          "Acceptance test EA",
		DataType:             "String",
		Enabled:              &enabled,
		InventoryDisplayType: "General",
		InputType:            "TEXT",
	}
	created, createResp, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Equal(t, 201, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	eaID := created.ID
	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByIDV1(cleanupCtx, eaID)
		acc.LogCleanupDeleteError(t, "computer extension attribute", eaID, delErr)
	})

	acc.LogTestStage(t, "List", "Listing computer extension attributes")
	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListV1(ctx2, map[string]string{"page": "0", "page-size": "200"})
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, r := range list.Results {
		if r.ID == eaID {
			found = true
			assert.Equal(t, createReq.Name, r.Name)
			break
		}
	}
	assert.True(t, found)

	acc.LogTestStage(t, "GetByID", "Fetching computer extension attribute by ID=%s", eaID)
	fetched, fetchResp, err := svc.GetByIDV1(ctx, eaID)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, eaID, fetched.ID)
	assert.Equal(t, createReq.Name, fetched.Name)

	acc.LogTestStage(t, "Update", "Updating computer extension attribute ID=%s", eaID)
	updatedName := acc.UniqueName("acc-test-ea-updated")
	updateReq := &computer_extension_attributes.RequestComputerExtensionAttribute{
		Name:                 updatedName,
		Description:          "Updated description",
		DataType:             "String",
		Enabled:              &enabled,
		InventoryDisplayType: "General",
		InputType:            "TEXT",
	}
	updated, updateResp, err := svc.UpdateByIDV1(ctx, eaID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 202}, updateResp.StatusCode)
	assert.Equal(t, eaID, updated.ID)

	verified, _, err := svc.GetByIDV1(ctx, eaID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, updatedName, verified.Name)

	acc.LogTestStage(t, "Delete", "Deleting computer extension attribute ID=%s", eaID)
	deleteResp, err := svc.DeleteByIDV1(ctx, eaID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
}

func TestAcceptance_ComputerExtensionAttributes_DeleteMultiple(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ComputerExtensionAttributes
	ctx := context.Background()

	enabled := true
	createReq := &computer_extension_attributes.RequestComputerExtensionAttribute{
		Name:                 acc.UniqueName("acc-delmulti-ea-1"),
		DataType:             "String",
		InventoryDisplayType: "General",
		InputType:            "TEXT",
		Enabled:              &enabled,
	}
	c1, _, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, c1)

	createReq2 := &computer_extension_attributes.RequestComputerExtensionAttribute{
		Name:                 acc.UniqueName("acc-delmulti-ea-2"),
		DataType:             "String",
		InventoryDisplayType: "General",
		InputType:            "TEXT",
		Enabled:              &enabled,
	}
	c2, _, err := svc.CreateV1(ctx, createReq2)
	require.NoError(t, err)
	require.NotNil(t, c2)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, _ = svc.DeleteByIDV1(cleanupCtx, c1.ID)
		_, _ = svc.DeleteByIDV1(cleanupCtx, c2.ID)
	})

	resp, err := svc.DeleteComputerExtensionAttributesByIDV1(ctx, &computer_extension_attributes.DeleteComputerExtensionAttributesByIDRequest{
		IDs: []string{c1.ID, c2.ID},
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

// =============================================================================
// TestAcceptance_ComputerExtensionAttributes_ListWithRSQLFilter
// =============================================================================

func TestAcceptance_ComputerExtensionAttributes_ListWithRSQLFilter(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ComputerExtensionAttributes
	ctx := context.Background()

	name := acc.UniqueName("acc-rsql-ea")
	enabled := true
	createReq := &computer_extension_attributes.RequestComputerExtensionAttribute{
		Name:                 name,
		DataType:             "String",
		InventoryDisplayType: "General",
		InputType:            "TEXT",
		Enabled:              &enabled,
	}

	created, _, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	eaID := created.ID
	acc.LogTestSuccess(t, "Created computer extension attribute ID=%s name=%q", eaID, name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByIDV1(cleanupCtx, eaID)
		acc.LogCleanupDeleteError(t, "computer extension attribute", eaID, delErr)
	})

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`name=="%s"`, name),
	}

	list, listResp, err := svc.ListV1(ctx, rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, ea := range list.Results {
		if ea.ID == eaID {
			found = true
			assert.Equal(t, name, ea.Name)
			break
		}
	}
	assert.True(t, found, "computer extension attribute should appear in RSQL-filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target EA found=%v", list.TotalCount, found)
}
