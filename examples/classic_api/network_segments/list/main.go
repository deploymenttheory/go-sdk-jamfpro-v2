// Package main demonstrates ListNetworkSegments — returns all network segments from the Classic API.
//
// Run with: go run ./examples/classic_api/network_segments/list
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

	list, resp, err := client.NetworkSegments.ListNetworkSegments(context.Background())
	if err != nil {
		log.Fatalf("ListNetworkSegments failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Total network segments: %d\n", list.Size)
	for _, s := range list.Results {
		fmt.Printf("  ID=%-5d  Name=%s\n", s.ID, s.Name)
	}
}
