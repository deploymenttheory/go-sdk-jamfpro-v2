// Package main demonstrates DeleteDockItemByIDV1 - removes a dock item by ID.
//
// Run with: go run ./examples/jamf_pro_api/dock_items/delete
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates a dock item then deletes it.
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

	createReq := &dock_items.RequestDockItem{
		Name: fmt.Sprintf("example-delete-%d", time.Now().UnixMilli()),
		Path: "/Applications/Safari.app",
		Type: dock_items.TypeApp,
	}
	created, _, err := client.DockItems.CreateDockItemV1(ctx, createReq)
	if err != nil {
		log.Fatalf("CreateDockItemV1 failed: %v", err)
	}
	id := created.ID
	fmt.Printf("Created dock item ID: %s\n", id)

	resp, err := client.DockItems.DeleteDockItemByIDV1(ctx, id)
	if err != nil {
		log.Fatalf("DeleteDockItemByIDV1 failed: %v", err)
	}

	fmt.Printf("Status: %d (204 = success)\n", resp.StatusCode)
	fmt.Println("Dock item deleted successfully")
}
