package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	// Initialize the Jamf Pro client from environment variables
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	ctx := context.Background()

	// Get current LAPS settings
	settings, resp, err := client.LocalAdminPassword.GetSettingsV2(ctx)
	if err != nil {
		log.Fatalf("Failed to get LAPS settings: %v (HTTP %d)", err, resp.StatusCode)
	}

	fmt.Printf("LAPS Settings:\n")
	fmt.Printf("  Auto Deploy Enabled: %v\n", settings.AutoDeployEnabled)
	fmt.Printf("  Password Rotation Time: %d days\n", settings.PasswordRotationTime)
	fmt.Printf("  Auto Rotate Enabled: %v\n", settings.AutoRotateEnabled)
	fmt.Printf("  Auto Rotate Expiration Time: %d days\n", settings.AutoRotateExpirationTime)
}
