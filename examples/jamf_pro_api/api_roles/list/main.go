// Package main demonstrates ListAPIRolesV1 - lists API roles.
//
// Run with: go run ./examples/jamf_pro_api/api_roles/list
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

	result, resp, err := client.APIRoles.ListAPIRolesV1(ctx, map[string]string{"page": "0", "page-size": "50"})
	if err != nil {
		log.Fatalf("ListAPIRolesV1 failed: %v", err)
	}
	fmt.Printf("Status: %d Total: %d\n", resp.StatusCode, result.TotalCount)
	for i, r := range result.Results {
		if i >= 10 {
			break
		}
		fmt.Printf("  ID=%s DisplayName=%q Privileges=%d\n", r.ID, r.DisplayName, len(r.Privileges))
	}
}
