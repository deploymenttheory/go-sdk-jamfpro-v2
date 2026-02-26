package sso_failover

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/sso_failover/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.SSOFailoverMock) {
	t.Helper()
	mock := mocks.NewSSOFailoverMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnitGetV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "https://sso.example.com/failover?key=abc123", result.FailoverURL)
}

func TestUnitRegenerateV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.RegenerateV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.NotEmpty(t, result.FailoverURL)
}

func TestUnit_SSOFailover_GetV1_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetErrorMock()
	result, resp, err := svc.GetV1(context.Background())
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 500, resp.StatusCode)
}
