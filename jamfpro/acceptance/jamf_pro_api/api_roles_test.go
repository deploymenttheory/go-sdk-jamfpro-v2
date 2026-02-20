package jamf_pro_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/api_roles"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAcceptance_APIRoles_ListV1(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.APIRoles
	ctx := context.Background()

	result, resp, err := svc.ListV1(ctx, map[string]string{"page": "0", "page-size": "100"})
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.GreaterOrEqual(t, result.TotalCount, 0)
}

func TestAcceptance_APIRoles_Lifecycle(t *testing.T) {
	acc.RequireClient(t)
	svc := acc.Client.APIRoles
	ctx := context.Background()

	name := fmt.Sprintf("acc-api-role-%d", time.Now().UnixMilli())
	createReq := &api_roles.RequestAPIRole{
		DisplayName: name,
		Privileges:  []string{"Read Computers"},
	}
	created, createResp, err := svc.CreateV1(ctx, createReq)
	require.NoError(t, err)
	require.NotNil(t, created)
	require.NotNil(t, createResp)
	assert.Contains(t, []int{200, 201}, createResp.StatusCode)
	assert.NotEmpty(t, created.ID)

	roleID := created.ID
	acc.Cleanup(t, func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, _ = svc.DeleteByIDV1(cleanupCtx, roleID)
	})

	fetched, _, err := svc.GetByIDV1(ctx, roleID)
	require.NoError(t, err)
	require.NotNil(t, fetched)
	assert.Equal(t, name, fetched.DisplayName)

	updatedName := name + "-updated"
	updateReq := &api_roles.RequestAPIRole{DisplayName: updatedName, Privileges: []string{"Read Computers"}}
	_, updateResp, err := svc.UpdateByIDV1(ctx, roleID, updateReq)
	require.NoError(t, err)
	require.NotNil(t, updateResp)
	assert.Contains(t, []int{200, 202}, updateResp.StatusCode)

	delResp, err := svc.DeleteByIDV1(ctx, roleID)
	require.NoError(t, err)
	require.NotNil(t, delResp)
	assert.Equal(t, 204, delResp.StatusCode)
}
