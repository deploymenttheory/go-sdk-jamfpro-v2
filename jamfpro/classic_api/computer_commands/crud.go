package computer_commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/apilifecycle"
	"resty.dev/v3"
)

type (
	// ComputerCommands handles the Classic API computer MDM command endpoints.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createcomputercommandbycommand
	ComputerCommands struct {
		client client.Client
	}
)

// NewComputerCommands returns a new ComputerCommands service backed by the provided HTTP client.
func NewComputerCommands(client client.Client) *ComputerCommands {
	return &ComputerCommands{client: client}
}

// removedIn1128 is the Jamf Pro version that removed the BlankPush and
// DeleteUser computer commands from the Classic API.
var removedIn1128 = apilifecycle.MustParse("11.28.0")

// -----------------------------------------------------------------------------
// Classic API - Computer Command Operations
// -----------------------------------------------------------------------------

// SendCommand sends an MDM command to one or more computers by ID.
// URL: POST /JSSResource/computercommands/command/{command}/id/{ids}
//
// The BlankPush and DeleteUser commands were removed in Jamf Pro 11.28.0; when
// the connected server is at or above that version this method returns an
// *apilifecycle.RemovedError without contacting the API. Detect it with
// apilifecycle.IsRemoved(err).
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createcomputercommandbycommandandid
func (s *ComputerCommands) SendCommand(ctx context.Context, command string, ids ...string) (*resty.Response, error) {
	if err := validateComputerCommand(command); err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return nil, fmt.Errorf("at least one computer id is required")
	}

	// Removals guard: BlankPush and DeleteUser were removed in Jamf Pro 11.28.0.
	if command == CommandBlankPush || command == CommandDeleteUser {
		label := "classic_api/computer_commands.ComputerCommands.SendCommand:" + command
		if err := apilifecycle.EnsureSupported(ctx, s.client, label, removedIn1128); err != nil {
			return nil, err
		}
	}

	endpoint := fmt.Sprintf("%s/command/%s/id/%s",
		constants.EndpointClassicComputerCommands, command, strings.Join(ids, ","))

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Post(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to send computer command %s: %w", command, err)
	}

	return resp, nil
}
