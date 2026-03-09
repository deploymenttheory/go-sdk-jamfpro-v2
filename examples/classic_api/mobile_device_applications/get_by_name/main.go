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

	appName := "Sample iOS App" // Replace with the desired mobile device application name
	app, _, err := jamfClient.ClassicAPI.MobileDeviceApplications.GetByName(context.Background(), appName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	appXML, err := xml.MarshalIndent(app, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling mobile device application data: %v", err)
	}
	fmt.Printf("Mobile Device Application %q:\n%s\n", appName, string(appXML))
}
