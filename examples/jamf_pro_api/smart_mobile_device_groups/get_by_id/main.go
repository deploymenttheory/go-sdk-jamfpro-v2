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

	groupID := "1"

	result, _, err := jamfClient.JamfProAPI.SmartMobileDeviceGroups.GetByID(context.Background(), groupID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("ID: %s\n", result.GroupID)
	fmt.Printf("Name: %s\n", result.GroupName)
	fmt.Printf("Description: %s\n", result.GroupDescription)
	fmt.Printf("Member Count: %d\n\n", result.Count)

	fmt.Printf("Criteria: %d\n", len(result.Criteria))
	for _, c := range result.Criteria {
		fmt.Printf("  [%d] %s %s %q (%s)\n", c.Priority, c.Name, c.SearchType, c.Value, c.AndOr)
	}
}
