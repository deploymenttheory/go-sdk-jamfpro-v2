// Package main demonstrates CreateNetworkSegment — creates a new network segment via the Classic API.
//
// Run with: go run ./examples/classic_api/network_segments/create
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates a network segment then deletes it.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/network_segments"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	req := &network_segments.RequestNetworkSegment{
		Name:            fmt.Sprintf("example-netseg-%d", time.Now().UnixMilli()),
		StartingAddress: "10.10.10.0",
		EndingAddress:   "10.10.10.255",
	}

	created, resp, err := client.NetworkSegments.CreateNetworkSegment(ctx, req)
	if err != nil {
		log.Fatalf("CreateNetworkSegment failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Created network segment ID: %d\n", created.ID)

	// Cleanup: delete the created network segment
	if _, err := client.NetworkSegments.DeleteNetworkSegmentByID(ctx, created.ID); err != nil {
		fmt.Printf("Note: cleanup delete failed: %v\n", err)
	} else {
		fmt.Println("Cleanup: network segment deleted")
	}
}
