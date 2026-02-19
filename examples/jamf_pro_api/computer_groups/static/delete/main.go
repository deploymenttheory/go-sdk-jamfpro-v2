// Package main demonstrates DeleteStaticGroupByIDV2 - deletes a static computer group.
//
// Run with: go run ./examples/jamf_pro_api/computer_groups/static/delete
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, auth env vars, and GROUP_ID env var.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
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

	resp, err := client.ComputerGroups.DeleteStaticGroupByIDV2(ctx, groupID)
	if err != nil {
		log.Fatalf("DeleteStaticGroupByIDV2 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Println("Static group deleted successfully")
}
