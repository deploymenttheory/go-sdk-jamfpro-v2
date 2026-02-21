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

	syncStates, resp, err := jamfClient.DeviceEnrollments.GetAllSyncStatesV1(ctx)
	if err != nil {
		log.Fatalf("Failed to get all device enrollment sync states: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Total Sync States: %d\n", len(syncStates))
	fmt.Printf("\nAll Sync States:\n")
	for i, state := range syncStates {
		fmt.Printf("  [%d] State: %s, Instance: %s, Timestamp: %s\n",
			i+1, state.SyncState, state.InstanceID, state.Timestamp)
	}
}
