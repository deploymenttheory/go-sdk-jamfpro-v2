package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/environment_type"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Environment Type
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetV2(ctx) - Gets the cloud services environment type for this instance
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Information
//     -- Reason: Singleton read-only endpoint with no mutating counterpart
//     -- Tests: TestAcceptance_EnvironmentType_GetV2
//
// Notes
// -----------------------------------------------------------------------------
//   • Added in Jamf Pro 11.31; older instances return 404
//
// =============================================================================

func TestAcceptance_EnvironmentType_GetV2(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.JamfProAPI.EnvironmentType
	ctx := context.Background()

	result, resp, err := svc.GetV2(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())

	assert.Contains(t, []string{
		environment_type.EnvironmentStaging,
		environment_type.EnvironmentProduction,
		environment_type.EnvironmentSandbox,
	}, result.Environment, "unexpected environment type returned by the server")
}
