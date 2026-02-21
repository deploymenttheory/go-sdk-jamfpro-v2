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

	// List all groups
	groups, resp, err := client.Groups.ListV1(ctx, nil)
	if err != nil {
		log.Fatalf("Error listing groups: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Total groups: %d\n", groups.TotalCount)
	for _, g := range groups.Results {
		fmt.Printf("- ID: %s, Name: %s, Type: %s, Jamf Pro ID: %s, Members: %d\n",
			g.GroupPlatformId, g.GroupName, g.GroupType, g.GroupJamfProId, g.MembershipCount)
	}

	// List with pagination and filtering
	rsqlQuery := map[string]string{
		"page":      "0",
		"page-size": "10",
		"sort":      "groupName:asc",
		"filter":    `groupType=="COMPUTER"`, // Only computer groups
	}

	computerGroups, paginatedResp, err := client.Groups.ListV1(ctx, rsqlQuery)
	if err != nil {
		log.Fatalf("Error listing computer groups with pagination: %v", err)
	}

	fmt.Printf("\nPaginated Computer Groups Status: %d\n", paginatedResp.StatusCode)
	fmt.Printf("Computer Groups Total: %d\n", computerGroups.TotalCount)
	for _, g := range computerGroups.Results {
		fmt.Printf("- %s (%s) - %d members\n", g.GroupName, g.GroupJamfProId, g.MembershipCount)
	}
}
