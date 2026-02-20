package device_communication_settings

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/device_communication_settings/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.DeviceCommunicationSettingsMock) {
	t.Helper()
	mock := mocks.NewDeviceCommunicationSettingsMock()
	return NewService(mock), mock
}

func TestUnitGetV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, result.AutoRenewMobileDeviceMdmProfileWhenCaRenewed)
	assert.Equal(t, 30, result.MdmProfileMobileDeviceExpirationLimitInDays)
}

func TestUnitUpdateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnitUpdateV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterPutMock()

	request := &ResourceDeviceCommunicationSettings{
		MdmProfileMobileDeviceExpirationLimitInDays: 30,
		MdmProfileComputerExpirationLimitInDays:    30,
	}
	result, resp, err := svc.UpdateV1(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
}
