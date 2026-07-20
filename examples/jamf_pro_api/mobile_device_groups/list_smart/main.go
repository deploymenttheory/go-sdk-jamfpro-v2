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

	result, _, err := jamfClient.JamfProAPI.MobileDeviceGroups.ListSmartV2(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Total Smart Groups: %d\n\n", result.TotalCount)
	for _, group := range result.Results {
		fmt.Printf("ID: %s\n", group.ID)
		fmt.Printf("  Name: %s\n", group.Name)
		fmt.Printf("  Site ID: %s\n", group.SiteId)
		fmt.Printf("  Member Count: %d\n", group.Count)
		fmt.Println()
	}
}
