// Package main demonstrates UpdateSiteByID — updates an existing site via the Classic API.
//
// Run with: go run ./examples/classic_api/sites/update
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
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

	// Create a site to update
	createReq := &sites.RequestSite{
		Name: fmt.Sprintf("example-update-%d", time.Now().UnixMilli()),
	}
	created, _, err := client.Sites.CreateSite(ctx, createReq)
	if err != nil {
		log.Fatalf("CreateSite failed: %v", err)
	}
	fmt.Printf("Created site ID: %d name: %s\n", created.ID, created.Name)

	// Update the site
	updateReq := &sites.RequestSite{
		Name: fmt.Sprintf("example-updated-%d", time.Now().UnixMilli()),
	}
	updated, resp, err := client.Sites.UpdateSiteByID(ctx, created.ID, updateReq)
	if err != nil {
		_, _ = client.Sites.DeleteSiteByID(ctx, created.ID)
		log.Fatalf("UpdateSiteByID failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Updated site ID: %d\n", updated.ID)
	fmt.Printf("New name: %s\n", updated.Name)

	_, _ = client.Sites.DeleteSiteByID(ctx, created.ID)
	fmt.Println("Cleanup: site deleted")
}
