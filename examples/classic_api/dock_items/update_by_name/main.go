package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/dock_items"
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
	updateReq := &dock_items.Request{
		Name:     "Updated Safari",
		Type:     "App",
		Path:     "/Applications/Safari.app",
		Contents: "",
	}

	updated, _, err := jamfClient.ClassicDockItems.UpdateByName(context.Background(), itemName, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Dock Item Updated: ID=%d name=%q\n", updated.ID, updated.Name)
}
