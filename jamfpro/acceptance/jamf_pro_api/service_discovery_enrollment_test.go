package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_ServiceDiscoveryEnrollment_GetV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ServiceDiscoveryEnrollment
	ctx := context.Background()

	result, resp, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestAcceptance_ServiceDiscoveryEnrollment_UpdateV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ServiceDiscoveryEnrollment
	ctx := context.Background()

	current, _, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, current)

	request := *current
	if len(request.WellKnownSettings) > 0 {
		origType := request.WellKnownSettings[0].EnrollmentType
		if request.WellKnownSettings[0].EnrollmentType == "USER_ENROLLMENT" {
			request.WellKnownSettings[0].EnrollmentType = "DEVICE_ENROLLMENT"
		} else {
			request.WellKnownSettings[0].EnrollmentType = "USER_ENROLLMENT"
		}
		_, resp, err := svc.UpdateV1(ctx, &request)
		require.NoError(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 204, resp.StatusCode)
		// Restore
		request.WellKnownSettings[0].EnrollmentType = origType
		_, _, _ = svc.UpdateV1(ctx, &request)
	}
}
