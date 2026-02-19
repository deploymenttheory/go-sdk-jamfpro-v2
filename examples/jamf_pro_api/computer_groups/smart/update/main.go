// Package main demonstrates UpdateSmartGroupV2 - updates a smart computer group.
//
// Run with: go run ./examples/jamf_pro_api/computer_groups/smart/update
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, auth env vars, and GROUP_ID env var.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_groups"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	groupID := os.Getenv("GROUP_ID")
	if groupID == "" {
		log.Fatal("GROUP_ID env var required")
	}

	ctx := context.Background()

	req := &computer_groups.RequestSmartGroup{
		Name: "Updated Smart Group Name",
		Criteria: []computer_groups.Criterion{
			{Name: "Computer Name", Priority: 0, AndOr: "and", SearchType: "is", Value: "MyMac"},
		},
	}

	result, resp, err := client.ComputerGroups.UpdateSmartGroupV2(ctx, groupID, req)
	if err != nil {
		log.Fatalf("UpdateSmartGroupV2 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Updated smart group ID: %s\n", result.ID)
	fmt.Printf("Name: %s\n", result.Name)
}
