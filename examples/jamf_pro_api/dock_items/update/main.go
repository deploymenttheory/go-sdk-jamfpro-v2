// Package main demonstrates UpdateDockItemByIDV1 - updates an existing dock item.
//
// Run with: go run ./examples/jamf_pro_api/dock_items/update
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

	createReq := &dock_items.RequestDockItem{
		Name: fmt.Sprintf("example-update-%d", time.Now().UnixMilli()),
		Path: "/Applications/Safari.app",
		Type: dock_items.TypeApp,
	}
	created, _, err := client.DockItems.CreateDockItemV1(ctx, createReq)
	if err != nil {
		log.Fatalf("CreateDockItemV1 failed: %v", err)
	}
	id := created.ID

	updateReq := &dock_items.RequestDockItem{
		Name: fmt.Sprintf("example-updated-%d", time.Now().UnixMilli()),
		Path: "/Applications/Google Chrome.app",
		Type: dock_items.TypeApp,
	}
	result, resp, err := client.DockItems.UpdateDockItemByIDV1(ctx, id, updateReq)
	if err != nil {
		_, _ = client.DockItems.DeleteDockItemByIDV1(ctx, id)
		log.Fatalf("UpdateDockItemByIDV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Updated dock item ID: %s Name=%q Path=%q\n", result.ID, result.Name, result.Path)

	fetched, _, _ := client.DockItems.GetDockItemByIDV1(ctx, id)
	if fetched != nil {
		fmt.Printf("Verified: name=%q path=%q\n", fetched.Name, fetched.Path)
	}

	_, _ = client.DockItems.DeleteDockItemByIDV1(ctx, id)
	fmt.Println("Cleanup: dock item deleted")
}
