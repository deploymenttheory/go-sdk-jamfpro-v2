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

	itemName := "Safari" // Replace with the desired dock item name
	item, _, err := jamfClient.ClassicAPI.DockItems.GetByName(context.Background(), itemName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	itemXML, err := xml.MarshalIndent(item, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling dock item data: %v", err)
	}
	fmt.Printf("Dock Item %q:\n%s\n", itemName, string(itemXML))
}
