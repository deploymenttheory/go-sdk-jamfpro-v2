package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_groups"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Computer Groups
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   Smart Groups:
//   • ListSmartV2(ctx, rsqlQuery) - Lists smart computer groups with optional RSQL filtering
//   • GetSmartByIDV2(ctx, id) - Retrieves a smart computer group by ID
//   • CreateSmartV2(ctx, request) - Creates a new smart computer group
//   • UpdateSmartV2(ctx, id, request) - Updates an existing smart computer group
//   • DeleteSmartV2(ctx, id) - Deletes a smart computer group by ID
//
//   Static Groups:
//   • ListStaticV2(ctx, rsqlQuery) - Lists static computer groups with optional RSQL filtering
//   • GetStaticByIDV2(ctx, id) - Retrieves a static computer group by ID
//   • CreateStaticV2(ctx, request) - Creates a new static computer group
//   • UpdateStaticByIDV2(ctx, id, request) - Updates an existing static group (PATCH)
//   • DeleteStaticByIDV2(ctx, id) - Deletes a static computer group by ID
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 1: Full CRUD Lifecycle (Smart Groups)
//     -- Reason: Service supports complete Create, Read, Update, Delete operations for smart groups
//     -- Tests: TestAcceptance_ComputerGroups_Smart_Lifecycle
//     -- Flow: Create → List → GetByID → Update → Verify → Delete
//
//   ✓ Pattern 1: Full CRUD Lifecycle (Static Groups)
//     -- Reason: Service supports complete Create, Read, Update, Delete operations for static groups
//     -- Tests: TestAcceptance_ComputerGroups_Static_Lifecycle
//     -- Flow: Create → List → GetByID → Update (PATCH) → Verify → Delete
//
//   ✗ Pattern 5: RSQL Filter Testing [MANDATORY - MISSING]
//     -- Reason: Both ListSmartV2 and ListStaticV2 accept rsqlQuery parameter for filtering
//     -- Tests: MISSING - Should be added as separate tests for both types
//     -- Flow: Create unique group → Filter with RSQL → Verify filtered results
//     -- Status: MANDATORY tests not implemented for either smart or static groups
//
//   ✓ Pattern 7: Validation Errors
//     -- Reason: Client-side validation prevents invalid API calls
//     -- Tests: TestAcceptance_ComputerGroups_ValidationErrors
//     -- Cases: Empty IDs, nil requests for both smart and static groups
//
// Test Coverage
// -----------------------------------------------------------------------------
//   Smart Groups:
//   ✓ Create operations (smart group creation with criteria)
//   ✓ Read operations (GetByID, List with pagination)
//   ✗ List with RSQL filtering [MANDATORY - MISSING]
//   ✓ Update operations (update criteria)
//   ✓ Delete operations (single delete)
//   ✓ Input validation and error handling
//
//   Static Groups:
//   ✓ Create operations (static group creation with computer IDs)
//   ✓ Read operations (GetByID, List with pagination)
//   ✗ List with RSQL filtering [MANDATORY - MISSING]
//   ✓ Update operations (PATCH membership)
//   ✓ Delete operations (single delete)
//   ✓ Input validation and error handling
//
//   General:
//   ✓ Cleanup and resource management
//
// Notes
// -----------------------------------------------------------------------------
//   • RSQL testing is MANDATORY because both List operations support filtering - currently missing
//   • Smart groups use criteria-based membership (dynamic)
//   • Static groups use explicit computer ID lists (manual)
//   • Static group update uses PATCH (partial update of membership)
//   • Smart group update uses PUT (full resource replacement)
//   • All tests register cleanup handlers to remove test groups
//   • Tests use acc.UniqueName() to avoid conflicts in shared test environments
//   • Static group creation may return 500 in some environments (test handles gracefully)
//   • TODO: Add TestAcceptance_ComputerGroups_Smart_ListWithRSQLFilter (MANDATORY)
//   • TODO: Add TestAcceptance_ComputerGroups_Static_ListWithRSQLFilter (MANDATORY)
//
// =============================================================================
// TestAcceptance_ComputerGroups_Smart_Lifecycle
// =============================================================================

