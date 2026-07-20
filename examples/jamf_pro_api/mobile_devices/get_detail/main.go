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

	// section may be repeated to request specific sections; all are returned
	// when it is omitted.
	result, _, err := jamfClient.JamfProAPI.MobileDevices.GetDetailV2(context.Background(), map[string]string{
		"section": "GENERAL",
	})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Total Mobile Devices: %d\n\n", result.TotalCount)
	for _, device := range result.Results {
		fmt.Printf("ID: %s\n", device.MobileDeviceID)
		fmt.Printf("  Device Type: %s\n", device.DeviceType)
		if device.General != nil {
			fmt.Printf("  Display Name: %s\n", device.General.DisplayName)
			fmt.Printf("  OS Version: %s\n", device.General.OsVersion)
			// lastContactDate was added to the general section in Jamf Pro 11.30.
			fmt.Printf("  Last Contact: %s\n", device.General.LastContactDate)
		}
		fmt.Println()
	}
}
