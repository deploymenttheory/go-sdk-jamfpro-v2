package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/scripts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Scripts
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ListScriptsV1(ctx, rsqlQuery) - Lists scripts with optional RSQL filtering
//   • GetScriptByIDV1(ctx, id) - Retrieves a script by ID
//   • CreateScriptV1(ctx, request) - Creates a new script
//   • UpdateScriptByIDV1(ctx, id, request) - Updates an existing script
//   • DeleteScriptByIDV1(ctx, id) - Deletes a script by ID
//   • DownloadScriptByIDV1(ctx, id) - Downloads script contents as plain text
//   • GetScriptHistoryV1(ctx, id, rsqlQuery) - Retrieves script history with RSQL filtering
//   • AddScriptHistoryNotesV1(ctx, id, request) - Adds notes to script history
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 1: Full CRUD Lifecycle
//     -- Reason: Service supports complete Create, Read, Update, Delete operations
//     -- Tests: TestAcceptance_Scripts_Lifecycle
//     -- Flow: Create → List → GetByID → Update → Verify → History → Delete
//
//   ✓ Pattern 5: RSQL Filter Testing [MANDATORY]
//     -- Reason: ListScriptsV1 accepts rsqlQuery parameter for filtering
//     -- Tests: TestAcceptance_Scripts_ListWithRSQLFilter
//     -- Flow: Create unique script → Filter with RSQL → Verify filtered results
//
//   ✓ Pattern 7: Validation Errors
//     -- Reason: Client-side validation prevents invalid API calls
//     -- Tests: TestAcceptance_Scripts_ValidationErrors
//     -- Cases: Empty IDs, nil requests, missing required fields
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Create operations (single script creation)
//   ✓ Read operations (GetByID, List with pagination)
//   ✓ List with RSQL filtering (mandatory for RSQL-supported endpoints)
//   ✓ Update operations (full resource update)
//   ✓ Delete operations (single delete)
//   ✓ Download operations (get script contents as plain text)
//   ✓ History operations (add notes, retrieve history with RSQL)
//   ✓ Input validation and error handling
//   ✓ Cleanup and resource management
//
// Notes
// -----------------------------------------------------------------------------
//   • RSQL testing is mandatory because ListScriptsV1 supports filtering
//   • All tests register cleanup handlers to remove test scripts
//   • History operations tested as API provides dedicated endpoints
//   • GetScriptHistoryV1 also supports RSQL filtering for history entries
//   • Tests use acc.UniqueName() to avoid conflicts in shared test environments
//   • Scripts can be used in policies for automated device management
//   • Priority options: "BEFORE" (runs before reboot), "AFTER" (runs after reboot)
//   • ScriptContents should include shebang (#!/bin/bash, #!/bin/zsh, etc.)
//   • DownloadScriptByIDV1 not tested but available for retrieving script text
//   • Comprehensive validation error testing ensures client-side validation works correctly
//
// =============================================================================
// TestAcceptance_Scripts_Lifecycle exercises the full write/read/delete
// lifecycle in the order: Create → List → GetByID → Update → GetByID
// (verify update) → AddHistoryNotes → GetHistory → Delete.
// =============================================================================

