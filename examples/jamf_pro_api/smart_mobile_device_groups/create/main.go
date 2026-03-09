package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/smart_mobile_device_groups"
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

	req := &smart_mobile_device_groups.RequestSmartMobileDeviceGroup{
		GroupName:        "go-sdk-v2-Smart-Mobile-Group",
		GroupDescription: "Created via SDK v2",
		Criteria: []smart_mobile_device_groups.SharedSubsetCriteriaJamfProAPI{
			{Name: "Model", Priority: 0, AndOr: "and", SearchType: "is", Value: "iPhone"},
		},
	}

	result, _, err := jamfClient.JamfProAPI.SmartMobileDeviceGroups.Create(context.Background(), req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Created smart mobile device group: %+v\n", result)
}
