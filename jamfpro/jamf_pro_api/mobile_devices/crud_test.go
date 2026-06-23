package mobile_devices

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/mobile_devices/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_MobileDevices_ListV2(t *testing.T) {
	mock := mocks.NewMobileDevicesMock()
	mock.RegisterListMock()

	svc := NewMobileDevices(mock)
	ctx := context.Background()

	result, resp, err := svc.ListV2(ctx, nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "1", result.Results[0].ID)
	assert.Equal(t, "iPad", result.Results[0].Name)
	assert.Equal(t, "DMQVGC0DHLA0", result.Results[0].SerialNumber)
	assert.Equal(t, "ios", result.Results[0].Type)
	assert.Equal(t, "tvos", result.Results[1].Type)
}

func TestUnit_MobileDevices_GetByIDV2(t *testing.T) {
	mock := mocks.NewMobileDevicesMock()
	mock.RegisterGetByIDMock("1")

	svc := NewMobileDevices(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByIDV2(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "iPad", result.Name)
	assert.Equal(t, "0dad565fb40b010a9e490440188063a378721069", result.Udid)
	assert.Equal(t, "73226fb6-61df-4c10-9552-eb9bc353d507", result.ManagementID)
}

func TestUnit_MobileDevices_GetByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewMobileDevicesMock()
	svc := NewMobileDevices(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByIDV2(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_MobileDevices_GetByIDV2_NotFound(t *testing.T) {
	mock := mocks.NewMobileDevicesMock()
	mock.RegisterGetByIDNotFoundMock("999")

	svc := NewMobileDevices(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByIDV2(ctx, "999")

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Equal(t, 404, resp.StatusCode())
}

func TestUnit_MobileDevices_GetDetailV2(t *testing.T) {
	mock := mocks.NewMobileDevicesMock()
	mock.RegisterGetDetailMock()

	svc := NewMobileDevices(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDetailV2(ctx, map[string]string{"exception-handling": "LENIENT"})

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)

	device := result.Results[0]
	assert.Equal(t, "1", device.MobileDeviceID)
	assert.Equal(t, "ios", device.DeviceType)
	require.NotNil(t, device.Hardware)
	assert.Equal(t, "5c28fdae", device.Hardware.SerialNumber)
	assert.Equal(t, "NORMAL", device.Hardware.BatteryHealth)
	require.NotNil(t, device.General)
	assert.Equal(t, "Banezicron", device.General.DisplayName)
	require.NotNil(t, device.UserAndLocation)
	assert.Equal(t, "admin", device.UserAndLocation.Username)
	// 11.29 additive fields
	assert.Equal(t, "mdmadmin", device.UserAndLocation.LastLoggedInUsernameMdm)
	assert.Equal(t, "2024-10-31T18:04:13Z", device.UserAndLocation.LastLoggedInUsernameMdmTimestamp)
}

func TestUnit_MobileDevices_GetDetailByIDV2(t *testing.T) {
	mock := mocks.NewMobileDevicesMock()
	mock.RegisterGetDetailByIDMock("1")

	svc := NewMobileDevices(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDetailByIDV2(ctx, "1", nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "iPad", result.Name)
	assert.Equal(t, "ios", result.Type)
	require.NotNil(t, result.Site)
	// 11.29 additive field on the site object
	assert.Equal(t, "5", result.Site.DivisionID)
	require.NotNil(t, result.Ios)
	assert.Equal(t, "iPad7,11", result.Ios.ModelIdentifier)
	require.NotNil(t, result.Ios.Security)
	assert.Equal(t, "NOT_SUPPORTED", result.Ios.Security.BootstrapTokenEscrowed)
}

func TestUnit_MobileDevices_GetDetailByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewMobileDevicesMock()
	svc := NewMobileDevices(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDetailByIDV2(ctx, "", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_MobileDevices_GetPairedDevicesByIDV2(t *testing.T) {
	mock := mocks.NewMobileDevicesMock()
	mock.RegisterGetPairedDevicesByIDMock("1")

	svc := NewMobileDevices(mock)
	ctx := context.Background()

	result, resp, err := svc.GetPairedDevicesByIDV2(ctx, "1", nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].MobileDeviceID)
}

func TestUnit_MobileDevices_GetPairedDevicesByIDV2_EmptyID(t *testing.T) {
	mock := mocks.NewMobileDevicesMock()
	svc := NewMobileDevices(mock)
	ctx := context.Background()

	result, resp, err := svc.GetPairedDevicesByIDV2(ctx, "", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_MobileDevices_ListV2_ClientError(t *testing.T) {
	mock := mocks.NewMobileDevicesMock()
	mock.RegisterListErrorMock()

	svc := NewMobileDevices(mock)
	ctx := context.Background()

	result, resp, err := svc.ListV2(ctx, nil)

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "simulated ListV2 API error")
}

func TestUnit_MobileDevices_GetDetailV2_ClientError(t *testing.T) {
	mock := mocks.NewMobileDevicesMock()
	mock.RegisterGetDetailErrorMock()

	svc := NewMobileDevices(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDetailV2(ctx, nil)

	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "simulated GetDetailV2 API error")
}
