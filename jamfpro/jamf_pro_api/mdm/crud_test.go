package mdm

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/mdm/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupMockService(t *testing.T) (*Mdm, *mocks.MDMMock) {
	t.Helper()
	mock := mocks.NewMDMMock()
	return NewMdm(mock), mock
}

func TestUnit_Mdm_BlankPush_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterBlankPushMock()

	result, resp, err := svc.BlankPush(context.Background(), []string{"device-001", "device-002"})
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
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

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "cmd-12345", result.ID)
	assert.Equal(t, "/api/v2/mdm/commands/cmd-12345", result.Href)
}

// TestUnit_Mdm_SendCommand_TriggerEnhancedLogCollection_Success exercises the
// TRIGGER_ENHANCED_LOG_COLLECTION command type added in Jamf Pro 11.29, verifying
// the appleCareToken field is serialised on the wire.
func TestUnit_Mdm_SendCommand_TriggerEnhancedLogCollection_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterSendCommandMock()

	req := &CommandRequest{
		CommandData: CommandData{
			CommandType:    CommandTypeTriggerEnhancedLogCollection,
			AppleCareToken: "a1b2c3d4-e5f6-7890-abcd-ef1234567890",
		},
		ClientData: []ClientData{
			{ManagementID: "device-001"},
		},
	}

	// The appleCareToken must be marshalled into the request body.
	body, err := json.Marshal(req.CommandData)
	require.NoError(t, err)
	assert.Contains(t, string(body), `"commandType":"TRIGGER_ENHANCED_LOG_COLLECTION"`)
	assert.Contains(t, string(body), `"appleCareToken":"a1b2c3d4-e5f6-7890-abcd-ef1234567890"`)

	result, resp, err := svc.SendCommand(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "cmd-12345", result.ID)
	assert.Equal(t, "/api/v2/mdm/commands/cmd-12345", result.Href)
}

// TestUnit_Mdm_SendCommand_CancelEnhancedLogCollection_Success exercises the
// CANCEL_ENHANCED_LOG_COLLECTION command type added in Jamf Pro 11.29, which
// carries no command-type-specific fields.
func TestUnit_Mdm_SendCommand_CancelEnhancedLogCollection_Success(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterSendCommandMock()

	req := &CommandRequest{
		CommandData: CommandData{
			CommandType: CommandTypeCancelEnhancedLogCollection,
		},
		ClientData: []ClientData{
			{ManagementID: "device-001"},
		},
	}

	// No appleCareToken should be emitted for the cancel command.
	body, err := json.Marshal(req.CommandData)
	require.NoError(t, err)
	assert.Contains(t, string(body), `"commandType":"CANCEL_ENHANCED_LOG_COLLECTION"`)
	assert.NotContains(t, string(body), "appleCareToken")

	result, resp, err := svc.SendCommand(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, resp)

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, "cmd-12345", result.ID)
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
	assert.Equal(t, 404, resp.StatusCode())
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

	assert.Equal(t, 200, resp.StatusCode())
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

	assert.Equal(t, 200, resp.StatusCode())
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

	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, 2, result.TotalCount)
	require.Len(t, result.Results, 2)
	assert.Equal(t, "cmd-uuid-001", result.Results[0].UUID)
	assert.Equal(t, "DeviceLock", result.Results[0].CommandType)
	assert.Equal(t, "Completed", result.Results[0].Status)
}

func TestUnit_Mdm_ListCommandsV2_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListCommandsErrorMock()

	result, resp, err := svc.ListCommandsV2(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
	assert.Contains(t, err.Error(), "failed to list MDM commands")
}

func TestUnit_Mdm_ListCommandsV2_NoMock(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListCommandsNoResponseErrorMock()

	result, resp, err := svc.ListCommandsV2(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.NotNil(t, resp)
}

func TestUnit_Mdm_ListCommandsV2_InvalidJSON(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterListCommandsInvalidJSONMock()

	result, resp, err := svc.ListCommandsV2(context.Background(), nil)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Contains(t, err.Error(), "mergePage failed")
}

func TestUnit_Mdm_BlankPush_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterBlankPushErrorMock()

	result, resp, err := svc.BlankPush(context.Background(), []string{"device-001"})
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
}

func TestUnit_Mdm_DeployPackage_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterDeployPackageErrorMock()

	req := &DeployPackageRequest{
		Manifest:         PackageManifest{HashType: "SHA256", URL: "https://example.com/pkg.dmg", Hash: "abc", Title: "Test"},
		InstallAsManaged: true,
		Devices:          []int{1001},
		GroupID:          "g1",
	}
	result, resp, err := svc.DeployPackage(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
}

func TestUnit_Mdm_RenewProfile_Error(t *testing.T) {
	svc, mock := setupMockService(t)
	mock.RegisterRenewProfileErrorMock()

	req := &RenewProfileRequest{UDIDs: []string{"udid-001"}}
	result, resp, err := svc.RenewProfile(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, result)
	require.NotNil(t, resp)
	assert.Equal(t, 500, resp.StatusCode())
}

// -----------------------------------------------------------------------------
// Jamf Pro 11.30
// -----------------------------------------------------------------------------

func TestUnit_Mdm_ListCommandsV1_Success(t *testing.T) {
	mock := mocks.NewMDMMock()
	mock.RegisterListCommandsV1Mock()

	svc := NewMdm(mock)

	result, resp, err := svc.ListCommandsV1(context.Background(), map[string]string{
		"client-management-id": "bbbbbbbb-3f1e-4b3a-a5b3-ca0cd7430937",
	})

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode())
	require.Len(t, result, 2)

	assert.Equal(t, "aaaaaaaa-3f1e-4b3a-a5b3-ca0cd7430937", result[0].UUID)
	assert.Equal(t, MdmCommandStateAcknowledged, result[0].CommandState)
	assert.Equal(t, "2019-05-16T20:44:01.112Z", result[0].DateCompleted)
	require.NotNil(t, result[0].Client)
	assert.Equal(t, MdmClientTypeComputer, result[0].Client.ClientType)
	assert.Equal(t, 1, result[0].ProfileID)

	// 11.30 added COMMAND_FORMAT_ERROR to the command state enum.
	assert.Equal(t, MdmCommandStateCommandFormatError, result[1].CommandState)
	assert.Equal(t, CommandTypeTriggerEnhancedLogCollection, result[1].CommandType)
	require.NotNil(t, result[1].CommandError)
	assert.Equal(t, 1234, result[1].CommandError.ErrorCode)
}

func TestUnit_Mdm_ListCommandsV1_Error(t *testing.T) {
	mock := mocks.NewMDMMock()
	mock.RegisterListCommandsV1ErrorMock()

	svc := NewMdm(mock)

	result, _, err := svc.ListCommandsV1(context.Background(), nil)

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestUnit_Mdm_ListCommandsV1_NoMock(t *testing.T) {
	svc := NewMdm(mocks.NewMDMMock())

	_, _, err := svc.ListCommandsV1(context.Background(), nil)

	assert.Error(t, err)
}
