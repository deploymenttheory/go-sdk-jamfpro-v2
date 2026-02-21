package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/groups"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	ctx := context.Background()

	// Replace with actual group platform ID
	platformID := "platform-id-here"

	// First, get the existing group
	existing, _, err := client.Groups.GetByIDV1(ctx, platformID)
	if err != nil {
		log.Fatalf("Error getting existing group: %v", err)
	}

	fmt.Printf("Current group name: %s\n", existing.GroupName)
	fmt.Printf("Current description: %s\n", existing.GroupDescription)

	// Update the group
	updateReq := &groups.RequestUpdateGroup{
		GroupName:        existing.GroupName,
		GroupDescription: "Updated description via Go SDK",
	}

	updated, resp, err := client.Groups.UpdateByIDV1(ctx, platformID, updateReq)
	if err != nil {
		log.Fatalf("Error updating group: %v", err)
	}

	fmt.Printf("\nStatus: %d\n", resp.StatusCode)
	fmt.Printf("Updated group name: %s\n", updated.GroupName)
	fmt.Printf("Updated description: %s\n", updated.GroupDescription)
}
