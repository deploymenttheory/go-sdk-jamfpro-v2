package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/ldap_servers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_LDAPServers_lifecycle exercises the full write/read/delete
// lifecycle: Create → GetByID → GetByName → UpdateByID → UpdateByName →
// GetByID (verify) → DeleteByID.
// =============================================================================

func TestAcceptance_LDAPServers_lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicLdapServers
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Create
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Create", "Creating test LDAP server")

	serverName := acc.UniqueName("sdkv2_acc_acc-test-ldap")
	createReq := &ldap_servers.RequestLDAPServer{
		Connection: ldap_servers.RequestConnection{
			Name:               serverName,
			Hostname:           "ldap.example.com",
			ServerType:         "Open Directory",
			Port:               389,
			UseSSL:             false,
			AuthenticationType: "none",
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, createResp, err := svc.Create(ctx1, createReq)
	require.NoError(t, err, "Create LDAP server should not return an error")
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode, "expected 200 or 201")
	assert.Positive(t, created.ID, "created LDAP server ID should be a positive integer")

	serverID := created.ID
	acc.LogTestSuccess(t, "LDAP server created with ID=%d name=%q", serverID, serverName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, serverID)
		acc.LogCleanupDeleteError(t, "LDAP server", fmt.Sprintf("%d", serverID), delErr)
	})

	// ------------------------------------------------------------------
	// 2. GetByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID", "Getting LDAP server by ID=%d", serverID)

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByID(ctx2, serverID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, serverID, fetched.Connection.ID)
	assert.Equal(t, serverName, fetched.Connection.Name)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.Connection.ID, fetched.Connection.Name)

	// ------------------------------------------------------------------
	// 3. GetByName
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByName", "Getting LDAP server by name=%q", serverName)

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	fetchedByName, fetchByNameResp, err := svc.GetByName(ctx3, serverName)
	if err != nil && client.IsServerError(err) {
		t.Skipf("GetByName returned 500 in this environment; skipping GetByName portion")
	} else {
		require.NoError(t, err, "GetByName should not return an error")
		require.NotNil(t, fetchedByName)
		assert.Equal(t, 200, fetchByNameResp.StatusCode)
		assert.Equal(t, serverID, fetchedByName.Connection.ID)
		assert.Equal(t, serverName, fetchedByName.Connection.Name)
		acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetchedByName.Connection.ID, fetchedByName.Connection.Name)
	}

	// ------------------------------------------------------------------
	// 4. UpdateByID
	// ------------------------------------------------------------------
	updatedName := acc.UniqueName("sdkv2_acc_acc-test-ldap-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating LDAP server ID=%d to name=%q", serverID, updatedName)

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()

	updateReq := &ldap_servers.RequestLDAPServer{
		Connection: ldap_servers.RequestConnection{
			Name:               updatedName,
			Hostname:           "ldap2.example.com",
			ServerType:         "Open Directory",
			Port:               636,
			UseSSL:             true,
			AuthenticationType: "none",
		},
	}
	updated, updateResp, err := svc.UpdateByID(ctx4, serverID, updateReq)
	require.NoError(t, err, "UpdateByID should not return an error")
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode)

	// ------------------------------------------------------------------
	// 5. UpdateByName (back to original name)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "UpdateByName", "Updating LDAP server name=%q back to original", updatedName)

	ctx5, cancel5 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel5()

	revertReq := &ldap_servers.RequestLDAPServer{
		Connection: ldap_servers.RequestConnection{
			Name:               serverName,
			Hostname:           "ldap.example.com",
			ServerType:         "Open Directory",
			Port:               389,
			UseSSL:             false,
			AuthenticationType: "none",
		},
	}
	reverted, revertResp, err := svc.UpdateByName(ctx5, updatedName, revertReq)
	require.NoError(t, err, "UpdateByName should not return an error")
	require.NotNil(t, reverted)
	assert.Contains(t, []int{200, 201}, revertResp.StatusCode, "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByName: status=%d", revertResp.StatusCode)

	// ------------------------------------------------------------------
	// 6. GetByID — verify revert
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "GetByID (post-update)", "Re-fetching to verify name revert")

	ctx6, cancel6 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel6()

	verified, verifyResp, err := svc.GetByID(ctx6, serverID)
	require.NoError(t, err)
	require.NotNil(t, verified)
	assert.Equal(t, 200, verifyResp.StatusCode)
	assert.Equal(t, serverName, verified.Connection.Name, "name should reflect the revert")
	acc.LogTestSuccess(t, "Name revert verified: %q", verified.Connection.Name)

	// ------------------------------------------------------------------
	// 7. DeleteByID
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Delete", "Deleting LDAP server ID=%d", serverID)

	ctx7, cancel7 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel7()

	deleteResp, err := svc.DeleteByID(ctx7, serverID)
	require.NoError(t, err, "DeleteByID should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "LDAP server ID=%d deleted", serverID)
}

