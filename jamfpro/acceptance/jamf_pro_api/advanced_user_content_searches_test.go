package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/advanced_user_content_searches"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"resty.dev/v3"
)

// =============================================================================
// TestAcceptance_AdvancedUserContentSearches_lifecycle
// =============================================================================

func TestAcceptance_AdvancedUserContentSearches_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.AdvancedUserContentSearches
	ctx := context.Background()

	// 1. Create
	acc.LogTestStage(t, "Create", "Creating test advanced user content search")

	searchName := acc.UniqueName("sdkv2_acc_adv-user-search")
	createReq := &advanced_user_content_searches.ResourceAdvancedUserContentSearch{
		Name: searchName,
		Criteria: []advanced_user_content_searches.CriteriaJamfProAPI{
			{Name: "Username", AndOr: "and", SearchType: "like", Value: "%"},
		},
	}

	created, createResp, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err, "CreateV1 should not return an error")
	require.NotNil(t, created)
	assert.Equal(t, 201, createResp.StatusCode())
	assert.NotEmpty(t, created.ID)

	searchID := created.ID
	acc.LogTestSuccess(t, "Advanced user content search created with ID=%s name=%q", searchID, searchName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByIDV1(cleanupCtx, searchID)
		acc.LogCleanupDeleteError(t, "advanced user content search", searchID, delErr)
	})

	// 2. GetByID (with retry for eventual consistency)
	acc.LogTestStage(t, "GetByID", "Getting advanced user content search by ID=%s", searchID)

	var fetched *advanced_user_content_searches.ResourceAdvancedUserContentSearch
	var fetchResp *resty.Response
	err = acc.RetryOnNotFound(t, 3, 500*time.Millisecond, func() error {
		var getErr error
		fetched, fetchResp, getErr = svc.GetByIDV1(ctx, searchID)
		return getErr
	})
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode())
	assert.Equal(t, searchID, fetched.ID)
	assert.Equal(t, searchName, fetched.Name)
	acc.LogTestSuccess(t, "GetByIDV1: ID=%s name=%q", fetched.ID, fetched.Name)

	// 3. Update
	acc.LogTestStage(t, "Update", "Updating advanced user content search ID=%s", searchID)

	updatedName := acc.UniqueName("sdkv2_acc_adv-user-search-updated")
	updateReq := &advanced_user_content_searches.ResourceAdvancedUserContentSearch{
		Name: updatedName,
		Criteria: []advanced_user_content_searches.CriteriaJamfProAPI{
			{Name: "Username", AndOr: "and", SearchType: "like", Value: "%"},
		},
	}
	updated, updateResp, err := svc.UpdateByIDV1(ctx, searchID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, updateResp.StatusCode())
	acc.LogTestSuccess(t, "Advanced user content search updated: ID=%s", searchID)

	// 4. Re-fetch to verify
	fetched2, _, err := svc.GetByIDV1(ctx, searchID)
	require.NoError(t, err)
	assert.Equal(t, updatedName, fetched2.Name)
	acc.LogTestSuccess(t, "Update verified: name=%q", fetched2.Name)

	// 5. Delete
	acc.LogTestStage(t, "Delete", "Deleting advanced user content search ID=%s", searchID)

	deleteResp, err := svc.DeleteByIDV1(ctx, searchID)
	require.NoError(t, err)
	require.NotNil(t, deleteResp)
	assert.Equal(t, 204, deleteResp.StatusCode())
	acc.LogTestSuccess(t, "Advanced user content search ID=%s deleted", searchID)
}

// =============================================================================
// TestAcceptance_AdvancedUserContentSearches_list_with_rsql_filter
// =============================================================================

func TestAcceptance_AdvancedUserContentSearches_list_with_rsql_filter(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.AdvancedUserContentSearches
	ctx := context.Background()

	name := acc.UniqueName("sdkv2_acc_rsql-adv-user-search")
	createReq := &advanced_user_content_searches.ResourceAdvancedUserContentSearch{
		Name: name,
		Criteria: []advanced_user_content_searches.CriteriaJamfProAPI{
			{Name: "Username", AndOr: "and", SearchType: "like", Value: "%"},
		},
	}

	created, _, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	searchID := created.ID
	acc.LogTestSuccess(t, "Created advanced user content search ID=%s name=%q", searchID, name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByIDV1(cleanupCtx, searchID)
		acc.LogCleanupDeleteError(t, "advanced user content search", searchID, delErr)
	})

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`name=="%s"`, name),
	}

	list, listResp, err := svc.ListV1(ctx, rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode())

	found := false
	for _, s := range list.Results {
		if s.ID == searchID {
			found = true
			assert.Equal(t, name, s.Name)
			break
		}
	}
	assert.True(t, found, "advanced user content search should appear in RSQL-filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target search found=%v", list.TotalCount, found)
}

// =============================================================================
// TestAcceptance_AdvancedUserContentSearches_validation_errors
// =============================================================================

func TestAcceptance_AdvancedUserContentSearches_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.AdvancedUserContentSearches

	t.Run("GetByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("CreateV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "search is required")
	})

	t.Run("UpdateByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateByIDV1(context.Background(), "", &advanced_user_content_searches.ResourceAdvancedUserContentSearch{
			Name: "x",
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})

	t.Run("UpdateByIDV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.UpdateByIDV1(context.Background(), "1", nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "search is required")
	})

	t.Run("DeleteByIDV1_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})
}
