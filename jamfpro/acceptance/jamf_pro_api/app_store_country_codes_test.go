package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: App Store Country Codes
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ListV1(ctx) - Returns all App Store country codes (no pagination/rsql)
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Information
//     -- Reason: Service is read-only; returns static list of country codes
//     -- Tests: TestAcceptance_AppStoreCountryCodes_list_v1
//     -- Flow: List → Verify at least one code returned with non-empty Code/Name
//
// Notes
// -----------------------------------------------------------------------------
//   • Returns a static list from Jamf Pro — always non-empty in any environment
//   • No RSQL, pagination, or mutation operations
//
// =============================================================================

func TestAcceptance_AppStoreCountryCodes_list_v1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.AppStoreCountryCodes
	ctx := context.Background()

	result, resp, err := svc.ListV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotEmpty(t, result.CountryCodes, "country codes list should not be empty")

	if len(result.CountryCodes) > 0 {
		assert.NotEmpty(t, result.CountryCodes[0].Code)
		assert.NotEmpty(t, result.CountryCodes[0].Name)
	}

	acc.LogTestSuccess(t, "ListV1: returned %d country codes", len(result.CountryCodes))
}
