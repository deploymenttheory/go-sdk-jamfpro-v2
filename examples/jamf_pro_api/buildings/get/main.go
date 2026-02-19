// Package main demonstrates GetBuildingByIDV1 - retrieves a single building by ID.
//
// Run with: go run ./examples/jamf_pro_api/buildings/get
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

	building, resp, err := client.Buildings.GetBuildingByIDV1(ctx, id)
	if err != nil {
		log.Fatalf("GetBuildingByIDV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("ID: %s\n", building.ID)
	fmt.Printf("Name: %s\n", building.Name)
	fmt.Printf("Address: %s, %s, %s %s\n", building.StreetAddress1, building.City, building.StateProvince, building.ZipPostalCode)
}
