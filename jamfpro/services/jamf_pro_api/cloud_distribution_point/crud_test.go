package cloud_distribution_point

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/cloud_distribution_point/mocks"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.CloudDistributionPointMock) {
	t.Helper()
	mock := mocks.NewCloudDistributionPointMock()
	mock.RegisterMocks()
	return NewService(mock), mock
}

func TestUnitGetV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.True(t, result.HasConnectionSucceeded)
	require.Equal(t, "JCDS", result.CdnType)
}

func TestUnitGetUploadCapabilityV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetUploadCapabilityV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.True(t, result.DirectUploadCapable)
}

func TestUnitGetTestConnectionV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetTestConnectionV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.True(t, result.HasConnectionSucceeded)
}

func TestUnitDeleteV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	resp, err := svc.DeleteV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, resp)
	require.Equal(t, 204, resp.StatusCode)
}
