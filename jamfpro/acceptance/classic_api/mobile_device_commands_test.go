package classic_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_device_commands"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/apilifecycle"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAcceptance_MobileDeviceCommands_RemovedCommandsGated verifies that the
// removed ClearPasscode / BlankPush / RestartDevice / DisableLostMode mobile
// device commands are blocked client-side by the API-lifecycle removal guard
// when the connected server is Jamf Pro 11.28.0 or newer. No MDM command is
// ever sent to a real device.
func TestAcceptance_MobileDeviceCommands_RemovedCommandsGated(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	server := serverVersionOrSkip(t, ctx)
	if !server.AtLeast(apilifecycle.MustParse("11.28.0")) {
		t.Skipf("server %s predates the 11.28 removals; nothing to assert", server)
	}

	svc := acc.Client.ClassicAPI.MobileDeviceCommands
	for _, command := range []string{
		mobile_device_commands.CommandClearPasscode,
		mobile_device_commands.CommandBlankPush,
		mobile_device_commands.CommandRestartDevice,
		mobile_device_commands.CommandDisableLostMode,
	} {
		t.Run(command, func(t *testing.T) {
			acc.LogTestStage(t, "Removed", "Expecting removal guard to block %s", command)
			_, err := svc.SendCommand(ctx, command, "1")
			require.Error(t, err)
			assert.True(t, apilifecycle.IsRemoved(err),
				"expected RemovedError for %s on server %s, got %v", command, server, err)
		})
	}
}
