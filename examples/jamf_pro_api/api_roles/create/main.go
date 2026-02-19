// Package main demonstrates CreateAPIRoleV1 - creates an API role.
//
// Run with: go run ./examples/jamf_pro_api/api_roles/create
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/api_roles"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	ctx := context.Background()

	req := &api_roles.RequestAPIRole{
		DisplayName: fmt.Sprintf("example-role-%d", time.Now().UnixMilli()),
		Privileges:  []string{"Read Computers"},
	}
	result, resp, err := client.APIRoles.CreateAPIRoleV1(ctx, req)
	if err != nil {
		log.Fatalf("CreateAPIRoleV1 failed: %v", err)
	}
	fmt.Printf("Status: %d Created ID: %s DisplayName: %s\n", resp.StatusCode, result.ID, result.DisplayName)
	_, _ = client.APIRoles.DeleteAPIRoleByIDV1(ctx, result.ID)
}
