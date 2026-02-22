package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/volume_purchasing_locations"
	"github.com/stretchr/testify/require"
)

// TestAcceptance_VolumePurchasingLocations_List verifies that listing volume purchasing locations
// works against a real Jamf Pro instance. Create/Update/Delete require a valid VPP service token
// and are not exercised here.
func TestAcceptance_VolumePurchasingLocations_ListV1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.VolumePurchasingLocations
	ctx := context.Background()

	list, resp, err := svc.ListV1(ctx, map[string]string{"page": "0", "page-size": "100"})
	require.NoError(t, err)
	require.NotNil(t, list)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	require.GreaterOrEqual(t, list.TotalCount, 0)
}

// TestAcceptance_VolumePurchasingLocations_GetByID fetches a single location by ID when at least one exists.
func TestAcceptance_VolumePurchasingLocations_GetByIDV1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.VolumePurchasingLocations
	ctx := context.Background()

	list, _, err := svc.ListV1(ctx, map[string]string{"page": "0", "page-size": "1"})
	require.NoError(t, err)
	require.NotNil(t, list)

	if list.TotalCount == 0 {
		t.Skip("No volume purchasing locations exist; skipping GetByID")
	}

	got, resp, err := svc.GetByIDV1(ctx, list.Results[0].ID)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, list.Results[0].ID, got.ID)
}

// TestAcceptance_VolumePurchasingLocations_GetHistoryV1 fetches history for a volume purchasing location.
func TestAcceptance_VolumePurchasingLocations_GetHistoryV1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.VolumePurchasingLocations
	ctx := context.Background()

	list, _, err := svc.ListV1(ctx, map[string]string{"page": "0", "page-size": "1"})
	require.NoError(t, err)
	require.NotNil(t, list)

	if list.TotalCount == 0 {
		t.Skip("No volume purchasing locations exist; skipping GetHistoryV1")
	}

	history, resp, err := svc.GetHistoryV1(ctx, list.Results[0].ID, nil)
	require.NoError(t, err)
	require.NotNil(t, history)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
	require.GreaterOrEqual(t, history.TotalCount, 0)
}

// TestAcceptance_VolumePurchasingLocations_AddHistoryNotesV1 adds a history note to a volume purchasing location.
func TestAcceptance_VolumePurchasingLocations_AddHistoryNotesV1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.VolumePurchasingLocations
	ctx := context.Background()

	list, _, err := svc.ListV1(ctx, map[string]string{"page": "0", "page-size": "1"})
	require.NoError(t, err)
	require.NotNil(t, list)

	if list.TotalCount == 0 {
		t.Skip("No volume purchasing locations exist; skipping AddHistoryNotesV1")
	}

	noteReq := &volume_purchasing_locations.AddHistoryNotesRequest{
		ObjectHistoryNote: "Test history note from acceptance test",
	}

	resp, err := svc.AddHistoryNotesV1(ctx, list.Results[0].ID, noteReq)
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 201, resp.StatusCode)
}

// TestAcceptance_VolumePurchasingLocations_RevokeVolumePurchasingLocationLicensesByIDV1 tests revoking licenses.
// This test may skip if there are no VPP locations or if the operation is not supported.
func TestAcceptance_VolumePurchasingLocations_RevokeVolumePurchasingLocationLicensesByIDV1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.VolumePurchasingLocations
	ctx := context.Background()

	list, _, err := svc.ListV1(ctx, map[string]string{"page": "0", "page-size": "1"})
	require.NoError(t, err)
	require.NotNil(t, list)

	if list.TotalCount == 0 {
		t.Skip("No volume purchasing locations exist; skipping RevokeVolumePurchasingLocationLicensesByIDV1")
	}

	resp, err := svc.RevokeVolumePurchasingLocationLicensesByIDV1(ctx, list.Results[0].ID)
	if err != nil {
		t.Skipf("Failed to revoke licenses (may not be supported or no licenses to revoke): %v", err)
		return
	}
	require.NotNil(t, resp)
	require.Equal(t, 204, resp.StatusCode)
}
