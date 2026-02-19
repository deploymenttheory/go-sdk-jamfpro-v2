// Package main demonstrates GetCategoryByID - retrieves a single category by ID.
//
// Run with: go run ./examples/jamf_pro_api/categories/get
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Set CATEGORY_ID or uses first from list.
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
	id := os.Getenv("CATEGORY_ID")
	if id == "" {
		// Fallback: use first category from list
		list, _, err := client.Categories.ListCategories(ctx, map[string]string{"page": "0", "pageSize": "1"})
		if err != nil || len(list.Results) == 0 {
			log.Fatal("Set CATEGORY_ID or ensure at least one category exists")
		}
		id = list.Results[0].ID
		fmt.Printf("Using first category ID: %s\n", id)
	}

	category, resp, err := client.Categories.GetCategoryByID(ctx, id)
	if err != nil {
		log.Fatalf("GetCategoryByID failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("ID: %s\n", category.ID)
	fmt.Printf("Name: %s\n", category.Name)
	fmt.Printf("Priority: %d\n", category.Priority)
}
