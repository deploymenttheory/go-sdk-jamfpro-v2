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

	group, resp, err := client.Groups.GetByIDV1(ctx, platformID)
	if err != nil {
		log.Fatalf("Error getting group: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Group Platform ID: %s\n", group.GroupPlatformId)
	fmt.Printf("Group Jamf Pro ID: %s\n", group.GroupJamfProId)
	fmt.Printf("Group Name: %s\n", group.GroupName)
	fmt.Printf("Description: %s\n", group.GroupDescription)
	fmt.Printf("Type: %s\n", group.GroupType)
	fmt.Printf("Smart Group: %t\n", group.Smart)
	fmt.Printf("Membership Count: %d\n", group.MembershipCount)

	if group.Smart && len(group.Criteria) > 0 {
		fmt.Printf("\nSmart Group Criteria:\n")
		for i, criterion := range group.Criteria {
			fmt.Printf("  %d. %s %s %s (Priority: %d)\n",
				i+1, criterion.Name, criterion.SearchType, criterion.Value, criterion.Priority)
		}
	}

	if !group.Smart && len(group.Assignments) > 0 {
		fmt.Printf("\nStatic Group Assignments:\n")
		for i, assignment := range group.Assignments {
			fmt.Printf("  %d. Device ID: %s, Selected: %t\n",
				i+1, assignment.DeviceID, assignment.Selected)
		}
	}
}
