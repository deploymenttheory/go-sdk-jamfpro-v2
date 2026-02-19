package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_Icons_GetByID(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Icons
	ctx := context.Background()

	// Assume at least one icon exists (e.g. id 1)
	result, resp, err := svc.GetByIDV1(ctx, 1)
	if err != nil {
		t.Skipf("GetByIDV1 failed (no icon?): %v", err)
		return
	}
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.ID, 1)
}
