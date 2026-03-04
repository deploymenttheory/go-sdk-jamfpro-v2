package classic_api

import (
	"context"
	"testing"
	"time"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/command_flush"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// TestAcceptance_CommandFlush_flush_by_id_and_status tests the FlushByIDAndStatus
// operation for clearing MDM commands on individual devices or groups.
// =============================================================================

func TestAcceptance_CommandFlush_flush_by_id_and_status(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicCommandFlush
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Flush Failed commands for a computer
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "FlushByIDAndStatus", "Flushing failed computer commands")

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	resp, err := svc.FlushByIDAndStatus(ctx1, "computers", "1", "Failed")
	require.NoError(t, err, "FlushByIDAndStatus should not return an error")
	require.NotNil(t, resp)
	assert.Contains(t, []int{200, 201, 204}, resp.StatusCode(), "expected 200, 201, or 204")

	acc.LogTestSuccess(t, "Successfully flushed failed commands for computer 1")

	// ------------------------------------------------------------------
	// 2. Flush Pending commands for a mobile device
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "FlushByIDAndStatus", "Flushing pending mobile device commands")

	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	resp2, err := svc.FlushByIDAndStatus(ctx2, "mobiledevices", "1", "Pending")
	require.NoError(t, err, "FlushByIDAndStatus should not return an error")
	require.NotNil(t, resp2)
	assert.Contains(t, []int{200, 201, 204}, resp2.StatusCode(), "expected 200, 201, or 204")

	acc.LogTestSuccess(t, "Successfully flushed pending commands for mobile device 1")

	// ------------------------------------------------------------------
	// 3. Flush Pending+Failed commands for a computer group
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "FlushByIDAndStatus", "Flushing pending+failed commands for computer group")

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	resp3, err := svc.FlushByIDAndStatus(ctx3, "computergroups", "1", "Pending+Failed")
	require.NoError(t, err, "FlushByIDAndStatus should not return an error")
	require.NotNil(t, resp3)
	assert.Contains(t, []int{200, 201, 204}, resp3.StatusCode(), "expected 200, 201, or 204")

	acc.LogTestSuccess(t, "Successfully flushed pending+failed commands for computer group 1")
}

// =============================================================================
// TestAcceptance_CommandFlush_flush_with_xml tests the FlushWithXML operation
// for batch clearing MDM commands using XML request body.
// =============================================================================

func TestAcceptance_CommandFlush_flush_with_xml(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicCommandFlush
	ctx := context.Background()

	// ------------------------------------------------------------------
	// 1. Check if mobile devices exist
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "FlushWithXML", "Checking for mobile devices")

	mobileDeviceSvc := acc.Client.ClassicMobileDevices
	ctx0, cancel0 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel0()

	mobileDevicesList, _, err := mobileDeviceSvc.List(ctx0)
	if err != nil || mobileDevicesList == nil || len(mobileDevicesList.Results) == 0 {
		t.Skip("No mobile devices found - skipping FlushWithXML test")
	}

	deviceID := mobileDevicesList.Results[0].ID
	acc.LogTestSuccess(t, "Found mobile device with ID=%d", deviceID)

	// ------------------------------------------------------------------
	// 2. Flush Pending commands for mobile device
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "FlushWithXML", "Flushing pending commands for mobile device")

	req := &command_flush.RequestCommandFlush{
		Status: "Pending",
		MobileDevices: &command_flush.MobileDevices{
			MobileDevice: []command_flush.DeviceID{
				{ID: deviceID},
			},
		},
	}

	ctx1, cancel1 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel1()

	resp, err := svc.FlushWithXML(ctx1, req)
	require.NoError(t, err, "FlushWithXML should not return an error")
	require.NotNil(t, resp)
	assert.Contains(t, []int{200, 201, 204}, resp.StatusCode(), "expected 200, 201, or 204")

	acc.LogTestSuccess(t, "Successfully flushed pending commands for mobile device %d", deviceID)

	// ------------------------------------------------------------------
	// 3. Check if computers exist
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "FlushWithXML", "Checking for computers")

	computerSvc := acc.Client.ClassicComputers
	ctx2, cancel2 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel2()

	computersList, _, err := computerSvc.List(ctx2)
	if err != nil || computersList == nil || len(computersList.Results) == 0 {
		acc.LogTestSuccess(t, "No computers found - skipping computer flush test")
		return
	}

	computerID := computersList.Results[0].ID
	acc.LogTestSuccess(t, "Found computer with ID=%d", computerID)

	// ------------------------------------------------------------------
	// 4. Flush Failed commands for computer
	// ------------------------------------------------------------------
	acc.LogTestStage(t, "FlushWithXML", "Flushing failed commands for computer")

	req2 := &command_flush.RequestCommandFlush{
		Status: "Failed",
		Computers: &command_flush.Computers{
			Computer: []command_flush.DeviceID{
				{ID: computerID},
			},
		},
	}

	ctx3, cancel3 := context.WithTimeout(ctx, acc.Config.RequestTimeout)
	defer cancel3()

	resp2, err := svc.FlushWithXML(ctx3, req2)
	require.NoError(t, err, "FlushWithXML should not return an error")
	require.NotNil(t, resp2)
	assert.Contains(t, []int{200, 201, 204}, resp2.StatusCode(), "expected 200, 201, or 204")

	acc.LogTestSuccess(t, "Successfully flushed failed commands for computer %d", computerID)
}

// =============================================================================
// TestAcceptance_CommandFlush_validation_errors tests validation error handling.
// =============================================================================

func TestAcceptance_CommandFlush_validation_errors(t *testing.T) {
	acc.RequireClient(t)

	svc := acc.Client.ClassicCommandFlush
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Test invalid idType
	acc.LogTestStage(t, "ValidationErrors", "Testing invalid idType")
	_, err := svc.FlushByIDAndStatus(ctx, "invalidtype", "1", "Pending")
	require.Error(t, err, "should return error for invalid idType")
	assert.Contains(t, err.Error(), "invalid idType")
	acc.LogTestSuccess(t, "Validation correctly rejected invalid idType")

	// Test invalid status
	acc.LogTestStage(t, "ValidationErrors", "Testing invalid status")
	_, err = svc.FlushByIDAndStatus(ctx, "computers", "1", "InvalidStatus")
	require.Error(t, err, "should return error for invalid status")
	assert.Contains(t, err.Error(), "invalid status")
	acc.LogTestSuccess(t, "Validation correctly rejected invalid status")

	// Test XML request with no devices
	acc.LogTestStage(t, "ValidationErrors", "Testing XML request with no devices")
	req := &command_flush.RequestCommandFlush{
		Status: "Pending",
	}
	_, err = svc.FlushWithXML(ctx, req)
	require.Error(t, err, "should return error for request with no devices")
	assert.Contains(t, err.Error(), "at least one device list")
	acc.LogTestSuccess(t, "Validation correctly rejected empty device list")
}
