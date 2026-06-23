package mobile_device_commands_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_device_commands"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_device_commands/mocks"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/apilifecycle"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUnit_MobileDeviceCommands_SendCommand_RemovedOn1128(t *testing.T) {
	removed := []string{
		mobile_device_commands.CommandClearPasscode,
		mobile_device_commands.CommandBlankPush,
		mobile_device_commands.CommandRestartDevice,
		mobile_device_commands.CommandDisableLostMode,
	}
	for _, command := range removed {
		t.Run(command, func(t *testing.T) {
			mock := mocks.NewMobileDeviceCommandsMock()
			mock.ServerVersionStr = "11.28.1"
			svc := mobile_device_commands.NewMobileDeviceCommands(mock)

			_, err := svc.SendCommand(context.Background(), command, "1")
			require.Error(t, err)
			assert.True(t, apilifecycle.IsRemoved(err), "expected RemovedError for %s, got %v", command, err)
		})
	}
}

func TestUnit_MobileDeviceCommands_SendCommand_ClearPasscode_AllowedOn1127(t *testing.T) {
	mock := mocks.NewMobileDeviceCommandsMock()
	mock.ServerVersionStr = "11.27.0"
	mock.RegisterSendCommandMock(mobile_device_commands.CommandClearPasscode, "1")
	svc := mobile_device_commands.NewMobileDeviceCommands(mock)

	resp, err := svc.SendCommand(context.Background(), mobile_device_commands.CommandClearPasscode, "1")
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode())
}

func TestUnit_MobileDeviceCommands_SendCommand_NonGated_NotVersionChecked(t *testing.T) {
	mock := mocks.NewMobileDeviceCommandsMock()
	// ServerVersionStr left unset on purpose; a non-removed command must not
	// consult the removal guard.
	mock.RegisterSendCommandMock(mobile_device_commands.CommandEnableLostMode, "3")
	svc := mobile_device_commands.NewMobileDeviceCommands(mock)

	resp, err := svc.SendCommand(context.Background(), mobile_device_commands.CommandEnableLostMode, "3")
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode())
}

func TestUnit_MobileDeviceCommands_SendCommand_EmptyCommand(t *testing.T) {
	svc := mobile_device_commands.NewMobileDeviceCommands(mocks.NewMobileDeviceCommandsMock())
	_, err := svc.SendCommand(context.Background(), "", "1")
	require.Error(t, err)
}
