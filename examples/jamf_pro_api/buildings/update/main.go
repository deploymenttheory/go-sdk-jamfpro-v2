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

	buildingID := "1" // Replace with the desired building ID
	updateReq := &buildings.RequestBuilding{
		Name:           "go-sdk-v2-Building-Updated",
		StreetAddress1: "300 Updated Ave",
		City:           "Austin",
		StateProvince:  "TX",
		ZipPostalCode:  "78702",
		Country:        "United States",
	}

	result, _, err := jamfClient.Buildings.UpdateByIDV1(context.Background(), buildingID, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Updated building: %+v\n", result)
}
