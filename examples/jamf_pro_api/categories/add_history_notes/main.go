// Package main demonstrates AddCategoryHistoryNotesV1 - adds notes to a category's history.
//
// Run with: go run ./examples/jamf_pro_api/categories/add_history_notes
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates a category, adds a note, then deletes it.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/categories"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	// Create a category
	createReq := &categories.RequestCategory{
		Name:     fmt.Sprintf("example-history-%d", time.Now().UnixMilli()),
		Priority: 1,
	}
	created, _, err := client.Categories.CreateCategoryV1(ctx, createReq)
	if err != nil {
		log.Fatalf("CreateCategoryV1 failed: %v", err)
	}
	id := created.ID
	fmt.Printf("Created category ID: %s\n", id)

	// Add a history note
	noteReq := &categories.AddCategoryHistoryNotesRequest{
		Note: fmt.Sprintf("Example note added at %s", time.Now().Format(time.RFC3339)),
	}
	resp, err := client.Categories.AddCategoryHistoryNotesV1(ctx, id, noteReq)
	if err != nil {
		_, _ = client.Categories.DeleteCategoryByIDV1(ctx, id)
		log.Fatalf("AddCategoryHistoryNotesV1 failed: %v", err)
	}

	fmt.Printf("Status: %d (201 = success)\n", resp.StatusCode)
	fmt.Println("History note added")

	// Fetch history to verify
	history, _, err := client.Categories.GetCategoryHistoryV1(ctx, id, nil)
	if err == nil {
		fmt.Printf("History entries: %d\n", history.TotalCount)
		for _, e := range history.Results {
			if e.Note != "" {
				fmt.Printf("  Note: %s (by %s)\n", e.Note, e.Username)
			}
		}
	}

	_, _ = client.Categories.DeleteCategoryByIDV1(ctx, id)
	fmt.Println("Cleanup: category deleted")
}