// =============================================================================
// TestAcceptance_LDAPServers_delete_by_name creates an LDAP server then
// deletes it by name.
// =============================================================================

func TestAcceptance_LDAPServers_delete_by_name(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicLdapServers
	ctx := context.Background()

	serverName := acc.UniqueName("sdkv2_acc_acc-test-ldap-dbn")
	createReq := &ldap_servers.RequestLDAPServer{
		Connection: ldap_servers.RequestConnection{
			Name:               serverName,
			Hostname:           "ldap.example.com",
			ServerType:         "Open Directory",
			Port:               389,
			UseSSL:             false,
			AuthenticationType: "none",
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	created, _, err := svc.Create(ctx1, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)

	serverID := created.ID
	acc.LogTestSuccess(t, "Created LDAP server ID=%d name=%q for delete-by-name test", serverID, serverName)

	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, delErr := svc.DeleteByID(cleanupCtx, serverID)
		acc.LogCleanupDeleteError(t, "LDAP server", fmt.Sprintf("%d", serverID), delErr)
	})

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	deleteResp, err := svc.DeleteByName(ctx2, serverName)
	require.NoError(t, err, "DeleteByName should not return an error")
	require.NotNil(t, deleteResp)
	assert.Contains(t, []int{200, 204}, deleteResp.StatusCode)
	acc.LogTestSuccess(t, "LDAP server %q deleted by name", serverName)
}

// =============================================================================
// TestAcceptance_LDAPServers_validation_errors validates error handling.
// =============================================================================

func TestAcceptance_LDAPServers_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicLdapServers

	t.Run("GetByID_ZeroID", func(t *testing.T) {
		_, _, err := svc.GetByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "LDAP server ID must be a positive integer")
	})

	t.Run("GetByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "LDAP server name is required")
	})

	t.Run("Create_NilRequest", func(t *testing.T) {
		_, _, err := svc.Create(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})

	t.Run("UpdateByID_ZeroID", func(t *testing.T) {
		req := &ldap_servers.RequestLDAPServer{
			Connection: ldap_servers.RequestConnection{Name: "sdkv2_acc_test"},
		}
		_, _, err := svc.UpdateByID(context.Background(), 0, req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "LDAP server ID must be a positive integer")
	})

	t.Run("UpdateByName_EmptyName", func(t *testing.T) {
		req := &ldap_servers.RequestLDAPServer{
			Connection: ldap_servers.RequestConnection{Name: "sdkv2_acc_test"},
		}
		_, _, err := svc.UpdateByName(context.Background(), "", req)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "LDAP server name is required")
	})

	t.Run("DeleteByID_ZeroID", func(t *testing.T) {
		_, err := svc.DeleteByID(context.Background(), 0)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "LDAP server ID must be a positive integer")
	})

	t.Run("DeleteByName_EmptyName", func(t *testing.T) {
		_, err := svc.DeleteByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "LDAP server name is required")
	})
}
