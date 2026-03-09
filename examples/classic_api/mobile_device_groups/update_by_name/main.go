package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_device_groups"
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

	groupName := "go-sdk-v2-mobile-smart-group" // Replace with the desired mobile device group name
	updateReq := &mobile_device_groups.RequestMobileDeviceGroup{
		Name:    "go-sdk-v2-mobile-smart-group-renamed",
		IsSmart: false,
	}

	updated, _, err := jamfClient.ClassicAPI.MobileDeviceGroups.UpdateByName(context.Background(), groupName, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Mobile Device Group Updated: ID=%d\n", updated.ID)
}
