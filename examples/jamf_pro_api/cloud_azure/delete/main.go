package main

import (
	"context"
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

	resp, err := jamfClient.JamfProAPI.CloudAzure.DeleteByIDV1(context.Background(), cloudAzureID)
	if err != nil {
		fmt.Printf("Error deleting Azure Cloud IDP: %v\n", err)
		return
	}
	fmt.Printf("Successfully deleted Azure Cloud IDP (Status: %d)\n", resp.StatusCode())
}
