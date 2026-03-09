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

	groupName := "All Mobile Devices" // Replace with the desired mobile device group name
	group, _, err := jamfClient.ClassicAPI.MobileDeviceGroups.GetByName(context.Background(), groupName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	groupXML, err := xml.MarshalIndent(group, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling mobile device group data: %v", err)
	}
	fmt.Printf("Mobile Device Group %q:\n%s\n", groupName, string(groupXML))
}
