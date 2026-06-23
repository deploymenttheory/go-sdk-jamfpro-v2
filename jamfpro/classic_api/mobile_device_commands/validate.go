package mobile_device_commands

import "fmt"

// Mobile device command strings accepted by
// POST /JSSResource/mobiledevicecommands/command/{command}.
//
// ClearPasscode, BlankPush, RestartDevice and DisableLostMode were removed in
// Jamf Pro 11.28.0 (see SendCommand's removal guard). The remaining constants
// are common still-supported commands; the list is not exhaustive — any
// non-empty command string is accepted.
const (
	// CommandClearPasscode was deprecated 2025-04-03 and removed in Jamf Pro 11.28.0.
	CommandClearPasscode = "ClearPasscode"
	// CommandBlankPush was deprecated 2025-05-27 and removed in Jamf Pro 11.28.0.
	CommandBlankPush = "BlankPush"
	// CommandRestartDevice was deprecated 2025-02-25 and removed in Jamf Pro 11.28.0.
	CommandRestartDevice = "RestartDevice"
	// CommandDisableLostMode was deprecated 2025-05-30 and removed in Jamf Pro 11.28.0.
	CommandDisableLostMode = "DisableLostMode"

	CommandDeviceLock      = "DeviceLock"
	CommandEraseDevice     = "EraseDevice"
	CommandEnableLostMode  = "EnableLostMode"
	CommandUpdateInventory = "UpdateInventory"
	CommandUnmanageDevice  = "UnmanageDevice"
	CommandShutDownDevice  = "ShutDownDevice"
)

// validateMobileDeviceCommand performs the minimal client-side validation shared
// by all mobile device command sends.
func validateMobileDeviceCommand(command string) error {
	if command == "" {
		return fmt.Errorf("command is required")
	}
	return nil
}
