package cache_settings

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/cache_settings/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*CacheSettings, *mocks.CacheSettingsMock) {
	t.Helper()
	mock := mocks.NewCacheSettingsMock()
	return NewCacheSettings(mock), mock
}

func TestUnit_CacheSettings_NewService(t *testing.T) {
	mock := mocks.NewCacheSettingsMock()
	svc := NewCacheSettings(mock)
	require.NotNil(t, svc)
}

func TestUnit_CacheSettings_GetV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "ehcache", result.CacheType)
	assert.Equal(t, 3600, result.TimeToLiveSeconds)
}

func TestUnit_CacheSettings_GetV1_NoMock(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered

	result, resp, err := svc.GetV1(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response")
}

func TestUnit_CacheSettings_UpdateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.UpdateV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}

func TestUnit_CacheSettings_UpdateV1_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterPutMock()

	request := &ResourceCacheSettings{
		CacheType:         "ehcache",
		TimeToLiveSeconds: 3600,
		TimeToIdleSeconds: 1800,
	}
	result, resp, err := svc.UpdateV1(context.Background(), request)
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_CacheSettings_UpdateV1_NoMock(t *testing.T) {
	svc, _ := setupMockService(t)
	// No mock registered

	request := &ResourceCacheSettings{
		CacheType:         "ehcache",
		TimeToLiveSeconds: 3600,
	}
	result, resp, err := svc.UpdateV1(context.Background(), request)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "no response")
}
