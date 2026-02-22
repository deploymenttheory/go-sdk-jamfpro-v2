package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/local_admin_password"
)

func main() {
	// Initialize the Jamf Pro client from environment variables
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	ctx := context.Background()

	// Get current settings first
	currentSettings, _, err := client.LocalAdminPassword.GetSettingsV2(ctx)
	if err != nil {
		log.Fatalf("Failed to get current LAPS settings: %v", err)
	}

	fmt.Printf("Current LAPS Settings:\n")
	fmt.Printf("  Auto Deploy Enabled: %v\n", currentSettings.AutoDeployEnabled)
	fmt.Printf("  Password Rotation Time: %d days\n", currentSettings.PasswordRotationTime)
	fmt.Printf("  Auto Rotate Enabled: %v\n", currentSettings.AutoRotateEnabled)
	fmt.Printf("  Auto Rotate Expiration Time: %d days\n", currentSettings.AutoRotateExpirationTime)

	// Update settings
	updatedSettings := &local_admin_password.SettingsResource{
		AutoDeployEnabled:        true,
		PasswordRotationTime:     90,
		AutoRotateEnabled:        true,
		AutoRotateExpirationTime: 7,
	}

	resp, err := client.LocalAdminPassword.UpdateSettingsV2(ctx, updatedSettings)
	if err != nil {
		log.Fatalf("Failed to update LAPS settings: %v (HTTP %d)", err, resp.StatusCode)
	}

	fmt.Printf("\nLAPS settings updated successfully (HTTP %d)\n", resp.StatusCode)

	// Verify the update
	verifySettings, _, err := client.LocalAdminPassword.GetSettingsV2(ctx)
	if err != nil {
		log.Fatalf("Failed to verify updated LAPS settings: %v", err)
	}

	fmt.Printf("\nVerified Updated Settings:\n")
	fmt.Printf("  Auto Deploy Enabled: %v\n", verifySettings.AutoDeployEnabled)
	fmt.Printf("  Password Rotation Time: %d days\n", verifySettings.PasswordRotationTime)
	fmt.Printf("  Auto Rotate Enabled: %v\n", verifySettings.AutoRotateEnabled)
	fmt.Printf("  Auto Rotate Expiration Time: %d days\n", verifySettings.AutoRotateExpirationTime)
}
