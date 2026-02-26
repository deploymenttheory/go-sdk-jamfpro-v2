package self_service_plus_settings

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/self_service_plus_settings/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.SelfServicePlusSettingsMock) {
	t.Helper()
	mock := mocks.NewSelfServicePlusSettingsMock()
	return NewService(mock), mock
}

func TestUnit_SelfServicePlusSettings_Get_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()
	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.False(t, result.Enabled)
}

func TestUnit_SelfServicePlusSettings_Update_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()
	result, resp, err := svc.UpdateV1(context.Background(), &ResourceSelfServicePlusSettings{Enabled: true})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
	assert.Nil(t, result)
}

func TestUnit_SelfServicePlusSettings_Update_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UpdateV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_SelfServicePlusSettings_GetFeatureToggleEnabledV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterFeatureToggleMock()
	enabled, resp, err := svc.GetFeatureToggleEnabledV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, enabled)
}

func TestUnit_SelfServicePlusSettings_GetFeatureToggleEnabledV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	enabled, resp, err := svc.GetFeatureToggleEnabledV1(context.Background())
	require.Error(t, err)
	assert.False(t, enabled)
	_ = resp
}

func TestUnit_SelfServicePlusSettings_GetV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.GetV1(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_SelfServicePlusSettings_UpdateV1_Error(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateV1(context.Background(), &ResourceSelfServicePlusSettings{Enabled: true})
	require.Error(t, err)
	assert.Nil(t, result)
	_ = resp
}

func TestUnit_SelfServicePlusSettings_UpdateV1_Non204Response(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateNon204Mock()

	result, resp, err := svc.UpdateV1(context.Background(), &ResourceSelfServicePlusSettings{Enabled: true})
	require.NoError(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
}
