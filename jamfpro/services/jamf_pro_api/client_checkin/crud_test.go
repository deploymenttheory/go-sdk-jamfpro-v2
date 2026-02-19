package client_checkin

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/client_checkin/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.ClientCheckinMock) {
	t.Helper()
	mock := mocks.NewClientCheckinMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnitGetV3_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetV3(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, 15, result.CheckInFrequency)
	require.True(t, result.CreateHooks)
}

func TestUnitUpdateV3_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	settings := &ResourceClientCheckinSettings{CheckInFrequency: 30, CreateHooks: true}
	result, resp, err := svc.UpdateV3(context.Background(), settings)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
}

func TestUnitUpdateV3_NilSettings(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UpdateV3(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
}
