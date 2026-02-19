// Package main demonstrates ListPrivilegesV1 - lists API role privileges.
//
// Run with: go run ./examples/jamf_pro_api/api_role_privileges/list
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

	result, resp, err := client.APIRolePrivileges.ListPrivilegesV1(ctx)
	if err != nil {
		log.Fatalf("ListPrivilegesV1 failed: %v", err)
	}
	fmt.Printf("Status: %d Privileges: %d\n", resp.StatusCode, len(result.Privileges))
	for i, p := range result.Privileges {
		if i >= 10 {
			break
		}
		fmt.Printf("  %s\n", p)
	}
}
