// Package main demonstrates DeleteSmartGroupV2 - deletes a smart computer group.
//
// Run with: go run ./examples/jamf_pro_api/computer_groups/smart/delete
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

	resp, err := client.ComputerGroups.DeleteSmartGroupV2(ctx, groupID)
	if err != nil {
		log.Fatalf("DeleteSmartGroupV2 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Println("Smart group deleted successfully")
}
