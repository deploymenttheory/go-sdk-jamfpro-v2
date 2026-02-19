// Package main demonstrates DeleteBuildingsByIDV1 - deletes multiple buildings by their IDs.
//
// Run with: go run ./examples/jamf_pro_api/buildings/delete_multiple
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates two buildings then bulk deletes them.
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

	// Create two buildings
	ids := make([]string, 0, 2)
	for i := 0; i < 2; i++ {
		req := &buildings.RequestBuilding{
			Name:    fmt.Sprintf("example-bulk-%d-%d", i, time.Now().UnixMilli()),
			City:    "Austin",
			Country: "United States",
		}
		created, _, err := client.Buildings.CreateBuildingV1(ctx, req)
		if err != nil {
			log.Fatalf("CreateBuildingV1 %d failed: %v", i, err)
		}
		ids = append(ids, created.ID)
		fmt.Printf("Created building ID: %s\n", created.ID)
	}

	// Bulk delete
	bulkReq := &buildings.DeleteBuildingsByIDRequest{IDs: ids}
	resp, err := client.Buildings.DeleteBuildingsByIDV1(ctx, bulkReq)
	if err != nil {
		log.Fatalf("DeleteBuildingsByIDV1 failed: %v", err)
	}

	fmt.Printf("Status: %d (204 = success)\n", resp.StatusCode)
	fmt.Printf("Deleted %d buildings: %v\n", len(ids), ids)
}
