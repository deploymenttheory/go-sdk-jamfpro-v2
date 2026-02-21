package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	ctx := context.Background()

	// Replace with actual group platform ID
	platformID := "platform-id-here"

	// Note: Groups are typically managed through computer_groups or mobile_device_groups services.
	// Direct deletion via the Groups API may not be supported for all group types.

	resp, err := client.Groups.DeleteByIDV1(ctx, platformID)
	if err != nil {
		log.Fatalf("Error deleting group: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Group %s deleted successfully\n", platformID)
}
