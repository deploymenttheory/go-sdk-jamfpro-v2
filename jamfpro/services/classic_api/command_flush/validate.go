package command_flush

import (
	"fmt"
	"slices"
	"strings"
)

// validateIDType validates the device ID type.
func validateIDType(idType string) error {
	validTypes := []string{"computers", "computergroups", "mobiledevices", "mobiledevicegroups"}
	if !slices.Contains(validTypes, idType) {
		return fmt.Errorf("invalid idType: %s (must be one of: %s)", idType, strings.Join(validTypes, ", "))
	}
	return nil
}

// validateStatus validates the command status.
func validateStatus(status string) error {
	validStatuses := []string{"Pending", "Failed", "Pending+Failed"}
	if !slices.Contains(validStatuses, status) {
		return fmt.Errorf("invalid status: %s (must be one of: %s)", status, strings.Join(validStatuses, ", "))
	}
	return nil
}

// validateCommandFlushRequest validates the XML-based command flush request.
func validateCommandFlushRequest(req *RequestCommandFlush) error {
	if req == nil {
		return fmt.Errorf("request cannot be nil")
	}

	if err := validateStatus(req.Status); err != nil {
		return err
	}

	if req.MobileDevices == nil && req.Computers == nil {
		return fmt.Errorf("at least one device list (mobile_devices or computers) must be provided")
	}

	return nil
}
