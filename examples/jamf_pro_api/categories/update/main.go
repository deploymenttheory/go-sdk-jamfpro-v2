// Package main demonstrates UpdateCategoryByIDV1 - updates an existing category.
//
// Run with: go run ./examples/jamf_pro_api/categories/update
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
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

	// Create a category to update
	createReq := &categories.RequestCategory{
		Name:     fmt.Sprintf("example-update-%d", time.Now().UnixMilli()),
		Priority: 1,
	}
	created, _, err := client.Categories.CreateCategoryV1(ctx, createReq)
	if err != nil {
		log.Fatalf("CreateCategoryV1 failed: %v", err)
	}
	id := created.ID

	// Update the category
	updateReq := &categories.RequestCategory{
		Name:     fmt.Sprintf("example-updated-%d", time.Now().UnixMilli()),
		Priority: 9,
	}
	result, resp, err := client.Categories.UpdateCategoryByIDV1(ctx, id, updateReq)
	if err != nil {
		_, _ = client.Categories.DeleteCategoryByIDV1(ctx, id)
		log.Fatalf("UpdateCategoryByIDV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Updated category ID: %s\n", result.ID)
	fmt.Printf("Href: %s\n", result.Href)

	// Verify
	fetched, _, _ := client.Categories.GetCategoryByIDV1(ctx, id)
	if fetched != nil {
		fmt.Printf("Verified: name=%q priority=%d\n", fetched.Name, fetched.Priority)
	}

	_, _ = client.Categories.DeleteCategoryByIDV1(ctx, id)
	fmt.Println("Cleanup: category deleted")
}
