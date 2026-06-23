package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAcceptance_Groups_v2_list_and_get exercises the read side of the v2
// unified groups surface (List → GetByID) against live data. It is read-only to
// avoid mutating the instance's existing groups.
func TestAcceptance_Groups_v2_list_and_get(t *testing.T) {
	acc.RequireClient(t)
	acc.GreaterThanJamfProVersion(t, 11, 27, 9) // v2 unified groups added in 11.28

	svc := acc.Client.JamfProAPI.Groups
	ctx := context.Background()

	list, resp, err := svc.ListV2(ctx, nil)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	if len(list.Results) == 0 {
		t.Skip("no groups present to GET")
	}

	platformID := list.Results[0].GroupPlatformId
	require.NotEmpty(t, platformID, "v2 group list entry should carry a groupPlatformId")

	got, _, err := svc.GetByIDV2(ctx, platformID)
	require.NoError(t, err)
	assert.Equal(t, platformID, got.GroupPlatformId)
}
