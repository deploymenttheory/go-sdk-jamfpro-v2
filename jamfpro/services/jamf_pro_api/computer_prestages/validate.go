package computer_prestages

import (
	"fmt"
	"slices"
	"strings"
)

// validateRequest validates the computer prestage request for Create and Update operations.
func validateRequest(prestage *ResourceComputerPrestage) error {
	if prestage == nil {
		return fmt.Errorf("prestage cannot be nil")
	}

	// Validate required fields for create/update
	if prestage.DisplayName == "" {
		return fmt.Errorf("displayName is required")
	}

	// Validate enum fields
	if err := validateRecoveryLockPasswordType(prestage.RecoveryLockPasswordType); err != nil {
		return err
	}

	if err := validatePrestageMinimumOsTargetVersionType(prestage.PrestageMinimumOsTargetVersionType); err != nil {
		return err
	}

	// Validate AccountSettings if present
	if prestage.AccountSettings != nil {
		if err := validateUserAccountType(prestage.AccountSettings.UserAccountType); err != nil {
			return err
		}
		if err := validatePrefillType(prestage.AccountSettings.PrefillType); err != nil {
			return err
		}
	}

	return nil
}

// validateRecoveryLockPasswordType validates the recovery lock password type enum.
func validateRecoveryLockPasswordType(passwordType string) error {
	if passwordType == "" {
		return nil // Empty is allowed (optional field)
	}

	validTypes := []string{"MANUAL", "RANDOM"}
	if !slices.Contains(validTypes, passwordType) {
		return fmt.Errorf("invalid recoveryLockPasswordType: %s (must be one of: %s)",
			passwordType, strings.Join(validTypes, ", "))
	}
	return nil
}

// validatePrestageMinimumOsTargetVersionType validates the prestage minimum OS target version type enum.
func validatePrestageMinimumOsTargetVersionType(versionType string) error {
	if versionType == "" {
		return nil // Empty is allowed
	}

	validTypes := []string{
		"NO_ENFORCEMENT",
		"MINIMUM_OS_LATEST_VERSION",
		"MINIMUM_OS_LATEST_MAJOR_VERSION",
		"MINIMUM_OS_LATEST_MINOR_VERSION",
		"MINIMUM_OS_SPECIFIC_VERSION",
	}
	if !slices.Contains(validTypes, versionType) {
		return fmt.Errorf("invalid prestageMinimumOsTargetVersionType: %s (must be one of: %s)",
			versionType, strings.Join(validTypes, ", "))
	}
	return nil
}

// validateUserAccountType validates the user account type enum in account settings.
func validateUserAccountType(accountType string) error {
	if accountType == "" {
		return nil // Empty is allowed
	}

	validTypes := []string{"ADMINISTRATOR", "STANDARD", "SKIP"}
	if !slices.Contains(validTypes, accountType) {
		return fmt.Errorf("invalid userAccountType: %s (must be one of: %s)",
			accountType, strings.Join(validTypes, ", "))
	}
	return nil
}

// validatePrefillType validates the prefill type enum in account settings.
func validatePrefillType(prefillType string) error {
	if prefillType == "" {
		return nil // Empty is allowed
	}

	validTypes := []string{"CUSTOM", "DEVICE_OWNER"}
	if !slices.Contains(validTypes, prefillType) {
		return fmt.Errorf("invalid prefillType: %s (must be one of: CUSTOM, DEVICE_OWNER)", prefillType)
	}
	return nil
}
