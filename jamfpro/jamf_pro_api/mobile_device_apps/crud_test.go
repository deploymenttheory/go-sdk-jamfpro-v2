package mobile_device_apps

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/mobile_device_apps/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*MobileDeviceApps, *mocks.MobileDeviceAppsMock) {
	t.Helper()
	mock := mocks.NewMobileDeviceAppsMock()
	return NewMobileDeviceApps(mock), mock
}

func TestUnit_MobileDeviceApps_NewService(t *testing.T) {
	mock := mocks.NewMobileDeviceAppsMock()
	svc := NewMobileDeviceApps(mock)
	require.NotNil(t, svc)
}

func TestUnit_MobileDeviceApps_ReinstallAppConfigV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterReinstallAppConfigMock()

	req := &RequestReinstallAppConfig{ReinstallCode: "abc-123-reinstall-code"}
	resp, err := svc.ReinstallAppConfigV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_MobileDeviceApps_ReinstallAppConfigV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.ReinstallAppConfigV1(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MobileDeviceApps_ReinstallAppConfigV1_EmptyReinstallCode(t *testing.T) {
	svc, _ := setupMockService(t)

	req := &RequestReinstallAppConfig{ReinstallCode: ""}
	resp, err := svc.ReinstallAppConfigV1(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "reinstallCode is required")
}

func TestUnit_MobileDeviceApps_ReinstallAppConfigV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	req := &RequestReinstallAppConfig{ReinstallCode: "invalid-code"}
	resp, err := svc.ReinstallAppConfigV1(context.Background(), req)
	assert.Error(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode())
	assert.Contains(t, err.Error(), "failed to reinstall app config")
}

func TestUnit_MobileDeviceApps_ReinstallAppConfigV1_NoMockRegistered(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterReinstallAppConfigErrorMock()

	req := &RequestReinstallAppConfig{ReinstallCode: "valid-code"}
	resp, err := svc.ReinstallAppConfigV1(context.Background(), req)
	assert.Error(t, err)
	assert.NotNil(t, resp)
}
