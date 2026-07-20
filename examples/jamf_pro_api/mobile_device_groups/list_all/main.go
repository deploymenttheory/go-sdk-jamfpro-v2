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

	// Returns both smart and static groups in a single summary list.
	result, _, err := jamfClient.JamfProAPI.MobileDeviceGroups.ListAllV2(context.Background())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Total Groups: %d\n\n", len(result))
	for _, group := range result {
		kind := "static"
		if group.IsSmartGroup {
			kind = "smart"
		}
		fmt.Printf("ID: %d\n", group.ID)
		fmt.Printf("  Name: %s\n", group.Name)
		fmt.Printf("  Kind: %s\n", kind)
		fmt.Println()
	}
}
