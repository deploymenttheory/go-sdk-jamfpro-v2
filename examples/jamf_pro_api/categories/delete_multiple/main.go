// Package main demonstrates DeleteCategoriesByID - deletes multiple categories by their IDs.
//
// Run with: go run ./examples/jamf_pro_api/categories/delete_multiple
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates two categories then bulk deletes them.
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

	// Create two categories
	ids := make([]string, 0, 2)
	for i := 0; i < 2; i++ {
		req := &categories.RequestCategory{
			Name:     fmt.Sprintf("example-bulk-%d-%d", i, time.Now().UnixMilli()),
			Priority: 1,
		}
		created, _, err := client.Categories.CreateCategory(ctx, req)
		if err != nil {
			log.Fatalf("CreateCategory %d failed: %v", i, err)
		}
		ids = append(ids, created.ID)
		fmt.Printf("Created category ID: %s\n", created.ID)
	}

	// Bulk delete
	bulkReq := &categories.DeleteCategoriesByIDRequest{IDs: ids}
	resp, err := client.Categories.DeleteCategoriesByID(ctx, bulkReq)
	if err != nil {
		log.Fatalf("DeleteCategoriesByID failed: %v", err)
	}

	fmt.Printf("Status: %d (204 = success)\n", resp.StatusCode)
	fmt.Printf("Deleted %d categories: %v\n", len(ids), ids)
}
