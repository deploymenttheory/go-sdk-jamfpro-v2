// Package main demonstrates ListStaticGroupsV2 - retrieves all static computer groups.
//
// Run with: go run ./examples/jamf_pro_api/computer_groups/static/list
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

	result, resp, err := client.ComputerGroups.ListStaticGroupsV2(ctx, map[string]string{
		"page":     "0",
		"pageSize": "50",
	})
	if err != nil {
		log.Fatalf("ListStaticGroupsV2 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Total count: %d\n", result.TotalCount)
	for i, g := range result.Results {
		if i >= 5 {
			fmt.Printf("... and %d more\n", result.TotalCount-5)
			break
		}
		fmt.Printf("  ID=%s Name=%q (static, %d computers)\n", g.ID, g.Name, len(g.ComputerIds))
	}
}
