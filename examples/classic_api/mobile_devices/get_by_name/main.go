package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	authConfig, err := client.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	deviceName := "iPhone-01" // Replace with the desired mobile device name
	device, _, err := jamfClient.ClassicMobileDevices.GetByName(context.Background(), deviceName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	deviceXML, err := xml.MarshalIndent(device, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling mobile device data: %v", err)
	}
	fmt.Printf("Mobile Device %q:\n%s\n", deviceName, string(deviceXML))
}
