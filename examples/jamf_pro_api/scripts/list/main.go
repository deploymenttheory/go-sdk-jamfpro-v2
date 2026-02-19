// Package main demonstrates ListScripts — retrieves all script objects with optional pagination and RSQL filtering.
//
// Run with: go run ./examples/jamf_pro_api/scripts/list
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

	// List all scripts (first page, default page size)
	result, resp, err := client.Scripts.ListScripts(ctx, map[string]string{
		"page":      "0",
		"page-size": "50",
	})
	if err != nil {
		log.Fatalf("ListScripts failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Total count: %d\n", result.TotalCount)
	for i, s := range result.Results {
		if i >= 5 {
			fmt.Printf("... and %d more\n", result.TotalCount-5)
			break
		}
		fmt.Printf("  ID=%s Name=%q Priority=%s\n", s.ID, s.Name, s.Priority)
	}

	// Example: list with RSQL filter
	filtered, _, err := client.Scripts.ListScripts(ctx, map[string]string{
		"filter": `name=="Install Homebrew"`,
	})
	if err != nil {
		log.Fatalf("ListScripts with filter failed: %v", err)
	}
	fmt.Printf("\nFiltered (name==\"Install Homebrew\"): %d result(s)\n", filtered.TotalCount)
	for _, s := range filtered.Results {
		fmt.Printf("  ID=%s Name=%q\n", s.ID, s.Name)
	}
}
