package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/smart_mobile_device_groups"
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

	siteID := "-1"
	req := &smart_mobile_device_groups.RequestSmartMobileDeviceGroup{
		GroupName:        "Supervised iPads",
		GroupDescription: "All supervised iPads, created via the SDK",
		SiteId:           &siteID,
		Criteria: []smart_mobile_device_groups.SharedSubsetCriteriaJamfProAPI{
			{
				Name:       "Supervised",
				Priority:   0,
				AndOr:      "and",
				SearchType: "is",
				Value:      "true",
			},
		},
	}

	result, _, err := jamfClient.JamfProAPI.SmartMobileDeviceGroups.Create(context.Background(), req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Created smart mobile device group ID: %s\n", result.ID)
	fmt.Printf("Href: %s\n", result.Href)
}
