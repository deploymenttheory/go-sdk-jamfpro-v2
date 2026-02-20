package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_SsoSettings_GetV3(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.SsoSettings
	ctx := context.Background()

	result, resp, err := svc.GetV3(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestAcceptance_SsoSettings_GetEnrollmentCustomizationDependencies(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.SsoSettings
	ctx := context.Background()

	result, resp, err := svc.GetEnrollmentCustomizationDependenciesV3(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestAcceptance_SsoSettings_UpdateV3(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.SsoSettings
	ctx := context.Background()

	current, _, err := svc.GetV3(ctx)
	require.NoError(t, err)
	require.NotNil(t, current)

	request := *current
	request.SsoBypassAllowed = !request.SsoBypassAllowed
	updated, resp, err := svc.UpdateV3(ctx, &request)
	require.NoError(t, err)
	require.NotNil(t, updated)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Restore
	request.SsoBypassAllowed = current.SsoBypassAllowed
	_, _, _ = svc.UpdateV3(ctx, &request)
}
