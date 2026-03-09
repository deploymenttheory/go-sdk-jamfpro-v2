package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Ebooks
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ListV1(ctx, rsqlQuery) - Lists all ebook objects with optional filtering
//   • GetByIDV1(ctx, id) - Retrieves an ebook by ID
//   • GetScopeByIDV1(ctx, id) - Retrieves the scope for an ebook by ID
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 4: Read-Only with Existing Data
//     -- Reason: Service is read-only (no create/update/delete); tests use
//                pre-existing data from the environment
//     -- Tests: TestAcceptance_Ebooks_list_v1, TestAcceptance_Ebooks_get_by_id_v1,
//               TestAcceptance_Ebooks_get_scope_by_id_v1
//     -- Flow: List → skip if empty → GetByID on first result → GetScopeByID
//
//   ✓ Pattern 7: Validation Errors
//     -- Reason: Service validates empty IDs before making network calls
//     -- Tests: TestAcceptance_Ebooks_validation_errors
//     -- Cases: GetByIDV1("") → "ebook ID is required"
//               GetScopeByIDV1("") → "ebook ID is required"
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ List operations (ListV1 with pagination)
//   ✓ Read operations (GetByIDV1, GetScopeByIDV1 using existing data)
//   ✓ Input validation and error handling
//
// Notes
// -----------------------------------------------------------------------------
//   • Ebooks are managed via App Catalog/VPP; no create/update/delete in this API
//   • Tests skip gracefully if no ebooks exist in the environment
//   • RSQL filter test omitted — no create operation to seed a known resource
//
// =============================================================================

// =============================================================================
// TestAcceptance_Ebooks_list_v1
// =============================================================================

func TestAcceptance_Ebooks_list_v1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.Ebooks
	ctx := context.Background()

	list, resp, err := svc.ListV1(ctx, map[string]string{"page": "0", "page-size": "100"})
	require.NoError(t, err)
	require.NotNil(t, list)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.GreaterOrEqual(t, list.TotalCount, 0)

	acc.LogTestSuccess(t, "ListV1: totalCount=%d", list.TotalCount)
}

// =============================================================================
// TestAcceptance_Ebooks_get_by_id_v1
// =============================================================================

func TestAcceptance_Ebooks_get_by_id_v1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.Ebooks
	ctx := context.Background()

	list, _, err := svc.ListV1(ctx, map[string]string{"page": "0", "page-size": "1"})
	require.NoError(t, err)
	require.NotNil(t, list)

	if list.TotalCount == 0 {
		t.Skip("No ebooks exist in this environment; skipping GetByID")
	}

	firstID := list.Results[0].ID

	got, resp, err := svc.GetByIDV1(ctx, firstID)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, firstID, got.ID)
	assert.NotEmpty(t, got.Name)

	acc.LogTestSuccess(t, "GetByIDV1: ID=%s name=%q", got.ID, got.Name)
}

// =============================================================================
// TestAcceptance_Ebooks_get_scope_by_id_v1
// =============================================================================

func TestAcceptance_Ebooks_get_scope_by_id_v1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.Ebooks
	ctx := context.Background()

	list, _, err := svc.ListV1(ctx, map[string]string{"page": "0", "page-size": "1"})
	require.NoError(t, err)
	require.NotNil(t, list)

	if list.TotalCount == 0 {
		t.Skip("No ebooks exist in this environment; skipping GetScopeByID")
	}

	firstID := list.Results[0].ID

	scope, resp, err := svc.GetScopeByIDV1(ctx, firstID)
	require.NoError(t, err)
	require.NotNil(t, scope)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())

	acc.LogTestSuccess(t, "GetScopeByIDV1: ID=%s allComputers=%v allMobileDevices=%v", firstID, scope.AllComputers, scope.AllMobileDevices)
}

// =============================================================================
// TestAcceptance_Ebooks_validation_errors
// =============================================================================

func TestAcceptance_Ebooks_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.Ebooks

	t.Run("GetByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "ebook ID is required")
	})

	t.Run("GetScopeByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetScopeByIDV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "ebook ID is required")
	})
}
