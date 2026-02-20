package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_APIRolePrivileges_ListV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.APIRolePrivileges
	ctx := context.Background()

	result, resp, err := svc.ListV1(ctx)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, result.Privileges)
}

func TestAcceptance_APIRolePrivileges_SearchByNameV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.APIRolePrivileges
	ctx := context.Background()

	result, resp, err := svc.SearchPrivilegesByNameV1(ctx, "Read", 20)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}
