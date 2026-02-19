package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/require"
)

// TestAcceptance_VolumePurchasingLocations_List verifies that listing volume purchasing locations
// works against a real Jamf Pro instance. Create/Update/Delete require a valid VPP service token
// and are not exercised here.
func TestAcceptance_VolumePurchasingLocations_List(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.VolumePurchasingLocations
	ctx := context.Background()

	list, resp, err := svc.ListVolumePurchasingLocationsV1(ctx, map[string]string{"page": "0", "page-size": "100"})
	require.NoError(t, err)
	require.NotNil(t, list)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	require.GreaterOrEqual(t, list.TotalCount, 0)
}

// TestAcceptance_VolumePurchasingLocations_GetByID fetches a single location by ID when at least one exists.
func TestAcceptance_VolumePurchasingLocations_GetByID(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.VolumePurchasingLocations
	ctx := context.Background()

	list, _, err := svc.ListVolumePurchasingLocationsV1(ctx, map[string]string{"page": "0", "page-size": "1"})
	require.NoError(t, err)
	require.NotNil(t, list)

	if list.TotalCount == 0 {
		t.Skip("No volume purchasing locations exist; skipping GetByID")
	}

	got, resp, err := svc.GetVolumePurchasingLocationByIDV1(ctx, list.Results[0].ID)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, list.Results[0].ID, got.ID)
}
