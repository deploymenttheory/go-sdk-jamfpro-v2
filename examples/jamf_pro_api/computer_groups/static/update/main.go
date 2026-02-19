// Package main demonstrates UpdateStaticGroupByIDV2 - updates static group membership (PATCH).
//
// Run with: go run ./examples/jamf_pro_api/computer_groups/static/update
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

	req := &computer_groups.RequestStaticGroup{
		Name:        "Updated Static Group",
		ComputerIds: []string{"1", "2", "3"},
	}

	result, resp, err := client.ComputerGroups.UpdateStaticGroupByIDV2(ctx, groupID, req)
	if err != nil {
		log.Fatalf("UpdateStaticGroupByIDV2 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Updated static group ID: %s\n", result.ID)
	fmt.Printf("Computer IDs: %v\n", result.ComputerIds)
}
