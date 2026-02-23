package jcds

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jcds/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnitGetPackagesV1_Success(t *testing.T) {
	mock := mocks.NewJCDSMock()
	mock.RegisterGetPackagesMock()
	service := NewService(mock)

	result, resp, err := service.GetPackagesV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Len(t, result, 1)
	assert.Equal(t, "test-package.pkg", result[0].FileName)
	assert.Equal(t, int64(1024000), result[0].Length)
}

func TestUnitGetPackageURIByNameV1_Success(t *testing.T) {
	mock := mocks.NewJCDSMock()
	mock.RegisterGetPackageURIByNameMock()
	service := NewService(mock)

	result, resp, err := service.GetPackageURIByNameV1(context.Background(), "test-package.pkg")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "s3://jamf-bucket/path/test-package.pkg", result.URI)
}

func TestUnitGetPackageURIByNameV1_EmptyName(t *testing.T) {
	mock := mocks.NewJCDSMock()
	service := NewService(mock)

	result, resp, err := service.GetPackageURIByNameV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "package name is required")
}

func TestUnitRenewCredentialsV1_Success(t *testing.T) {
	mock := mocks.NewJCDSMock()
	mock.RegisterRenewCredentialsMock()
	service := NewService(mock)

	result, resp, err := service.RenewCredentialsV1(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "test-access-key", result.AccessKeyID)
	assert.Equal(t, "us-east-1", result.Region)
	assert.Equal(t, "jamf-bucket", result.BucketName)
}

func TestUnitRefreshInventoryV1_Success(t *testing.T) {
	mock := mocks.NewJCDSMock()
	mock.RegisterRefreshInventoryMock()
	service := NewService(mock)

	resp, err := service.RefreshInventoryV1(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode)
}

// Note: CreatePackageV1 and DeletePackageV1 are not unit tested here because they
// involve AWS SDK operations that are difficult to properly mock without significant
// infrastructure. These operations are better tested through integration or acceptance tests
// with actual AWS credentials and S3 buckets, or through manual testing with Jamf Pro.
