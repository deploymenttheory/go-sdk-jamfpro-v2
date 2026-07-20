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

	result, _, err := jamfClient.JamfProAPI.MobileDevices.ListV2(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Total Mobile Devices: %d\n\n", result.TotalCount)
	for _, device := range result.Results {
		fmt.Printf("ID: %s\n", device.ID)
		fmt.Printf("  Name: %s\n", device.Name)
		fmt.Printf("  Serial: %s\n", device.SerialNumber)
		fmt.Printf("  Type: %s\n", device.Type)
		fmt.Printf("  Management ID: %s\n", device.ManagementID)
		fmt.Println()
	}
}
