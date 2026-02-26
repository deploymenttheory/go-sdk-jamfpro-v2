package disk_encryption_configurations

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/disk_encryption_configurations/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.DiskEncryptionConfigurationsMock) {
	t.Helper()
	mock := mocks.NewDiskEncryptionConfigurationsMock()
	return NewService(mock), mock
}

// =============================================================================
// ListDiskEncryptionConfigurations
// =============================================================================

func TestUnit_DiskEncryptionConfigurations_List_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListMock()

	result, resp, err := svc.List(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.Size)
	require.Len(t, result.Results, 2)
	assert.Equal(t, 1, result.Results[0].ID)
	assert.Equal(t, "FileVault Config", result.Results[0].Name)
}

// =============================================================================
// GetDiskEncryptionConfigurationByID
// =============================================================================

func TestUnit_DiskEncryptionConfigurations_GetByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByIDMock()

	result, resp, err := svc.GetByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "FileVault Config", result.Name)
	assert.Equal(t, "Individual", result.KeyType)
	assert.Equal(t, "Management Account", result.FileVaultEnabledUsers)
}

func TestUnit_DiskEncryptionConfigurations_GetByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetByID(context.Background(), 0)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "disk encryption configuration ID must be a positive integer")
}

func TestUnit_DiskEncryptionConfigurations_GetByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetByID(context.Background(), -1)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "disk encryption configuration ID must be a positive integer")
}

func TestUnit_DiskEncryptionConfigurations_GetByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	_, _, err := svc.GetByID(context.Background(), 999)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "404")
}

// =============================================================================
// GetDiskEncryptionConfigurationByName
// =============================================================================

func TestUnit_DiskEncryptionConfigurations_GetByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetByNameMock()

	result, resp, err := svc.GetByName(context.Background(), "FileVault Config")
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "FileVault Config", result.Name)
}

func TestUnit_DiskEncryptionConfigurations_GetByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetByName(context.Background(), "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "disk encryption configuration name is required")
}

// =============================================================================
// CreateDiskEncryptionConfiguration
// =============================================================================

func TestUnit_DiskEncryptionConfigurations_Create_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateMock()

	req := &RequestDiskEncryptionConfiguration{
		Name:                  "FileVault Config",
		KeyType:               "Individual",
		FileVaultEnabledUsers: "Management Account",
	}
	result, resp, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnit_DiskEncryptionConfigurations_Create_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.Create(context.Background(), nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_DiskEncryptionConfigurations_Create_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestDiskEncryptionConfiguration{Name: "FileVault Config"}
	_, _, err := svc.Create(context.Background(), req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "409")
}

// =============================================================================
// UpdateDiskEncryptionConfigurationByID
// =============================================================================

func TestUnit_DiskEncryptionConfigurations_UpdateByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateByIDMock()

	req := &RequestDiskEncryptionConfiguration{Name: "FileVault Config"}
	result, resp, err := svc.UpdateByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnit_DiskEncryptionConfigurations_UpdateByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByID(context.Background(), 0, &RequestDiskEncryptionConfiguration{Name: "x"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "disk encryption configuration ID must be a positive integer")
}

func TestUnit_DiskEncryptionConfigurations_UpdateByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByID(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateDiskEncryptionConfigurationByName
// =============================================================================

func TestUnit_DiskEncryptionConfigurations_UpdateByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateByNameMock()

	req := &RequestDiskEncryptionConfiguration{Name: "FileVault Config"}
	result, resp, err := svc.UpdateByName(context.Background(), "FileVault Config", req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnit_DiskEncryptionConfigurations_UpdateByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByName(context.Background(), "", &RequestDiskEncryptionConfiguration{Name: "x"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "disk encryption configuration name is required")
}

func TestUnit_DiskEncryptionConfigurations_UpdateByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateByName(context.Background(), "FileVault Config", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteDiskEncryptionConfigurationByID
// =============================================================================

func TestUnit_DiskEncryptionConfigurations_DeleteByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteByIDMock()

	resp, err := svc.DeleteByID(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnit_DiskEncryptionConfigurations_DeleteByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.DeleteByID(context.Background(), 0)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "disk encryption configuration ID must be a positive integer")
}

// =============================================================================
// DeleteDiskEncryptionConfigurationByName
// =============================================================================

func TestUnit_DiskEncryptionConfigurations_DeleteByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteByNameMock()

	resp, err := svc.DeleteByName(context.Background(), "FileVault Config")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnit_DiskEncryptionConfigurations_DeleteByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.DeleteByName(context.Background(), "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "disk encryption configuration name is required")
}
