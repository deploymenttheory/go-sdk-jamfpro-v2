package computer_commands_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/computer_commands"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/computer_commands/mocks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/apilifecycle"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_ComputerCommands_SendCommand_BlankPush_RemovedOn1128(t *testing.T) {
	mock := mocks.NewComputerCommandsMock()
	mock.ServerVersionStr = "11.28.0"
	svc := computer_commands.NewComputerCommands(mock)

	resp, err := svc.SendCommand(context.Background(), computer_commands.CommandBlankPush, "1")
	require.Error(t, err)
	assert.Nil(t, resp)
	assert.True(t, apilifecycle.IsRemoved(err), "expected a RemovedError, got %v", err)
}

func TestUnit_ComputerCommands_SendCommand_DeleteUser_RemovedOn1128(t *testing.T) {
	mock := mocks.NewComputerCommandsMock()
	mock.ServerVersionStr = "11.28.1"
	svc := computer_commands.NewComputerCommands(mock)

	_, err := svc.SendCommand(context.Background(), computer_commands.CommandDeleteUser, "1", "2")
	require.Error(t, err)
	assert.True(t, apilifecycle.IsRemoved(err))
}

func TestUnit_ComputerCommands_SendCommand_BlankPush_AllowedOn1127(t *testing.T) {
	mock := mocks.NewComputerCommandsMock()
	mock.ServerVersionStr = "11.27.9"
	mock.RegisterSendCommandMock(computer_commands.CommandBlankPush, "1")
	svc := computer_commands.NewComputerCommands(mock)

	resp, err := svc.SendCommand(context.Background(), computer_commands.CommandBlankPush, "1")
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode())
}

func TestUnit_ComputerCommands_SendCommand_NonGated_NotVersionChecked(t *testing.T) {
	mock := mocks.NewComputerCommandsMock()
	// ServerVersionStr deliberately left unset: if the guard were consulted for a
	// non-removed command, ServerVersion would error and the call would fail open
	// but still log. The command must succeed regardless.
	mock.RegisterSendCommandMock(computer_commands.CommandDeviceLock, "5")
	svc := computer_commands.NewComputerCommands(mock)

	resp, err := svc.SendCommand(context.Background(), computer_commands.CommandDeviceLock, "5")
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode())
}

func TestUnit_ComputerCommands_SendCommand_EmptyCommand(t *testing.T) {
	svc := computer_commands.NewComputerCommands(mocks.NewComputerCommandsMock())
	_, err := svc.SendCommand(context.Background(), "", "1")
	require.Error(t, err)
}

func TestUnit_ComputerCommands_SendCommand_NoIDs(t *testing.T) {
	svc := computer_commands.NewComputerCommands(mocks.NewComputerCommandsMock())
	_, err := svc.SendCommand(context.Background(), computer_commands.CommandDeviceLock)
	require.Error(t, err)
}
