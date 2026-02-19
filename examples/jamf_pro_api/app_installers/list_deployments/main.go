// Package main demonstrates ListDeploymentsV1 - lists Jamf App Catalog deployments.
//
// Run with: go run ./examples/jamf_pro_api/app_installers/list_deployments
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

	result, resp, err := client.AppInstallers.ListDeploymentsV1(ctx, map[string]string{"page": "0", "page-size": "50"})
	if err != nil {
		log.Fatalf("ListDeploymentsV1 failed: %v", err)
	}
	fmt.Printf("Status: %d Total deployments: %d\n", resp.StatusCode, result.TotalCount)
	for i, r := range result.Results {
		if i >= 5 {
			break
		}
		fmt.Printf("  ID=%s Name=%q\n", r.ID, r.Name)
	}
}
