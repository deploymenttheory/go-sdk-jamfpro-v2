package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Sites
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ListV1(ctx) - Lists all sites
//   • GetObjectsByIDV1(ctx, id, rsqlQuery) - Retrieves objects for a site
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 2: Read-Only Operations
//     -- Reason: Sites are system-managed resources (no Create/Update/Delete)
//     -- Tests: TestAcceptance_Sites_list_and_get_objects
//     -- Flow: List sites → Get objects for first site
//
//   ✓ Pattern 5: RSQL Filter Testing [MANDATORY]
//     -- Reason: GetObjectsByIDV1 accepts rsqlQuery parameter for filtering
//     -- Tests: TestAcceptance_Sites_get_objects_with_rsql_filter
//     -- Flow: Get objects with RSQL filter → Verify filtered results
//
//   ✓ Pattern 7: Validation Errors
//     -- Reason: Client-side validation prevents invalid API calls
//     -- Tests: TestAcceptance_Sites_validation_errors
//     -- Cases: Empty IDs
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Read operations (List, GetObjectsByID)
//   ✓ List with pagination
//   ✓ GetObjectsByID with RSQL filtering (mandatory for RSQL-supported endpoints)
//   ✓ Input validation and error handling
//
// Notes
// -----------------------------------------------------------------------------
//   • Sites are system-managed resources, so no CRUD lifecycle tests
//   • RSQL testing is mandatory because GetObjectsByIDV1 supports filtering
//   • Tests require at least one site to exist in the Jamf Pro instance
//   • GetObjectsByIDV1 may return empty results if no objects assigned to site
//
// =============================================================================
// TestAcceptance_Sites_list_and_get_objects exercises read-only operations:
// List → GetObjectsByID.
// =============================================================================

func TestAcceptance_Sites_list_and_get_objects(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Sites
	ctx := context.Background()

	// 1. List all sites
	acc.LogTestStage(t, "List", "Listing all sites")

	sites, listResp, err := svc.ListV1(ctx)
	require.NoError(t, err, "ListV1 should not return an error")
	require.NotNil(t, sites)
	assert.Equal(t, 200, listResp.StatusCode)

	if len(sites) == 0 {
		t.Skip("No sites found in Jamf Pro instance")
	}

	acc.LogTestSuccess(t, "Found %d site(s)", len(sites))

	// 2. Get objects for the first site
	firstSiteID := sites[0].ID
	acc.LogTestStage(t, "GetObjectsByID", "Getting objects for site ID=%s", firstSiteID)

	rsqlQuery := map[string]string{
		"page":      "0",
		"page-size": "100",
	}

	objects, objResp, err := svc.GetObjectsByIDV1(ctx, firstSiteID, rsqlQuery)
	require.NoError(t, err, "GetObjectsByIDV1 should not return an error")
	require.NotNil(t, objects)
	assert.Equal(t, 200, objResp.StatusCode)

	acc.LogTestSuccess(t, "Site ID=%s has %d object(s)", firstSiteID, objects.TotalCount)

	// Verify structure if objects exist
	if objects.TotalCount > 0 {
		assert.NotEmpty(t, objects.Results)
		firstObj := objects.Results[0]
		assert.NotEmpty(t, firstObj.SiteID)
		assert.NotEmpty(t, firstObj.ObjectType)
		assert.NotEmpty(t, firstObj.ObjectID)
		acc.LogTestSuccess(t, "First object: siteId=%s, type=%s, objectId=%s",
			firstObj.SiteID, firstObj.ObjectType, firstObj.ObjectID)
	}
}

// =============================================================================
// TestAcceptance_Sites_get_objects_with_rsql_filter
// =============================================================================

func TestAcceptance_Sites_get_objects_with_rsql_filter(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Sites
	ctx := context.Background()

	// Get first site
	sites, _, err := svc.ListV1(ctx)
	require.NoError(t, err)
	if len(sites) == 0 {
		t.Skip("No sites found in Jamf Pro instance")
	}

	siteID := sites[0].ID
	acc.LogTestStage(t, "RSQL Filter", "Testing RSQL filter on site ID=%s", siteID)

	// Test with RSQL filter for Computer objects
	rsqlQuery := map[string]string{
		"filter":    `objectType=="Computer"`,
		"page":      "0",
		"page-size": "50",
	}

	objects, objResp, err := svc.GetObjectsByIDV1(ctx, siteID, rsqlQuery)
	require.NoError(t, err)
	require.NotNil(t, objects)
	assert.Equal(t, 200, objResp.StatusCode)

	// Verify all returned objects match the filter
	for _, obj := range objects.Results {
		assert.Equal(t, "Computer", obj.ObjectType,
			"all objects should match objectType==Computer filter")
	}

	acc.LogTestSuccess(t, "RSQL filter returned %d Computer object(s)", objects.TotalCount)
}

// =============================================================================
// TestAcceptance_Sites_validation_errors
// =============================================================================

func TestAcceptance_Sites_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Sites

	t.Run("GetObjectsByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetObjectsByIDV1(context.Background(), "", nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "id is required")
	})
}
