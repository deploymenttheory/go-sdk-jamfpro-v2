// Package main demonstrates CreateCategoryV1 - creates a new category.
//
// Run with: go run ./examples/jamf_pro_api/categories/create
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

	req := &categories.RequestCategory{
		Name:     fmt.Sprintf("example-category-%d", time.Now().UnixMilli()),
		Priority: 5,
	}

	result, resp, err := client.Categories.CreateCategoryV1(ctx, req)
	if err != nil {
		log.Fatalf("CreateCategoryV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Created category ID: %s\n", result.ID)
	fmt.Printf("Href: %s\n", result.Href)

	// Cleanup: delete the created category
	if _, err := client.Categories.DeleteCategoryByIDV1(ctx, result.ID); err != nil {
		fmt.Printf("Note: cleanup delete failed: %v\n", err)
	} else {
		fmt.Println("Cleanup: category deleted")
	}
}
