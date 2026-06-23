package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAcceptance_M2M_GetTenantIdV1 verifies the new (Jamf Pro 11.28)
// GET /api/v1/m2m/tenant-id endpoint returns a tenant ID.
func TestAcceptance_M2M_GetTenantIdV1(t *testing.T) {
	acc.RequireClient(t)
	acc.GreaterThanJamfProVersion(t, 11, 27, 9) // m2m/tenant-id added in 11.28

	result, resp, err := acc.Client.JamfProAPI.M2M.GetTenantIdV1(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotEmpty(t, result.TenantId)
}
