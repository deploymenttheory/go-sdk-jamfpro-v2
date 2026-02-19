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

func uniqueDepartmentName(base string) string {
	return fmt.Sprintf("%s-%d", base, time.Now().UnixMilli())
}

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
		Name: uniqueDepartmentName("acc-test-dept"),
	}
	created, createResp, err := svc.CreateDepartmentV1(ctx, createReq)
	require.NoError(t, err, "CreateDepartmentV1 should not return an error")
	require.NotNil(t, created)
	assert.Equal(t, 201, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	departmentID := created.ID
	acc.LogTestSuccess(t, "Department created with ID=%s", departmentID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteDepartmentByIDV1(cleanupCtx, departmentID)
		acc.LogCleanupDeleteError(t, "department", departmentID, delErr)
	})

	// 2. List — verify creation
	acc.LogTestStage(t, "List", "Listing departments to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListDepartmentsV1(ctx2, map[string]string{"page": "0", "page-size": "200"})
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

	fetched, fetchResp, err := svc.GetDepartmentByIDV1(ctx, departmentID)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, departmentID, fetched.ID)
	assert.Equal(t, createReq.Name, fetched.Name)
	acc.LogTestSuccess(t, "GetByID: name=%q", fetched.Name)

	// 4. Update
	acc.LogTestStage(t, "Update", "Updating department ID=%s", departmentID)

	updateReq := &departments.RequestDepartment{
		Name: uniqueDepartmentName("acc-test-dept-updated"),
	}
	updated, updateResp, err := svc.UpdateDepartmentByIDV1(ctx, departmentID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode)
	assert.Equal(t, updateReq.Name, updated.Name)
	acc.LogTestSuccess(t, "Department updated: ID=%s", departmentID)

	// 5. Re-fetch to verify
	fetched2, _, err := svc.GetDepartmentByIDV1(ctx, departmentID)
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

	deleteResp, err := svc.DeleteDepartmentByIDV1(ctx, departmentID)
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

	name := uniqueDepartmentName("acc-rsql-dept")
	createReq := &departments.RequestDepartment{Name: name}

	created, _, err := svc.CreateDepartmentV1(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	departmentID := created.ID
	acc.LogTestSuccess(t, "Created department ID=%s name=%q", departmentID, name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteDepartmentByIDV1(cleanupCtx, departmentID)
		acc.LogCleanupDeleteError(t, "department", departmentID, delErr)
	})

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`name=="%s"`, name),
	}

	list, listResp, err := svc.ListDepartmentsV1(ctx, rsqlQuery)
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
		_, _, err := svc.GetDepartmentByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "department ID is required")
	})

	t.Run("CreateDepartmentV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateDepartmentV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateDepartmentByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateDepartmentByIDV1(context.Background(), "", &departments.RequestDepartment{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("DeleteDepartmentByIDV1_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteDepartmentByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "department ID is required")
	})
}
