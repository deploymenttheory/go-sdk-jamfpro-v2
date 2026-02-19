package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/require"
)

// TestAcceptance_CloudDistributionPoint_Get verifies GET cloud distribution point.
// Create/Update/Delete are not run to avoid changing instance CDP configuration.
func TestAcceptance_CloudDistributionPoint_Get(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.CloudDistributionPoint
	ctx := context.Background()

	result, resp, err := svc.GetV1(ctx)
	// 404 is acceptable when no CDP is configured
	if err != nil && resp != nil && resp.StatusCode == 404 {
		t.Skip("No cloud distribution point configured")
		return
	}
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
}
