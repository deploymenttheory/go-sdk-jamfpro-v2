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

	req := &mobile_device_groups.RequestStaticMobileDeviceGroup{
		Name:        "Loaner iPads",
		Description: "Static group created via the SDK",
		SiteId:      "-1",
		Assignments: []mobile_device_groups.StaticMobileDeviceGroupAssignment{
			{MobileDeviceID: "1", Selected: true},
		},
	}

	result, _, err := jamfClient.JamfProAPI.MobileDeviceGroups.CreateStaticV2(context.Background(), req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Created static group ID: %s\n", result.ID)
	fmt.Printf("Href: %s\n", result.Href)
}
