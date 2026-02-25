package computer_inventory

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_inventory/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_ComputerInventory_ListV3(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterListMock()

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.ListV3(ctx, nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	assert.Len(t, result.Results, 2)
	assert.Equal(t, "Test-Mac-001", result.Results[0].General.Name)
	assert.Equal(t, "C02ABC123DEF", result.Results[0].Hardware.SerialNumber)
}

func TestUnit_ComputerInventory_GetByIDV3(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterGetByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByIDV3(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Test-Mac-001", result.General.Name)
	assert.Equal(t, "C02ABC123DEF", result.Hardware.SerialNumber)
	assert.Equal(t, "MacBook Pro", result.Hardware.Model)
	assert.Equal(t, "testuser", result.UserAndLocation.Username)
}

func TestUnit_ComputerInventory_GetByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetByIDV3(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_ComputerInventory_UpdateByIDV3(t *testing.T) {
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

	result, resp, err := svc.UpdateByIDV3(ctx, "1", request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, result)
}

func TestUnit_ComputerInventory_UpdateByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	request := &ResourceComputerInventory{}

	result, resp, err := svc.UpdateByIDV3(ctx, "", request)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_ComputerInventory_UpdateByIDV3_NilRequest(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.UpdateByIDV3(ctx, "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_ComputerInventory_DeleteByIDV3(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterDeleteByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteByIDV3(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnit_ComputerInventory_DeleteByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteByIDV3(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_ComputerInventory_ListFileVaultV3(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterListFileVaultMock()

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.ListFileVaultV3(ctx)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 1, result.TotalCount)
	assert.Len(t, result.Results, 1)
	assert.Equal(t, "1", result.Results[0].ComputerId)
	assert.Equal(t, "VALID", result.Results[0].IndividualRecoveryKeyValidityStatus)
}

func TestUnit_ComputerInventory_GetFileVaultByIDV3(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterGetFileVaultByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetFileVaultByIDV3(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ComputerId)
	assert.Equal(t, "Test-Mac-001", result.Name)
	assert.NotEmpty(t, result.PersonalRecoveryKey)
	assert.Equal(t, "ENCRYPTED", result.BootPartitionEncryptionDetails.PartitionFileVault2State)
}

func TestUnit_ComputerInventory_GetFileVaultByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetFileVaultByIDV3(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_ComputerInventory_GetRecoveryLockPasswordByIDV3(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterGetRecoveryLockPasswordByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetRecoveryLockPasswordByIDV3(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, result.RecoveryLockPassword)
}

func TestUnit_ComputerInventory_GetRecoveryLockPasswordByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetRecoveryLockPasswordByIDV3(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_ComputerInventory_DeleteAttachmentByIDV3(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterDeleteAttachmentMock("1", "100")

	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteAttachmentByIDV3(ctx, "1", "100")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 204, resp.StatusCode)
}

func TestUnit_ComputerInventory_DeleteAttachmentByIDV3_EmptyComputerID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteAttachmentByIDV3(ctx, "", "100")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "computerID is required")
}

func TestUnit_ComputerInventory_DeleteAttachmentByIDV3_EmptyAttachmentID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.DeleteAttachmentByIDV3(ctx, "1", "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "attachmentID is required")
}

func TestUnit_ComputerInventory_RemoveMDMProfileByIDV1(t *testing.T) {
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

func TestUnit_ComputerInventory_RemoveMDMProfileByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.RemoveMDMProfileByIDV1(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_ComputerInventory_EraseByIDV1(t *testing.T) {
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

func TestUnit_ComputerInventory_EraseByIDV1_EmptyID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	request := &RequestEraseDeviceComputer{}

	resp, err := svc.EraseByIDV1(ctx, "", request)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_ComputerInventory_EraseByIDV1_NilRequest(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.EraseByIDV1(ctx, "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_ComputerInventory_CreateV3(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterCreateMock()

	svc := NewService(mock)
	ctx := context.Background()

	request := &ResourceComputerInventory{
		General: ComputerInventorySubsetGeneral{
			Name: "Test-Mac-001",
		},
		Hardware: ComputerInventorySubsetHardware{
			SerialNumber: "C02ABC123DEF",
		},
	}

	result, resp, err := svc.CreateV3(ctx, request)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "/api/v3/computers-inventory/1", result.HREF)
}

func TestUnit_ComputerInventory_CreateV3_NilRequest(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.CreateV3(ctx, nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_ComputerInventory_GetDetailByIDV3(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterGetDetailByIDMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDetailByIDV3(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Test-Mac-001", result.General.Name)
}

func TestUnit_ComputerInventory_GetDetailByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDetailByIDV3(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_ComputerInventory_UploadAttachmentByIDV3(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterUploadAttachmentMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	attachment := []byte("test attachment data")

	resp, err := svc.UploadAttachmentByIDV3(ctx, "1", attachment)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 201, resp.StatusCode)
}

func TestUnit_ComputerInventory_UploadAttachmentByIDV3_EmptyComputerID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	attachment := []byte("test attachment data")

	resp, err := svc.UploadAttachmentByIDV3(ctx, "", attachment)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "computerID is required")
}

func TestUnit_ComputerInventory_UploadAttachmentByIDV3_EmptyAttachment(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	resp, err := svc.UploadAttachmentByIDV3(ctx, "1", []byte{})

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "attachment data is required")
}

func TestUnit_ComputerInventory_GetAttachmentByIDV3(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterGetAttachmentMock("1", "100")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetAttachmentByIDV3(ctx, "1", "100")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, result)
}

func TestUnit_ComputerInventory_GetAttachmentByIDV3_EmptyComputerID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetAttachmentByIDV3(ctx, "", "100")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "computerID is required")
}

func TestUnit_ComputerInventory_GetAttachmentByIDV3_EmptyAttachmentID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetAttachmentByIDV3(ctx, "1", "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "attachmentID is required")
}

func TestUnit_ComputerInventory_GetDeviceLockPinByIDV3(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterGetDeviceLockPinMock("1")

	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDeviceLockPinByIDV3(ctx, "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, result.Pin)
	assert.Equal(t, "123456", result.Pin)
}

func TestUnit_ComputerInventory_GetDeviceLockPinByIDV3_EmptyID(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	svc := NewService(mock)
	ctx := context.Background()

	result, resp, err := svc.GetDeviceLockPinByIDV3(ctx, "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}
