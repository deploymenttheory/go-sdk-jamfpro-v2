package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_Notifications_list(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.JamfProAPI.Notifications
	ctx := context.Background()

	result, resp, err := svc.ListV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode())
}

// TestAcceptance_Notifications_delete_by_type_and_id_v1 tests the DELETE /v1/notifications/{type}/{id}
// endpoint added in Jamf Pro 11.25. It attempts to delete a notification of a known type with ID -1
// (which is always valid per the API spec) and expects a 204 response.
func TestAcceptance_Notifications_delete_by_type_and_id_v1(t *testing.T) {
	acc.RequireClient(t)
	acc.GreaterThanJamfProVersion(t, 11, 25, 0)
	svc := acc.Client.JamfProAPI.Notifications
	ctx := context.Background()

	// ID -1 is a valid special value accepted by the API (clears all notifications of this type).
	resp, err := svc.DeleteByTypeAndIDV1(ctx, "APNS_CERT_REVOKED", "-1")
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 204, resp.StatusCode())
}
