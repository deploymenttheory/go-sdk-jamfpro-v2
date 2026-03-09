package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	authConfig, err := jamfpro.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Replace "1" with the actual Azure Cloud IDP ID
	cloudAzureID := "1"

	// Get existing configuration first (UpdateByIDV1 requires full body)
	existing, _, err := jamfClient.JamfProAPI.CloudAzure.GetByIDV1(context.Background(), cloudAzureID)
	if err != nil {
		fmt.Printf("Error retrieving existing Azure Cloud IDP: %v\n", err)
		return
	}

	// Update fields as needed
	existing.CloudIdPCommon.DisplayName = "Updated Azure Cloud IDP"
	existing.Server.Enabled = true
	existing.Server.SearchTimeout = 60

	result, _, err := jamfClient.JamfProAPI.CloudAzure.UpdateByIDV1(context.Background(), cloudAzureID, existing)
	if err != nil {
		fmt.Printf("Error updating Azure Cloud IDP: %v\n", err)
		return
	}
	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("Updated Azure Cloud IDP:\n%s\n", string(out))
}
