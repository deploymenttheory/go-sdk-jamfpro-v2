package computer_inventory

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/computer_inventory/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_ComputerInventory_ListV4(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterListV4Mock()

	svc := NewComputerInventory(mock)

	result, resp, err := svc.ListV4(context.Background(), nil)

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "Test-Mac-001", result.Results[0].General.Name)
	assert.Equal(t, "C02ABC123DEF", result.Results[0].Hardware.SerialNumber)
	// 11.30: lastContactTime renamed to lastContact, lastCheckIn added.
	assert.Equal(t, "2018-10-31T18:04:13Z", result.Results[0].General.LastContact)
	assert.Equal(t, "2018-10-31T19:12:44Z", result.Results[0].General.LastCheckIn)
}

func TestUnit_ComputerInventory_ListV4_Error(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterListV4ErrorMock()

	svc := NewComputerInventory(mock)

	result, _, err := svc.ListV4(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUnit_ComputerInventory_ListV4_InvalidJSON(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterListV4InvalidJSONMock()

	svc := NewComputerInventory(mock)

	result, _, err := svc.ListV4(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUnit_ComputerInventory_GetByIDV4(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterGetByIDV4Mock("1")

	svc := NewComputerInventory(mock)

	result, resp, err := svc.GetByIDV4(context.Background(), "1")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
	assert.Equal(t, "Test-Mac-001", result.General.Name)
	assert.Equal(t, "2018-10-31T18:04:13Z", result.General.LastContact)
	assert.Equal(t, "2018-10-31T19:12:44Z", result.General.LastCheckIn)
	assert.Equal(t, "247.185.82.186", result.General.LastReportedIpV4)
	assert.Equal(t, "testuser", result.UserAndLocation.Username)
}

func TestUnit_ComputerInventory_GetByIDV4_EmptyID(t *testing.T) {
	svc := NewComputerInventory(mocks.NewComputerInventoryMock())

	result, resp, err := svc.GetByIDV4(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_ComputerInventory_GetByIDV4_Error(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterGetByIDV4ErrorMock("1")

	svc := NewComputerInventory(mock)

	result, _, err := svc.GetByIDV4(context.Background(), "1")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUnit_ComputerInventory_GetByIDV4_NoMock(t *testing.T) {
	svc := NewComputerInventory(mocks.NewComputerInventoryMock())

	_, _, err := svc.GetByIDV4(context.Background(), "1")

	assert.Error(t, err)
}

func TestUnit_ComputerInventory_GetDetailByIDV4(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterGetDetailByIDV4Mock("1")

	svc := NewComputerInventory(mock)

	result, resp, err := svc.GetDetailByIDV4(context.Background(), "1")

	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "1", result.ID)
}

func TestUnit_ComputerInventory_CreateV4(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterCreateV4Mock()

	svc := NewComputerInventory(mock)

	request := &ResourceComputerInventoryV4{
		General: ComputerInventorySubsetGeneralV4{Name: "New-Mac"},
	}

	result, resp, err := svc.CreateV4(context.Background(), request)

	require.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode())
	assert.NotNil(t, result)
}

func TestUnit_ComputerInventory_CreateV4_NilRequest(t *testing.T) {
	svc := NewComputerInventory(mocks.NewComputerInventoryMock())

	result, resp, err := svc.CreateV4(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_ComputerInventory_UpdateByIDV4(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterUpdateByIDV4Mock("1")

	svc := NewComputerInventory(mock)

	request := &ResourceComputerInventoryV4{
		UserAndLocation: ComputerInventorySubsetUserAndLocation{Username: "updateduser"},
	}

	result, resp, err := svc.UpdateByIDV4(context.Background(), "1", request)

	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotNil(t, result)
}

func TestUnit_ComputerInventory_UpdateByIDV4_EmptyID(t *testing.T) {
	svc := NewComputerInventory(mocks.NewComputerInventoryMock())

	_, _, err := svc.UpdateByIDV4(context.Background(), "", &ResourceComputerInventoryV4{})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_ComputerInventory_DeleteByIDV4(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterDeleteByIDV4Mock("1")

	svc := NewComputerInventory(mock)

	resp, err := svc.DeleteByIDV4(context.Background(), "1")

	require.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_ComputerInventory_DeleteByIDV4_EmptyID(t *testing.T) {
	svc := NewComputerInventory(mocks.NewComputerInventoryMock())

	resp, err := svc.DeleteByIDV4(context.Background(), "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "id is required")
}

func TestUnit_ComputerInventory_ListFileVaultV4(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterListFileVaultV4Mock()

	svc := NewComputerInventory(mock)

	result, resp, err := svc.ListFileVaultV4(context.Background())

	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotEmpty(t, result.Results)
}

func TestUnit_ComputerInventory_GetFileVaultByIDV4(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterGetFileVaultByIDV4Mock("1")

	svc := NewComputerInventory(mock)

	result, resp, err := svc.GetFileVaultByIDV4(context.Background(), "1")

	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotNil(t, result)
}

func TestUnit_ComputerInventory_GetDeviceLockPinByIDV4(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterGetDeviceLockPinV4Mock("1")

	svc := NewComputerInventory(mock)

	result, resp, err := svc.GetDeviceLockPinByIDV4(context.Background(), "1")

	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotNil(t, result)
}

func TestUnit_ComputerInventory_GetRecoveryLockPasswordByIDV4(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterGetRecoveryLockPasswordByIDV4Mock("1")

	svc := NewComputerInventory(mock)

	result, resp, err := svc.GetRecoveryLockPasswordByIDV4(context.Background(), "1")

	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotNil(t, result)
}

func TestUnit_ComputerInventory_UploadAttachmentByIDV4(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterUploadAttachmentV4Mock("1")

	svc := NewComputerInventory(mock)

	resp, err := svc.UploadAttachmentByIDV4(context.Background(), "1", []byte("file-bytes"))

	require.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode())
}

func TestUnit_ComputerInventory_UploadAttachmentByIDV4_EmptyAttachment(t *testing.T) {
	svc := NewComputerInventory(mocks.NewComputerInventoryMock())

	resp, err := svc.UploadAttachmentByIDV4(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "attachment data is required")
}

func TestUnit_ComputerInventory_DeleteAttachmentByIDV4(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterDeleteAttachmentV4Mock("1", "2")

	svc := NewComputerInventory(mock)

	resp, err := svc.DeleteAttachmentByIDV4(context.Background(), "1", "2")

	require.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_ComputerInventory_DeleteAttachmentByIDV4_EmptyAttachmentID(t *testing.T) {
	svc := NewComputerInventory(mocks.NewComputerInventoryMock())

	resp, err := svc.DeleteAttachmentByIDV4(context.Background(), "1", "")

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "attachmentID is required")
}

// RemoveMDMProfileByIDV4 and EraseByIDV4 moved from the singular
// /api/v1/computer-inventory path onto the v4 base path in Jamf Pro 11.30.

func TestUnit_ComputerInventory_RemoveMDMProfileByIDV4(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterRemoveMDMProfileV4Mock("1")

	svc := NewComputerInventory(mock)

	result, resp, err := svc.RemoveMDMProfileByIDV4(context.Background(), "1")

	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.NotNil(t, result)
}

func TestUnit_ComputerInventory_EraseByIDV4(t *testing.T) {
	mock := mocks.NewComputerInventoryMock()
	mock.RegisterEraseV4Mock("1")

	svc := NewComputerInventory(mock)

	resp, err := svc.EraseByIDV4(context.Background(), "1", &RequestEraseDeviceComputer{})

	require.NoError(t, err)
	assert.Equal(t, 204, resp.StatusCode())
}

func TestUnit_ComputerInventory_EraseByIDV4_NilRequest(t *testing.T) {
	svc := NewComputerInventory(mocks.NewComputerInventoryMock())

	resp, err := svc.EraseByIDV4(context.Background(), "1", nil)

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}
