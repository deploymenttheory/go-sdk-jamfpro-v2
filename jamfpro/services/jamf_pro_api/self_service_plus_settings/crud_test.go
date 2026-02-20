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

func TestUnitGet_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()
	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.False(t, result.Enabled)
}

func TestUnitUpdate_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()
	result, resp, err := svc.UpdateV1(context.Background(), &ResourceSelfServicePlusSettings{Enabled: true})
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
	assert.Nil(t, result)
}

func TestUnitUpdate_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UpdateV1(context.Background(), nil)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
}
