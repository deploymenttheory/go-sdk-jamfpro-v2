package computer_commands

import "fmt"

// Computer command strings accepted by POST /JSSResource/computercommands/command/{command}.
//
// BlankPush and DeleteUser were removed in Jamf Pro 11.28.0 (see SendCommand's
// removal guard). The remaining constants are common still-supported commands;
// the list is not exhaustive — any non-empty command string is accepted.
const (
	// CommandBlankPush was deprecated 2025-05-27 and removed in Jamf Pro 11.28.0.
	CommandBlankPush = "BlankPush"
	// CommandDeleteUser was deprecated 2025-02-25 and removed in Jamf Pro 11.28.0.
	CommandDeleteUser = "DeleteUser"

	CommandDeviceLock           = "DeviceLock"
	CommandEraseDevice          = "EraseDevice"
	CommandUnmanageDevice       = "UnmanageDevice"
	CommandScheduleOSUpdate     = "ScheduleOSUpdate"
	CommandEnableRemoteDesktop  = "EnableRemoteDesktop"
	CommandDisableRemoteDesktop = "DisableRemoteDesktop"
)

// validateComputerCommand performs the minimal client-side validation shared by
// all computer command sends.
func validateComputerCommand(command string) error {
	if command == "" {
		return fmt.Errorf("command is required")
	}
	return nil
}
