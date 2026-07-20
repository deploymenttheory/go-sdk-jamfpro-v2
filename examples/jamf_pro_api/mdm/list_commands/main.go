package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	// Initialize the Jamf Pro client from environment variables
	jamfClient, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	ctx := context.Background()

	// NOTE: /v2/mdm/commands requires a filter; the API returns HTTP 400 without one.
	// Valid statuses include Pending and Acknowledged.
	result, _, err := jamfClient.JamfProAPI.Mdm.ListCommandsV2(ctx, map[string]string{
		"filter": "status==Pending",
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Total Commands: %d\n\n", result.TotalCount)
	for _, cmd := range result.Results {
		fmt.Printf("UUID: %s\n", cmd.UUID)
		fmt.Printf("  Type: %s\n", cmd.CommandType)
		fmt.Printf("  Status: %s\n", cmd.Status)
		fmt.Printf("  Date Sent: %s\n", cmd.DateSent)
		fmt.Println()
	}
}
