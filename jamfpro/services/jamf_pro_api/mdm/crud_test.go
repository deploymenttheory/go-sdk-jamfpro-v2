package mdm

import (
	"context"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/mdm/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Service, *mocks.MDMMock) {
	t.Helper()
	mock := mocks.NewMDMMock()
	return NewService(mock), mock
}

func TestUnit_Mdm_BlankPush_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterBlankPushMock()

	result, resp, err := svc.BlankPush(context.Background(), []string{"device-001", "device-002"})
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, result.ErrorUUIDs)
	assert.Empty(t, result.ErrorUUIDs)
}

func TestUnit_Mdm_BlankPush_EmptyClientManagementIDs(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.BlankPush(context.Background(), []string{})
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "clientManagementIDs is required")
}

func TestUnit_Mdm_BlankPush_NilClientManagementIDs(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.BlankPush(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "clientManagementIDs is required")
}

func TestUnit_Mdm_SendCommand_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterSendCommandMock()

	req := &CommandRequest{
		CommandData: CommandData{
			CommandType: "DeviceLock",
		},
		ClientData: []ClientData{
			{ManagementID: "device-001"},
		},
	}
	result, resp, err := svc.SendCommand(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "cmd-12345", result.ID)
	assert.Equal(t, "/api/v2/mdm/commands/cmd-12345", result.Href)
}

func TestUnit_Mdm_SendCommand_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.SendCommand(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Mdm_SendCommand_NotFound(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterNotFoundErrorMock()

	req := &CommandRequest{
		CommandData: CommandData{CommandType: "DeviceLock"},
		ClientData:  []ClientData{{ManagementID: "device-999"}},
	}
	result, resp, err := svc.SendCommand(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 404, resp.StatusCode)
}

func TestUnit_Mdm_DeployPackage_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeployPackageMock()

	req := &DeployPackageRequest{
		Manifest: PackageManifest{
			HashType: "SHA256",
			URL:      "https://example.com/pkg.dmg",
			Hash:     "abc123",
			Title:    "Test Package",
		},
		InstallAsManaged: true,
		Devices:          []int{1001, 1002},
		GroupID:          "group-1",
	}
	result, resp, err := svc.DeployPackage(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	require.Len(t, result.QueuedCommands, 2)
	assert.Equal(t, 1001, result.QueuedCommands[0].Device)
	assert.Equal(t, "uuid-abc-123", result.QueuedCommands[0].CommandUUID)
	assert.Equal(t, 1002, result.QueuedCommands[1].Device)
	assert.Equal(t, "uuid-def-456", result.QueuedCommands[1].CommandUUID)
	assert.Empty(t, result.Errors)
}

func TestUnit_Mdm_DeployPackage_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.DeployPackage(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_Mdm_RenewProfile_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterRenewProfileMock()

	req := &RenewProfileRequest{
		UDIDs: []string{"udid-001", "udid-002"},
	}
	result, resp, err := svc.RenewProfile(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.NotNil(t, result.UDIDsNotProcessed.UDIDs)
	assert.Empty(t, result.UDIDsNotProcessed.UDIDs)
}

func TestUnit_Mdm_RenewProfile_NilRequest(t *testing.T) {
	svc, _ := setupMockService(t)

	result, resp, err := svc.RenewProfile(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Nil(t, resp)
	assert.Contains(t, err.Error(), "request is required")
}

func TestUnit_MDM_ListCommandsV2_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListCommandsMock()

	result, resp, err := svc.ListCommandsV2(context.Background(), nil)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "cmd-uuid-001", result.Results[0].UUID)
	assert.Equal(t, "DeviceLock", result.Results[0].CommandType)
	assert.Equal(t, "Completed", result.Results[0].Status)
}
