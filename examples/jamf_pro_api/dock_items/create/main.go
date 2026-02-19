// Package main demonstrates CreateDockItemV1 - creates a new dock item.
//
// Run with: go run ./examples/jamf_pro_api/dock_items/create
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/dock_items"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	req := &dock_items.RequestDockItem{
		Name: fmt.Sprintf("example-dock-%d", time.Now().UnixMilli()),
		Path: "/Applications/Safari.app",
		Type: dock_items.TypeApp,
	}

	result, resp, err := client.DockItems.CreateDockItemV1(ctx, req)
	if err != nil {
		log.Fatalf("CreateDockItemV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Created dock item ID: %s\n", result.ID)
	fmt.Printf("Href: %s\n", result.Href)

	if _, err := client.DockItems.DeleteDockItemByIDV1(ctx, result.ID); err != nil {
		fmt.Printf("Note: cleanup delete failed: %v\n", err)
	} else {
		fmt.Println("Cleanup: dock item deleted")
	}
}
