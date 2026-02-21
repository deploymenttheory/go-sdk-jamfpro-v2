package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
)

func main() {
	authConfig := client.AuthConfigFromEnv()
	if err := authConfig.Validate(); err != nil {
		log.Fatalf("Invalid auth config: %v", err)
	}

	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	ctx := context.Background()
	enrollmentID := "1"

	syncStates, resp, err := jamfClient.DeviceEnrollments.GetSyncStatesV1(ctx, enrollmentID)
	if err != nil {
		log.Fatalf("Failed to get device enrollment sync states: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Sync States:\n")
	for _, state := range syncStates {
		fmt.Printf("  - State: %s, Instance: %s, Timestamp: %s\n", state.SyncState, state.InstanceID, state.Timestamp)
	}

	latestSync, latestResp, err := jamfClient.DeviceEnrollments.GetLatestSyncStateV1(ctx, enrollmentID)
	if err != nil {
		log.Fatalf("Failed to get latest sync state: %v", err)
	}

	fmt.Printf("\nLatest Sync State (Status Code: %d):\n", latestResp.StatusCode)
	fmt.Printf("  State: %s\n", latestSync.SyncState)
	fmt.Printf("  Instance: %s\n", latestSync.InstanceID)
	fmt.Printf("  Timestamp: %s\n", latestSync.Timestamp)
}
