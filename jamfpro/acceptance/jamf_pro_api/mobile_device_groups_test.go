package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/require"
)

// TestAcceptance_MobileDeviceGroups_Smart_List verifies listing smart mobile device groups.
func TestAcceptance_MobileDeviceGroups_Smart_List(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.MobileDeviceGroups
	ctx := context.Background()

	list, resp, err := svc.ListSmartV1(ctx, map[string]string{"page": "0", "page-size": "100"})
	require.NoError(t, err)
	require.NotNil(t, list)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	require.GreaterOrEqual(t, list.TotalCount, 0)
}

// TestAcceptance_MobileDeviceGroups_Smart_GetByID fetches a smart group by ID when at least one exists.
func TestAcceptance_MobileDeviceGroups_Smart_GetByID(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.MobileDeviceGroups
	ctx := context.Background()

	list, _, err := svc.ListSmartV1(ctx, map[string]string{"page": "0", "page-size": "1"})
	require.NoError(t, err)
	require.NotNil(t, list)

	if list.TotalCount == 0 {
		t.Skip("No smart mobile device groups exist; skipping GetByID")
	}

	got, resp, err := svc.GetSmartByIDV1(ctx, list.Results[0].ID)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, list.Results[0].ID, got.ID)
}

// TestAcceptance_MobileDeviceGroups_Static_List verifies listing static mobile device groups.
func TestAcceptance_MobileDeviceGroups_Static_List(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.MobileDeviceGroups
	ctx := context.Background()

	list, resp, err := svc.ListStaticV1(ctx, map[string]string{"page": "0", "page-size": "100"})
	require.NoError(t, err)
	require.NotNil(t, list)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	require.GreaterOrEqual(t, list.TotalCount, 0)
}

// TestAcceptance_MobileDeviceGroups_Static_GetByID fetches a static group by ID when at least one exists.
func TestAcceptance_MobileDeviceGroups_Static_GetByID(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.MobileDeviceGroups
	ctx := context.Background()

	list, _, err := svc.ListStaticV1(ctx, map[string]string{"page": "0", "page-size": "1"})
	require.NoError(t, err)
	require.NotNil(t, list)

	if list.TotalCount == 0 {
		t.Skip("No static mobile device groups exist; skipping GetByID")
	}

	got, resp, err := svc.GetStaticByIDV1(ctx, list.Results[0].ID)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, list.Results[0].ID, got.ID)
}
