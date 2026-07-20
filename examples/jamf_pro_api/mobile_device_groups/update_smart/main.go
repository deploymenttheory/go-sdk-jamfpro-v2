package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/mobile_device_groups"
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

	groupID := "1"

	siteID := "-1"
	req := &mobile_device_groups.RequestSmartMobileDeviceGroup{
		Name:        "Supervised iPads (updated)",
		Description: "Updated via the SDK",
		SiteId:      &siteID,
		Criteria: []mobile_device_groups.CriteriaJamfProAPI{
			{
				Name:       "Supervised",
				Priority:   0,
				AndOr:      "and",
				SearchType: "is",
				Value:      "true",
			},
		},
	}

	result, _, err := jamfClient.JamfProAPI.MobileDeviceGroups.UpdateSmartByIDV2(context.Background(), groupID, req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Updated smart group: %s\n", result.ID)
	fmt.Printf("  Name: %s\n", result.Name)
}
