package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/computer_groups"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAcceptance_ComputerGroups_v3_smart_lifecycle exercises the v3 smart
// computer group surface: Create → GetByID → Update → Membership → Delete.
func TestAcceptance_ComputerGroups_v3_smart_lifecycle(t *testing.T) {
	acc.RequireClient(t)
	acc.GreaterThanJamfProVersion(t, 11, 27, 9) // v3 computer-groups added in 11.28

	svc := acc.Client.JamfProAPI.ComputerGroups
	ctx := context.Background()

	name := acc.UniqueName("sdkv2_acc_cg_v3_smart")
	created, _, err := svc.CreateSmartV3(ctx, &computer_groups.RequestSmartGroupV3{
		Name:   name,
		SiteID: "-1",
		Criteria: []computer_groups.CriterionV3{
			{Name: "Operating System", Priority: 0, AndOr: "and", SearchType: "like", Value: "macOS"},
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, created.ID)

	acc.Cleanup(t, func() {
		_, delErr := svc.DeleteSmartByIDV3(context.Background(), created.ID)
		acc.LogCleanupDeleteError(t, "computer smart group v3", created.ID, delErr)
	})

	var fetched *computer_groups.ResourceSmartGroupV3
	err = acc.RetryOnNotFound(t, 3, 500_000_000, func() error {
		var getErr error
		fetched, _, getErr = svc.GetSmartByIDV3(ctx, created.ID)
		return getErr
	})
	require.NoError(t, err)
	assert.Equal(t, name, fetched.Name)
	require.NotEmpty(t, fetched.Criteria)

	updatedName := acc.UniqueName("sdkv2_acc_cg_v3_smart_upd")
	_, _, err = svc.UpdateSmartByIDV3(ctx, created.ID, &computer_groups.RequestSmartGroupV3{
		Name:     updatedName,
		SiteID:   "-1",
		Criteria: fetched.Criteria,
	})
	require.NoError(t, err)

	_, _, err = svc.GetSmartGroupMembershipByIDV3(ctx, created.ID)
	require.NoError(t, err)

	delResp, err := svc.DeleteSmartByIDV3(ctx, created.ID)
	require.NoError(t, err)
	assert.Equal(t, 204, delResp.StatusCode())
}

// TestAcceptance_ComputerGroups_v3_static_lifecycle exercises the v3 static
// computer group surface: Create → GetByID → Update → Delete.
func TestAcceptance_ComputerGroups_v3_static_lifecycle(t *testing.T) {
	acc.RequireClient(t)
	acc.GreaterThanJamfProVersion(t, 11, 27, 9)

	svc := acc.Client.JamfProAPI.ComputerGroups
	ctx := context.Background()

	name := acc.UniqueName("sdkv2_acc_cg_v3_static")
	created, _, err := svc.CreateStaticV3(ctx, &computer_groups.RequestStaticGroupV3{
		Name:   name,
		SiteID: "-1",
	})
	require.NoError(t, err)
	require.NotEmpty(t, created.ID)

	acc.Cleanup(t, func() {
		_, delErr := svc.DeleteStaticByIDV3(context.Background(), created.ID)
		acc.LogCleanupDeleteError(t, "computer static group v3", created.ID, delErr)
	})

	var fetched *computer_groups.ResourceStaticGroupV3
	err = acc.RetryOnNotFound(t, 3, 500_000_000, func() error {
		var getErr error
		fetched, _, getErr = svc.GetStaticByIDV3(ctx, created.ID)
		return getErr
	})
	require.NoError(t, err)
	assert.Equal(t, name, fetched.Name)

	_, _, err = svc.UpdateStaticByIDV3(ctx, created.ID, &computer_groups.RequestStaticGroupV3{
		Name:   acc.UniqueName("sdkv2_acc_cg_v3_static_upd"),
		SiteID: "-1",
	})
	require.NoError(t, err)

	delResp, err := svc.DeleteStaticByIDV3(ctx, created.ID)
	require.NoError(t, err)
	assert.Equal(t, 204, delResp.StatusCode())
}
