// Package main demonstrates DeleteNetworkSegmentByID — removes a network segment via the Classic API.
//
// Run with: go run ./examples/classic_api/network_segments/delete
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

	// Create a network segment to delete
	createReq := &network_segments.RequestNetworkSegment{
		Name:            fmt.Sprintf("example-delete-%d", time.Now().UnixMilli()),
		StartingAddress: "10.30.30.0",
		EndingAddress:   "10.30.30.255",
	}
	created, _, err := client.NetworkSegments.CreateNetworkSegment(ctx, createReq)
	if err != nil {
		log.Fatalf("CreateNetworkSegment failed: %v", err)
	}
	fmt.Printf("Created network segment ID: %d\n", created.ID)

	resp, err := client.NetworkSegments.DeleteNetworkSegmentByID(ctx, created.ID)
	if err != nil {
		log.Fatalf("DeleteNetworkSegmentByID failed: %v", err)
	}

	fmt.Printf("Status: %d (200 = success)\n", resp.StatusCode)
	fmt.Println("Network segment deleted successfully")
}
