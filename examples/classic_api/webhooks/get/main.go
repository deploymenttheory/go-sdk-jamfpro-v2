// Package main demonstrates GetWebhookByID — retrieves a single webhook by ID.
//
// Run with: go run ./examples/classic_api/webhooks/get
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Set WEBHOOK_ID or uses first from list.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	var id int
	if raw := os.Getenv("WEBHOOK_ID"); raw != "" {
		id, err = strconv.Atoi(raw)
		if err != nil {
			log.Fatalf("invalid WEBHOOK_ID %q: %v", raw, err)
		}
	} else {
		list, _, err := client.Webhooks.ListWebhooks(ctx)
		if err != nil || len(list.Results) == 0 {
			log.Fatal("Set WEBHOOK_ID or ensure at least one webhook exists")
		}
		id = list.Results[0].ID
		fmt.Printf("Using first webhook ID: %d\n", id)
	}

	webhook, resp, err := client.Webhooks.GetWebhookByID(ctx, id)
	if err != nil {
		log.Fatalf("GetWebhookByID failed: %v", err)
	}

	fmt.Printf("Status:  %d\n", resp.StatusCode)
	fmt.Printf("ID:      %d\n", webhook.ID)
	fmt.Printf("Name:    %s\n", webhook.Name)
	fmt.Printf("Enabled: %v\n", webhook.Enabled)
	fmt.Printf("URL:     %s\n", webhook.URL)
	fmt.Printf("Event:   %s\n", webhook.Event)
}