func TestAcceptance_ComputerGroups_Smart_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ComputerGroups
	ctx := context.Background()

	// 1. Create
	acc.LogTestStage(t, "Create", "Creating test smart computer group")

	createReq := &computer_groups.RequestSmartGroup{
		Name: acc.UniqueName("acc-test-smart"),
		Criteria: []computer_groups.Criterion{
			{Name: "Computer Name", Priority: 0, AndOr: "and", SearchType: "like", Value: "test-%"},
		},
	}
	created, createResp, err := svc.CreateSmartV2(ctx, createReq)
	require.NoError(t, err, "CreateSmartV2 should not return an error")
	require.NotNil(t, created)
	assert.Equal(t, 201, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	groupID := created.ID
	acc.LogTestSuccess(t, "Smart group created with ID=%s", groupID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteSmartV2(cleanupCtx, groupID)
		acc.LogCleanupDeleteError(t, "smart computer group", groupID, delErr)
	})

	// 2. List — verify creation
	acc.LogTestStage(t, "List", "Listing smart groups to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListSmartV2(ctx2, map[string]string{"page": "0", "page-size": "200"})
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, g := range list.Results {
		if g.ID == groupID {
			found = true
			assert.Equal(t, createReq.Name, g.Name)
			break
		}
	}
	assert.True(t, found, "newly created smart group should appear in list")
	acc.LogTestSuccess(t, "Smart group ID=%s found in list (%d total)", groupID, list.TotalCount)

	// 3. GetByID
	acc.LogTestStage(t, "GetByID", "Fetching smart group by ID=%s", groupID)

	fetched, fetchResp, err := svc.GetSmartByIDV2(ctx, groupID)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	if fetched.ID != "" {
		assert.Equal(t, groupID, fetched.ID, "GetByID response ID when present")
	}
	assert.Equal(t, createReq.Name, fetched.Name)
	acc.LogTestSuccess(t, "GetByID: name=%q", fetched.Name)

	// 4. Update
	acc.LogTestStage(t, "Update", "Updating smart group ID=%s", groupID)

	updateReq := &computer_groups.RequestSmartGroup{
		Name: acc.UniqueName("acc-test-smart-updated"),
		Criteria: []computer_groups.Criterion{
			{Name: "Computer Name", Priority: 0, AndOr: "and", SearchType: "like", Value: "updated-%"},
		},
	}
	updated, updateResp, err := svc.UpdateSmartV2(ctx, groupID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 202}, updateResp.StatusCode, "Update returns 200 or 202")
	assert.Equal(t, updateReq.Name, updated.Name)
	acc.LogTestSuccess(t, "Smart group updated: ID=%s", groupID)

	// 5. Re-fetch to verify
	fetched2, _, err := svc.GetSmartByIDV2(ctx, groupID)
	require.NoError(t, err)
	assert.Equal(t, updateReq.Name, fetched2.Name)
	acc.LogTestSuccess(t, "Update verified: name=%q", fetched2.Name)

	// 6. Delete
	acc.LogTestStage(t, "Delete", "Deleting smart group ID=%s", groupID)

	deleteResp, err := svc.DeleteSmartV2(ctx, groupID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Smart group ID=%s deleted", groupID)
}

// =============================================================================
// TestAcceptance_ComputerGroups_Static_Lifecycle
// =============================================================================

func TestAcceptance_ComputerGroups_Static_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ComputerGroups
	ctx := context.Background()

	// 1. Create — static group with empty or minimal membership (computer IDs may not exist)
	acc.LogTestStage(t, "Create", "Creating test static computer group")

	createReq := &computer_groups.RequestStaticGroup{
		Name:        acc.UniqueName("acc-test-static"),
		ComputerIds: []string{},
	}
	created, createResp, err := svc.CreateStaticV2(ctx, createReq)
	if err != nil && createResp != nil && createResp.StatusCode == 500 {
		t.Skip("Static computer group create returned 500 in this environment; skipping lifecycle")
	}
	require.NoError(t, err, "CreateStaticV2 should not return an error")
	require.NotNil(t, created)
	assert.Equal(t, 201, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	groupID := created.ID
	acc.LogTestSuccess(t, "Static group created with ID=%s", groupID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteStaticByIDV2(cleanupCtx, groupID)
		acc.LogCleanupDeleteError(t, "static computer group", groupID, delErr)
	})

	// 2. List — verify creation
	acc.LogTestStage(t, "List", "Listing static groups to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListStaticV2(ctx2, map[string]string{"page": "0", "page-size": "200"})
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, g := range list.Results {
		if g.ID == groupID {
			found = true
			assert.Equal(t, createReq.Name, g.Name)
			break
		}
	}
	assert.True(t, found, "newly created static group should appear in list")
	acc.LogTestSuccess(t, "Static group ID=%s found in list (%d total)", groupID, list.TotalCount)

	// 3. GetByID
	acc.LogTestStage(t, "GetByID", "Fetching static group by ID=%s", groupID)

	fetched, fetchResp, err := svc.GetStaticByIDV2(ctx, groupID)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, groupID, fetched.ID)
	assert.Equal(t, createReq.Name, fetched.Name)
	assert.False(t, fetched.IsSmart)
	acc.LogTestSuccess(t, "GetByID: name=%q", fetched.Name)

	// 4. Update membership (PATCH)
	acc.LogTestStage(t, "Update", "Updating static group membership ID=%s", groupID)

	updateReq := &computer_groups.RequestStaticGroup{
		Name:        createReq.Name,
		ComputerIds: []string{},
	}
	updated, updateResp, err := svc.UpdateStaticByIDV2(ctx, groupID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode)
	acc.LogTestSuccess(t, "Static group membership updated: ID=%s", groupID)

	// 5. Re-fetch to verify
	fetched2, _, err := svc.GetStaticByIDV2(ctx, groupID)
	require.NoError(t, err)
	assert.Equal(t, groupID, fetched2.ID)
	acc.LogTestSuccess(t, "Update verified")

	// 6. Delete
	acc.LogTestStage(t, "Delete", "Deleting static group ID=%s", groupID)

	deleteResp, err := svc.DeleteStaticByIDV2(ctx, groupID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Static group ID=%s deleted", groupID)
}

