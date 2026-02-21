package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/categories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Categories
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ListV1(ctx, rsqlQuery) - Lists categories with optional RSQL filtering
//   • GetByIDV1(ctx, id) - Retrieves a category by ID
//   • CreateV1(ctx, request) - Creates a new category
//   • UpdateByIDV1(ctx, id, request) - Updates an existing category
//   • DeleteByIDV1(ctx, id) - Deletes a category by ID
//   • DeleteCategoriesByIDV1(ctx, request) - Deletes multiple categories by IDs
//   • GetCategoryHistoryV1(ctx, id, rsqlQuery) - Retrieves category history
//   • AddCategoryHistoryNotesV1(ctx, id, request) - Adds notes to category history
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 1: Full CRUD Lifecycle
//     -- Reason: Service supports complete Create, Read, Update, Delete operations
//     -- Tests: TestAcceptance_Categories_Lifecycle
//     -- Flow: Create → List → GetByID → Update → Verify → History → Delete
//
//   ✓ Pattern 5: RSQL Filter Testing [MANDATORY]
//     -- Reason: ListV1 accepts rsqlQuery parameter for filtering
//     -- Tests: TestAcceptance_Categories_ListWithRSQLFilter
//     -- Flow: Create unique category → Filter with RSQL → Verify filtered results
//
//   ✓ Pattern 6: Bulk Operations
//     -- Reason: Service provides DeleteCategoriesByIDV1 for bulk deletion
//     -- Tests: TestAcceptance_Categories_BulkDelete
//     -- Flow: Create multiple → Bulk delete → Verify deletion
//
//   ✓ Pattern 7: Validation Errors
//     -- Reason: Client-side validation prevents invalid API calls
//     -- Tests: TestAcceptance_Categories_ValidationErrors
//     -- Cases: Empty IDs, nil requests, missing required fields
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Create operations (single category creation)
//   ✓ Read operations (GetByID, List with pagination)
//   ✓ List with RSQL filtering (mandatory for RSQL-supported endpoints)
//   ✓ Update operations (full resource update)
//   ✓ Delete operations (single delete)
//   ✓ Bulk delete operations (multiple categories)
//   ✓ History operations (add notes, retrieve history)
//   ✓ Input validation and error handling
//   ✓ Cleanup and resource management
//
// Notes
// -----------------------------------------------------------------------------
//   • RSQL testing is mandatory because ListV1 supports filtering
//   • All tests register cleanup handlers to remove test categories
//   • History operations tested as API provides dedicated endpoints
//   • Tests use acc.UniqueName() to avoid conflicts in shared test environments
//   • Categories have a Priority field that can be tested for update verification
//
// =============================================================================
// TestAcceptance_Categories_Lifecycle exercises the full write/read/delete
// lifecycle in the order: Create → List → GetByID → Update → GetByID
// (verify update) → AddHistoryNotes → GetHistory → Delete.
// =============================================================================

