package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_AppInstallers_list_titles_v1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.JamfProAPI.AppInstallers
	ctx := context.Background()

	result, resp, err := svc.ListTitlesV1(ctx, map[string]string{"page": "0", "page-size": "50"})
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode())
	require.GreaterOrEqual(t, result.TotalCount, 0)
}

func TestAcceptance_AppInstallers_list_deployments_v1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.JamfProAPI.AppInstallers
	ctx := context.Background()

	result, resp, err := svc.ListDeploymentsV1(ctx, map[string]string{"page": "0", "page-size": "50"})
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode())
	require.GreaterOrEqual(t, result.TotalCount, 0)
}
