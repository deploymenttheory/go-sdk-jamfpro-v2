// Package main demonstrates CreateSite — creates a new site via the Classic API.
//
// Run with: go run ./examples/classic_api/sites/create
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates a site then deletes it.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/sites"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	req := &sites.RequestSite{
		Name: fmt.Sprintf("example-site-%d", time.Now().UnixMilli()),
	}

	created, resp, err := client.Sites.CreateSite(ctx, req)
	if err != nil {
		log.Fatalf("CreateSite failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Created site ID: %d\n", created.ID)
	fmt.Printf("Name: %s\n", created.Name)

	// Cleanup: delete the created site
	if _, err := client.Sites.DeleteSiteByID(ctx, created.ID); err != nil {
		fmt.Printf("Note: cleanup delete failed: %v\n", err)
	} else {
		fmt.Println("Cleanup: site deleted")
	}
}
