package jamf_pro_api

import (
	"context"
	"strconv"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/api_integrations"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/api_roles"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_ApiIntegrations_ListV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ApiIntegrations
	ctx := context.Background()

	result, resp, err := svc.ListV1(ctx, nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 0)
}

func TestAcceptance_ApiIntegrations_CreateGetUpdateDelete(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.ApiIntegrations
	roleSvc := acc.Client.APIRoles
	ctx := context.Background()
	name := acc.UniqueName("acc-api-integration")

	// Create a valid API role first (dependency chain: api_roles → api_integrations)
	roleName := acc.UniqueName("acc-api-integration-role")
	roleReq := &api_roles.RequestAPIRole{
		DisplayName: roleName,
		Privileges:  []string{"Read Computers"},
	}
	createdRole, roleResp, err := roleSvc.CreateV1(ctx, roleReq)
	require.NoError(t, err)
	require.NotNil(t, createdRole)
	require.NotNil(t, roleResp)
	assert.Contains(t, []int{200, 201}, roleResp.StatusCode)

	roleID := createdRole.ID
	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, _ = roleSvc.DeleteByIDV1(cleanupCtx, roleID)
	})

	// Create API integration using the valid role name
	created, resp, err := svc.CreateV1(ctx, &api_integrations.ResourceApiIntegration{
		DisplayName:                name,
		Enabled:                    true,
		AuthorizationScopes:        []string{roleName},
		AccessTokenLifetimeSeconds: 3600,
	})
	require.NoError(t, err)
	require.NotNil(t, created)
	assert.Equal(t, 200, resp.StatusCode)

	idStr := strconv.Itoa(created.ID)
	acc.Cleanup(t, func() {
		_, _ = svc.DeleteByIDV1(ctx, idStr)
	})

	getByID, resp, err := svc.GetByIDV1(ctx, idStr)
	require.NoError(t, err)
	require.NotNil(t, getByID)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, name, getByID.DisplayName)

	byName, resp, err := svc.GetByNameV1(ctx, name)
	require.NoError(t, err)
	require.NotNil(t, byName)
	assert.Equal(t, name, byName.DisplayName)

	creds, resp, err := svc.RefreshClientCredentialsByIDV1(ctx, idStr)
	require.NoError(t, err)
	require.NotNil(t, creds)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, creds.ClientID)
	assert.NotEmpty(t, creds.ClientSecret)

	updated, resp, err := svc.UpdateByIDV1(ctx, idStr, &api_integrations.ResourceApiIntegration{
		DisplayName:                name,
		Enabled:                    false,
		AuthorizationScopes:        created.AuthorizationScopes,
		AccessTokenLifetimeSeconds: 3600,
	})
	require.NoError(t, err)
	require.NotNil(t, updated)
	assert.Equal(t, 200, resp.StatusCode)

	delResp, err := svc.DeleteByIDV1(ctx, idStr)
	require.NoError(t, err)
	assert.Equal(t, 204, delResp.StatusCode)
}
