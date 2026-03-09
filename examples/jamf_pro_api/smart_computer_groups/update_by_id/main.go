package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/smart_computer_groups"
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

	groupID := "1"
	req := &smart_computer_groups.RequestSmartGroup{
		Name:        "go-sdk-v2-Smart-Computer-Group-Updated",
		Description: "Updated via Jamf Pro API SDK v2",
		Criteria: []smart_computer_groups.SubsetCriteria{
			{Name: "Computer Name", Priority: 1, AndOr: "and", SearchType: "is", Value: "*"},
		},
	}

	result, _, err := jamfClient.JamfProAPI.SmartComputerGroups.UpdateByID(context.Background(), groupID, req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Updated smart computer group: ID=%s Name=%s\n", result.ID, result.Name)
}
