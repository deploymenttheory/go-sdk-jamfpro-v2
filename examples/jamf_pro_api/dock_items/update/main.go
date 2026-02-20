package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/dock_items"
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

	dockItemID := "1" // Replace with the desired dock item ID
	updateReq := &dock_items.RequestDockItem{
		Name: "go-sdk-v2-Dock-Item-Updated",
		Path: "/Applications/Google Chrome.app",
		Type: dock_items.TypeApp,
	}

	result, _, err := jamfClient.DockItems.UpdateDockItemByIDV1(context.Background(), dockItemID, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Updated dock item: %+v\n", result)
}
