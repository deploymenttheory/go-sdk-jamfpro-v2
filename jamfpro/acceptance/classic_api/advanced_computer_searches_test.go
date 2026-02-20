package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/advanced_computer_searches"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_AdvancedComputerSearches_Lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → GetByName → UpdateByID →
// UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_AdvancedComputerSearches_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.AdvancedComputerSearches
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test advanced computer search")

	searchName := uniqueName("acc-test-search")
	createReq := &advanced_computer_searches.RequestAdvancedComputerSearch{
		Name:   searchName,
		ViewAs: "Standard Web Page",
		Sort1:  "Computer Name",
		Criteria: advanced_computer_searches.CriteriaContainer{
			Size: 1,
			Criterion: []advanced_computer_searches.Criterion{
				{
					Name:       "Operating System",
					Priority:   0,
					AndOr:      "and",
					SearchType: "like",
					Value:      "macOS",
				},
			},
		},
		DisplayFields: []advanced_computer_searches.DisplayField{
			{Name: "Computer Name"},
			{Name: "Serial Number"},
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.CreateAdvancedComputerSearch(ctx1, createReq)
	require.NoError(t, err, "CreateAdvancedComputerSearch should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created search ID should be a positive integer")

	searchID := created.ID
	acc.LogTestSuccess(t, "Advanced computer search created with ID=%d name=%q", searchID, searchName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteAdvancedComputerSearchByID(cleanupCtx, searchID)
		acc.LogCleanupDeleteError(t, "advanced computer search", fmt.Sprintf("%d", searchID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new search appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing advanced computer searches to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListAdvancedComputerSearches(ctx2)
	require.NoError(t, err, "ListAdvancedComputerSearches should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, s := range list.Results {
		if s.ID == searchID {
			found = true
			assert.Equal(t, searchName, s.Name)
			break
		}
	}
	assert.True(t, found, "newly created search should appear in list")
	acc.LogTestSuccess(t, "Search ID=%d found in list (%d total)", searchID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching search by ID=%d", searchID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetAdvancedComputerSearchByID(ctx3, searchID)
	require.NoError(t, err, "GetAdvancedComputerSearchByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, searchID, fetched.ID)
	assert.Equal(t, searchName, fetched.Name)
	assert.Equal(t, "Standard Web Page", fetched.ViewAs)
	assert.Equal(t, 1, fetched.Criteria.Size)
	require.Len(t, fetched.Criteria.Criterion, 1)
	assert.Equal(t, "Operating System", fetched.Criteria.Criterion[0].Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.ID, fetched.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching search by name=%q", searchName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetAdvancedComputerSearchByName(ctx4, searchName)
	require.NoError(t, err, "GetAdvancedComputerSearchByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, searchID, fetchedByName.ID)
	assert.Equal(t, searchName, fetchedByName.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.ID, fetchedByName.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := uniqueName("acc-test-search-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating search ID=%d to name=%q", searchID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &advanced_computer_searches.RequestAdvancedComputerSearch{
		Name:   updatedName,
		ViewAs: "Standard Web Page",
		Sort1:  "Serial Number",
		Criteria: advanced_computer_searches.CriteriaContainer{
			Size: 1,
			Criterion: []advanced_computer_searches.Criterion{
				{
					Name:       "Computer Name",
					Priority:   0,
					SearchType: "like",
					Value:      "Mac",
				},
			},
		},
		DisplayFields: []advanced_computer_searches.DisplayField{
			{Name: "Computer Name"},
			{Name: "Operating System"},
		},
	}
	updated, updateResp, err := svc.UpdateAdvancedComputerSearchByID(ctx5, searchID, updateReq)
	require.NoError(t, err, "UpdateAdvancedComputerSearchByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating search name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &advanced_computer_searches.RequestAdvancedComputerSearch{
		Name:   searchName,
		ViewAs: "Standard Web Page",
		Criteria: advanced_computer_searches.CriteriaContainer{
			Size: 1,
			Criterion: []advanced_computer_searches.Criterion{
				{
					Name:       "Operating System",
					Priority:   0,
					SearchType: "like",
					Value:      "macOS",
				},
			},
		},
	}
	reverted, revertResp, err := svc.UpdateAdvancedComputerSearchByName(ctx6, updatedName, revertReq)
	require.NoError(t, err, "UpdateAdvancedComputerSearchByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode)

	// ------------------------------------------------------------------
	// 7. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	verified, verifyResp, err := svc.GetAdvancedComputerSearchByID(ctx7, searchID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, searchName, verified.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.Name)

	// ------------------------------------------------------------------
	// 8. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting search ID=%d", searchID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteAdvancedComputerSearchByID(ctx8, searchID)
	require.NoError(t, err, "DeleteAdvancedComputerSearchByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Search ID=%d deleted", searchID)
}

// =============================================================================
// TestAcceptance_AdvancedComputerSearches_DeleteByName creates a search then deletes by name.
// =============================================================================

func TestAcceptance_AdvancedComputerSearches_DeleteByName(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.AdvancedComputerSearches
	ctx := context.Background()

	searchName := uniqueName("acc-test-search-del")
	createReq := &advanced_computer_searches.RequestAdvancedComputerSearch{
		Name:   searchName,
		ViewAs: "Standard Web Page",
		Criteria: advanced_computer_searches.CriteriaContainer{
			Size: 1,
			Criterion: []advanced_computer_searches.Criterion{
				{
					Name:       "Computer Name",
					Priority:   0,
					AndOr:      "and",
					SearchType: "like",
					Value:      "Test",
				},
			},
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.CreateAdvancedComputerSearch(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	searchID := created.ID
	acc.LogTestSuccess(t, "Created search ID=%d name=%q for delete-by-name test", searchID, searchName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteAdvancedComputerSearchByID(cleanupCtx, searchID)
		acc.LogCleanupDeleteError(t, "advanced computer search", fmt.Sprintf("%d", searchID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteAdvancedComputerSearchByName(ctx2, searchName)
	require.NoError(t, err, "DeleteAdvancedComputerSearchByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Search %q deleted by name", searchName)
}

// =============================================================================
// TestAcceptance_AdvancedComputerSearches_ValidationErrors tests client-side validation
// without making any network calls.
// =============================================================================

func TestAcceptance_AdvancedComputerSearches_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.AdvancedComputerSearches

	t.Run("GetAdvancedComputerSearchByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetAdvancedComputerSearchByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "advanced computer search ID must be a positive integer")
	})

	t.Run("GetAdvancedComputerSearchByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetAdvancedComputerSearchByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "advanced computer search name is required")
	})

	t.Run("CreateAdvancedComputerSearch_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateAdvancedComputerSearch(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateAdvancedComputerSearchByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateAdvancedComputerSearchByID(context.Background(), 0, &advanced_computer_searches.RequestAdvancedComputerSearch{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "advanced computer search ID must be a positive integer")
	})

	t.Run("UpdateAdvancedComputerSearchByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateAdvancedComputerSearchByName(context.Background(), "", &advanced_computer_searches.RequestAdvancedComputerSearch{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "advanced computer search name is required")
	})

	t.Run("DeleteAdvancedComputerSearchByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteAdvancedComputerSearchByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "advanced computer search ID must be a positive integer")
	})

	t.Run("DeleteAdvancedComputerSearchByName_EmptyName", func(t *testing.T) {
		_, err := svc.DeleteAdvancedComputerSearchByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "advanced computer search name is required")
	})
}
