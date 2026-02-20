package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_LoginCustomization_GetV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.LoginCustomization
	ctx := context.Background()

	result, resp, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestAcceptance_LoginCustomization_UpdateV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.LoginCustomization
	ctx := context.Background()

	current, _, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, current)

	request := *current
	request.DisclaimerHeading = "Acceptance test" // max 20 chars per API
	if request.DisclaimerMainText == "" {
		request.DisclaimerMainText = "Acceptance test disclaimer text"
	}
	if request.ActionText == "" {
		request.ActionText = "Accept" // max 20 chars per API
	}
	updated, resp, err := svc.UpdateV1(ctx, &request)
	require.NoError(t, err)
	require.NotNil(t, updated)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	// Restore original
	request.DisclaimerHeading = current.DisclaimerHeading
	_, _, _ = svc.UpdateV1(ctx, &request)
}
