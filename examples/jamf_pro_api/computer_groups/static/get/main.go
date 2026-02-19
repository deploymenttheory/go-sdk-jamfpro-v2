// Package main demonstrates GetStaticGroupByIDV2 - retrieves a static computer group by ID.
//
// Run with: go run ./examples/jamf_pro_api/computer_groups/static/get
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, auth env vars, and GROUP_ID env var (or pass as arg).
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
	if groupID == "" && len(os.Args) > 1 {
		groupID = os.Args[1]
	}
	if groupID == "" {
		log.Fatal("GROUP_ID env var or argument required")
	}

	ctx := context.Background()

	result, resp, err := client.ComputerGroups.GetStaticGroupByIDV2(ctx, groupID)
	if err != nil {
		log.Fatalf("GetStaticGroupByIDV2 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("ID: %s\n", result.ID)
	fmt.Printf("Name: %s\n", result.Name)
	fmt.Printf("IsSmart: %v\n", result.IsSmart)
	fmt.Printf("Computer IDs: %v\n", result.ComputerIds)
}
