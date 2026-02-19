// Package main demonstrates GetCategoryHistory - retrieves the history object for a category.
//
// Run with: go run ./examples/jamf_pro_api/categories/get_history
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
		list, _, err := client.Categories.ListCategories(ctx, map[string]string{"page": "0", "pageSize": "1"})
		if err != nil || len(list.Results) == 0 {
			log.Fatal("Set CATEGORY_ID or ensure at least one category exists")
		}
		id = list.Results[0].ID
		fmt.Printf("Using first category ID: %s\n", id)
	}

	history, resp, err := client.Categories.GetCategoryHistory(ctx, id)
	if err != nil {
		log.Fatalf("GetCategoryHistory failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Total history entries: %d\n", history.TotalCount)
	for i, entry := range history.Results {
		if i >= 10 {
			fmt.Printf("... and %d more\n", history.TotalCount-10)
			break
		}
		fmt.Printf("  [%s] %s: %s (%s)\n", entry.Date, entry.Username, entry.Note, entry.Details)
	}
}
