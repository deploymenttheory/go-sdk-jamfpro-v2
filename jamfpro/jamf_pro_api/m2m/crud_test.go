package m2m

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/m2m/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*M2M, *mocks.M2MMock) {
	t.Helper()
	mock := mocks.NewM2MMock()
	mock.RegisterMocks()
	return NewM2M(mock), mock
}

func TestUnit_M2M_GetTenantIdV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetTenantIdV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode())
	require.Equal(t, "5c4425d9-8181-42c6-b44a-0c0ea350614f", result.TenantId)
}

func TestUnit_M2M_GetTenantIdV1_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundError("GET", "/api/v1/m2m/tenant-id")
	result, resp, err := svc.GetTenantIdV1(context.Background())
	require.Error(t, err)
	require.Nil(t, result)
	require.NotNil(t, resp)
	require.Equal(t, 404, resp.StatusCode())
}
