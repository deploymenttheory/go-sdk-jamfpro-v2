// Package main demonstrates GetSiteByID — retrieves a single site by ID.
//
// Run with: go run ./examples/classic_api/sites/get
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Set SITE_ID or uses first from list.
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
	if raw := os.Getenv("SITE_ID"); raw != "" {
		id, err = strconv.Atoi(raw)
		if err != nil {
			log.Fatalf("invalid SITE_ID %q: %v", raw, err)
		}
	} else {
		list, _, err := client.Sites.ListSites(ctx)
		if err != nil || len(list.Results) == 0 {
			log.Fatal("Set SITE_ID or ensure at least one site exists")
		}
		id = list.Results[0].ID
		fmt.Printf("Using first site ID: %d\n", id)
	}

	site, resp, err := client.Sites.GetSiteByID(ctx, id)
	if err != nil {
		log.Fatalf("GetSiteByID failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("ID:   %d\n", site.ID)
	fmt.Printf("Name: %s\n", site.Name)
}
