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

	// Replace with actual computer group name
	groupName := "All Managed Clients"

	group, resp, err := client.Groups.GetComputerGroupByNameV1(ctx, groupName)
	if err != nil {
		log.Fatalf("Error getting computer group by name: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Group Platform ID: %s\n", group.GroupPlatformId)
	fmt.Printf("Group Jamf Pro ID: %s\n", group.GroupJamfProId)
	fmt.Printf("Group Name: %s\n", group.GroupName)
	fmt.Printf("Description: %s\n", group.GroupDescription)
	fmt.Printf("Type: %s\n", group.GroupType)
	fmt.Printf("Smart Group: %t\n", group.Smart)
	fmt.Printf("Membership Count: %d\n", group.MembershipCount)
}
