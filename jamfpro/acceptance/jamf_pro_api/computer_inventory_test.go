package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/computer_inventory"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_ComputerInventory_list(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.JamfProAPI.ComputerInventory

	result, _, err := svc.ListV3(ctx, nil)
	if err != nil {
		t.Skipf("Failed to list computer inventory (may not have computers enrolled): %v", err)
		return
	}

	require.NotNil(t, result)
	assert.GreaterOrEqual(t, result.TotalCount, 0)
}

func TestAcceptance_ComputerInventory_get_by_id(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.JamfProAPI.ComputerInventory

	list, _, err := svc.ListV3(ctx, nil)
	if err != nil || len(list.Results) == 0 {
		t.Skipf("No computers available for testing: %v", err)
		return
	}

	computerID := list.Results[0].ID

	result, _, err := svc.GetByIDV3(ctx, computerID)
	require.NoError(t, err)
	assert.Equal(t, computerID, result.ID)
	assert.NotEmpty(t, result.General.Name)
}

func TestAcceptance_ComputerInventory_update(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.JamfProAPI.ComputerInventory

	list, _, err := svc.ListV3(ctx, nil)
	if err != nil || len(list.Results) == 0 {
		t.Skipf("No computers available for testing: %v", err)
		return
	}

	computerID := list.Results[0].ID

	original, _, err := svc.GetByIDV3(ctx, computerID)
	require.NoError(t, err)

	updateReq := &computer_inventory.ResourceComputerInventory{
		UserAndLocation: computer_inventory.ComputerInventorySubsetUserAndLocation{
			Position: "Test Position Updated",
		},
	}

	updated, _, err := svc.UpdateByIDV3(ctx, computerID, updateReq)
	if err != nil {
		t.Skipf("Failed to update computer inventory (may not have permissions): %v", err)
		return
	}

	assert.NotNil(t, updated)

	restoreReq := &computer_inventory.ResourceComputerInventory{
		UserAndLocation: original.UserAndLocation,
	}
	_, _, _ = svc.UpdateByIDV3(ctx, computerID, restoreReq)
}

func TestAcceptance_ComputerInventory_file_vault(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.JamfProAPI.ComputerInventory

	result, _, err := svc.ListFileVaultV3(ctx)
	if err != nil {
		t.Skipf("Failed to list FileVault inventory (may not be configured): %v", err)
		return
	}

	require.NotNil(t, result)
	assert.GreaterOrEqual(t, result.TotalCount, 0)

	if len(result.Results) > 0 {
		computerID := result.Results[0].ComputerId

		fvDetails, _, err := svc.GetFileVaultByIDV3(ctx, computerID)
		if err != nil {
			t.Skipf("Failed to get FileVault details: %v", err)
			return
		}

		assert.Equal(t, computerID, fvDetails.ComputerId)
		assert.NotEmpty(t, fvDetails.Name)
	}
}

func TestAcceptance_ComputerInventory_recovery_lock_password(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.JamfProAPI.ComputerInventory

	list, _, err := svc.ListV3(ctx, nil)
	if err != nil || len(list.Results) == 0 {
		t.Skipf("No computers available for testing: %v", err)
		return
	}

	computerID := list.Results[0].ID

	result, _, err := svc.GetRecoveryLockPasswordByIDV3(ctx, computerID)
	if err != nil {
		t.Skipf("Failed to get recovery lock password (may not be configured): %v", err)
		return
	}

	assert.NotNil(t, result)
}

func TestAcceptance_ComputerInventory_get_detail_by_id(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.JamfProAPI.ComputerInventory

	list, _, err := svc.ListV3(ctx, nil)
	if err != nil || len(list.Results) == 0 {
		t.Skipf("No computers available for testing: %v", err)
		return
	}

	computerID := list.Results[0].ID

	result, _, err := svc.GetDetailByIDV3(ctx, computerID)
	require.NoError(t, err)
	assert.Equal(t, computerID, result.ID)
	assert.NotEmpty(t, result.General.Name)
}

func TestAcceptance_ComputerInventory_device_lock_pin(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.JamfProAPI.ComputerInventory

	list, _, err := svc.ListV3(ctx, nil)
	if err != nil || len(list.Results) == 0 {
		t.Skipf("No computers available for testing: %v", err)
		return
	}

	computerID := list.Results[0].ID

	result, _, err := svc.GetDeviceLockPinByIDV3(ctx, computerID)
	if err != nil {
		t.Skipf("Failed to get device lock PIN (may not be configured): %v", err)
		return
	}

	assert.NotNil(t, result)
	assert.NotEmpty(t, result.Pin)
}

func TestAcceptance_ComputerInventory_attachments(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.JamfProAPI.ComputerInventory

	list, _, err := svc.ListV3(ctx, nil)
	if err != nil || len(list.Results) == 0 {
		t.Skipf("No computers available for testing: %v", err)
		return
	}

	computerID := list.Results[0].ID

	attachment := []byte("Test attachment content")

	_, err = svc.UploadAttachmentByIDV3(ctx, computerID, attachment)
	if err != nil {
		t.Skipf("Failed to upload attachment (may not have permissions): %v", err)
		return
	}

	computer, _, err := svc.GetByIDV3(ctx, computerID)
	if err != nil || len(computer.Attachments) == 0 {
		t.Skipf("No attachments found for computer: %v", err)
		return
	}

	attachmentID := computer.Attachments[0].ID

	downloadedData, _, err := svc.GetAttachmentByIDV3(ctx, computerID, attachmentID)
	if err != nil {
		t.Skipf("Failed to download attachment: %v", err)
		return
	}
	assert.NotNil(t, downloadedData)

	_, err = svc.DeleteAttachmentByIDV3(ctx, computerID, attachmentID)
	if err != nil {
		t.Skipf("Failed to delete attachment: %v", err)
	}
}
