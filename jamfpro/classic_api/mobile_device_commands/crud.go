package mobile_device_commands

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
	// MobileDeviceCommands handles the Classic API mobile device MDM command endpoints.
	//
	// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmobiledevicecommand
	MobileDeviceCommands struct {
		client client.Client
	}
)

// NewMobileDeviceCommands returns a new MobileDeviceCommands service backed by the provided HTTP client.
func NewMobileDeviceCommands(client client.Client) *MobileDeviceCommands {
	return &MobileDeviceCommands{client: client}
}

// removedIn1128 is the Jamf Pro version that removed the ClearPasscode,
// BlankPush, RestartDevice and DisableLostMode mobile device commands from the
// Classic API.
var removedIn1128 = apilifecycle.MustParse("11.28.0")

// -----------------------------------------------------------------------------
// Classic API - Mobile Device Command Operations
// -----------------------------------------------------------------------------

// SendCommand sends an MDM command to one or more mobile devices by ID.
// URL: POST /JSSResource/mobiledevicecommands/command/{command}/id/{ids}
//
// The ClearPasscode, BlankPush, RestartDevice and DisableLostMode commands were
// removed in Jamf Pro 11.28.0; when the connected server is at or above that
// version this method returns an *apilifecycle.RemovedError without contacting
// the API. Detect it with apilifecycle.IsRemoved(err).
//
// Classic API docs: https://developer.jamf.com/jamf-pro/reference/createmobiledevicecommandbycommandandid
func (s *MobileDeviceCommands) SendCommand(ctx context.Context, command string, ids ...string) (*resty.Response, error) {
	if err := validateMobileDeviceCommand(command); err != nil {
		return nil, err
	}
	if len(ids) == 0 {
		return nil, fmt.Errorf("at least one mobile device id is required")
	}

	// Removals guard: these commands were removed in Jamf Pro 11.28.0.
	switch command {
	case CommandClearPasscode, CommandBlankPush, CommandRestartDevice, CommandDisableLostMode:
		label := "classic_api/mobile_device_commands.MobileDeviceCommands.SendCommand:" + command
		if err := apilifecycle.EnsureSupported(ctx, s.client, label, removedIn1128); err != nil {
			return nil, err
		}
	}

	endpoint := fmt.Sprintf("%s/command/%s/id/%s",
		constants.EndpointClassicMobileDeviceCommands, command, strings.Join(ids, ","))

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationXML).
		Post(endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to send mobile device command %s: %w", command, err)
	}

	return resp, nil
}
