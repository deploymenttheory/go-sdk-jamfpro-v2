// Package main demonstrates ListWebhooks — returns all webhooks from the Classic API.
//
// Run with: go run ./examples/classic_api/webhooks/list
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

	list, resp, err := client.Webhooks.ListWebhooks(context.Background())
	if err != nil {
		log.Fatalf("ListWebhooks failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Total webhooks: %d\n", list.Size)
	for _, w := range list.Results {
		fmt.Printf("  ID=%-5d  Name=%s\n", w.ID, w.Name)
	}
}
