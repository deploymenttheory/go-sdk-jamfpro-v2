package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/software_update_servers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_SoftwareUpdateServers_Lifecycle exercises the full
// write/read/delete lifecycle: Create → List → GetByID → GetByName →
// UpdateByID → UpdateByName → GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_SoftwareUpdateServers_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.SoftwareUpdateServers
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test software update server")

	serverName := uniqueName("acc-test-sus")
	createReq := &software_update_servers.RequestSoftwareUpdateServer{
		Name:      serverName,
		IPAddress: "192.168.200.10",
		Port:      8088,
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.CreateSoftwareUpdateServer(ctx1, createReq)
	require.NoError(t, err, "CreateSoftwareUpdateServer should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created software update server ID should be a positive integer")

	serverID := created.ID
	acc.LogTestSuccess(t, "Software update server created with ID=%d name=%q", serverID, serverName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteSoftwareUpdateServerByID(cleanupCtx, serverID)
		acc.LogCleanupDeleteError(t, "software update server", fmt.Sprintf("%d", serverID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. List — verify the new server appears
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "List", "Listing software update servers to verify creation")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	list, listResp, err := svc.ListSoftwareUpdateServers(ctx2)
	require.NoError(t, err, "ListSoftwareUpdateServers should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	assert.Positive(t, list.Size, "size should be positive")

	found := false
	for _, s := range list.Results {
		if s.ID == serverID {
			found = true
			assert.Equal(t, serverName, s.Name)
			break
		}
	}
	assert.True(t, found, "newly created software update server should appear in list")
	acc.LogTestSuccess(t, "Software update server ID=%d found in list (%d total)", serverID, list.Size)

	// ------------------------------------------------------------------
	// 3. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Fetching software update server by ID=%d", serverID)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetched, fetchResp, err := svc.GetSoftwareUpdateServerByID(ctx3, serverID)
	require.NoError(t, err, "GetSoftwareUpdateServerByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, serverID, fetched.ID)
	assert.Equal(t, serverName, fetched.Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.ID, fetched.Name)

	// ------------------------------------------------------------------
	// 4. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Fetching software update server by name=%q", serverName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	fetchedByName, fetchByNameResp, err := svc.GetSoftwareUpdateServerByName(ctx4, serverName)
	require.NoError(t, err, "GetSoftwareUpdateServerByName should not return an error")
	require.NotNil(t, fetchedByName)
	assert.Equal(t, 200, fetchByNameResp.StatusCode)
	assert.Equal(t, serverID, fetchedByName.ID)
	assert.Equal(t, serverName, fetchedByName.Name)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.ID, fetchedByName.Name)

	// ------------------------------------------------------------------
	// 5. UpdateByID
	// ------------------------------------------------------------------
	updatedName := uniqueName("acc-test-sus-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating software update server ID=%d to name=%q", serverID, updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	updateReq := &software_update_servers.RequestSoftwareUpdateServer{
		Name:      updatedName,
		IPAddress: "192.168.200.10",
		Port:      8088,
	}
	updated, updateResp, err := svc.UpdateSoftwareUpdateServerByID(ctx5, serverID, updateReq)
	require.NoError(t, err, "UpdateSoftwareUpdateServerByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating software update server name=%q back to original", updatedName)

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	revertReq := &software_update_servers.RequestSoftwareUpdateServer{
		Name:      serverName,
		IPAddress: "192.168.200.10",
		Port:      8088,
	}
	reverted, revertResp, err := svc.UpdateSoftwareUpdateServerByName(ctx6, updatedName, revertReq)
	require.NoError(t, err, "UpdateSoftwareUpdateServerByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode)

	// ------------------------------------------------------------------
	// 7. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	verified, verifyResp, err := svc.GetSoftwareUpdateServerByID(ctx7, serverID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, serverName, verified.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.Name)

	// ------------------------------------------------------------------
	// 8. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting software update server ID=%d", serverID)

	ctx8, cancel8 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel8()

	deleteResp, err := svc.DeleteSoftwareUpdateServerByID(ctx8, serverID)
	require.NoError(t, err, "DeleteSoftwareUpdateServerByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Software update server ID=%d deleted", serverID)
}

// =============================================================================
// TestAcceptance_SoftwareUpdateServers_DeleteByName creates a server then deletes by name.
// =============================================================================

func TestAcceptance_SoftwareUpdateServers_DeleteByName(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.SoftwareUpdateServers
	ctx := context.Background()

	serverName := uniqueName("acc-test-sus-dbn")
	createReq := &software_update_servers.RequestSoftwareUpdateServer{
		Name:      serverName,
		IPAddress: "172.16.100.10",
		Port:      8088,
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.CreateSoftwareUpdateServer(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	serverID := created.ID
	acc.LogTestSuccess(t, "Created software update server ID=%d name=%q for delete-by-name test", serverID, serverName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteSoftwareUpdateServerByID(cleanupCtx, serverID)
		acc.LogCleanupDeleteError(t, "software update server", fmt.Sprintf("%d", serverID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteSoftwareUpdateServerByName(ctx2, serverName)
	require.NoError(t, err, "DeleteSoftwareUpdateServerByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "Software update server %q deleted by name", serverName)
}

// =============================================================================
// TestAcceptance_SoftwareUpdateServers_ValidationErrors tests client-side validation.
// =============================================================================

func TestAcceptance_SoftwareUpdateServers_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.SoftwareUpdateServers

	t.Run("GetSoftwareUpdateServerByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetSoftwareUpdateServerByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "software update server ID must be a positive integer")
	})

	t.Run("GetSoftwareUpdateServerByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetSoftwareUpdateServerByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "software update server name is required")
	})

	t.Run("CreateSoftwareUpdateServer_NilRequest", func(t *testing.T) {
		_, _, err := svc.CreateSoftwareUpdateServer(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateSoftwareUpdateServerByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.UpdateSoftwareUpdateServerByID(context.Background(), 0, &software_update_servers.RequestSoftwareUpdateServer{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "software update server ID must be a positive integer")
	})

	t.Run("UpdateSoftwareUpdateServerByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateSoftwareUpdateServerByName(context.Background(), "", &software_update_servers.RequestSoftwareUpdateServer{Name: "x"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "software update server name is required")
	})

	t.Run("DeleteSoftwareUpdateServerByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteSoftwareUpdateServerByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "software update server ID must be a positive integer")
	})

	t.Run("DeleteSoftwareUpdateServerByName_EmptyName", func(t *testing.T) {
		_, err := svc.DeleteSoftwareUpdateServerByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "software update server name is required")
	})
}
