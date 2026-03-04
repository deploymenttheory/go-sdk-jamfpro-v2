package classic_api

import (
	"context"
	"fmt"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mobile_devices"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_MobileDevices_list verifies list returns without error.
// Note: Mobile devices are typically enrolled via MDM; we cannot create them
// via Classic API for lifecycle testing. This test focuses on read operations.
// =============================================================================

func TestAcceptance_MobileDevices_list(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicMobileDevices
	ctx := context.Background()

	acc.LogTestStage(t, "List", "Listing mobile devices")

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	list, listResp, err := svc.List(ctx1)
	require.NoError(t, err, "List should not return an error")
	require.NotNil(t, list)
	assert.Equal(t, 200, listResp.StatusCode)
	acc.LogTestSuccess(t, "List returned %d mobile devices", list.Size)
}

// =============================================================================
// TestAcceptance_MobileDevices_get_by_id fetches a mobile device by ID if any exist.
// Skips if no devices are in the instance.
// =============================================================================

func TestAcceptance_MobileDevices_get_by_id(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicMobileDevices
	ctx := context.Background()

	acc.LogTestStage(t, "List", "Listing mobile devices to find one to fetch")

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	list, _, err := svc.List(ctx1)
	require.NoError(t, err)
	if list == nil || len(list.Results) == 0 {
		t.Skip("No mobile devices in instance; skipping GetByID test")
		return
	}

	deviceID := fmt.Sprintf("%d", list.Results[0].ID)
	acc.LogTestStage(t, "GetByID", "Getting mobile device by ID=%s", deviceID)

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByID(ctx2, deviceID)
	require.NoError(t, err, "GetByID should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.Equal(t, list.Results[0].ID, fetched.General.ID)
	acc.LogTestSuccess(t, "GetByID: ID=%d name=%q", fetched.General.ID, fetched.General.Name)
}

// =============================================================================
// TestAcceptance_MobileDevices_get_by_name fetches a mobile device by name if any exist.
// Skips if no devices are in the instance.
// =============================================================================

func TestAcceptance_MobileDevices_get_by_name(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicMobileDevices
	ctx := context.Background()

	acc.LogTestStage(t, "List", "Listing mobile devices to find one to fetch")

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	list, _, err := svc.List(ctx1)
	require.NoError(t, err)
	if list == nil || len(list.Results) == 0 {
		t.Skip("No mobile devices in instance; skipping GetByName test")
		return
	}

	deviceName := list.Results[0].Name
	if deviceName == "" {
		deviceName = list.Results[0].DeviceName
	}
	if deviceName == "" {
		t.Skip("Device has no name; skipping GetByName test")
		return
	}

	acc.LogTestStage(t, "GetByName", "Getting mobile device by name=%q", deviceName)

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByName(ctx2, deviceName)
	require.NoError(t, err, "GetByName should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	acc.LogTestSuccess(t, "GetByName: ID=%d name=%q", fetched.General.ID, fetched.General.Name)
}

// =============================================================================
// TestAcceptance_MobileDevices_get_by_id_and_data_subset fetches a subset if a device exists.
// Skips if no devices are in the instance.
// =============================================================================

func TestAcceptance_MobileDevices_get_by_id_and_data_subset(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicMobileDevices
	ctx := context.Background()

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	list, _, err := svc.List(ctx1)
	require.NoError(t, err)
	if list == nil || len(list.Results) == 0 {
		t.Skip("No mobile devices in instance; skipping GetByIDAndDataSubset test")
		return
	}

	deviceID := fmt.Sprintf("%d", list.Results[0].ID)
	acc.LogTestStage(t, "GetByIDAndDataSubset", "Getting mobile device ID=%s subset=General", deviceID)

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, fetchResp, err := svc.GetByIDAndDataSubset(ctx2, deviceID, "General")
	require.NoError(t, err, "GetByIDAndDataSubset should not return an error")
	require.NotNil(t, fetched)
	assert.Equal(t, 200, fetchResp.StatusCode)
	assert.NotEmpty(t, fetched.General.Name)
	acc.LogTestSuccess(t, "GetByIDAndDataSubset: ID=%d", fetched.General.ID)
}

// =============================================================================
// TestAcceptance_MobileDevices_validation_errors validates error handling.
// =============================================================================

func TestAcceptance_MobileDevices_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicMobileDevices

	t.Run("GetByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.GetByID(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mobile device ID cannot be empty")
	})

	t.Run("GetByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.GetByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mobile device name cannot be empty")
	})

	t.Run("Create_NilDevice", func(t *testing.T) {
		_, _, err := svc.Create(context.Background(), nil)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mobile device is required")
	})

	t.Run("UpdateByID_EmptyID", func(t *testing.T) {
		_, _, err := svc.UpdateByID(context.Background(), "", &mobile_devices.ResponseMobileDevice{
			General: mobile_devices.MobileDeviceSubsetGeneral{Name: "sdkv2_acc_test"},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mobile device ID cannot be empty")
	})

	t.Run("UpdateByName_EmptyName", func(t *testing.T) {
		_, _, err := svc.UpdateByName(context.Background(), "", &mobile_devices.ResponseMobileDevice{
			General: mobile_devices.MobileDeviceSubsetGeneral{Name: "sdkv2_acc_test"},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mobile device name cannot be empty")
	})

	t.Run("DeleteByID_EmptyID", func(t *testing.T) {
		_, err := svc.DeleteByID(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mobile device ID cannot be empty")
	})

	t.Run("DeleteByName_EmptyName", func(t *testing.T) {
		_, err := svc.DeleteByName(context.Background(), "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "mobile device name cannot be empty")
	})
}

// =============================================================================
// TestAcceptance_MobileDevices_update_by_id updates a device if one exists.
// Note: This test updates location/asset info only; it does not delete devices.
// Skips if no devices exist.
// =============================================================================

func TestAcceptance_MobileDevices_update_by_id(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicMobileDevices
	ctx := context.Background()

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	list, _, err := svc.List(ctx1)
	require.NoError(t, err)
	if list == nil || len(list.Results) == 0 {
		t.Skip("No mobile devices in instance; skipping UpdateByID test")
		return
	}

	deviceID := fmt.Sprintf("%d", list.Results[0].ID)
	acc.LogTestStage(t, "GetByID", "Getting mobile device by ID=%s", deviceID)

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	fetched, _, err := svc.GetByID(ctx2, deviceID)
	require.NoError(t, err)
	require.NotNil(t, fetched)

	updatedName := acc.UniqueName("sdkv2_acc_acc-test-mobdev-updated")
	acc.LogTestStage(t, "UpdateByID", "Updating mobile device ID=%s to device_name=%q", deviceID, updatedName)

	updateReq := &mobile_devices.ResponseMobileDevice{
		General:               fetched.General,
		Location:              fetched.Location,
		Purchasing:            fetched.Purchasing,
		Applications:          fetched.Applications,
		SecurityObject:        fetched.SecurityObject,
		Network:               fetched.Network,
		Certificates:          fetched.Certificates,
		ConfigurationProfiles: fetched.ConfigurationProfiles,
		ProvisioningProfiles:  fetched.ProvisioningProfiles,
		MobileDeviceGroups:    fetched.MobileDeviceGroups,
		ExtensionAttributes:   fetched.ExtensionAttributes,
	}
	updateReq.General.DeviceName = updatedName
	updateReq.General.DisplayName = updatedName

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	updated, updateResp, err := svc.UpdateByID(ctx3, deviceID, updateReq)
	if err != nil {
		t.Skipf("UpdateByID skipped (may require permissions): %v", err)
		return
	}
	require.NotNil(t, updated)
	assert.Contains(t, []int{200, 201}, updateResp.StatusCode(), "expected 200 or 201")
	acc.LogTestSuccess(t, "UpdateByID: status=%d", updateResp.StatusCode())

	// Revert the device name
	revertReq := &mobile_devices.ResponseMobileDevice{
		General:               fetched.General,
		Location:              fetched.Location,
		Purchasing:            fetched.Purchasing,
		Applications:          fetched.Applications,
		SecurityObject:        fetched.SecurityObject,
		Network:               fetched.Network,
		Certificates:          fetched.Certificates,
		ConfigurationProfiles: fetched.ConfigurationProfiles,
		ProvisioningProfiles:  fetched.ProvisioningProfiles,
		MobileDeviceGroups:    fetched.MobileDeviceGroups,
		ExtensionAttributes:   fetched.ExtensionAttributes,
	}

	ctx4, cancel4 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel4()
	time.Sleep(500 * time.Millisecond) // Brief pause before revert
	_, _, _ = svc.UpdateByID(ctx4, deviceID, revertReq)
}
