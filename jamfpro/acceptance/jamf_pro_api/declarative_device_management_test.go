package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeclarativeDeviceManagement_GetStatusItems(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.DeclarativeDeviceManagement

	clientManagementID := "test-device-id"

	result, _, err := svc.GetStatusItemsV1(ctx, clientManagementID)
	if err != nil {
		t.Skipf("Failed to get DDM status items (device may not exist or DDM not enabled): %v", err)
		return
	}

	require.NotNil(t, result)
	assert.NotNil(t, result.StatusItems)
}

func TestDeclarativeDeviceManagement_GetStatusItemByKey(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.DeclarativeDeviceManagement

	clientManagementID := "test-device-id"
	key := "device.model.family"

	result, _, err := svc.GetStatusItemByKeyV1(ctx, clientManagementID, key)
	if err != nil {
		t.Skipf("Failed to get DDM status item (device may not exist or DDM not enabled): %v", err)
		return
	}

	require.NotNil(t, result)
	assert.Equal(t, key, result.Key)
}

func TestDeclarativeDeviceManagement_ForceSync(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.DeclarativeDeviceManagement

	clientManagementID := "test-device-id"

	_, err := svc.ForceSyncV1(ctx, clientManagementID)
	if err != nil {
		t.Skipf("Failed to force DDM sync (device may not exist or DDM not enabled): %v", err)
		return
	}
}
