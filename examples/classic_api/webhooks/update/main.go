// Package main demonstrates UpdateWebhookByID — updates an existing webhook via the Classic API.
//
// Run with: go run ./examples/classic_api/webhooks/update
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
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

	// Create a webhook to update
	createReq := &webhooks.RequestWebhook{
		Name:               fmt.Sprintf("example-update-%d", time.Now().UnixMilli()),
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
	fmt.Printf("Created webhook ID: %d name: %s\n", created.ID, created.Name)

	// Update the webhook
	updateReq := &webhooks.RequestWebhook{
		Name:               fmt.Sprintf("example-updated-%d", time.Now().UnixMilli()),
		Enabled:            true,
		URL:                "https://hooks.example.com/jamf-updated",
		ContentType:        "application/json",
		Event:              "ComputerAdded",
		AuthenticationType: "NONE",
	}
	updated, resp, err := client.Webhooks.UpdateWebhookByID(ctx, created.ID, updateReq)
	if err != nil {
		_, _ = client.Webhooks.DeleteWebhookByID(ctx, created.ID)
		log.Fatalf("UpdateWebhookByID failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Updated webhook ID: %d\n", updated.ID)
	fmt.Printf("New name: %s\n", updated.Name)
	fmt.Printf("Enabled: %v\n", updated.Enabled)

	_, _ = client.Webhooks.DeleteWebhookByID(ctx, created.ID)
	fmt.Println("Cleanup: webhook deleted")
}
