package jamf_pro_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/mobile_device_groups"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAcceptance_MobileDeviceGroups_v2_static_lifecycle exercises the v2 static
// mobile device group surface: Create → GetByID → Update → Delete. A static
// group with empty assignments is used so the test does not depend on real
// enrolled devices.
func TestAcceptance_MobileDeviceGroups_v2_static_lifecycle(t *testing.T) {
	acc.RequireClient(t)
	acc.GreaterThanJamfProVersion(t, 11, 27, 9) // v2 mobile-device-groups added in 11.28

	svc := acc.Client.JamfProAPI.MobileDeviceGroups
	ctx := context.Background()

	name := acc.UniqueName("sdkv2_acc_mdg_v2_static")
	created, _, err := svc.CreateStaticV2(ctx, &mobile_device_groups.RequestStaticMobileDeviceGroup{
		Name:        name,
		SiteId:      "-1",
		Assignments: []mobile_device_groups.StaticMobileDeviceGroupAssignment{},
	})
	require.NoError(t, err)
	require.NotEmpty(t, created.ID)

	acc.Cleanup(t, func() {
		_, delErr := svc.DeleteStaticByIDV2(context.Background(), created.ID)
		acc.LogCleanupDeleteError(t, "mobile static group v2", created.ID, delErr)
	})

	var fetched *mobile_device_groups.ResourceStaticMobileDeviceGroup
	err = acc.RetryOnNotFound(t, 3, 500_000_000, func() error {
		var getErr error
		fetched, _, getErr = svc.GetStaticByIDV2(ctx, created.ID)
		return getErr
	})
	require.NoError(t, err)
	assert.Equal(t, name, fetched.Name)

	updatedName := acc.UniqueName("sdkv2_acc_mdg_v2_static_upd")
	updated, _, err := svc.UpdateStaticByIDV2(ctx, created.ID, &mobile_device_groups.RequestStaticMobileDeviceGroup{
		Name:        updatedName,
		SiteId:      "-1",
		Assignments: []mobile_device_groups.StaticMobileDeviceGroupAssignment{},
	})
	require.NoError(t, err)
	assert.Equal(t, updatedName, updated.Name)

	delResp, err := svc.DeleteStaticByIDV2(ctx, created.ID)
	require.NoError(t, err)
	assert.Equal(t, 204, delResp.StatusCode())
}
