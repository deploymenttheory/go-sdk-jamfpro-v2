package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/sites"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// uniqueName returns a name unique to the test run to avoid conflicts with
// existing data and between parallel test runs.
func uniqueName(prefix string) string {
	return fmt.Sprintf("%s-%d", prefix, time.Now().UnixMilli())
}

// =============================================================================
// TestAcceptance_Sites_Lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → GetByName → UpdateByID →
// UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_Sites_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Sites
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test site")

	siteName := uniqueName("acc-test-site")
	createReq := &sites.RequestSite{Name: siteName}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.CreateSite(ctx1, createReq)
	require.NoError(t, err, "CreateSite should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created site ID should be a positive integer")
	// Classic API POST responses return only the assigned ID, not the full resource.

	siteID := created.ID
	acc.LogTestSuccess(t, "Site created with ID=%d name=%q", siteID, siteName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteSiteByID(cleanupCtx, siteID)
		acc.LogCleanupDeleteError(t, "site", fmt.Sprintf("%d", siteID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new site appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing sites to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListSites(ctx2)
	require.NoError(t, err, "ListSites should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, s := range list.Results {
		if s.ID == siteID {
			found = true
			assert.Equal(t, siteName, s.Name)
			break
		}
	}
	assert.True(t, found, "newly created site should appear in list")
	acc.LogTestSuccess(t, "Site ID=%d found in list (%d total)", siteID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching site by ID=%d", siteID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetSiteByID(ctx3, siteID)
	require.NoError(t, err, "GetSiteByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, siteID, fetched.ID)
	assert.Equal(t, siteName, fetched.Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.ID, fetched.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching site by name=%q", siteName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetSiteByName(ctx4, siteName)
	require.NoError(t, err, "GetSiteByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, siteID, fetchedByName.ID)
	assert.Equal(t, siteName, fetchedByName.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.ID, fetchedByName.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := uniqueName("acc-test-site-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating site ID=%d to name=%q", siteID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &sites.RequestSite{Name: updatedName}
	updated, updateResp, err := svc.UpdateSiteByID(ctx5, siteID, updateReq)
	require.NoError(t, err, "UpdateSiteByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	// Classic API PUT responses return only the resource ID, not the full resource.
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. UpdateByName (back to original name for next step)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating site name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &sites.RequestSite{Name: siteName}
	reverted, revertResp, err := svc.UpdateSiteByName(ctx6, updatedName, revertReq)
	require.NoError(t, err, "UpdateSiteByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode, "expected 200 or 201")
	// Classic API PUT responses return only the resource ID, not the full resource.
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode)

	// ------------------------------------------------------------------
	// 7. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	verified, verifyResp, err := svc.GetSiteByID(ctx7, siteID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, siteName, verified.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.Name)

	// ------------------------------------------------------------------
	// 8. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting site ID=%d", siteID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteSiteByID(ctx8, siteID)
	require.NoError(t, err, "DeleteSiteByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Site ID=%d deleted", siteID)
}

// =============================================================================
// TestAcceptance_Sites_DeleteByName creates a site then deletes it by name.
// =============================================================================

func TestAcceptance_Sites_DeleteByName(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Sites
	ctx := context.Background()

	siteName := uniqueName("acc-test-site-dbn")
	createReq := &sites.RequestSite{Name: siteName}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.CreateSite(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	siteID := created.ID
	acc.LogTestSuccess(t, "Created site ID=%d name=%q for delete-by-name test", siteID, siteName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteSiteByID(cleanupCtx, siteID)
		acc.LogCleanupDeleteError(t, "site", fmt.Sprintf("%d", siteID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteSiteByName(ctx2, siteName)
	require.NoError(t, err, "DeleteSiteByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Site %q deleted by name", siteName)
}

// =============================================================================
// TestAcceptance_Sites_ValidationErrors tests client-side validation
// without making any network calls.
// =============================================================================

func TestAcceptance_Sites_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Sites

	t.Run("GetSiteByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetSiteByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "site ID must be a positive integer")
	})

	t.Run("GetSiteByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetSiteByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "site name is required")
	})

	t.Run("CreateSite_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateSite(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateSiteByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateSiteByID(context.Background(), 0, &sites.RequestSite{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "site ID must be a positive integer")
	})

	t.Run("UpdateSiteByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateSiteByName(context.Background(), "", &sites.RequestSite{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "site name is required")
	})

	t.Run("DeleteSiteByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteSiteByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "site ID must be a positive integer")
	})

	t.Run("DeleteSiteByName_EmptyName", func(t *testing.T) {
		_, err := svc.DeleteSiteByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "site name is required")
	})
}
