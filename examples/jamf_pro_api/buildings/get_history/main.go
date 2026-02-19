// Package main demonstrates GetBuildingHistoryV1 - retrieves the history object for a building.
//
// Run with: go run ./examples/jamf_pro_api/buildings/get_history
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Set BUILDING_ID or uses first from list.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()
	id := os.Getenv("BUILDING_ID")
	if id == "" {
		list, _, err := client.Buildings.ListBuildingsV1(ctx, map[string]string{"page": "0", "pageSize": "1"})
		if err != nil || len(list.Results) == 0 {
			log.Fatal("Set BUILDING_ID or ensure at least one building exists")
		}
		id = list.Results[0].ID
		fmt.Printf("Using first building ID: %s\n", id)
	}

	history, resp, err := client.Buildings.GetBuildingHistoryV1(ctx, id, nil)
	if err != nil {
		log.Fatalf("GetBuildingHistoryV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Total history entries: %d\n", history.TotalCount)
	for i, entry := range history.Results {
		if i >= 10 {
			fmt.Printf("... and %d more\n", history.TotalCount-10)
			break
		}
		fmt.Printf("  [%s] %s: %s (%s)\n", entry.Date, entry.Username, entry.Note, entry.Details)
	}
}
