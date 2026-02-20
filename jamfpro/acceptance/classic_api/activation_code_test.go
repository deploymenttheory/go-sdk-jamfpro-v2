package classic_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_ActivationCode_GetAndUpdate gets the current activation code
// only. Update is not exercised in acceptance tests to avoid breaking the server
// with an invalid code.
// =============================================================================

func TestAcceptance_ActivationCode_GetActivationCode(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ActivationCode
	ctx := context.Background()

	acc.LogTestStage(t, "Get", "Fetching current activation code")

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	got, getResp, err := svc.GetActivationCode(ctx1)
	require.NoError(t, err, "GetActivationCode should not return an error")
	require.NotNil(t, got)
	require.NotNil(t, getResp)
	assert.Equal(t, 200, getResp.StatusCode)

	acc.LogTestSuccess(t, "Retrieved activation code: org=%q code=%q",
		got.OrganizationName, got.Code)
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
