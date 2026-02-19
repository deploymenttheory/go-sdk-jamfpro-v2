// Package main demonstrates ListV1 - lists notifications for the current user and site.
//
// Run with: go run ./examples/jamf_pro_api/notifications/list
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

	result, resp, err := client.Notifications.ListV1(ctx)
	if err != nil {
		log.Fatalf("ListV1 failed: %v", err)
	}
	fmt.Printf("Status: %d Notifications: %d\n", resp.StatusCode, len(result))
	for i, n := range result {
		if i >= 5 {
			break
		}
		fmt.Printf("  Type=%s ID=%s\n", n.Type, n.ID)
	}
}
