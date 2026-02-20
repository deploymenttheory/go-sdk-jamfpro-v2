package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_SelfServicePlusSettings_Get(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.SelfServicePlusSettings
	ctx := context.Background()
	result, resp, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestAcceptance_SelfServicePlusSettings_Update(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.SelfServicePlusSettings
	ctx := context.Background()
	current, _, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, current)
	request := *current
	request.Enabled = !request.Enabled
	_, resp, err := svc.UpdateV1(ctx, &request)
	require.NoError(t, err)
	require.NotNil(t, resp)
	request.Enabled = current.Enabled
	_, _, _ = svc.UpdateV1(ctx, &request)
}
