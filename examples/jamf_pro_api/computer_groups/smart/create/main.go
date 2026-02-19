// Package main demonstrates CreateSmartGroupV2 - creates a new smart computer group.
//
// Run with: go run ./examples/jamf_pro_api/computer_groups/smart/create
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_groups"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	req := &computer_groups.RequestSmartGroup{
		Name: fmt.Sprintf("example-smart-group-%d", time.Now().UnixMilli()),
		Criteria: []computer_groups.Criterion{
			{Name: "Computer Name", Priority: 0, AndOr: "and", SearchType: "like", Value: "Mac-%"},
		},
	}

	result, resp, err := client.ComputerGroups.CreateSmartGroupV2(ctx, req)
	if err != nil {
		log.Fatalf("CreateSmartGroupV2 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Created smart group ID: %s\n", result.ID)
	fmt.Printf("Href: %s\n", result.Href)

	if _, err := client.ComputerGroups.DeleteSmartGroupV2(ctx, result.ID); err != nil {
		fmt.Printf("Note: cleanup delete failed: %v\n", err)
	} else {
		fmt.Println("Cleanup: smart group deleted")
	}
}