// =============================================================================
// TestAcceptance_ComputerGroups_Smart_ListWithRSQLFilter
// =============================================================================

func TestAcceptance_ComputerGroups_Smart_ListWithRSQLFilter(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ComputerGroups
	ctx := context.Background()

	name := acc.UniqueName("acc-rsql-smart")
	createReq := &computer_groups.RequestSmartGroup{
		Name: name,
		Criteria: []computer_groups.Criterion{
			{Name: "Computer Name", Priority: 0, AndOr: "and", SearchType: "like", Value: "rsql-test-%"},
		},
	}

	created, _, err := svc.CreateSmartV2(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	groupID := created.ID
	acc.LogTestSuccess(t, "Created smart group ID=%s name=%q", groupID, name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteSmartV2(cleanupCtx, groupID)
		acc.LogCleanupDeleteError(t, "smart computer group", groupID, delErr)
	})

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`name=="%s"`, name),
	}

	list, listResp, err := svc.ListSmartV2(ctx, rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, g := range list.Results {
		if g.ID == groupID {
			found = true
			assert.Equal(t, name, g.Name)
			break
		}
	}
	assert.True(t, found, "smart group should appear in RSQL-filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target smart group found=%v", list.TotalCount, found)
}

// =============================================================================
// TestAcceptance_ComputerGroups_Static_ListWithRSQLFilter
// =============================================================================

func TestAcceptance_ComputerGroups_Static_ListWithRSQLFilter(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ComputerGroups
	ctx := context.Background()

	name := acc.UniqueName("acc-rsql-static")
	createReq := &computer_groups.RequestStaticGroup{
		Name:        name,
		ComputerIds: []string{},
	}

	created, createResp, err := svc.CreateStaticV2(ctx, createReq)
	if err != nil && createResp != nil && createResp.StatusCode == 500 {
		t.Skip("Static computer group create returned 500 in this environment; skipping RSQL filter test")
	}
	require.NoError(t, err)
	require.NotNil(t, created)

	groupID := created.ID
	acc.LogTestSuccess(t, "Created static group ID=%s name=%q", groupID, name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteStaticByIDV2(cleanupCtx, groupID)
		acc.LogCleanupDeleteError(t, "static computer group", groupID, delErr)
	})

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`name=="%s"`, name),
	}

	list, listResp, err := svc.ListStaticV2(ctx, rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, g := range list.Results {
		if g.ID == groupID {
			found = true
			assert.Equal(t, name, g.Name)
			break
		}
	}
	assert.True(t, found, "static group should appear in RSQL-filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target static group found=%v", list.TotalCount, found)
}

// =============================================================================
// TestAcceptance_ComputerGroups_ValidationErrors
// =============================================================================

func TestAcceptance_ComputerGroups_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ComputerGroups

	t.Run("GetSmartByIDV2_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetSmartByIDV2(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "smart group ID is required")
	})

	t.Run("CreateSmartV2_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateSmartV2(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateSmartV2_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateSmartV2(context.Background(), "", &computer_groups.RequestSmartGroup{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("DeleteSmartV2_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteSmartV2(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "smart group ID is required")
	})

	t.Run("GetStaticByIDV2_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetStaticByIDV2(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "static group ID is required")
	})

	t.Run("CreateStaticV2_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateStaticV2(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateStaticByIDV2_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateStaticByIDV2(context.Background(), "", &computer_groups.RequestStaticGroup{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("DeleteStaticByIDV2_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteStaticByIDV2(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "static group ID is required")
	})
}
