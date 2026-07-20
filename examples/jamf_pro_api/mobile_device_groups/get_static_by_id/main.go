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

	groupID := "8"

	result, _, err := jamfClient.JamfProAPI.MobileDeviceGroups.GetStaticByIDV2(context.Background(), groupID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("ID: %s\n", result.ID)
	fmt.Printf("Name: %s\n", result.Name)
	fmt.Printf("Description: %s\n", result.Description)
	fmt.Printf("Member Count: %d\n", result.Count)
}
