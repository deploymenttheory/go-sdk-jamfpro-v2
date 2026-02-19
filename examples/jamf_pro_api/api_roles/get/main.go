// Package main demonstrates GetAPIRoleByIDV1 - gets an API role by ID.
//
// Run with: go run ./examples/jamf_pro_api/api_roles/get <id>
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: go run ./examples/jamf_pro_api/api_roles/get <id>")
	}
	id := os.Args[1]

	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	ctx := context.Background()

	result, resp, err := client.APIRoles.GetAPIRoleByIDV1(ctx, id)
	if err != nil {
		log.Fatalf("GetAPIRoleByIDV1 failed: %v", err)
	}
	fmt.Printf("Status: %d ID: %s DisplayName: %s Privileges: %v\n", resp.StatusCode, result.ID, result.DisplayName, result.Privileges)
}
