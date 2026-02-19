// Package main demonstrates ListBuildingsV1 - retrieves all building objects.
//
// Run with: go run ./examples/jamf_pro_api/buildings/list
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

	result, resp, err := client.Buildings.ListBuildingsV1(ctx, map[string]string{
		"page":     "0",
		"pageSize": "50",
	})
	if err != nil {
		log.Fatalf("ListBuildingsV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Total count: %d\n", result.TotalCount)
	for i, b := range result.Results {
		if i >= 5 {
			fmt.Printf("... and %d more\n", result.TotalCount-5)
			break
		}
		fmt.Printf("  ID=%s Name=%q City=%s\n", b.ID, b.Name, b.City)
	}
}
