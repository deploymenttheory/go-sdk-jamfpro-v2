// Package main demonstrates CreateStaticGroupV2 - creates a new static computer group.
//
// Run with: go run ./examples/jamf_pro_api/computer_groups/static/create
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

	req := &computer_groups.RequestStaticGroup{
		Name:        fmt.Sprintf("example-static-group-%d", time.Now().UnixMilli()),
		ComputerIds: []string{},
	}

	result, resp, err := client.ComputerGroups.CreateStaticGroupV2(ctx, req)
	if err != nil {
		log.Fatalf("CreateStaticGroupV2 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Created static group ID: %s\n", result.ID)
	fmt.Printf("Href: %s\n", result.Href)

	if _, err := client.ComputerGroups.DeleteStaticGroupByIDV2(ctx, result.ID); err != nil {
		fmt.Printf("Note: cleanup delete failed: %v\n", err)
	} else {
		fmt.Println("Cleanup: static group deleted")
	}
}
