package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/disk_encryption_configurations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_DiskEncryptionConfigurations_Lifecycle exercises the full
// write/read/delete lifecycle: Create → List → GetByID → GetByName →
// UpdateByID → UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_DiskEncryptionConfigurations_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.DiskEncryptionConfigurations
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test disk encryption configuration")

	configName := uniqueName("acc-test-diskenc")
	createReq := &disk_encryption_configurations.RequestDiskEncryptionConfiguration{
		Name:                  configName,
		KeyType:               "Individual",
		FileVaultEnabledUsers: "Management Account",
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.CreateDiskEncryptionConfiguration(ctx1, createReq)
	require.NoError(t, err, "CreateDiskEncryptionConfiguration should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created disk encryption configuration ID should be a positive integer")

	configID := created.ID
	acc.LogTestSuccess(t, "Disk encryption configuration created with ID=%d name=%q", configID, configName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteDiskEncryptionConfigurationByID(cleanupCtx, configID)
		acc.LogCleanupDeleteError(t, "disk encryption configuration", fmt.Sprintf("%d", configID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new configuration appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing disk encryption configurations to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListDiskEncryptionConfigurations(ctx2)
	require.NoError(t, err, "ListDiskEncryptionConfigurations should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, c := range list.Results {
		if c.ID == configID {
			found = true
			assert.Equal(t, configName, c.Name)
			break
		}
	}
	assert.True(t, found, "newly created disk encryption configuration should appear in list")
	acc.LogTestSuccess(t, "Disk encryption configuration ID=%d found in list (%d total)", configID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching disk encryption configuration by ID=%d", configID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetDiskEncryptionConfigurationByID(ctx3, configID)
	require.NoError(t, err, "GetDiskEncryptionConfigurationByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, configID, fetched.ID)
	assert.Equal(t, configName, fetched.Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.ID, fetched.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching disk encryption configuration by name=%q", configName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetDiskEncryptionConfigurationByName(ctx4, configName)
	require.NoError(t, err, "GetDiskEncryptionConfigurationByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, configID, fetchedByName.ID)
	assert.Equal(t, configName, fetchedByName.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.ID, fetchedByName.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := uniqueName("acc-test-diskenc-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating disk encryption configuration ID=%d to name=%q", configID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &disk_encryption_configurations.RequestDiskEncryptionConfiguration{
		Name:                  updatedName,
		KeyType:               "Individual",
		FileVaultEnabledUsers: "Management Account",
	}
	updated, updateResp, err := svc.UpdateDiskEncryptionConfigurationByID(ctx5, configID, updateReq)
	require.NoError(t, err, "UpdateDiskEncryptionConfigurationByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating disk encryption configuration name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &disk_encryption_configurations.RequestDiskEncryptionConfiguration{
		Name:                  configName,
		KeyType:               "Individual",
		FileVaultEnabledUsers: "Management Account",
	}
	reverted, revertResp, err := svc.UpdateDiskEncryptionConfigurationByName(ctx6, updatedName, revertReq)
	require.NoError(t, err, "UpdateDiskEncryptionConfigurationByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode)

	// ------------------------------------------------------------------
	// 7. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	verified, verifyResp, err := svc.GetDiskEncryptionConfigurationByID(ctx7, configID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, configName, verified.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.Name)

	// ------------------------------------------------------------------
	// 8. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting disk encryption configuration ID=%d", configID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteDiskEncryptionConfigurationByID(ctx8, configID)
	require.NoError(t, err, "DeleteDiskEncryptionConfigurationByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Disk encryption configuration ID=%d deleted", configID)
}

// =============================================================================
// TestAcceptance_DiskEncryptionConfigurations_DeleteByName creates a config then deletes by name.
// =============================================================================

func TestAcceptance_DiskEncryptionConfigurations_DeleteByName(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.DiskEncryptionConfigurations
	ctx := context.Background()

	configName := uniqueName("acc-test-diskenc-dbn")
	createReq := &disk_encryption_configurations.RequestDiskEncryptionConfiguration{
		Name:    configName,
		KeyType: "Individual",
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.CreateDiskEncryptionConfiguration(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	configID := created.ID
	acc.LogTestSuccess(t, "Created disk encryption configuration ID=%d name=%q for delete-by-name test", configID, configName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteDiskEncryptionConfigurationByID(cleanupCtx, configID)
		acc.LogCleanupDeleteError(t, "disk encryption configuration", fmt.Sprintf("%d", configID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteDiskEncryptionConfigurationByName(ctx2, configName)
	require.NoError(t, err, "DeleteDiskEncryptionConfigurationByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Disk encryption configuration %q deleted by name", configName)
}

// =============================================================================
// TestAcceptance_DiskEncryptionConfigurations_ValidationErrors tests client-side validation.
// =============================================================================

func TestAcceptance_DiskEncryptionConfigurations_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.DiskEncryptionConfigurations

	t.Run("GetDiskEncryptionConfigurationByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetDiskEncryptionConfigurationByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "disk encryption configuration ID must be a positive integer")
	})

	t.Run("GetDiskEncryptionConfigurationByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetDiskEncryptionConfigurationByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "disk encryption configuration name is required")
	})

	t.Run("CreateDiskEncryptionConfiguration_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateDiskEncryptionConfiguration(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateDiskEncryptionConfigurationByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateDiskEncryptionConfigurationByID(context.Background(), 0, &disk_encryption_configurations.RequestDiskEncryptionConfiguration{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "disk encryption configuration ID must be a positive integer")
	})

	t.Run("UpdateDiskEncryptionConfigurationByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateDiskEncryptionConfigurationByName(context.Background(), "", &disk_encryption_configurations.RequestDiskEncryptionConfiguration{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "disk encryption configuration name is required")
	})

	t.Run("DeleteDiskEncryptionConfigurationByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteDiskEncryptionConfigurationByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "disk encryption configuration ID must be a positive integer")
	})

	t.Run("DeleteDiskEncryptionConfigurationByName_EmptyName", func(t *testing.T) {
		_, err := svc.DeleteDiskEncryptionConfigurationByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "disk encryption configuration name is required")
	})
}
