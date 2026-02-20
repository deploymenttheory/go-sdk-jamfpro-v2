package jamf_pro_api

import (
	"context"
	"math/rand"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/icons"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// findValidIconID tries IDs in random order (1–100) until GetByIDV1 returns 200.
// Logs success to test notifications. Returns (id, result, true) when found, (0, nil, false) when none exist.
func findValidIconID(t *testing.T) (int, *icons.ResourceIcon, bool) {
	t.Helper()
	acc.RequireClient(t)
	svc := acc.Client.Icons
	ctx := context.Background()

	acc.LogTestStage(t, "Icons", "Finding icon: trying IDs 1–100 until GetByID returns 200")
	perm := rand.Perm(100)
	for i := range perm {
		id := perm[i] + 1
		result, resp, err := svc.GetByIDV1(ctx, id)
		if err != nil {
			continue
		}
		if resp != nil && resp.StatusCode == 200 && result != nil {
			acc.LogTestSuccess(t, "GetByID succeeded for icon ID=%d", id)
			return id, result, true
		}
	}
	return 0, nil, false
}

func TestAcceptance_Icons_GetByID(t *testing.T) {
	id, result, ok := findValidIconID(t)
	if !ok {
		t.Skip("no icon with ID in 1–100 found in this environment")
		return
	}
	require.NotNil(t, result)
	assert.GreaterOrEqual(t, id, 1)
	assert.Equal(t, id, result.ID)
}

func TestAcceptance_Icons_Download(t *testing.T) {
	id, _, ok := findValidIconID(t)
	if !ok {
		t.Skip("no icon with ID in 1–100 found in this environment")
		return
	}

	acc.RequireClient(t)
	svc := acc.Client.Icons
	ctx := context.Background()

	acc.LogTestStage(t, "Icons", "Attempting download for icon ID=%d", id)
	body, resp, err := svc.DownloadV1(ctx, id, "original", "0")
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, body)
}
