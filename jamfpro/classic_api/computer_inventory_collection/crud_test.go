package computer_inventory_collection_test

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/computer_inventory_collection"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/computer_inventory_collection/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupMockService creates a Service wired to a fresh ComputerInventoryCollectionMock.
func setupMockService(t *testing.T) (*computer_inventory_collection.ComputerInventoryCollection, *mocks.ComputerInventoryCollectionMock) {
	t.Helper()
	mock := mocks.NewComputerInventoryCollectionMock()
	return computer_inventory_collection.NewComputerInventoryCollection(mock), mock
}

// =============================================================================
// Get
// =============================================================================

func TestUnit_ComputerInventoryCollection_Get_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterGetMock()

	result, resp, err := svc.Get(context.Background())
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.True(t, result.LocalUserAccounts)
	assert.True(t, result.HiddenAccounts)
	assert.True(t, result.Printers)
	assert.True(t, result.InclueApplications)
	assert.False(t, result.InclueFonts)
	assert.False(t, result.IncluePlugins)
	assert.Len(t, result.Applications, 2)
	assert.Equal(t, "/Applications/Safari.app", result.Applications[0].Path)
	assert.Equal(t, "Mac", result.Applications[0].Platform)
}

// =============================================================================
// Update
// =============================================================================

func TestUnit_ComputerInventoryCollection_Update_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterUpdateMock()

	settings := &computer_inventory_collection.ResourceComputerInventoryCollection{
		LocalUserAccounts:             true,
		HomeDirectorySizes:           true,
		HiddenAccounts:               true,
		Printers:                     true,
		ActiveServices:               true,
		ComputerLocationInformation:  true,
		PackageReceipts:              true,
		AvailableSoftwareUpdates:     true,
		InclueApplications:           true,
		InclueFonts:                  true,
		IncluePlugins:                true,
	}

	resp, err := svc.Update(context.Background(), settings)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
}

func TestUnit_ComputerInventoryCollection_Update_NilSettings(t *testing.T) {
	svc, _ := setupMockService(t)

	resp, err := svc.Update(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "settings is required")
}

func TestUnit_ComputerInventoryCollection_Get_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	_, _, err := svc.Get(context.Background())
	require.Error(t, err)
}

func TestUnit_ComputerInventoryCollection_Update_Error(t *testing.T) {
	svc, _ := setupMockService(t)
	settings := &computer_inventory_collection.ResourceComputerInventoryCollection{LocalUserAccounts: true}
	_, err := svc.Update(context.Background(), settings)
	require.Error(t, err)
}
