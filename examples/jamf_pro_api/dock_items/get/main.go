// Package main demonstrates GetDockItemByIDV1 - retrieves a single dock item by ID.
//
// Run with: go run ./examples/jamf_pro_api/dock_items/get
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, DOCK_ITEM_ID, and auth env vars.
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

	ctx := context.Background()
	id := os.Getenv("DOCK_ITEM_ID")
	if id == "" {
		log.Fatal("DOCK_ITEM_ID environment variable is required (dock items have no list endpoint)")
	}

	item, resp, err := client.DockItems.GetDockItemByIDV1(ctx, id)
	if err != nil {
		log.Fatalf("GetDockItemByIDV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("ID: %s\n", item.ID)
	fmt.Printf("Name: %s\n", item.Name)
	fmt.Printf("Path: %s\n", item.Path)
	fmt.Printf("Type: %s\n", item.Type)
	if item.Contents != "" {
		fmt.Printf("Contents: %s\n", item.Contents)
	}
}
