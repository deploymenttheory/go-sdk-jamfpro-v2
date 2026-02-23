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
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	ctx := context.Background()

	// Get client management IDs from environment or use example values
	deviceID := os.Getenv("JAMF_DEVICE_MANAGEMENT_ID")
	if deviceID == "" {
		deviceID = "device-management-id-001"
	}
	deviceIDs := []string{deviceID}

	result, resp, err := client.MDM.BlankPush(ctx, deviceIDs)
	if err != nil {
		log.Fatalf("Failed to send blank push: %v (HTTP %d)", err, resp.StatusCode)
	}

	fmt.Printf("Blank push sent successfully (HTTP %d)\n", resp.StatusCode)
	if len(result.ErrorUUIDs) > 0 {
		fmt.Printf("Devices with errors: %v\n", result.ErrorUUIDs)
	} else {
		fmt.Println("All devices received the blank push command.")
	}
}
