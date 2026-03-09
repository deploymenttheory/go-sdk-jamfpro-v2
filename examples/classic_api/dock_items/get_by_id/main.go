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

	itemID := 1 // Replace with the desired dock item ID
	item, _, err := jamfClient.ClassicAPI.DockItems.GetByID(context.Background(), itemID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	itemXML, err := xml.MarshalIndent(item, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling dock item data: %v", err)
	}
	fmt.Printf("Dock Item ID %d:\n%s\n", itemID, string(itemXML))
}
