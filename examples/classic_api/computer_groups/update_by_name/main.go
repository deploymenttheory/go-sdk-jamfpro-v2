package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/computer_groups"
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

	groupName := "go-sdk-v2-smart-group" // Replace with the desired computer group name
	updateReq := &computer_groups.RequestComputerGroup{
		Name:    "go-sdk-v2-smart-group-renamed",
		IsSmart: false,
	}

	updated, _, err := jamfClient.ClassicAPI.ComputerGroups.UpdateByName(context.Background(), groupName, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Computer Group Updated: ID=%d\n", updated.ID)
}
