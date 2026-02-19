// Package main demonstrates UpdateBuildingByIDV1 - updates an existing building.
//
// Run with: go run ./examples/jamf_pro_api/buildings/update
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
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
		Name:           fmt.Sprintf("example-update-%d", time.Now().UnixMilli()),
		StreetAddress1: "200 Original St",
		City:           "Austin",
		StateProvince:  "TX",
		Country:        "United States",
	}
	created, _, err := client.Buildings.CreateBuildingV1(ctx, createReq)
	if err != nil {
		log.Fatalf("CreateBuildingV1 failed: %v", err)
	}
	id := created.ID

	updateReq := &buildings.RequestBuilding{
		Name:           fmt.Sprintf("example-updated-%d", time.Now().UnixMilli()),
		StreetAddress1: "300 Updated Ave",
		City:           "Austin",
		StateProvince:  "TX",
		ZipPostalCode:  "78702",
		Country:        "United States",
	}
	result, resp, err := client.Buildings.UpdateBuildingByIDV1(ctx, id, updateReq)
	if err != nil {
		_, _ = client.Buildings.DeleteBuildingByIDV1(ctx, id)
		log.Fatalf("UpdateBuildingByIDV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Updated building ID: %s Name=%q\n", result.ID, result.Name)

	fetched, _, _ := client.Buildings.GetBuildingByIDV1(ctx, id)
	if fetched != nil {
		fmt.Printf("Verified: name=%q address=%s\n", fetched.Name, fetched.StreetAddress1)
	}

	_, _ = client.Buildings.DeleteBuildingByIDV1(ctx, id)
	fmt.Println("Cleanup: building deleted")
}
