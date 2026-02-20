package classic_api

import (
	"context"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/activation_code"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_ActivationCode_GetAndUpdate tests the Get and Update workflow:
// Get current activation code → Update with test values → Get to verify →
// Restore original values.
// =============================================================================

func TestAcceptance_ActivationCode_GetAndUpdate(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ActivationCode
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Get current activation code (to restore later)
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Get", "Fetching current activation code")

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	original, getResp, err := svc.GetActivationCode(ctx1)
	require.NoError(t, err, "GetActivationCode should not return an error")
	require.NotNil(t, original)
	require.NotNil(t, getResp)
	assert.Equal(t, 200, getResp.StatusCode)

	acc.LogTestSuccess(t, "Retrieved original activation code: org=%q code=%q",
		original.OrganizationName, original.Code)

	// Store original values to restore later
	originalOrgName := original.OrganizationName
	originalCode := original.Code

	// Register cleanup to restore original values
	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		restoreReq := &activation_code.RequestActivationCode{
			OrganizationName: originalOrgName,
			Code:             originalCode,
		}
		_, restoreErr := svc.UpdateActivationCode(cleanupCtx, restoreReq)
		if restoreErr != nil {
			acc.LogTestWarning(t, "Failed to restore original activation code: %v", restoreErr)
		} else {
			acc.LogTestSuccess(t, "Restored original activation code")
		}
	})

	// ------------------------------------------------------------------
	// 2. Update activation code with test values
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Update", "Updating activation code with test values")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	testOrgName := "Example Organization"
	testCode := "ABCD-1234-EFGH-5678"

	updateReq := &activation_code.RequestActivationCode{
		OrganizationName: testOrgName,
		Code:             testCode,
	}

	updateResp, err := svc.UpdateActivationCode(ctx2, updateReq)
	require.NoError(t, err, "UpdateActivationCode should not return an error")
	require.NotNil(t, updateResp)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode, "expected 200 or 201")

	acc.LogTestSuccess(t, "Updated activation code: org=%q code=%q", testOrgName, testCode)

	// ------------------------------------------------------------------
	// 3. Get activation code again to verify update
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "Get (verify)", "Fetching activation code to verify update")

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	updated, verifyResp, err := svc.GetActivationCode(ctx3)
	require.NoError(t, err, "GetActivationCode should not return an error")
	require.NotNil(t, updated)
	require.NotNil(t, verifyResp)
	assert.Equal(t, 200, verifyResp.StatusCode)

	assert.Equal(t, testOrgName, updated.OrganizationName, "organization name should match updated value")
	assert.Equal(t, testCode, updated.Code, "code should match updated value")

	acc.LogTestSuccess(t, "Verified updated activation code: org=%q code=%q",
		updated.OrganizationName, updated.Code)
}

// =============================================================================
// TestAcceptance_ActivationCode_ValidationErrors tests client-side validation
// without making any network calls.
// =============================================================================

func TestAcceptance_ActivationCode_ValidationErrors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ActivationCode

	t.Run("UpdateActivationCode_NilRequest", func(t *testing.T) {
		_, err := svc.UpdateActivationCode(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "request is required")
	})
}
