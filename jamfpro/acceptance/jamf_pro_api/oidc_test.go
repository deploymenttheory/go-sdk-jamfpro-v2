package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/oidc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Acceptance Tests: OIDC (OpenID Connect)
// =============================================================================
//
// Service Operations Available
// -----------------------------------------------------------------------------
//   • GetDirectIdPLoginURLV1(ctx) - Gets the direct IdP login URL
//   • GetPublicKeyV1(ctx) - Gets the OIDC public key (JWKS)
//   • GetPublicFeaturesV1(ctx) - Gets public OIDC configuration features
//   • GenerateCertificateV1(ctx) - Generates a new OIDC signing certificate
//   • GetRedirectURLV1(ctx, request) - Gets the OIDC redirect URL for authentication
//
// Test Strategies Applied
// -----------------------------------------------------------------------------
//   ✓ Pattern 3: Read-Only Information
//     -- Tests: TestAcceptance_OIDC_public_endpoints_v1
//     -- Note: GetPublicFeaturesV1 and GetPublicKeyV1 always succeed
//     -- Note: GetDirectIdPLoginURLV1 may fail if OIDC is not configured
//
//   ✓ Pattern 7: Validation Errors
//     -- Tests: TestAcceptance_OIDC_validation_errors
//     -- Cases: GetRedirectURLV1(nil) → "OIDC redirect URL request cannot be nil"
//
// Notes
// -----------------------------------------------------------------------------
//   • GenerateCertificateV1 is intentionally not called in acceptance tests
//     because it would invalidate the existing OIDC signing certificate
//   • GetRedirectURLV1 requires a valid configuration; uses validation test only
//
// =============================================================================

// TestAcceptance_OIDC_public_endpoints_v1 verifies the public OIDC endpoints.
func TestAcceptance_OIDC_public_endpoints_v1(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.OIDC
	ctx := context.Background()

	// GetPublicFeaturesV1 — always returns whether Jamf ID auth is enabled
	acc.LogTestStage(t, "GetPublicFeaturesV1", "Fetching public OIDC features")

	features, featuresResp, err := svc.GetPublicFeaturesV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, features)
	require.NotNil(t, featuresResp)
	assert.Equal(t, 200, featuresResp.StatusCode)
	acc.LogTestSuccess(t, "GetPublicFeaturesV1: jamfIdAuthEnabled=%v", features.JamfIdAuthenticationEnabled)

	// GetPublicKeyV1 — returns the JWKS public key set
	acc.LogTestStage(t, "GetPublicKeyV1", "Fetching OIDC public key (JWKS)")

	pubKey, pubKeyResp, err := svc.GetPublicKeyV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, pubKey)
	require.NotNil(t, pubKeyResp)
	assert.Equal(t, 200, pubKeyResp.StatusCode)
	acc.LogTestSuccess(t, "GetPublicKeyV1: %d key(s) in JWKS", len(pubKey.Keys))

	// GetDirectIdPLoginURLV1 — may fail if OIDC IdP is not configured
	acc.LogTestStage(t, "GetDirectIdPLoginURLV1", "Fetching direct IdP login URL")

	loginURL, loginResp, err := svc.GetDirectIdPLoginURLV1(ctx)
	if err != nil {
		acc.LogTestWarning(t, "GetDirectIdPLoginURLV1 returned error (OIDC IdP may not be configured): %v", err)
	} else {
		require.NotNil(t, loginURL)
		require.NotNil(t, loginResp)
		assert.Equal(t, 200, loginResp.StatusCode)
		acc.LogTestSuccess(t, "GetDirectIdPLoginURLV1: url=%s", loginURL.URL)
	}
}

// TestAcceptance_OIDC_validation_errors verifies input validation.
func TestAcceptance_OIDC_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.OIDC

	t.Run("GetRedirectURLV1_NilRequest", func(t *testing.T) {
		_, _, err := svc.GetRedirectURLV1(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "OIDC redirect URL request cannot be nil")
	})

	t.Run("GetRedirectURLV1_EmptyRequest", func(t *testing.T) {
		// A non-nil but empty request should reach the server (not validation error)
		// This test verifies that an empty (but non-nil) request is accepted by the validator
		_, _, err := svc.GetRedirectURLV1(context.Background(), &oidc.RequestOIDCRedirectURL{})
		// May succeed or return server error, but should NOT return the nil validation error
		if err != nil {
			assert.NotContains(t, err.Error(), "OIDC redirect URL request cannot be nil")
		}
	})
}
