package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	// Initialize the Jamf Pro client from environment variables
	jamfClient, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	ctx := context.Background()

	// Look up commands either by client management ID or by command UUIDs
	// (max 40 UUIDs; more returns HTTP 414).
	managementID := os.Getenv("JAMF_DEVICE_MANAGEMENT_ID")
	if managementID == "" {
		managementID = "device-management-id-001"
	}

	// Deprecated: Jamf deprecated GET /v1/mdm/commands on 2023-10-16. Prefer
	// ListCommandsV2. It is shown here because it returns richer per-command
	// detail (dateCompleted, client, commandError) than the v2 list shape.
	result, _, err := jamfClient.JamfProAPI.Mdm.ListCommandsV1(ctx, map[string]string{
		"client-management-id": managementID,
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Commands: %d\n\n", len(result))
	for _, cmd := range result {
		fmt.Printf("UUID: %s\n", cmd.UUID)
		fmt.Printf("  Type: %s\n", cmd.CommandType)
		fmt.Printf("  State: %s\n", cmd.CommandState)
		fmt.Printf("  Sent: %s\n", cmd.DateSent)
		fmt.Printf("  Completed: %s\n", cmd.DateCompleted)
		if cmd.Client != nil {
			fmt.Printf("  Client: %s (%s)\n", cmd.Client.ManagementID, cmd.Client.ClientType)
		}
		if cmd.CommandError != nil {
			fmt.Printf("  Error %d: %s\n", cmd.CommandError.ErrorCode, cmd.CommandError.ErrorEnglishDescription)
		}
		fmt.Println()
	}
}
