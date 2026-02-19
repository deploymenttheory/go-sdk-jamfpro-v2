// Package main demonstrates DeleteWebhookByID — removes a webhook via the Classic API.
//
// Run with: go run ./examples/classic_api/webhooks/delete
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates a webhook then deletes it.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/webhooks"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	// Create a webhook to delete
	createReq := &webhooks.RequestWebhook{
		Name:               fmt.Sprintf("example-delete-%d", time.Now().UnixMilli()),
		Enabled:            false,
		URL:                "https://hooks.example.com/jamf",
		ContentType:        "application/json",
		Event:              "ComputerAdded",
		AuthenticationType: "NONE",
	}
	created, _, err := client.Webhooks.CreateWebhook(ctx, createReq)
	if err != nil {
		log.Fatalf("CreateWebhook failed: %v", err)
	}
	fmt.Printf("Created webhook ID: %d\n", created.ID)

	resp, err := client.Webhooks.DeleteWebhookByID(ctx, created.ID)
	if err != nil {
		log.Fatalf("DeleteWebhookByID failed: %v", err)
	}

	fmt.Printf("Status: %d (200 = success)\n", resp.StatusCode)
	fmt.Println("Webhook deleted successfully")
}
