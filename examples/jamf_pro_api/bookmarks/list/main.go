// Package main demonstrates ListV1 - lists bookmarks.
//
// Run with: go run ./examples/jamf_pro_api/bookmarks/list
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
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

	result, resp, err := client.Bookmarks.ListV1(ctx, map[string]string{"page": "0", "page-size": "50"})
	if err != nil {
		log.Fatalf("ListV1 failed: %v", err)
	}
	fmt.Printf("Status: %d Total: %d\n", resp.StatusCode, result.TotalCount)
	for i, b := range result.Results {
		if i >= 10 {
			break
		}
		fmt.Printf("  ID=%s Name=%q URL=%s\n", b.ID, b.Name, b.URL)
	}
}
