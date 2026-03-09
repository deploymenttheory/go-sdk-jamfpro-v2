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

	profileID := 1
	profile, _, err := jamfClient.ClassicAPI.MobileDeviceConfigurationProfiles.GetByID(context.Background(), profileID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	profileXML, err := xml.MarshalIndent(profile, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling mobile device configuration profile data: %v", err)
	}
	fmt.Printf("Mobile Device Configuration Profile ID %d:\n%s\n", profileID, string(profileXML))
}
