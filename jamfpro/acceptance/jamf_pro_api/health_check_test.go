package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Health Check
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetV1(ctx) - Returns whether the Jamf Pro API is healthy (bool)
//   • GetHealthStatusV1(ctx) - Returns request acceptance ratios per concurrency
//     group (Jamf Cloud only; returns 404 on non-cloud nodes)
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Information
//     -- Tests: TestAcceptance_HealthCheck_get_v1
//     -- Flow: GetV1 → verify healthy == true
//
//   ✓ Pattern 3: Read-Only Information (cloud-only, graceful skip on 404)
//     -- Tests: TestAcceptance_HealthCheck_get_health_status_v1
//     -- Flow: GetHealthStatusV1 → verify metrics are non-negative; skip on 404
//
// =============================================================================

// TestAcceptance_HealthCheck_get_v1 verifies the API health endpoint returns healthy.
func TestAcceptance_HealthCheck_get_v1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.HealthCheck
	ctx := context.Background()

	acc.LogTestStage(t, "Get", "Checking Jamf Pro API health")

	healthy, resp, err := svc.GetV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Contains(t, []int{200, 204}, resp.StatusCode())
	if resp.StatusCode() == 200 {
		assert.True(t, healthy, "Jamf Pro API should report as healthy")
	}

	acc.LogTestSuccess(t, "HealthCheck: healthy=%v status=%d", healthy, resp.StatusCode())
}

// TestAcceptance_HealthCheck_get_health_status_v1 verifies the health status endpoint.
// Only available on Jamf Cloud; returns 404 on non-cloud nodes.
func TestAcceptance_HealthCheck_get_health_status_v1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.HealthCheck
	ctx := context.Background()

	acc.LogTestStage(t, "GetHealthStatus", "Getting Jamf Pro API health status metrics")

	result, resp, err := svc.GetHealthStatusV1(ctx)
	if err != nil && resp != nil && resp.StatusCode() == 404 {
		t.Log("GetHealthStatusV1 returned 404 - not available on non-cloud nodes, skipping")
		return
	}
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())

	// Metric values are acceptance ratios between 0 and 1
	assert.GreaterOrEqual(t, result.API.OneMinute, 0.0)
	assert.GreaterOrEqual(t, result.UI.OneMinute, 0.0)

	acc.LogTestSuccess(t, "HealthStatus: api.1m=%.2f ui.1m=%.2f enrollment.1m=%.2f",
		result.API.OneMinute, result.UI.OneMinute, result.Enrollment.OneMinute)
}
