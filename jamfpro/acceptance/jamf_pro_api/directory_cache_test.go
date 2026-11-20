package jamf_pro_api

import (
	"context"
	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAcceptance_DirectoryCache_GetSyncStatusV1(t *testing.T) {
	acc.RequireClient(t)
	acc.AtLeastJamfProVersion(t, "11.34.0")
	result, resp, err := acc.Client.JamfProAPI.DirectoryCache.GetSyncStatusV1(context.Background())
	require.NoError(t, err)
	require.Equal(t, 200, resp.StatusCode())
	require.NotEmpty(t, result.State)
}
