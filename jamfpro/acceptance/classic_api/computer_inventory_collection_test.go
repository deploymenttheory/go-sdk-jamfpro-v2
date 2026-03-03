package classic_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_ComputerInventoryCollection_get tests retrieving computer
// inventory collection settings.
// =============================================================================

func TestAcceptance_ComputerInventoryCollection_get(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicComputerInventoryCollection
	ctx := context.Background()

	acc.LogTestStage(t, "Get", "Getting computer inventory collection settings")

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	got, getResp, err := svc.Get(ctx1)
	require.NoError(t, err, "Get should not return an error")
	require.NotNil(t, got)
	require.NotNil(t, getResp)
	assert.Equal(t, 200, getResp.StatusCode)

	acc.LogTestSuccess(t, "Retrieved computer inventory collection: local_user_accounts=%v inclue_applications=%v",
		got.LocalUserAccounts, got.InclueApplications)
}

// =============================================================================
// TestAcceptance_ComputerInventoryCollection_update tests updating computer
// inventory collection settings. Uses Get to fetch current state, then Update
// to restore it to avoid changing server configuration.
// =============================================================================

func TestAcceptance_ComputerInventoryCollection_update(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicComputerInventoryCollection
	ctx := context.Background()

	acc.LogTestStage(t, "Get", "Getting current computer inventory collection settings")

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	current, _, err := svc.Get(ctx1)
	require.NoError(t, err, "Get should not return an error")
	require.NotNil(t, current)

	acc.LogTestStage(t, "Update", "Updating computer inventory collection settings")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	updateResp, err := svc.Update(ctx2, current)
	require.NoError(t, err, "Update should not return an error")
	require.NotNil(t, updateResp)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")

	acc.LogTestSuccess(t, "Updated computer inventory collection (restored original settings)")
}

// =============================================================================
// TestAcceptance_ComputerInventoryCollection_validation_errors tests client-side
// validation without making any network calls.
// =============================================================================

func TestAcceptance_ComputerInventoryCollection_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicComputerInventoryCollection

	t.Run("Update_NilSettings", func(t *testing.T) {
		_, err := svc.Update(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "settings is required")
	})
}
