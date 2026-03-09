package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/static_computer_groups"
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

	groupID := "1" // Replace with the desired static computer group ID
	req := &static_computer_groups.RequestStaticGroup{
		Name:        "go-sdk-v2-Static-Group-Updated",
		Description: "Updated via static_computer_groups API",
		Assignments: []string{},
	}

	result, _, err := jamfClient.JamfProAPI.StaticComputerGroups.UpdateByIDV2(context.Background(), groupID, req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Updated static computer group: %+v\n", result)
}
