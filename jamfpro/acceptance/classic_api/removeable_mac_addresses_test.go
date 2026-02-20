package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/removeable_mac_addresses"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_RemoveableMacAddresses_Lifecycle exercises the full write/read/delete
// lifecycle: Create → List → GetByID → GetByName → UpdateByID →
// UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_RemoveableMacAddresses_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.RemoveableMacAddresses
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test removeable MAC address")

	macAddress := uniqueName("AA:BB:CC:DD:EE")
	createReq := &removeable_mac_addresses.RequestRemoveableMacAddress{Name: macAddress}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.CreateRemoveableMacAddress(ctx1, createReq)
	require.NoError(t, err, "CreateRemoveableMacAddress should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created removeable MAC address ID should be a positive integer")

	macID := created.ID
	acc.LogTestSuccess(t, "Removeable MAC address created with ID=%d name=%q", macID, macAddress)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteRemoveableMacAddressByID(cleanupCtx, macID)
		acc.LogCleanupDeleteError(t, "removeable MAC address", fmt.Sprintf("%d", macID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new removeable MAC address appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing removeable MAC addresses to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListRemoveableMacAddresses(ctx2)
	require.NoError(t, err, "ListRemoveableMacAddresses should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, m := range list.Results {
		if m.ID == macID {
			found = true
			assert.Equal(t, macAddress, m.Name)
			break
		}
	}
	assert.True(t, found, "newly created removeable MAC address should appear in list")
	acc.LogTestSuccess(t, "Removeable MAC address ID=%d found in list (%d total)", macID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching removeable MAC address by ID=%d", macID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetRemoveableMacAddressByID(ctx3, macID)
	require.NoError(t, err, "GetRemoveableMacAddressByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, macID, fetched.ID)
	assert.Equal(t, macAddress, fetched.Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.ID, fetched.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching removeable MAC address by name=%q", macAddress)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetRemoveableMacAddressByName(ctx4, macAddress)
	require.NoError(t, err, "GetRemoveableMacAddressByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, macID, fetchedByName.ID)
	assert.Equal(t, macAddress, fetchedByName.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.ID, fetchedByName.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := uniqueName("AA:BB:CC:DD:FF")
	acc.LogTestStage(t, "UpdateByID", "Updating removeable MAC address ID=%d to name=%q", macID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &removeable_mac_addresses.RequestRemoveableMacAddress{Name: updatedName}
	updated, updateResp, err := svc.UpdateRemoveableMacAddressByID(ctx5, macID, updateReq)
	require.NoError(t, err, "UpdateRemoveableMacAddressByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating removeable MAC address name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &removeable_mac_addresses.RequestRemoveableMacAddress{Name: macAddress}
	reverted, revertResp, err := svc.UpdateRemoveableMacAddressByName(ctx6, updatedName, revertReq)
	require.NoError(t, err, "UpdateRemoveableMacAddressByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode)

	// ------------------------------------------------------------------
	// 7. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	verified, verifyResp, err := svc.GetRemoveableMacAddressByID(ctx7, macID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, macAddress, verified.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.Name)

	// ------------------------------------------------------------------
	// 8. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting removeable MAC address ID=%d", macID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteRemoveableMacAddressByID(ctx8, macID)
	require.NoError(t, err, "DeleteRemoveableMacAddressByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Removeable MAC address ID=%d deleted", macID)
}

// =============================================================================
// TestAcceptance_RemoveableMacAddresses_DeleteByName creates a removeable MAC address then deletes by name.
// =============================================================================

func TestAcceptance_RemoveableMacAddresses_DeleteByName(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.RemoveableMacAddresses
	ctx := context.Background()

	macAddress := uniqueName("AA:BB:CC:DD:EE")
	createReq := &removeable_mac_addresses.RequestRemoveableMacAddress{Name: macAddress}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.CreateRemoveableMacAddress(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	macID := created.ID
	acc.LogTestSuccess(t, "Created removeable MAC address ID=%d name=%q for delete-by-name test", macID, macAddress)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteRemoveableMacAddressByID(cleanupCtx, macID)
		acc.LogCleanupDeleteError(t, "removeable MAC address", fmt.Sprintf("%d", macID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteRemoveableMacAddressByName(ctx2, macAddress)
	require.NoError(t, err, "DeleteRemoveableMacAddressByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Removeable MAC address %q deleted by name", macAddress)
}

// =============================================================================
// TestAcceptance_RemoveableMacAddresses_ValidationErrors tests client-side validation
// without making any network calls.
// =============================================================================

func TestAcceptance_RemoveableMacAddresses_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.RemoveableMacAddresses

	t.Run("GetRemoveableMacAddressByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetRemoveableMacAddressByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "removeable MAC address ID must be a positive integer")
	})

	t.Run("GetRemoveableMacAddressByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetRemoveableMacAddressByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "removeable MAC address name is required")
	})

	t.Run("CreateRemoveableMacAddress_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateRemoveableMacAddress(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateRemoveableMacAddressByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateRemoveableMacAddressByID(context.Background(), 0, &removeable_mac_addresses.RequestRemoveableMacAddress{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "removeable MAC address ID must be a positive integer")
	})

	t.Run("UpdateRemoveableMacAddressByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateRemoveableMacAddressByName(context.Background(), "", &removeable_mac_addresses.RequestRemoveableMacAddress{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "removeable MAC address name is required")
	})

	t.Run("DeleteRemoveableMacAddressByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteRemoveableMacAddressByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "removeable MAC address ID must be a positive integer")
	})

	t.Run("DeleteRemoveableMacAddressByName_EmptyName", func(t *testing.T) {
		_, err := svc.DeleteRemoveableMacAddressByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "removeable MAC address name is required")
	})
}
