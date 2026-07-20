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

	mobileDeviceID := "1"

	result, _, err := jamfClient.JamfProAPI.MobileDevices.GetPairedDevicesByIDV2(context.Background(), mobileDeviceID, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Paired Devices: %d\n\n", result.TotalCount)
	for _, device := range result.Results {
		fmt.Printf("ID: %s\n", device.MobileDeviceID)
		if device.General != nil {
			fmt.Printf("  Display Name: %s\n", device.General.DisplayName)
			fmt.Printf("  Last Contact: %s\n", device.General.LastContactDate)
		}
		fmt.Println()
	}
}
