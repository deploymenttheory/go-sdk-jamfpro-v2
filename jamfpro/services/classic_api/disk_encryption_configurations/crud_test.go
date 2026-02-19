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

func TestUnitListDiskEncryptionConfigurations_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListDiskEncryptionConfigurationsMock()

	result, resp, err := svc.ListDiskEncryptionConfigurations(context.Background())
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

func TestUnitGetDiskEncryptionConfigurationByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDiskEncryptionConfigurationByIDMock()

	result, resp, err := svc.GetDiskEncryptionConfigurationByID(context.Background(), 1)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "FileVault Config", result.Name)
	assert.Equal(t, "Individual", result.KeyType)
	assert.Equal(t, "Management Account", result.FileVaultEnabledUsers)
}

func TestUnitGetDiskEncryptionConfigurationByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetDiskEncryptionConfigurationByID(context.Background(), 0)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "disk encryption configuration ID must be a positive integer")
}

func TestUnitGetDiskEncryptionConfigurationByID_NegativeID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetDiskEncryptionConfigurationByID(context.Background(), -1)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "disk encryption configuration ID must be a positive integer")
}

func TestUnitGetDiskEncryptionConfigurationByID_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	_, _, err := svc.GetDiskEncryptionConfigurationByID(context.Background(), 999)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "404")
}

// =============================================================================
// GetDiskEncryptionConfigurationByName
// =============================================================================

func TestUnitGetDiskEncryptionConfigurationByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetDiskEncryptionConfigurationByNameMock()

	result, resp, err := svc.GetDiskEncryptionConfigurationByName(context.Background(), "FileVault Config")
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "FileVault Config", result.Name)
}

func TestUnitGetDiskEncryptionConfigurationByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.GetDiskEncryptionConfigurationByName(context.Background(), "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "disk encryption configuration name is required")
}

// =============================================================================
// CreateDiskEncryptionConfiguration
// =============================================================================

func TestUnitCreateDiskEncryptionConfiguration_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterCreateDiskEncryptionConfigurationMock()

	req := &RequestDiskEncryptionConfiguration{
		Name:                  "FileVault Config",
		KeyType:               "Individual",
		FileVaultEnabledUsers: "Management Account",
	}
	result, resp, err := svc.CreateDiskEncryptionConfiguration(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnitCreateDiskEncryptionConfiguration_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.CreateDiskEncryptionConfiguration(context.Background(), nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnitCreateDiskEncryptionConfiguration_Conflict(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterConflictErrorMock()

	req := &RequestDiskEncryptionConfiguration{Name: "FileVault Config"}
	_, _, err := svc.CreateDiskEncryptionConfiguration(context.Background(), req)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "409")
}

// =============================================================================
// UpdateDiskEncryptionConfigurationByID
// =============================================================================

func TestUnitUpdateDiskEncryptionConfigurationByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateDiskEncryptionConfigurationByIDMock()

	req := &RequestDiskEncryptionConfiguration{Name: "FileVault Config"}
	result, resp, err := svc.UpdateDiskEncryptionConfigurationByID(context.Background(), 1, req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnitUpdateDiskEncryptionConfigurationByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateDiskEncryptionConfigurationByID(context.Background(), 0, &RequestDiskEncryptionConfiguration{Name: "x"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "disk encryption configuration ID must be a positive integer")
}

func TestUnitUpdateDiskEncryptionConfigurationByID_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateDiskEncryptionConfigurationByID(context.Background(), 1, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// UpdateDiskEncryptionConfigurationByName
// =============================================================================

func TestUnitUpdateDiskEncryptionConfigurationByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateDiskEncryptionConfigurationByNameMock()

	req := &RequestDiskEncryptionConfiguration{Name: "FileVault Config"}
	result, resp, err := svc.UpdateDiskEncryptionConfigurationByName(context.Background(), "FileVault Config", req)
	require.NoError(t, err)
	require.NotNil(t, result)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.ID)
}

func TestUnitUpdateDiskEncryptionConfigurationByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateDiskEncryptionConfigurationByName(context.Background(), "", &RequestDiskEncryptionConfiguration{Name: "x"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "disk encryption configuration name is required")
}

func TestUnitUpdateDiskEncryptionConfigurationByName_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.UpdateDiskEncryptionConfigurationByName(context.Background(), "FileVault Config", nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "request is required")
}

// =============================================================================
// DeleteDiskEncryptionConfigurationByID
// =============================================================================

func TestUnitDeleteDiskEncryptionConfigurationByID_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteDiskEncryptionConfigurationByIDMock()

	resp, err := svc.DeleteDiskEncryptionConfigurationByID(context.Background(), 1)
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteDiskEncryptionConfigurationByID_ZeroID(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.DeleteDiskEncryptionConfigurationByID(context.Background(), 0)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "disk encryption configuration ID must be a positive integer")
}

// =============================================================================
// DeleteDiskEncryptionConfigurationByName
// =============================================================================

func TestUnitDeleteDiskEncryptionConfigurationByName_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeleteDiskEncryptionConfigurationByNameMock()

	resp, err := svc.DeleteDiskEncryptionConfigurationByName(context.Background(), "FileVault Config")
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
}

func TestUnitDeleteDiskEncryptionConfigurationByName_EmptyName(t *testing.T) {
	svc, _ := setupMockService(t)
	_, err := svc.DeleteDiskEncryptionConfigurationByName(context.Background(), "")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "disk encryption configuration name is required")
}
