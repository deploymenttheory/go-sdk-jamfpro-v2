package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_Notifications_List(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Notifications
	ctx := context.Background()

	result, resp, err := svc.ListV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 200, resp.StatusCode)
}
