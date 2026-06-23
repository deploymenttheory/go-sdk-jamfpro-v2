package classic_api

import (
	"context"
	"testing"

	acc "github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/acceptance"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/computer_commands"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/apilifecycle"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestAcceptance_ComputerCommands_RemovedCommandsGated verifies that the
// removed BlankPush / DeleteUser computer commands are blocked client-side by
// the API-lifecycle removal guard when the connected server is Jamf Pro
// 11.28.0 or newer. No MDM command is ever sent to a real device.
func TestAcceptance_ComputerCommands_RemovedCommandsGated(t *testing.T) {
	acc.RequireClient(t)
	ctx := context.Background()

	server := serverVersionOrSkip(t, ctx)
	if !server.AtLeast(apilifecycle.MustParse("11.28.0")) {
		t.Skipf("server %s predates the 11.28 removals; nothing to assert", server)
	}

	svc := acc.Client.ClassicAPI.ComputerCommands
	for _, command := range []string{
		computer_commands.CommandBlankPush,
		computer_commands.CommandDeleteUser,
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
