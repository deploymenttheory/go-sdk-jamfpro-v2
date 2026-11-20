package jamf_pro_api

import (
	"context"
	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAcceptance_ActionLogs_ReadV1(t *testing.T) {
	acc.RequireClient(t)
	acc.AtLeastJamfProVersion(t, "11.34.0")
	svc := acc.Client.JamfProAPI.ActionLogs
	ctx := context.Background()
	list, resp, err := svc.ListV1(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode())
	if len(list.Results) == 0 {
		t.Skip("no action history available")
	}
	id := list.Results[0].ID
	detail, resp, err := svc.GetByIDV1(ctx, id)
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode())
	require.Equal(t, id, detail.ID)
	devices, resp, err := svc.GetDevicesV1(ctx, id, nil)
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode())
	require.NotNil(t, devices)
}
