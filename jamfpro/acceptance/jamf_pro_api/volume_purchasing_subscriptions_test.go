package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/require"
)

// TestAcceptance_VolumePurchasingSubscriptions_List verifies that listing volume purchasing
// subscriptions works against a real Jamf Pro instance. Create/Update/Delete require VPL
// configuration and are not exercised here.
func TestAcceptance_VolumePurchasingSubscriptions_List(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.VolumePurchasingSubscriptions
	ctx := context.Background()

	list, resp, err := svc.ListVolumePurchasingSubscriptionsV1(ctx, map[string]string{"page": "0", "page-size": "100"})
	require.NoError(t, err)
	require.NotNil(t, list)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	require.GreaterOrEqual(t, list.TotalCount, 0)
}

// TestAcceptance_VolumePurchasingSubscriptions_GetByID fetches a single subscription by ID when at least one exists.
func TestAcceptance_VolumePurchasingSubscriptions_GetByID(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.VolumePurchasingSubscriptions
	ctx := context.Background()

	list, _, err := svc.ListVolumePurchasingSubscriptionsV1(ctx, map[string]string{"page": "0", "page-size": "1"})
	require.NoError(t, err)
	require.NotNil(t, list)

	if list.TotalCount == 0 {
		t.Skip("No volume purchasing subscriptions exist; skipping GetByID")
	}

	got, resp, err := svc.GetVolumePurchasingSubscriptionByIDV1(ctx, list.Results[0].ID)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, list.Results[0].ID, got.ID)
}
