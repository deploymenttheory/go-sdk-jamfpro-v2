package computer_inventory

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_inventory/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListV1(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.ListV1(ctx, nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	assert.Len(t, result.Results, 2)
	assert.Equal(t, "Test-Mac-001", result.Results[0].General.Name)
	assert.Equal(t, "C02ABC123DEF", result.Results[0].Hardware.SerialNumber)
}

func TestGetByIDV1(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterGetByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByIDV1(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Test-Mac-001", result.General.Name)
	assert.Equal(t, "C02ABC123DEF", result.Hardware.SerialNumber)
	assert.Equal(t, "MacBook Pro", result.Hardware.Model)
	assert.Equal(t, "testuser", result.UserAndLocation.Username)
}

func TestGetByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByIDV1(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUpdateByIDV1(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterUpdateByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	request := &ResourceComputerInventory{
		UserAndLocation: ComputerInventorySubsetUserAndLocation{
			Username: "updateduser",
			Email:    "updated@example.com",
		},
	}

	result, resp, err := svc.UpdateByIDV1(ctx, "1", request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, result)
}

func TestUpdateByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	request := &ResourceComputerInventory{}

	result, resp, err := svc.UpdateByIDV1(ctx, "", request)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUpdateByIDV1_NilRequest(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.UpdateByIDV1(ctx, "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestDeleteByIDV1(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterDeleteByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteByIDV1(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestDeleteByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteByIDV1(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestListFileVaultV1(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterListFileVaultMock()

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.ListFileVaultV1(ctx, nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ComputerId)
	assert.Equal(t, "VALID", result.Results[0].IndividualRecoveryKeyValidityStatus)
}

func TestGetFileVaultByIDV1(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterGetFileVaultByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetFileVaultByIDV1(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ComputerId)
	assert.Equal(t, "Test-Mac-001", result.Name)
	assert.NotEmpty(t, result.PersonalRecoveryKey)
	assert.Equal(t, "ENCRYPTED", result.BootPartitionEncryptionDetails.PartitionFileVault2State)
}

func TestGetFileVaultByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetFileVaultByIDV1(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestGetRecoveryLockPasswordByIDV1(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterGetRecoveryLockPasswordByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetRecoveryLockPasswordByIDV1(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, result.RecoveryLockPassword)
}

func TestGetRecoveryLockPasswordByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetRecoveryLockPasswordByIDV1(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestDeleteAttachmentByIDV1(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterDeleteAttachmentMock("1", "100")

	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteAttachmentByIDV1(ctx, "1", "100")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestDeleteAttachmentByIDV1_EmptyComputerID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteAttachmentByIDV1(ctx, "", "100")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "computerID is required")
}

func TestDeleteAttachmentByIDV1_EmptyAttachmentID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteAttachmentByIDV1(ctx, "1", "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "attachmentID is required")
}

func TestRemoveMDMProfileByIDV1(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterRemoveMDMProfileMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.RemoveMDMProfileByIDV1(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.DeviceID)
	assert.NotEmpty(t, result.CommandUUID)
}

func TestRemoveMDMProfileByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.RemoveMDMProfileByIDV1(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestEraseByIDV1(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterEraseMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	pin := "123456"
	request := &RequestEraseDeviceComputer{
		Pin: &pin,
	}

	resp, err := svc.EraseByIDV1(ctx, "1", request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestEraseByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	request := &RequestEraseDeviceComputer{}

	resp, err := svc.EraseByIDV1(ctx, "", request)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestEraseByIDV1_NilRequest(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.EraseByIDV1(ctx, "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}
