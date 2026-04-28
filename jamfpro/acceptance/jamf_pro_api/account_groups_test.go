package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Account Groups (v1)
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   - ListV1(ctx, rsqlQuery) - Lists account groups with optional RSQL filtering
//   - GetByIDV1(ctx, id)     - Retrieves an account group by ID
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   - Pattern 4: Read-Only with Existing Data
//   - Pattern 5: RSQL Filter Testing [MANDATORY]
//   - Pattern 7: Validation Errors
//
// =============================================================================

// TestAcceptance_AccountGroups_read_existing_data verifies listing and getting account groups
// using read-only operations on existing tenant data.
func TestAcceptance_AccountGroups_read_existing_data(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()
	svc := acc.Client.JamfProAPI.AccountGroups

	// 1. LIST — must have at least one group
	acc.LogTestStage(t, "list", "listing account groups")
	list, resp, err := svc.ListV1(ctx, nil)
	if err != nil {
		t.Skipf("account-groups endpoint returned an error (may require elevated API client privileges): %v", err)
	}
	require.NotNil(t, list)
	assert.Equal(t, 200, resp.StatusCode())

	if len(list.Results) == 0 {
		t.Skip("no account groups found on this instance; skipping read-only test")
	}
	acc.LogTestSuccess(t, "found %d account group(s)", list.TotalCount)

	// 2. GET by ID — use first result
	first := list.Results[0]
	acc.LogTestStage(t, "get", "getting account group id=%s", first.ID)

	group, resp, err := svc.GetByIDV1(ctx, first.ID)
	require.NoError(t, err)
	require.NotNil(t, group)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, first.ID, group.ID)
	assert.NotEmpty(t, group.Name)
	acc.LogTestSuccess(t, "got account group name=%q accessLevel=%q privilegeLevel=%q",
		group.Name, group.AccessLevel, group.PrivilegeLevel)
}

// TestAcceptance_AccountGroups_list_with_rsql_filter tests RSQL filtering support.
func TestAcceptance_AccountGroups_list_with_rsql_filter(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()
	svc := acc.Client.JamfProAPI.AccountGroups

	// Get the full list first to find a group to filter on
	all, resp, err := svc.ListV1(ctx, nil)
	if err != nil {
		t.Skipf("account-groups endpoint returned an error (may require elevated API client privileges): %v", err)
	}
	assert.Equal(t, 200, resp.StatusCode())

	if len(all.Results) == 0 {
		t.Skip("no account groups found on this instance; skipping RSQL filter test")
	}

	targetName := all.Results[0].Name
	acc.LogTestStage(t, "filter", "filtering by name=%q", targetName)

	rsqlQuery := map[string]string{
		"filter": `name=="` + targetName + `"`,
	}

	filtered, resp, err := svc.ListV1(ctx, rsqlQuery)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	require.GreaterOrEqual(t, len(filtered.Results), 1, "RSQL filter should return at least one result")

	found := false
	for _, g := range filtered.Results {
		if g.Name == targetName {
			found = true
			break
		}
	}
	assert.True(t, found, "target account group should appear in RSQL-filtered results")
	acc.LogTestSuccess(t, "RSQL filter returned %d result(s); target found=%v", len(filtered.Results), found)
}

// TestAcceptance_AccountGroups_validation_errors tests parameter validation.
func TestAcceptance_AccountGroups_validation_errors(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.JamfProAPI.AccountGroups
	ctx := context.Background()

	t.Run("GetByIDV1_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByIDV1(ctx, "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "account group ID is required")
	})
}
