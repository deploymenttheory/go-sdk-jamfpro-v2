// Package main demonstrates CreateWebhook — creates a new webhook via the Classic API.
//
// Run with: go run ./examples/classic_api/webhooks/create
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

	req := &webhooks.RequestWebhook{
		Name:               fmt.Sprintf("example-webhook-%d", time.Now().UnixMilli()),
		Enabled:            false,
		URL:                "https://hooks.example.com/jamf",
		ContentType:        "application/json",
		Event:              "ComputerAdded",
		AuthenticationType: "NONE",
	}

	created, resp, err := client.Webhooks.CreateWebhook(ctx, req)
	if err != nil {
		log.Fatalf("CreateWebhook failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Created webhook ID: %d\n", created.ID)
	fmt.Printf("Name: %s\n", created.Name)

	// Cleanup: delete the created webhook
	if _, err := client.Webhooks.DeleteWebhookByID(ctx, created.ID); err != nil {
		fmt.Printf("Note: cleanup delete failed: %v\n", err)
	} else {
		fmt.Println("Cleanup: webhook deleted")
	}
}
