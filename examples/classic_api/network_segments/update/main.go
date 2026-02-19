// Package main demonstrates UpdateNetworkSegmentByID — updates an existing network segment via the Classic API.
//
// Run with: go run ./examples/classic_api/network_segments/update
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
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

	// Create a network segment to update
	createReq := &network_segments.RequestNetworkSegment{
		Name:            fmt.Sprintf("example-update-%d", time.Now().UnixMilli()),
		StartingAddress: "10.20.20.0",
		EndingAddress:   "10.20.20.255",
	}
	created, _, err := client.NetworkSegments.CreateNetworkSegment(ctx, createReq)
	if err != nil {
		log.Fatalf("CreateNetworkSegment failed: %v", err)
	}
	fmt.Printf("Created network segment ID: %d\n", created.ID)

	// Update the network segment
	updateReq := &network_segments.RequestNetworkSegment{
		Name:            fmt.Sprintf("example-updated-%d", time.Now().UnixMilli()),
		StartingAddress: "10.20.20.0",
		EndingAddress:   "10.20.20.128",
	}
	updated, resp, err := client.NetworkSegments.UpdateNetworkSegmentByID(ctx, created.ID, updateReq)
	if err != nil {
		_, _ = client.NetworkSegments.DeleteNetworkSegmentByID(ctx, created.ID)
		log.Fatalf("UpdateNetworkSegmentByID failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Updated network segment ID: %d\n", updated.ID)

	_, _ = client.NetworkSegments.DeleteNetworkSegmentByID(ctx, created.ID)
	fmt.Println("Cleanup: network segment deleted")
}
