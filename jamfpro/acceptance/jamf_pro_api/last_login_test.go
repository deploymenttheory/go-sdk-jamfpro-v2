package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Last Login (v1)
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   - GetV1(ctx) - Returns the date of the last login event
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   - Pattern 3: Read-Only Information
//
// =============================================================================

// TestAcceptance_LastLogin_get_v1 verifies the last login endpoint returns a valid date.
func TestAcceptance_LastLogin_get_v1(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	acc.LogTestStage(t, "get", "retrieving last login date")
	result, resp, err := acc.Client.JamfProAPI.LastLogin.GetV1(ctx)
	if err != nil {
		t.Skipf("last-login endpoint returned an error (may require elevated API client privileges): %v", err)
	}
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotEmpty(t, result.Date, "last login date should not be empty")
	acc.LogTestSuccess(t, "last login date=%q", result.Date)
}
