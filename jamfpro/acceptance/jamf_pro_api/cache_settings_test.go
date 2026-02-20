package jamf_pro_api

import (
	"context"
	"errors"
	"strings"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_CacheSettings_GetV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.CacheSettings
	ctx := context.Background()

	result, resp, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, result.CacheType)
}

func TestAcceptance_CacheSettings_UpdateV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.CacheSettings
	ctx := context.Background()

	current, _, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, current)

	request := *current
	request.TimeToLiveSeconds = current.TimeToLiveSeconds
	updated, resp, err := svc.UpdateV1(ctx, &request)
	if err != nil {
		var apiErr *client.APIError
		if errors.As(err, &apiErr) && apiErr.StatusCode == 403 && strings.Contains(strings.ToLower(apiErr.Message), "hosted") {
			t.Skipf("PUT cache-settings not available in hosted environments: %s", apiErr.Message)
		}
		require.NoError(t, err)
	}
	require.NotNil(t, updated)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}