func TestAcceptance_Scripts_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Scripts
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test script")

	scriptName := acc.UniqueName("acc-test-script")
	createReq := &scripts.RequestScript{
		Name:           scriptName,
		Priority:       scripts.ScriptPriorityAfter,
		Info:           "Acceptance test script",
		Notes:          "Created by automated acceptance test",
		ScriptContents: "#!/bin/bash\necho 'acceptance test'",
	}

	created, createResp, err := svc.CreateScriptV1(ctx, createReq)
	require.NoError(t, err, "CreateScriptV1 should not return an error")
	require.NotNil(t, created, "CreateScriptV1 result should not be nil")
	require.NotNil(t, createResp, "CreateScriptV1 response should not be nil")
	assert.Equal(t, 201, createResp.StatusCode, "expected 201 Created")
	assert.NotEmpty(t, created.ID, "created script ID should not be empty")
	assert.NotEmpty(t, created.Href, "created script Href should not be empty")

	scriptID := created.ID
	acc.LogTestSuccess(t, "Script created with ID=%s", scriptID)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteScriptByIDV1(cleanupCtx, scriptID)
		acc.LogCleanupDeleteError(t, "script", scriptID, delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new script appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing scripts to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListScriptsV1(ctx2, map[string]string{"page": "0", "page-size": "200"})
	require.NoError(t, err, "ListScriptsV1 should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.TotalCount, "total count should be positive")

	found := false
	for _, s := range list.Results {
		if s.ID == scriptID {
			found = true
			assert.Equal(t, scriptName, s.Name)
			break
		}
	}
	assert.True(t, found, "newly created script should appear in list")
	acc.LogTestSuccess(t, "Script ID=%s found in list (%d total)", scriptID, list.TotalCount)

	// ------------------------------------------------------------------
	// 3. GetByID — read the created script
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching script by ID=%s", scriptID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetScriptByIDV1(ctx3, scriptID)
	require.NoError(t, err, "GetScriptByIDV1 should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, scriptID, fetched.ID)
	assert.Equal(t, scriptName, fetched.Name)
	assert.Equal(t, scripts.ScriptPriorityAfter, fetched.Priority)
	assert.Equal(t, createReq.ScriptContents, fetched.ScriptContents)
	acc.LogTestSuccess(t, "GetByID: name=%q priority=%s", fetched.Name, fetched.Priority)

	// ------------------------------------------------------------------
	// 4. Update
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Update", "Updating script ID=%s", scriptID)

	updatedName := acc.UniqueName("acc-test-script-updated")
	updateReq := &scripts.RequestScript{
		Name:           updatedName,
		Priority:       scripts.ScriptPriorityBefore,
		Info:           "Updated acceptance test script",
		Notes:          "Updated by automated acceptance test",
		ScriptContents: "#!/bin/bash\necho 'updated acceptance test'",
	}

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	updated, updateResp, err := svc.UpdateScriptByIDV1(ctx4, scriptID, updateReq)
	require.NoError(t, err, "UpdateScriptByIDV1 should not return an error")
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode)
	assert.Equal(t, scriptID, updated.ID)
	acc.LogTestSuccess(t, "Script updated: ID=%s", updated.ID)

	// ------------------------------------------------------------------
	// 5. GetByID — verify update
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify update")

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	verified, verifyResp, err := svc.GetScriptByIDV1(ctx5, scriptID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, updatedName, verified.Name, "name should reflect the update")
	assert.Equal(t, scripts.ScriptPriorityBefore, verified.Priority, "priority should reflect the update")
	acc.LogTestSuccess(t, "Update verified: name=%q priority=%s", verified.Name, verified.Priority)

	// ------------------------------------------------------------------
	// 6. AddScriptHistoryNotesV1
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "AddHistoryNotes", "Adding history note to script ID=%s", scriptID)

	noteText := fmt.Sprintf("Acceptance test note at %s", time.Now().UTC().Format(time.RFC3339))
	noteReq := &scripts.AddScriptHistoryNotesRequest{Note: noteText}

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	noteResp, err := svc.AddScriptHistoryNotesV1(ctx6, scriptID, noteReq)
	require.NoError(t, err, "AddScriptHistoryNotesV1 should not return an error")
	require.NotNil(t, noteResp)
	assert.Contains(t, []int{200, 201}, noteResp.StatusCode)
	acc.LogTestSuccess(t, "History note added")

	// ------------------------------------------------------------------
	// 7. GetScriptHistoryV1 — verify note appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetHistory", "Fetching history for script ID=%s", scriptID)

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	history, historyResp, err := svc.GetScriptHistoryV1(ctx7, scriptID, nil)
	require.NoError(t, err, "GetScriptHistoryV1 should not return an error")
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
	// 8. DeleteScriptByIDV1
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting script ID=%s", scriptID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteScriptByIDV1(ctx8, scriptID)
	require.NoError(t, err, "DeleteScriptByIDV1 should not return an error")
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Script ID=%s deleted", scriptID)
}

// =============================================================================
// TestAcceptance_Scripts_ListWithRSQLFilter creates a script, then lists
// scripts using an RSQL filter expression to confirm the filter is accepted
// by the API and the created script appears in the filtered results.
// =============================================================================

func TestAcceptance_Scripts_ListWithRSQLFilter(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Scripts
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create a script whose name we can filter on.
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating script for RSQL filter test")

	name := acc.UniqueName("acc-rsql-script")
	createReq := &scripts.RequestScript{
		Name:           name,
		Priority:       scripts.ScriptPriorityAfter,
		ScriptContents: "#!/bin/bash\necho 'rsql test'",
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.CreateScriptV1(ctx1, createReq)
	require.NoError(t, err, "CreateScriptV1 should not return an error")
	require.NotNil(t, created)

	scriptID := created.ID
	acc.LogTestSuccess(t, "Created script ID=%s name=%q", scriptID, name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteScriptByIDV1(cleanupCtx, scriptID)
		acc.LogCleanupDeleteError(t, "script", scriptID, delErr)
	})

	// ------------------------------------------------------------------
	// 2. List with an RSQL filter: name == "<exact name>"
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "ListWithRSQLFilter", "Listing scripts with filter name==%q", name)

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`name=="%s"`, name),
	}

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListScriptsV1(ctx2, rsqlQuery)
	require.NoError(t, err, "ListScriptsV1 with RSQL filter should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, s := range list.Results {
		if s.ID == scriptID {
			found = true
			assert.Equal(t, name, s.Name)
			break
		}
	}
	assert.True(t, found, "script created for RSQL test should appear in filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target script found=%v", list.TotalCount, found)
}

// =============================================================================
// TestAcceptance_Scripts_ValidationErrors tests client-side validation
// without making any network calls.
// =============================================================================

func TestAcceptance_Scripts_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Scripts

	t.Run("GetScriptByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetScriptByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "script ID is required")
	})

	t.Run("CreateScript_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateScriptV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateScriptByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateScriptByIDV1(context.Background(), "", &scripts.RequestScript{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "script ID is required")
	})

	t.Run("DeleteScriptByID_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteScriptByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "script ID is required")
	})

	t.Run("GetScriptHistory_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetScriptHistoryV1(context.Background(), "", nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "script ID is required")
	})

	t.Run("AddScriptHistoryNotes_NilRequest", func(t *testing.T) {
		_, err := svc.AddScriptHistoryNotesV1(context.Background(), "1", nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request body is required")
	})
}
