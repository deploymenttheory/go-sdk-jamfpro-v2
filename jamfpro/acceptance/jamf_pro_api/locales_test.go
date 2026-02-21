package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Locales
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ListV1(ctx) - Retrieves all available locales
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Information
//     -- Reason: Service only provides read access to available locales
//     -- Tests: TestAcceptance_Locales_List
//     -- Flow: List locales → Verify response structure
//
//   Note: RSQL Filter Testing NOT applicable
//     -- ListV1 does not support filtering - returns all locales
//
// Test Coverage
// -----------------------------------------------------------------------------
//   ✓ List all locales
//   ✓ Verify response structure (200 status)
//   ✓ Verify non-empty result set
//
// Notes
// -----------------------------------------------------------------------------
//   • This is a read-only endpoint - no mutations possible
//   • Returns all available locales for the Jamf Pro instance
//   • Locales used for internationalization/localization
//   • No filtering or pagination - returns complete list
//   • No cleanup needed as this is a read-only operation
//
// =============================================================================

func TestAcceptance_Locales_List(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Locales
	ctx := context.Background()

	result, resp, err := svc.ListV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, result)
}
