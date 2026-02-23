package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/smart_mobile_device_groups"
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

	groupID := "1" // Replace with the desired smart mobile device group ID
	req := &smart_mobile_device_groups.RequestSmartMobileDeviceGroup{
		GroupName:        "go-sdk-v2-Smart-Mobile-Group-Updated",
		GroupDescription: "Updated via SDK v2",
		Criteria: []smart_mobile_device_groups.SharedSubsetCriteriaJamfProAPI{
			{Name: "Model", Priority: 0, AndOr: "and", SearchType: "is", Value: "iPhone"},
		},
	}

	result, _, err := jamfClient.SmartMobileDeviceGroups.UpdateByID(context.Background(), groupID, req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Updated smart mobile device group: %+v\n", result)
}
