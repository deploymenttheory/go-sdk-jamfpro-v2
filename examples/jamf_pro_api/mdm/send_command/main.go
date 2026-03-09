package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/mdm"
)

func main() {
	// Initialize the Jamf Pro client from environment variables
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	ctx := context.Background()

	// Get device management ID from environment or use example value
	managementID := os.Getenv("JAMF_DEVICE_MANAGEMENT_ID")
	if managementID == "" {
		managementID = "device-management-id-001"
	}

	// Send a DeviceLock command
	req := &mdm.CommandRequest{
		CommandData: mdm.CommandData{
			CommandType: "DeviceLock",
		},
		ClientData: []mdm.ClientData{
			{ManagementID: managementID},
		},
	}

	result, resp, err := client.JamfProAPI.Mdm.SendCommand(ctx, req)
	if err != nil {
		log.Fatalf("Failed to send MDM command: %v (HTTP %d)", err, resp.StatusCode())
	}

	fmt.Printf("MDM command sent successfully (HTTP %d)\n", resp.StatusCode())
	fmt.Printf("Command ID: %s\n", result.ID)
	fmt.Printf("Command href: %s\n", result.Href)
}
