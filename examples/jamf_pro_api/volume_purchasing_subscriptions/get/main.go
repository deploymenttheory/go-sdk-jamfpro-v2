// Package main demonstrates GetVolumePurchasingSubscriptionByIDV1 - retrieves a volume purchasing subscription by ID.
//
// Run with: go run ./examples/jamf_pro_api/volume_purchasing_subscriptions/get <id>
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
		log.Fatalf("Usage: go run ./examples/jamf_pro_api/volume_purchasing_subscriptions/get <id>")
	}
	id := os.Args[1]

	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	result, resp, err := client.VolumePurchasingSubscriptions.GetVolumePurchasingSubscriptionByIDV1(ctx, id)
	if err != nil {
		log.Fatalf("GetVolumePurchasingSubscriptionByIDV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("ID: %s Name: %s Enabled: %v\n", result.ID, result.Name, result.Enabled)
}
