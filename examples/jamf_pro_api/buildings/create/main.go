// Package main demonstrates CreateBuildingV1 - creates a new building.
//
// Run with: go run ./examples/jamf_pro_api/buildings/create
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

	req := &buildings.RequestBuilding{
		Name:           fmt.Sprintf("example-building-%d", time.Now().UnixMilli()),
		StreetAddress1: "100 Example St",
		City:           "Austin",
		StateProvince:  "TX",
		ZipPostalCode:  "78701",
		Country:        "United States",
	}

	result, resp, err := client.Buildings.CreateBuildingV1(ctx, req)
	if err != nil {
		log.Fatalf("CreateBuildingV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Created building ID: %s\n", result.ID)
	fmt.Printf("Href: %s\n", result.Href)

	if _, err := client.Buildings.DeleteBuildingByIDV1(ctx, result.ID); err != nil {
		fmt.Printf("Note: cleanup delete failed: %v\n", err)
	} else {
		fmt.Println("Cleanup: building deleted")
	}
}
