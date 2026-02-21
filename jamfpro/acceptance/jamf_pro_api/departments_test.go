package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/departments"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Departments
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ListV1(ctx, rsqlQuery) - Lists departments with optional RSQL filtering
//   • GetByIDV1(ctx, id) - Retrieves a department by ID
//   • CreateV1(ctx, request) - Creates a new department
//   • UpdateByIDV1(ctx, id, request) - Updates an existing department
//   • DeleteByIDV1(ctx, id) - Deletes a department by ID
//   • GetDepartmentHistoryV1(ctx, id, rsqlQuery) - Retrieves department history with RSQL filtering
//   • AddDepartmentHistoryNotesV1(ctx, id, request) - Adds notes to department history
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 1: Full CRUD Lifecycle
//     -- Reason: Service supports complete Create, Read, Update, Delete operations
//     -- Tests: TestAcceptance_Departments_Lifecycle
//     -- Flow: Create → List → GetByID → Update → Verify → History → Delete
//
//   ✓ Pattern 5: RSQL Filter Testing [MANDATORY]
//     -- Reason: ListV1 accepts rsqlQuery parameter for filtering
//     -- Tests: TestAcceptance_Departments_ListWithRSQLFilter
//     -- Flow: Create unique department → Filter with RSQL → Verify filtered results
//
//   ✓ Pattern 7: Validation Errors
//     -- Reason: Client-side validation prevents invalid API calls
//     -- Tests: TestAcceptance_Departments_ValidationErrors
//     -- Cases: Empty IDs, nil requests, missing required fields
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Create operations (single department creation)
//   ✓ Read operations (GetByID, List with pagination)
//   ✓ List with RSQL filtering (mandatory for RSQL-supported endpoints)
//   ✓ Update operations (full resource update)
//   ✓ Delete operations (single delete)
//   ✓ History operations (add notes, retrieve history)
//   ✓ Input validation and error handling
//   ✓ Cleanup and resource management
//
// Notes
// -----------------------------------------------------------------------------
//   • RSQL testing is mandatory because ListV1 supports filtering
//   • All tests register cleanup handlers to remove test departments
//   • History operations tested as API provides dedicated endpoints
//   • Tests use acc.UniqueName() to avoid conflicts in shared test environments
//   • Departments are organizational units used for device and user categorization
//   • GetDepartmentHistoryV1 also supports RSQL filtering for history entries
//   • Comprehensive validation error testing ensures client-side validation works correctly
//
// =============================================================================
// TestAcceptance_Departments_Lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → Update → GetByID (verify) → Delete.
// =============================================================================

func TestAcceptance_Departments_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Departments
	ctx := context.Background()

	// 1. Create
	acc.LogTestStage(t, "Create", "Creating test department")

	createReq := &departments.RequestDepartment{
		Name: acc.UniqueName("acc-test-dept"),
	}
	created, createResp, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err, "CreateDepartmentV1 should not return an error")
	require.NotNil(t, created)
	assert.Equal(t, 201, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	departmentID := created.ID
	acc.LogTestSuccess(t, "Department created with ID=%s", departmentID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByIDV1(cleanupCtx, departmentID)
		acc.LogCleanupDeleteError(t, "department", departmentID, delErr)
	})

	// 2. List — verify creation
	acc.LogTestStage(t, "List", "Listing departments to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListV1(ctx2, map[string]string{"page": "0", "page-size": "200"})
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, d := range list.Results {
		if d.ID == departmentID {
			found = true
			assert.Equal(t, createReq.Name, d.Name)
			break
		}
	}
	assert.True(t, found, "newly created department should appear in list")
	acc.LogTestSuccess(t, "Department ID=%s found in list (%d total)", departmentID, list.TotalCount)

	// 3. GetByID
	acc.LogTestStage(t, "GetByID", "Fetching department by ID=%s", departmentID)

	fetched, fetchResp, err := svc.GetByIDV1(ctx, departmentID)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, departmentID, fetched.ID)
	assert.Equal(t, createReq.Name, fetched.Name)
	acc.LogTestSuccess(t, "GetByID: name=%q", fetched.Name)

	// 4. Update
	acc.LogTestStage(t, "Update", "Updating department ID=%s", departmentID)

	updateReq := &departments.RequestDepartment{
		Name: acc.UniqueName("acc-test-dept-updated"),
	}
	updated, updateResp, err := svc.UpdateByIDV1(ctx, departmentID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode)
	assert.Equal(t, updateReq.Name, updated.Name)
	acc.LogTestSuccess(t, "Department updated: ID=%s", departmentID)

	// 5. Re-fetch to verify
	fetched2, _, err := svc.GetByIDV1(ctx, departmentID)
	require.NoError(t, err)
	assert.Equal(t, updateReq.Name, fetched2.Name)
	acc.LogTestSuccess(t, "Update verified: name=%q", fetched2.Name)

	// 5a. Add history note and fetch history
	acc.LogTestStage(t, "History", "Adding history note and fetching history for ID=%s", departmentID)

	noteReq := &departments.AddHistoryNotesRequest{
		Note: fmt.Sprintf("Acceptance test note at %s", time.Now().Format(time.RFC3339)),
	}
	noteResp, err := svc.AddDepartmentHistoryNotesV1(ctx, departmentID, noteReq)
	require.NoError(t, err)
	require.NotNil(t, noteResp)
	assert.Equal(t, 201, noteResp.StatusCode)
	acc.LogTestSuccess(t, "History note added")

	history, histResp, err := svc.GetDepartmentHistoryV1(ctx, departmentID, nil)
	require.NoError(t, err)
	require.NotNil(t, history)
	assert.Equal(t, 200, histResp.StatusCode)
	assert.GreaterOrEqual(t, history.TotalCount, 1)
	acc.LogTestSuccess(t, "History entries: %d", history.TotalCount)

	// 6. Delete
	acc.LogTestStage(t, "Delete", "Deleting department ID=%s", departmentID)

	deleteResp, err := svc.DeleteByIDV1(ctx, departmentID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Department ID=%s deleted", departmentID)
}

// =============================================================================
// TestAcceptance_Departments_ListWithRSQLFilter
// =============================================================================

func TestAcceptance_Departments_ListWithRSQLFilter(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Departments
	ctx := context.Background()

	name := acc.UniqueName("acc-rsql-dept")
	createReq := &departments.RequestDepartment{Name: name}

	created, _, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	departmentID := created.ID
	acc.LogTestSuccess(t, "Created department ID=%s name=%q", departmentID, name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByIDV1(cleanupCtx, departmentID)
		acc.LogCleanupDeleteError(t, "department", departmentID, delErr)
	})

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`name=="%s"`, name),
	}

	list, listResp, err := svc.ListV1(ctx, rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, d := range list.Results {
		if d.ID == departmentID {
			found = true
			assert.Equal(t, name, d.Name)
			break
		}
	}
	assert.True(t, found, "department should appear in RSQL-filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target department found=%v", list.TotalCount, found)
}

// =============================================================================
// TestAcceptance_Departments_ValidationErrors
// =============================================================================

func TestAcceptance_Departments_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Departments

	t.Run("GetDepartmentByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "department ID is required")
	})

	t.Run("CreateDepartmentV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateDepartmentByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateByIDV1(context.Background(), "", &departments.RequestDepartment{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("DeleteDepartmentByIDV1_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "department ID is required")
	})
}
