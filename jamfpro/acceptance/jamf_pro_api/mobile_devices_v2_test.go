package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_MobileDevices_v2_detail(t *testing.T) {
	acc.RequireClient(t)
	// The /detail endpoint predates 11.29; the exception-handling param is 11.29.
	// 11,27,9 lets this run on 11.28 while still gating very old instances.
	acc.GreaterThanJamfProVersion(t, 11, 27, 9)
	ctx := context.Background()

	svc := acc.Client.JamfProAPI.MobileDevices

	detail, resp, err := svc.GetDetailV2(ctx, map[string]string{
		"page-size":          "1",
		"exception-handling": "LENIENT",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	require.NotNil(t, detail)
	assert.GreaterOrEqual(t, detail.TotalCount, 0)

	list, listResp, err := svc.ListV2(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, listResp)
	assert.Equal(t, 200, listResp.StatusCode())
	require.NotNil(t, list)
	assert.GreaterOrEqual(t, list.TotalCount, 0)
}
