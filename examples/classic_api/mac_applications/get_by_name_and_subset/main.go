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

	appName := "Sample Mac App"     // Replace with the desired Mac application name
	subset := "General"            // Subset values: General, Scope, SelfService, VPPCodes, VPP
	app, _, err := jamfClient.ClassicAPI.MacApplications.GetByNameAndSubset(context.Background(), appName, subset)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	appXML, err := xml.MarshalIndent(app, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling Mac application data: %v", err)
	}
	fmt.Printf("Mac Application %q (subset %s):\n%s\n", appName, subset, string(appXML))
}
