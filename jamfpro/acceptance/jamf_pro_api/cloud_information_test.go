package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAcceptance_CloudInformation_GetV1 tests retrieving cloud information.
func TestAcceptance_CloudInformation_GetV1(t *testing.T) {
	acc.RequireClient(t)
	client := acc.Client

	result, resp, err := client.CloudInformation.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)

	t.Logf("Cloud Information:")
	t.Logf("  Cloud Instance: %v", result.CloudInstance)
	t.Logf("  Ramp Instance: %v", result.RampInstance)
	t.Logf("  Gov Cloud Instance: %v", result.GovCloudInstance)
	t.Logf("  Managed Service Provider Instance: %v", result.ManagedServiceProviderInstance)
}
