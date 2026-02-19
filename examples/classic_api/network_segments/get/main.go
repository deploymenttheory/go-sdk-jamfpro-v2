// Package main demonstrates GetNetworkSegmentByID — retrieves a single network segment by ID.
//
// Run with: go run ./examples/classic_api/network_segments/get
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Set NETWORK_SEGMENT_ID or uses first from list.
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
	if raw := os.Getenv("NETWORK_SEGMENT_ID"); raw != "" {
		id, err = strconv.Atoi(raw)
		if err != nil {
			log.Fatalf("invalid NETWORK_SEGMENT_ID %q: %v", raw, err)
		}
	} else {
		list, _, err := client.NetworkSegments.ListNetworkSegments(ctx)
		if err != nil || len(list.Results) == 0 {
			log.Fatal("Set NETWORK_SEGMENT_ID or ensure at least one network segment exists")
		}
		id = list.Results[0].ID
		fmt.Printf("Using first network segment ID: %d\n", id)
	}

	segment, resp, err := client.NetworkSegments.GetNetworkSegmentByID(ctx, id)
	if err != nil {
		log.Fatalf("GetNetworkSegmentByID failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("ID:              %d\n", segment.ID)
	fmt.Printf("Name:            %s\n", segment.Name)
	fmt.Printf("StartingAddress: %s\n", segment.StartingAddress)
	fmt.Printf("EndingAddress:   %s\n", segment.EndingAddress)
}
