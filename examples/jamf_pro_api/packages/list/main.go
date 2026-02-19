// Package main demonstrates ListPackagesV1 - retrieves all package objects.
//
// Run with: go run ./examples/jamf_pro_api/packages/list
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

	result, resp, err := client.Packages.ListPackagesV1(ctx, map[string]string{
		"page":     "0",
		"pageSize": "50",
	})
	if err != nil {
		log.Fatalf("ListPackagesV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Total count: %d\n", result.TotalCount)
	for i, p := range result.Results {
		if i >= 5 {
			fmt.Printf("... and %d more\n", result.TotalCount-5)
			break
		}
		fmt.Printf("  ID=%s Name=%q FileName=%s Category=%s\n", p.ID, p.PackageName, p.FileName, p.CategoryID)
	}
}
