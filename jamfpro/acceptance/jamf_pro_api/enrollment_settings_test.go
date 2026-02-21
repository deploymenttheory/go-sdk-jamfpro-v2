package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Enrollment Settings
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetV4(ctx) - Retrieves current enrollment settings (read-only)
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Information
//     -- Reason: Service only provides read access to enrollment settings
//     -- Tests: TestAcceptance_EnrollmentSettings_GetV4
//     -- Flow: Get settings → Verify response structure and status code
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ Get enrollment settings
//   ✓ Verify response structure (200 status)
//
// Notes
// -----------------------------------------------------------------------------
//   • This is a read-only endpoint - no update operations available
//   • Enrollment settings retrieved via V4 API
//   • No cleanup needed as this is a read-only operation
//   • Unlike other settings services, this one does not support updates
//
// =============================================================================

func TestAcceptance_EnrollmentSettings_GetV4(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.EnrollmentSettings
	ctx := context.Background()

	result, resp, err := svc.GetV4(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}
