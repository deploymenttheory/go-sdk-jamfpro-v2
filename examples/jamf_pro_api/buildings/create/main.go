// Package main demonstrates CreateBuildingV1 - creates a new building.
//
// Run with: go run ./examples/jamf_pro_api/buildings/create
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/buildings"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	authConfig, err := client.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	req := &buildings.RequestBuilding{
		Name:           "go-sdk-v2-Building",
		StreetAddress1: "100 Example St",
		City:           "Austin",
		StateProvince:  "TX",
		ZipPostalCode:  "78701",
		Country:        "United States",
	}

	result, _, err := jamfClient.Buildings.CreateV1(context.Background(), req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Created building: %+v\n", result)
}
