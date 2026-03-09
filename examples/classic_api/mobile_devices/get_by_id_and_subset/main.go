package main

import (
	"context"
	"encoding/xml"
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

	deviceID := "1"   // Replace with the desired mobile device ID
	subset := "General" // Subset: General, Location, Purchasing, etc.
	device, _, err := jamfClient.ClassicAPI.MobileDevices.GetByIDAndDataSubset(context.Background(), deviceID, subset)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	deviceXML, err := xml.MarshalIndent(device, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling mobile device data: %v", err)
	}
	fmt.Printf("Mobile Device ID %s (subset=%s):\n%s\n", deviceID, subset, string(deviceXML))
}
