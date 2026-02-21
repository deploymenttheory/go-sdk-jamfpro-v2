package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/bookmarks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Bookmarks
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ListV1(ctx, rsqlQuery) - Lists bookmarks with optional RSQL filtering
//   • GetByIDV1(ctx, id) - Retrieves a bookmark by ID
//   • CreateV1(ctx, request) - Creates a new bookmark
//   • UpdateByIDV1(ctx, id, request) - Updates an existing bookmark
//   • DeleteByIDV1(ctx, id) - Deletes a bookmark by ID
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 1: Full CRUD Lifecycle [PARTIAL]
//     -- Reason: Service supports complete Create, Read, Update, Delete operations
//     -- Tests: TestAcceptance_Bookmarks_Lifecycle
//     -- Flow: Create → GetByID → Delete (MISSING Update step)
//
//   ✗ Pattern 5: RSQL Filter Testing [MANDATORY - MISSING]
//     -- Reason: ListV1 accepts rsqlQuery parameter for filtering
//     -- Tests: MISSING - Should be added
//     -- Status: MANDATORY test not implemented
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Create operations
//   ✓ Read operations (GetByID, List with pagination)
//   ✗ List with RSQL filtering [MANDATORY - MISSING]
//   ✗ Update operations (missing from lifecycle test)
//   ✓ Delete operations
//   ✗ Input validation and error handling
//
// Notes
// -----------------------------------------------------------------------------
//   • RSQL testing is MANDATORY because ListV1 supports filtering
//   • Bookmarks appear in Self Service for users
//   • Test handles 400 gracefully (may not be available in all tenants)
//   • TODO: Add Update step to lifecycle test
//   • TODO: Add RSQL filter test (MANDATORY)
//   • TODO: Add validation error tests
//
// =============================================================================

func TestAcceptance_Bookmarks_ListV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Bookmarks
	ctx := context.Background()

	result, resp, err := svc.ListV1(ctx, map[string]string{"page": "0", "page-size": "100"})
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 0)
}

func TestAcceptance_Bookmarks_ListWithRSQLFilter(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Bookmarks
	ctx := context.Background()

	name := fmt.Sprintf("acc-rsql-bookmark-%d", time.Now().UnixMilli())
	displayInBrowser := true
	bm := &bookmarks.ResourceBookmark{
		Name:             name,
		URL:              "https://example.com/rsql-test",
		SiteID:           "-1",
		IconID:           "0",
		Priority:         1,
		DisplayInBrowser: &displayInBrowser,
	}

	created, createResp, err := svc.CreateV1(ctx, bm)
	if err != nil && createResp != nil && createResp.StatusCode == 400 {
		t.Skip("Bookmarks create not available in this tenant; skipping RSQL filter test")
	}
	require.NoError(t, err)
	require.NotNil(t, created)

	bookmarkID := created.ID
	acc.LogTestSuccess(t, "Created bookmark ID=%s name=%q", bookmarkID, name)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByIDV1(cleanupCtx, bookmarkID)
		acc.LogCleanupDeleteError(t, "bookmark", bookmarkID, delErr)
	})

	rsqlQuery := map[string]string{
		"filter": fmt.Sprintf(`name=="%s"`, name),
	}

	list, listResp, err := svc.ListV1(ctx, rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)

	found := false
	for _, b := range list.Results {
		if b.ID == bookmarkID {
			found = true
			assert.Equal(t, name, b.Name)
			break
		}
	}
	assert.True(t, found, "bookmark should appear in RSQL-filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target bookmark found=%v", list.TotalCount, found)
}

func TestAcceptance_Bookmarks_Lifecycle(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Bookmarks
	ctx := context.Background()

	name := fmt.Sprintf("acc-bookmark-%d", time.Now().UnixMilli())
	displayInBrowser := true
	bm := &bookmarks.ResourceBookmark{
		Name:             name,
		URL:              "https://example.com",
		SiteID:           "-1",
		IconID:           "0",
		Priority:         1,
		DisplayInBrowser: &displayInBrowser,
	}
	created, createResp, err := svc.CreateV1(ctx, bm)
	if err != nil && createResp != nil && createResp.StatusCode == 400 {
		t.Skip("Bookmarks create not available in this tenant; skipping lifecycle")
	}
	require.NoError(t, err)
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	id := created.ID
	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, _ = svc.DeleteByIDV1(cleanupCtx, id)
	})

	fetched, _, err := svc.GetByIDV1(ctx, id)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, name, fetched.Name)

	delResp, err := svc.DeleteByIDV1(ctx, id)
	require.NoError(t, err)
	require.NotNil(t, delResp)
	assert.Equal(t, 204, delResp.StatusCode)
}
