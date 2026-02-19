// Package main demonstrates DeleteBuildingByIDV1 - removes a building by ID.
//
// Run with: go run ./examples/jamf_pro_api/buildings/delete
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates a building then deletes it.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/buildings"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	createReq := &buildings.RequestBuilding{
		Name:    fmt.Sprintf("example-delete-%d", time.Now().UnixMilli()),
		City:    "Austin",
		Country: "United States",
	}
	created, _, err := client.Buildings.CreateBuildingV1(ctx, createReq)
	if err != nil {
		log.Fatalf("CreateBuildingV1 failed: %v", err)
	}
	id := created.ID
	fmt.Printf("Created building ID: %s\n", id)

	resp, err := client.Buildings.DeleteBuildingByIDV1(ctx, id)
	if err != nil {
		log.Fatalf("DeleteBuildingByIDV1 failed: %v", err)
	}

	fmt.Printf("Status: %d (204 = success)\n", resp.StatusCode)
	fmt.Println("Building deleted successfully")
}
