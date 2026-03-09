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

	bundleID := "com.apple.mobilesafari" // Replace with the desired bundle ID
	app, _, err := jamfClient.ClassicAPI.MobileDeviceApplications.GetByBundleID(context.Background(), bundleID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	appXML, err := xml.MarshalIndent(app, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling mobile device application data: %v", err)
	}
	fmt.Printf("Mobile Device Application (bundle ID %q):\n%s\n", bundleID, string(appXML))
}
