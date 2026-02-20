package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_SsoCertificate_GetV2(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.SsoCertificate
	ctx := context.Background()
	result, resp, err := svc.GetV2(ctx)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	if result != nil {
		_ = result.Keystore.Type
	}
}
