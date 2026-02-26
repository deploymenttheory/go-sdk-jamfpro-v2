package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/require"
)

// TestAcceptance_SmartMobileDeviceGroups_List verifies listing smart mobile device groups.
func TestAcceptance_SmartMobileDeviceGroups_list(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.SmartMobileDeviceGroups
	ctx := context.Background()

	list, resp, err := svc.List(ctx, map[string]string{"page": "0", "page-size": "100"})
	require.NoError(t, err)
	require.NotNil(t, list)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	require.GreaterOrEqual(t, list.TotalCount, 0)
}

// TestAcceptance_SmartMobileDeviceGroups_GetByID fetches a smart group by ID when at least one exists.
func TestAcceptance_SmartMobileDeviceGroups_get_by_id(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.SmartMobileDeviceGroups
	ctx := context.Background()

	list, _, err := svc.List(ctx, map[string]string{"page": "0", "page-size": "1"})
	require.NoError(t, err)
	require.NotNil(t, list)

	if list.TotalCount == 0 {
		t.Skip("No smart mobile device groups exist; skipping GetByID")
	}

	got, resp, err := svc.GetByID(ctx, list.Results[0].GroupID)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, list.Results[0].GroupID, got.GroupID)
}

// TestAcceptance_SmartMobileDeviceGroups_GetMembership fetches membership when at least one group exists.
func TestAcceptance_SmartMobileDeviceGroups_get_membership(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.SmartMobileDeviceGroups
	ctx := context.Background()

	list, _, err := svc.List(ctx, map[string]string{"page": "0", "page-size": "1"})
	require.NoError(t, err)
	require.NotNil(t, list)

	if list.TotalCount == 0 {
		t.Skip("No smart mobile device groups exist; skipping GetMembership")
	}

	membership, resp, err := svc.GetMembership(ctx, list.Results[0].GroupID, map[string]string{"page": "0", "page-size": "10"})
	require.NoError(t, err)
	require.NotNil(t, membership)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	require.GreaterOrEqual(t, membership.TotalCount, 0)
}
