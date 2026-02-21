package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConditionalAccess_GetDeviceComplianceFeatureToggle(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.ConditionalAccess

	result, _, err := svc.GetDeviceComplianceFeatureToggleV1(ctx)
	if err != nil {
		t.Skipf("Failed to get conditional access device compliance feature toggle (may not be supported): %v", err)
		return
	}

	require.NotNil(t, result)
	assert.NotNil(t, result.SharedDeviceFeatureEnabled)
}
