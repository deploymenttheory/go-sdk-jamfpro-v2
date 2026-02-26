package health_check

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/health_check/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.HealthCheckMock) {
	t.Helper()
	mock := mocks.NewHealthCheckMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnit_HealthCheck_GetV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	healthy, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.True(t, healthy)
	require.Equal(t, 200, resp.StatusCode)
}

func TestUnit_HealthCheck_GetV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterErrorMock()
	healthy, resp, err := svc.GetV1(context.Background())
	require.Error(t, err)
	require.False(t, healthy)
	require.NotNil(t, resp)
}
