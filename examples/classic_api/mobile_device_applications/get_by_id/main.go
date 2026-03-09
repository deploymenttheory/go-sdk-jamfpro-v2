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

	appID := 1 // Replace with the desired mobile device application ID
	app, _, err := jamfClient.ClassicAPI.MobileDeviceApplications.GetByID(context.Background(), appID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	appXML, err := xml.MarshalIndent(app, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling mobile device application data: %v", err)
	}
	fmt.Printf("Mobile Device Application ID %d:\n%s\n", appID, string(appXML))
}
