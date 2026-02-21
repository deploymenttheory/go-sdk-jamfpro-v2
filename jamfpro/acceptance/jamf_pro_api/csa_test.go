package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCSA_GetTokenExchangeDetails(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.Csa

	result, _, err := svc.GetTokenExchangeDetailsV1(ctx)
	if err != nil {
		t.Skipf("Failed to get CSA token exchange details (may not be configured): %v", err)
		return
	}

	require.NotNil(t, result)
	assert.NotEmpty(t, result.TenantID)
}

func TestCSA_GetTenantID(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	svc := acc.Client.Csa

	result, _, err := svc.GetTenantIDV1(ctx)
	if err != nil {
		t.Skipf("Failed to get CSA tenant ID (may not be configured): %v", err)
		return
	}

	require.NotNil(t, result)
	assert.NotEmpty(t, result.TenantID)
}
