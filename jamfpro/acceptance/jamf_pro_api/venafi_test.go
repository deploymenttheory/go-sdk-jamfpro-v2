package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/venafi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestVenafi_Lifecycle(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.Venafi
	ctx := context.Background()

	name := acc.UniqueName("Test Venafi CA")
	revocationEnabled := true

	createReq := &venafi.ResourceVenafi{
		Name:               name,
		ProxyAddress:       "localhost:9443",
		ClientID:           "jamf-pro",
		RefreshToken:       "test-refresh-token",
		RevocationEnabled: &revocationEnabled,
	}

	created, _, err := svc.Create(ctx, createReq)
	if err != nil {
		t.Skipf("Failed to create Venafi config (may not be supported on this tenant): %v", err)
		return
	}
	require.NotEmpty(t, created.ID)

	acc.Cleanup(t, func() {
		svc.DeleteByID(ctx, created.ID)
	})

	fetched, _, err := svc.GetByID(ctx, created.ID)
	require.NoError(t, err)
	assert.Equal(t, created.ID, fmt.Sprintf("%d", fetched.ID))
	assert.Equal(t, name, fetched.Name)
	assert.Equal(t, "localhost:9443", fetched.ProxyAddress)
	assert.True(t, fetched.RevocationEnabled)

	updateReq := &venafi.ResourceVenafi{
		Name: name + " Updated",
	}

	updated, _, err := svc.UpdateByID(ctx, created.ID, updateReq)
	require.NoError(t, err)
	assert.Equal(t, name+" Updated", updated.Name)

	status, _, err := svc.GetConnectionStatus(ctx, created.ID)
	if err != nil {
		t.Skipf("Connection status check failed (proxy may not be running): %v", err)
	} else {
		assert.NotEmpty(t, status.Status)
	}

	deps, _, err := svc.GetDependentProfiles(ctx, created.ID)
	require.NoError(t, err)
	assert.NotNil(t, deps)
	assert.GreaterOrEqual(t, deps.TotalCount, 0)

	history, _, err := svc.GetHistory(ctx, created.ID, nil)
	require.NoError(t, err)
	assert.NotNil(t, history)
	assert.GreaterOrEqual(t, history.TotalCount, 0)

	noteReq := &venafi.HistoryNoteRequest{
		Note: "Test history note",
	}
	_, _, err = svc.AddHistoryNote(ctx, created.ID, noteReq)
	require.NoError(t, err)

	_, err = svc.DeleteByID(ctx, created.ID)
	require.NoError(t, err)
}
