package jcds

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jcds/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_Jcds_GetPackagesV1_Success(t *testing.T) {
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

func TestUnit_Jcds_GetPackagesV1_NoMockRegistered(t *testing.T) {
	mock := mocks.NewJCDSMock()
	service := NewService(mock)

	result, resp, err := service.GetPackagesV1(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to get JCDS packages")
}

func TestUnit_Jcds_GetPackagesV1_APICError(t *testing.T) {
	mock := mocks.NewJCDSMock()
	mock.RegisterErrorMock("GET", "/api/v1/jcds/files", "api error")
	service := NewService(mock)

	result, resp, err := service.GetPackagesV1(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to get JCDS packages")
}

func TestUnit_Jcds_GetPackageURIByNameV1_Success(t *testing.T) {
	mock := mocks.NewJCDSMock()
	mock.RegisterGetPackageURIByNameMock()
	service := NewService(mock)

	result, resp, err := service.GetPackageURIByNameV1(context.Background(), "test-package.pkg")
	require.NoError(t, err)
	require.NotNil(t, result)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "s3://jamf-bucket/path/test-package.pkg", result.URI)
}

func TestUnit_Jcds_GetPackageURIByNameV1_EmptyName(t *testing.T) {
	mock := mocks.NewJCDSMock()
	service := NewService(mock)

	result, resp, err := service.GetPackageURIByNameV1(context.Background(), "")
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "package name is required")
}

func TestUnit_Jcds_GetPackageURIByNameV1_APICError(t *testing.T) {
	mock := mocks.NewJCDSMock()
	mock.RegisterErrorMock("GET", "/api/v1/jcds/files/some.pkg", "api error")
	service := NewService(mock)

	result, resp, err := service.GetPackageURIByNameV1(context.Background(), "some.pkg")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to get JCDS package URI")
}

func TestUnit_Jcds_RenewCredentialsV1_Success(t *testing.T) {
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

func TestUnit_Jcds_RenewCredentialsV1_APICError(t *testing.T) {
	mock := mocks.NewJCDSMock()
	mock.RegisterErrorMock("POST", "/api/v1/jcds/renew-credentials", "api error")
	service := NewService(mock)

	result, resp, err := service.RenewCredentialsV1(context.Background())
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to renew JCDS credentials")
}

func TestUnit_Jcds_RefreshInventoryV1_Success(t *testing.T) {
	mock := mocks.NewJCDSMock()
	mock.RegisterRefreshInventoryMock()
	service := NewService(mock)

	resp, err := service.RefreshInventoryV1(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnit_Jcds_RefreshInventoryV1_APICError(t *testing.T) {
	mock := mocks.NewJCDSMock()
	mock.RegisterErrorMock("POST", "/api/v1/jcds/refresh-inventory", "api error")
	service := NewService(mock)

	resp, err := service.RefreshInventoryV1(context.Background())
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to refresh JCDS inventory")
}

func TestUnit_Jcds_CreatePackageV1_EmptyFilePath(t *testing.T) {
	mock := mocks.NewJCDSMock()
	service := NewService(mock)

	result, resp, err := service.CreatePackageV1(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "file path is required")
}

func TestUnit_Jcds_CreatePackageV1_CredentialError(t *testing.T) {
	mock := mocks.NewJCDSMock()
	mock.RegisterErrorMock("POST", "/api/v1/jcds/files", "credential error")
	service := NewService(mock)

	tmp := t.TempDir()
	pkgPath := filepath.Join(tmp, "test.pkg")
	require.NoError(t, os.WriteFile(pkgPath, []byte("content"), 0644))

	result, resp, err := service.CreatePackageV1(context.Background(), pkgPath)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to obtain upload credentials")
}

func TestUnit_Jcds_CreatePackageV1_IncompleteCredentials(t *testing.T) {
	mock := mocks.NewJCDSMock()
	mock.RegisterIncompleteCredentialsMock()
	service := NewService(mock)

	tmp := t.TempDir()
	pkgPath := filepath.Join(tmp, "test.pkg")
	require.NoError(t, os.WriteFile(pkgPath, []byte("content"), 0644))

	result, resp, err := service.CreatePackageV1(context.Background(), pkgPath)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Contains(t, err.Error(), "incomplete upload credentials")
}

func TestUnit_Jcds_CreatePackageV1_InvalidExtension(t *testing.T) {
	mock := mocks.NewJCDSMock()
	mock.RegisterUploadCredentialsMock()
	service := NewService(mock)

	tmp := t.TempDir()
	txtPath := filepath.Join(tmp, "test.txt")
	require.NoError(t, os.WriteFile(txtPath, []byte("content"), 0644))

	result, resp, err := service.CreatePackageV1(context.Background(), txtPath)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Contains(t, err.Error(), "failed to read package file")
}

func TestUnit_Jcds_CreatePackageV1_FileNotFound(t *testing.T) {
	mock := mocks.NewJCDSMock()
	mock.RegisterUploadCredentialsMock()
	service := NewService(mock)

	result, resp, err := service.CreatePackageV1(context.Background(), "/nonexistent/test.pkg")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Contains(t, err.Error(), "failed to read package file")
}

func TestUnit_Jcds_DeletePackageV1_EmptyFilePath(t *testing.T) {
	mock := mocks.NewJCDSMock()
	service := NewService(mock)

	resp, err := service.DeletePackageV1(context.Background(), "")
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "file path is required")
}

func TestUnit_Jcds_DeletePackageV1_CredentialError(t *testing.T) {
	mock := mocks.NewJCDSMock()
	mock.RegisterErrorMock("POST", "/api/v1/jcds/files", "credential error")
	service := NewService(mock)

	resp, err := service.DeletePackageV1(context.Background(), "/path/to/test.pkg")
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "failed to obtain deletion credentials")
}

func TestUnit_Jcds_DeletePackageV1_IncompleteCredentials(t *testing.T) {
	mock := mocks.NewJCDSMock()
	mock.RegisterIncompleteCredentialsMock()
	service := NewService(mock)

	resp, err := service.DeletePackageV1(context.Background(), "/path/to/test.pkg")
	require.Error(t, err)
	assert.NotNil(t, resp)
	assert.Contains(t, err.Error(), "incomplete deletion credentials")
}

func TestUnit_Jcds_CreatePackageV1_S3UploadFails(t *testing.T) {
	mock := mocks.NewJCDSMock()
	mock.RegisterUploadCredentialsMock()
	service := NewService(mock)

	tmp := t.TempDir()
	pkgPath := filepath.Join(tmp, "test.pkg")
	require.NoError(t, os.WriteFile(pkgPath, []byte("package content"), 0644))

	// Uses fake AWS credentials - S3 upload will fail with InvalidClientTokenId or similar
	result, resp, err := service.CreatePackageV1(context.Background(), pkgPath)
	require.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
	assert.Contains(t, err.Error(), "failed to upload file")
}

func TestUnit_Jcds_DeletePackageV1_S3DeleteFails(t *testing.T) {
	mock := mocks.NewJCDSMock()
	mock.RegisterUploadCredentialsMock()
	service := NewService(mock)

	// Uses fake AWS credentials - S3 delete will fail
	resp, err := service.DeletePackageV1(context.Background(), "/path/to/test.pkg")
	require.Error(t, err)
	assert.NotNil(t, resp)
	assert.Contains(t, err.Error(), "failed to delete file")
}
