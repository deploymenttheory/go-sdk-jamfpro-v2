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

	profileID := 1 // Replace with the desired mobile device enrollment profile ID
	profile, _, err := jamfClient.ClassicAPI.MobileDeviceEnrollmentProfiles.GetByID(context.Background(), profileID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	profileXML, err := xml.MarshalIndent(profile, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling mobile device enrollment profile data: %v", err)
	}
	fmt.Printf("Mobile Device Enrollment Profile ID %d:\n%s\n", profileID, string(profileXML))
}
