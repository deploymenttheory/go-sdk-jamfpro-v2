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

	// Convenience helper: lists groups and filters client-side by name.
	groupName := "All Managed iPads"

	result, _, err := jamfClient.JamfProAPI.SmartMobileDeviceGroups.GetByName(context.Background(), groupName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("ID: %s\n", result.GroupID)
	fmt.Printf("Name: %s\n", result.GroupName)
	fmt.Printf("Member Count: %d\n", result.Count)
}
