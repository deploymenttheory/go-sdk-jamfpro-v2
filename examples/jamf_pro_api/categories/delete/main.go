// Package main demonstrates DeleteCategoryByID - removes a category by ID.
//
// Run with: go run ./examples/jamf_pro_api/categories/delete
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates a category then deletes it.
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

	// Create a category to delete
	createReq := &categories.RequestCategory{
		Name:     fmt.Sprintf("example-delete-%d", time.Now().UnixMilli()),
		Priority: 1,
	}
	created, _, err := client.Categories.CreateCategory(ctx, createReq)
	if err != nil {
		log.Fatalf("CreateCategory failed: %v", err)
	}
	id := created.ID
	fmt.Printf("Created category ID: %s\n", id)

	resp, err := client.Categories.DeleteCategoryByID(ctx, id)
	if err != nil {
		log.Fatalf("DeleteCategoryByID failed: %v", err)
	}

	fmt.Printf("Status: %d (204 = success)\n", resp.StatusCode)
	fmt.Println("Category deleted successfully")
}
