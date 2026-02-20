package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_Ldap_GetLdapGroupsV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Ldap
	ctx := context.Background()

	result, resp, err := svc.GetLdapGroupsV1(ctx, map[string]string{"q": "test"})
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, result.Results)
}

func TestAcceptance_Ldap_GetLdapServersV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.Ldap
	ctx := context.Background()

	result, resp, err := svc.GetLdapServersV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}
