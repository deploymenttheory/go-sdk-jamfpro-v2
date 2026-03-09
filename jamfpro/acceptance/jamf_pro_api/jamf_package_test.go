package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: Jamf Package
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • ListV1(ctx, application) - Lists packages for "protect" or "connect"
//   • GetV2(ctx, application)  - Gets package details for "protect" or "connect"
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Information
//     -- Tests: TestAcceptance_JamfPackage_list_and_get
//     -- Note: May return empty if Jamf Protect/Connect is not configured
//
//   ✓ Pattern 7: Validation Errors
//     -- Tests: TestAcceptance_JamfPackage_validation_errors
//     -- Cases: ListV1("invalid") → application must be error
//
// Notes
// -----------------------------------------------------------------------------
//   • Valid application values: "protect" (Jamf Protect) or "connect" (Jamf Connect)
//   • Returns empty results if Jamf Protect/Connect is not licensed/configured
//   • The test uses "protect" as the application value for all checks
//
// =============================================================================

// TestAcceptance_JamfPackage_list_and_get verifies listing and fetching package info.
func TestAcceptance_JamfPackage_list_and_get(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.JamfPackage
	ctx := context.Background()

	// ListV1 with "protect" application
	acc.LogTestStage(t, "ListV1", "Listing Jamf Protect packages")

	packages, listResp, err := svc.ListV1(ctx, "protect")
	require.NoError(t, err, "ListV1 should not error for valid application")
	require.NotNil(t, listResp)
	assert.Equal(t, 200, listResp.StatusCode())
	acc.LogTestSuccess(t, "ListV1 (protect): %d package(s) returned", len(packages))

	// GetV2 with "protect" application
	acc.LogTestStage(t, "GetV2", "Getting Jamf Protect package details")

	pkg, getResp, err := svc.GetV2(ctx, "protect")
	require.NoError(t, err, "GetV2 should not error for valid application")
	require.NotNil(t, pkg)
	assert.Equal(t, 200, getResp.StatusCode())
	acc.LogTestSuccess(t, "GetV2 (protect): displayName=%q artifacts=%d", pkg.DisplayName, len(pkg.Artifacts))

	// Also test "connect" application returns without error
	acc.LogTestStage(t, "ListV1", "Listing Jamf Connect packages")

	connectPackages, connectResp, err := svc.ListV1(ctx, "connect")
	require.NoError(t, err, "ListV1 should not error for 'connect' application")
	require.NotNil(t, connectResp)
	assert.Equal(t, 200, connectResp.StatusCode())
	acc.LogTestSuccess(t, "ListV1 (connect): %d package(s) returned", len(connectPackages))
}

// TestAcceptance_JamfPackage_validation_errors verifies input validation.
func TestAcceptance_JamfPackage_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.JamfProAPI.JamfPackage

	t.Run("ListV1_InvalidApplication", func(t *testing.T) {
		_, _, err := svc.ListV1(context.Background(), "invalid")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "application must be")
	})

	t.Run("GetV2_InvalidApplication", func(t *testing.T) {
		_, _, err := svc.GetV2(context.Background(), "invalid")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "application must be")
	})

	t.Run("ListV1_EmptyApplication", func(t *testing.T) {
		_, _, err := svc.ListV1(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "application must be")
	})
}
