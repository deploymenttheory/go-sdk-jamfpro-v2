package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/advanced_user_searches"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// uniqueUserSearchName generates a unique name for test user searches to avoid conflicts.
func uniqueUserSearchName(prefix string) string {
	return fmt.Sprintf("%s-%d", prefix, time.Now().UnixMilli())
}

// =============================================================================
// TestAcceptance_AdvancedUserSearches_Lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → GetByName → UpdateByID →
// UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_AdvancedUserSearches_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.AdvancedUserSearches
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test advanced user search")

	searchName := uniqueUserSearchName("acc-test-user-search")
	createReq := &advanced_user_searches.RequestAdvancedUserSearch{
		Name: searchName,
		Criteria: advanced_user_searches.CriteriaContainer{
			Size: 1,
			Criterion: []advanced_user_searches.Criterion{
				{
					Name:       "Username",
					Priority:   0,
					AndOr:      "and",
					SearchType: "like",
					Value:      "test",
				},
			},
		},
		DisplayFields: []advanced_user_searches.DisplayField{
			{Name: "Username"},
			{Name: "Email Address"},
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.CreateAdvancedUserSearch(ctx1, createReq)
	require.NoError(t, err, "CreateAdvancedUserSearch should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created search ID should be a positive integer")

	searchID := created.ID
	acc.LogTestSuccess(t, "Advanced user search created with ID=%d name=%q", searchID, searchName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteAdvancedUserSearchByID(cleanupCtx, searchID)
		acc.LogCleanupDeleteError(t, "advanced user search", fmt.Sprintf("%d", searchID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new search appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing advanced user searches to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListAdvancedUserSearches(ctx2)
	require.NoError(t, err, "ListAdvancedUserSearches should not return an error")
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

	fetched, fetchResp, err := svc.GetAdvancedUserSearchByID(ctx3, searchID)
	require.NoError(t, err, "GetAdvancedUserSearchByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, searchID, fetched.ID)
	assert.Equal(t, searchName, fetched.Name)
	assert.Equal(t, 1, fetched.Criteria.Size)
	require.Len(t, fetched.Criteria.Criterion, 1)
	assert.Equal(t, "Username", fetched.Criteria.Criterion[0].Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.ID, fetched.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching search by name=%q", searchName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetAdvancedUserSearchByName(ctx4, searchName)
	require.NoError(t, err, "GetAdvancedUserSearchByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, searchID, fetchedByName.ID)
	assert.Equal(t, searchName, fetchedByName.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.ID, fetchedByName.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := uniqueUserSearchName("acc-test-user-search-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating search ID=%d to name=%q", searchID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &advanced_user_searches.RequestAdvancedUserSearch{
		Name: updatedName,
		Criteria: advanced_user_searches.CriteriaContainer{
			Size: 1,
			Criterion: []advanced_user_searches.Criterion{
				{
					Name:       "Email Address",
					Priority:   0,
					AndOr:      "and",
					SearchType: "like",
					Value:      "@example.com",
				},
			},
		},
		DisplayFields: []advanced_user_searches.DisplayField{
			{Name: "Username"},
			{Name: "Full Name"},
		},
	}
	updated, updateResp, err := svc.UpdateAdvancedUserSearchByID(ctx5, searchID, updateReq)
	require.NoError(t, err, "UpdateAdvancedUserSearchByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating search name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &advanced_user_searches.RequestAdvancedUserSearch{
		Name: searchName,
		Criteria: advanced_user_searches.CriteriaContainer{
			Size: 1,
			Criterion: []advanced_user_searches.Criterion{
				{
					Name:       "Username",
					Priority:   0,
					AndOr:      "and",
					SearchType: "like",
					Value:      "test",
				},
			},
		},
	}
	reverted, revertResp, err := svc.UpdateAdvancedUserSearchByName(ctx6, updatedName, revertReq)
	require.NoError(t, err, "UpdateAdvancedUserSearchByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode)

	// ------------------------------------------------------------------
	// 7. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	verified, verifyResp, err := svc.GetAdvancedUserSearchByID(ctx7, searchID)
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

	deleteResp, err := svc.DeleteAdvancedUserSearchByID(ctx8, searchID)
	require.NoError(t, err, "DeleteAdvancedUserSearchByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Search ID=%d deleted", searchID)
}

// =============================================================================
// TestAcceptance_AdvancedUserSearches_DeleteByName creates a search then deletes by name.
// =============================================================================

func TestAcceptance_AdvancedUserSearches_DeleteByName(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.AdvancedUserSearches
	ctx := context.Background()

	searchName := uniqueUserSearchName("acc-test-user-search-del")
	createReq := &advanced_user_searches.RequestAdvancedUserSearch{
		Name: searchName,
		Criteria: advanced_user_searches.CriteriaContainer{
			Size: 1,
			Criterion: []advanced_user_searches.Criterion{
				{
					Name:       "Username",
					Priority:   0,
					AndOr:      "and",
					SearchType: "like",
					Value:      "test",
				},
			},
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.CreateAdvancedUserSearch(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	searchID := created.ID
	acc.LogTestSuccess(t, "Created search ID=%d name=%q for delete-by-name test", searchID, searchName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteAdvancedUserSearchByID(cleanupCtx, searchID)
		acc.LogCleanupDeleteError(t, "advanced user search", fmt.Sprintf("%d", searchID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteAdvancedUserSearchByName(ctx2, searchName)
	require.NoError(t, err, "DeleteAdvancedUserSearchByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Search %q deleted by name", searchName)
}

// =============================================================================
// TestAcceptance_AdvancedUserSearches_ValidationErrors tests client-side validation
// without making any network calls.
// =============================================================================

func TestAcceptance_AdvancedUserSearches_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.AdvancedUserSearches

	t.Run("GetAdvancedUserSearchByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetAdvancedUserSearchByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "advanced user search ID must be a positive integer")
	})

	t.Run("GetAdvancedUserSearchByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetAdvancedUserSearchByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "advanced user search name is required")
	})

	t.Run("CreateAdvancedUserSearch_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateAdvancedUserSearch(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateAdvancedUserSearchByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateAdvancedUserSearchByID(context.Background(), 0, &advanced_user_searches.RequestAdvancedUserSearch{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "advanced user search ID must be a positive integer")
	})

	t.Run("UpdateAdvancedUserSearchByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateAdvancedUserSearchByName(context.Background(), "", &advanced_user_searches.RequestAdvancedUserSearch{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "advanced user search name is required")
	})

	t.Run("DeleteAdvancedUserSearchByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteAdvancedUserSearchByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "advanced user search ID must be a positive integer")
	})

	t.Run("DeleteAdvancedUserSearchByName_EmptyName", func(t *testing.T) {
		_, err := svc.DeleteAdvancedUserSearchByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "advanced user search name is required")
	})
}