func TestAcceptance_Categories_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Categories
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test category")

	createReq := &categories.RequestCategory{
		Name:     acc.UniqueName("acc-test-category"),
		Priority: 7,
	}
	created, createResp, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err, "CreateCategoryV1 should not return an error")
	require.NotNil(t, created, "CreateCategoryV1 result should not be nil")
	require.NotNil(t, createResp, "CreateCategoryV1 response should not be nil")
	assert.Equal(t, 201, createResp.StatusCode, "expected 201 Created")
	assert.NotEmpty(t, created.ID, "created category ID should not be empty")
	assert.NotEmpty(t, created.Href, "created category Href should not be empty")

	categoryID := created.ID
	acc.LogTestSuccess(t, "Category created with ID=%s", categoryID)

	// Register cleanup so the category is removed even if the test fails.
	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByIDV1(cleanupCtx, categoryID)
		acc.LogCleanupDeleteError(t, "category", categoryID, delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new category appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing categories to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListV1(ctx2, map[string]string{"page": "0", "page-size": "200"})
	require.NoError(t, err, "ListCategoriesV1 should not return an error")
	require.NotNil(t, list, "ListCategoriesV1 result should not be nil")
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.TotalCount, "total count should be positive")

	found := false
	for _, c := range list.Results {
		if c.ID == categoryID {
			found = true
			assert.Equal(t, createReq.Name, c.Name)
			assert.Equal(t, createReq.Priority, c.Priority)
			break
		}
	}
	assert.True(t, found, "newly created category should appear in list")
	acc.LogTestSuccess(t, "Category ID=%s found in list (%d total)", categoryID, list.TotalCount)

	// ------------------------------------------------------------------
	// 3. GetByID — read the created category
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching category by ID=%s", categoryID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetByIDV1(ctx3, categoryID)
	require.NoError(t, err, "GetCategoryByIDV1 should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, categoryID, fetched.ID)
	assert.Equal(t, createReq.Name, fetched.Name)
	assert.Equal(t, createReq.Priority, fetched.Priority)
	acc.LogTestSuccess(t, "GetByID: name=%q priority=%d", fetched.Name, fetched.Priority)

	// ------------------------------------------------------------------
	// 4. Update
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Update", "Updating category ID=%s", categoryID)

	updatedName := acc.UniqueName("acc-test-category-updated")
	updateReq := &categories.RequestCategory{
		Name:     updatedName,
		Priority: 2,
	}

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	updated, updateResp, err := svc.UpdateByIDV1(ctx4, categoryID, updateReq)
	require.NoError(t, err, "UpdateCategoryByIDV1 should not return an error")
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode)
	assert.Equal(t, categoryID, updated.ID)
	acc.LogTestSuccess(t, "Category updated: ID=%s", updated.ID)

	// ------------------------------------------------------------------
	// 5. GetByID — verify update
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify update")

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	verified, verifyResp, err := svc.GetByIDV1(ctx5, categoryID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, updatedName, verified.Name, "name should reflect the update")
	assert.Equal(t, 2, verified.Priority, "priority should reflect the update")
	acc.LogTestSuccess(t, "Update verified: name=%q priority=%d", verified.Name, verified.Priority)

	// ------------------------------------------------------------------
	// 6. AddCategoryHistoryNotesV1
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "AddHistoryNotes", "Adding history note to category ID=%s", categoryID)

	noteText := fmt.Sprintf("Acceptance test note at %s", time.Now().UTC().Format(time.RFC3339))
	noteReq := &categories.AddCategoryHistoryNotesRequest{Note: noteText}

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	noteResp, err := svc.AddCategoryHistoryNotesV1(ctx6, categoryID, noteReq)
	require.NoError(t, err, "AddCategoryHistoryNotesV1 should not return an error")
	require.NotNil(t, noteResp)
	// Jamf returns 201 for POST to history
	assert.Contains(t, []int{200, 201}, noteResp.StatusCode)
	acc.LogTestSuccess(t, "History note added")

	// ------------------------------------------------------------------
	// 7. GetCategoryHistoryV1 — verify note appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetHistory", "Fetching history for category ID=%s", categoryID)

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	history, historyResp, err := svc.GetCategoryHistoryV1(ctx7, categoryID, nil)
	require.NoError(t, err, "GetCategoryHistoryV1 should not return an error")
	require.NotNil(t, history)
	assert.Equal(t, 200, historyResp.StatusCode)
	assert.Positive(t, history.TotalCount, "history should have at least one entry")

	noteFound := false
	for _, entry := range history.Results {
		if entry.Note == noteText {
			noteFound = true
			assert.NotEmpty(t, entry.Username)
			assert.NotEmpty(t, entry.Date)
			break
		}
	}
	assert.True(t, noteFound, "the added note should appear in history")
	acc.LogTestSuccess(t, "History verified: %d entries, note present=%v", history.TotalCount, noteFound)

	// ------------------------------------------------------------------
	// 8. DeleteCategoryByIDV1
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting category ID=%s", categoryID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteByIDV1(ctx8, categoryID)
	require.NoError(t, err, "DeleteCategoryByIDV1 should not return an error")
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Category ID=%s deleted", categoryID)
}

// =============================================================================
// TestAcceptance_Categories_ListWithRSQLFilter creates a category, then lists
// categories using an RSQL filter expression to confirm the filter is accepted
// by the API and the created category appears in the filtered results.
// =============================================================================

