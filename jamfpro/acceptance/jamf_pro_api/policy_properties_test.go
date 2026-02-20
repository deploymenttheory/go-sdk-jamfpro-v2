package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_PolicyProperties_Get(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.PolicyProperties
	ctx := context.Background()
	result, resp, err := svc.Get(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestAcceptance_PolicyProperties_Update(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.PolicyProperties
	ctx := context.Background()
	current, _, err := svc.Get(ctx)
	require.NoError(t, err)
	require.NotNil(t, current)
	request := *current
	request.AllowNetworkStateChangeTriggers = !request.AllowNetworkStateChangeTriggers
	_, resp, err := svc.Update(ctx, &request)
	require.NoError(t, err)
	require.NotNil(t, resp)
	request.AllowNetworkStateChangeTriggers = current.AllowNetworkStateChangeTriggers
	_, _, _ = svc.Update(ctx, &request)
}
