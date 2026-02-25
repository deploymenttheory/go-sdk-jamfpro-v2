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

func TestUnitCreateV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestCloudDistributionPointV1{
		CdnType: "JAMF_CLOUD",
		Master:  true,
	}
	result, resp, err := svc.CreateV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 201, resp.StatusCode)
	require.True(t, result.HasConnectionSucceeded)
	require.Equal(t, "JCDS", result.CdnType)
}

func TestUnitUpdateV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	req := &RequestCloudDistributionPointV1{
		CdnType: "JAMF_CLOUD",
		Master:  true,
	}
	result, resp, err := svc.UpdateV1(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.True(t, result.HasConnectionSucceeded)
}

func TestUnitGetHistoryV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetHistoryV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, 1, result.TotalCount)
	require.Len(t, result.Results, 1)
	require.Equal(t, 1, result.Results[0].ID)
	require.Equal(t, "admin", result.Results[0].Username)
	require.Equal(t, "Sso settings update", result.Results[0].Note)
}

func TestUnit_CloudDistributionPoint_GetFilesV1_Success(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.GetFilesV1(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	require.Equal(t, "file-001", result.Results[0].ID)
	require.Equal(t, "package.pkg", result.Results[0].FileName)
	require.Equal(t, "AVAILABLE", result.Results[0].Status)
}

func TestUnit_CloudDistributionPoint_CreateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.CreateV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "request is required")
}

func TestUnit_CloudDistributionPoint_UpdateV1_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	result, resp, err := svc.UpdateV1(context.Background(), nil)
	require.Error(t, err)
	require.Nil(t, result)
	require.Nil(t, resp)
	require.Contains(t, err.Error(), "request is required")
}