func TestAcceptance_Categories_ListWithRSQLFilter(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Categories
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create a category whose name we can filter on.
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating category for RSQL filter test")

	name := acc.UniqueName("acc-rsql-test")
	createReq := &categories.RequestCategory{Name: name, Priority: 5}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.CreateV1(ctx1, createReq)
	require.NoError(t, err, "CreateCategoryV1 should not return an error")
	require.NotNil(t, created)

	categoryID := created.ID
	acc.LogTestSuccess(t, "Created category ID=%s name=%q", categoryID, name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByIDV1(cleanupCtx, categoryID)
		acc.LogCleanupDeleteError(t, "category", categoryID, delErr)
	})

	// ------------------------------------------------------------------
	// 2. List with an RSQL filter: name == "<exact name>"
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "ListWithRSQLFilter", "Listing categories with filter name==%q", name)

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`name=="%s"`, name),
	}

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListV1(ctx2, rsqlQuery)
	require.NoError(t, err, "ListCategoriesV1 with RSQL filter should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, c := range list.Results {
		if c.ID == categoryID {
			found = true
			assert.Equal(t, name, c.Name)
			assert.Equal(t, 5, c.Priority)
			break
		}
	}
	assert.True(t, found, "category created for RSQL test should appear in filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target category found=%v", list.TotalCount, found)
}

// =============================================================================
// TestAcceptance_Categories_BulkDelete creates two categories and removes
// them together via DeleteCategoriesByIDV1.
// =============================================================================

func TestAcceptance_Categories_BulkDelete(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Categories
	ctx := context.Background()

	// Create two categories
	ids := make([]string, 0, 2)
	for i := 0; i < 2; i++ {
		req := &categories.RequestCategory{
			Name:     acc.UniqueName(fmt.Sprintf("acc-bulk-delete-%d", i)),
			Priority: 9,
		}

		ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
		created, resp, err := svc.CreateV1(ctx1, req)
		cancel1()
		require.NoError(t, err, "CreateCategoryV1 %d should succeed", i)
		require.NotNil(t, created)
		assert.Equal(t, 201, resp.StatusCode)
		ids = append(ids, created.ID)
		acc.LogTestSuccess(t, "Bulk test: created category ID=%s", created.ID)
	}

	// Cleanup safety net in case bulk delete itself fails
	acc.Cleanup(t, func() {
		cleanCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		for _, id := range ids {
			_, delErr := svc.DeleteByIDV1(cleanCtx, id)
			acc.LogCleanupDeleteError(t, "category", id, delErr)
		}
	})

	// Bulk delete
	acc.LogTestStage(t, "BulkDelete", "Deleting %d categories: %v", len(ids), ids)

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	bulkResp, err := svc.DeleteCategoriesByIDV1(ctx2, &categories.DeleteCategoriesByIDRequest{IDs: ids})
	require.NoError(t, err, "DeleteCategoriesByIDV1 should not return an error")
	require.NotNil(t, bulkResp)
	assert.Equal(t, 204, bulkResp.StatusCode)
	acc.LogTestSuccess(t, "Bulk delete of %d categories succeeded", len(ids))

	// Verify both are gone
	for _, id := range ids {
		ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
		_, _, getErr := svc.GetByIDV1(ctx3, id)
		cancel3()
		assert.Error(t, getErr, "deleted category ID=%s should return error on Get", id)
	}
}

// =============================================================================
// TestAcceptance_Categories_ValidationErrors tests client-side validation
// without making any network calls.
// =============================================================================

func TestAcceptance_Categories_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Categories

	t.Run("GetCategoryByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "category ID is required")
	})

	t.Run("CreateCategory_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateCategoryByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateByIDV1(context.Background(), "", &categories.RequestCategory{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("DeleteCategoryByID_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "category ID is required")
	})

	t.Run("DeleteCategoriesByID_EmptyIDs", func(t *testing.T) {
		_, err := svc.DeleteCategoriesByIDV1(context.Background(), &categories.DeleteCategoriesByIDRequest{IDs: []string{}})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "ids are required")
	})

	t.Run("GetCategoryHistory_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetCategoryHistoryV1(context.Background(), "", nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "category ID is required")
	})

	t.Run("AddCategoryHistoryNotes_NilRequest", func(t *testing.T) {
		_, err := svc.AddCategoryHistoryNotesV1(context.Background(), "1", nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request body is required")
	})
}
