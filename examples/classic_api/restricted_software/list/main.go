// Package main demonstrates List — returns all restricted software from the Classic API.
//
// Run with: go run ./examples/classic_api/restricted_software/list
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

	list, resp, err := client.ClassicRestrictedSoftware.List(context.Background())
	if err != nil {
		log.Fatalf("List failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Total restricted software items: %d\n", list.Size)
	for _, item := range list.Results {
		fmt.Printf("  ID=%-5d  Name=%s\n", item.ID, item.Name)
	}
}
