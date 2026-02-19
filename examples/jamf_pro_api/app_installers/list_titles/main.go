// Package main demonstrates ListTitlesV1 - lists Jamf App Catalog titles.
//
// Run with: go run ./examples/jamf_pro_api/app_installers/list_titles
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

	result, resp, err := client.AppInstallers.ListTitlesV1(ctx, map[string]string{"page": "0", "page-size": "50"})
	if err != nil {
		log.Fatalf("ListTitlesV1 failed: %v", err)
	}
	fmt.Printf("Status: %d Total titles: %d\n", resp.StatusCode, result.TotalCount)
	for i, r := range result.Results {
		if i >= 5 {
			break
		}
		fmt.Printf("  ID=%s Title=%q\n", r.ID, r.TitleName)
	}
}
