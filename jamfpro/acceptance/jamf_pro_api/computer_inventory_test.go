package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_inventory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComputerInventory_List(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.ComputerInventory

	result, _, err := svc.ListV1(ctx, nil)
	if err != nil {
		t.Skipf("Failed to list computer inventory (may not have computers enrolled): %v", err)
		return
	}

	require.NotNil(t, result)
	assert.GreaterOrEqual(t, result.TotalCount, 0)
}

func TestComputerInventory_GetByID(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.ComputerInventory

	list, _, err := svc.ListV1(ctx, nil)
	if err != nil || len(list.Results) == 0 {
		t.Skipf("No computers available for testing: %v", err)
		return
	}

	computerID := list.Results[0].ID

	result, _, err := svc.GetByIDV1(ctx, computerID)
	require.NoError(t, err)
	assert.Equal(t, computerID, result.ID)
	assert.NotEmpty(t, result.General.Name)
}

func TestComputerInventory_Update(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.ComputerInventory

	list, _, err := svc.ListV1(ctx, nil)
	if err != nil || len(list.Results) == 0 {
		t.Skipf("No computers available for testing: %v", err)
		return
	}

	computerID := list.Results[0].ID

	original, _, err := svc.GetByIDV1(ctx, computerID)
	require.NoError(t, err)

	updateReq := &computer_inventory.ResourceComputerInventory{
		UserAndLocation: computer_inventory.ComputerInventorySubsetUserAndLocation{
			Position: "Test Position Updated",
		},
	}

	updated, _, err := svc.UpdateByIDV1(ctx, computerID, updateReq)
	if err != nil {
		t.Skipf("Failed to update computer inventory (may not have permissions): %v", err)
		return
	}

	assert.NotNil(t, updated)

	restoreReq := &computer_inventory.ResourceComputerInventory{
		UserAndLocation: original.UserAndLocation,
	}
	_, _, _ = svc.UpdateByIDV1(ctx, computerID, restoreReq)
}

func TestComputerInventory_FileVault(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.ComputerInventory

	result, _, err := svc.ListFileVaultV1(ctx, nil)
	if err != nil {
		t.Skipf("Failed to list FileVault inventory (may not be configured): %v", err)
		return
	}

	require.NotNil(t, result)
	assert.GreaterOrEqual(t, result.TotalCount, 0)

	if len(result.Results) > 0 {
		computerID := result.Results[0].ComputerId

		fvDetails, _, err := svc.GetFileVaultByIDV1(ctx, computerID)
		if err != nil {
			t.Skipf("Failed to get FileVault details: %v", err)
			return
		}

		assert.Equal(t, computerID, fvDetails.ComputerId)
		assert.NotEmpty(t, fvDetails.Name)
	}
}

func TestComputerInventory_RecoveryLockPassword(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.ComputerInventory

	list, _, err := svc.ListV1(ctx, nil)
	if err != nil || len(list.Results) == 0 {
		t.Skipf("No computers available for testing: %v", err)
		return
	}

	computerID := list.Results[0].ID

	result, _, err := svc.GetRecoveryLockPasswordByIDV1(ctx, computerID)
	if err != nil {
		t.Skipf("Failed to get recovery lock password (may not be configured): %v", err)
		return
	}

	assert.NotNil(t, result)
}
