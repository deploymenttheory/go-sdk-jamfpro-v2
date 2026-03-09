package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/dock_items"
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
	updateReq := &dock_items.Request{
		Name:     "Updated App Name",
		Type:     "App",
		Path:     "/Applications/MyApp.app",
		Contents: "",
	}

	updated, _, err := jamfClient.ClassicAPI.DockItems.UpdateByID(context.Background(), itemID, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Dock Item Updated: ID=%d name=%q\n", updated.ID, updated.Name)
}
