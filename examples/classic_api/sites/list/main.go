// Package main demonstrates ListSites — returns all sites from the Classic API.
//
// Run with: go run ./examples/classic_api/sites/list
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

	list, resp, err := client.Sites.ListSites(context.Background())
	if err != nil {
		log.Fatalf("ListSites failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Total sites: %d\n", list.Size)
	for _, s := range list.Results {
		fmt.Printf("  ID=%-5d  Name=%s\n", s.ID, s.Name)
	}
}
