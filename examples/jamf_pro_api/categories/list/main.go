// Package main demonstrates ListCategories - retrieves all category objects with optional pagination and RSQL filtering.
//
// Run with: go run ./examples/jamf_pro_api/categories/list
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and OAuth2 or Basic auth env vars.
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	// List all categories (first page, default page size)
	result, resp, err := client.Categories.ListCategories(ctx, map[string]string{
		"page":     "0",
		"pageSize": "50",
	})
	if err != nil {
		log.Fatalf("ListCategories failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Total count: %d\n", result.TotalCount)
	for i, c := range result.Results {
		if i >= 5 {
			fmt.Printf("... and %d more\n", result.TotalCount-5)
			break
		}
		fmt.Printf("  ID=%s Name=%q Priority=%d\n", c.ID, c.Name, c.Priority)
	}

	// Example: list with RSQL filter
	filtered, _, err := client.Categories.ListCategories(ctx, map[string]string{
		"filter": `name=="Critical"`,
	})
	if err != nil {
		log.Fatalf("ListCategories with filter failed: %v", err)
	}
	fmt.Printf("\nFiltered (name==\"Critical\"): %d result(s)\n", filtered.TotalCount)
	for _, c := range filtered.Results {
		fmt.Printf("  ID=%s Name=%q\n", c.ID, c.Name)
	}
}
