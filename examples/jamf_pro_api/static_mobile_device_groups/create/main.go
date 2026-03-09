package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/static_mobile_device_groups"
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

	req := &static_mobile_device_groups.RequestStaticMobileDeviceGroup{
		Name:        "go-sdk-v2-Static-Mobile-Device-Group",
		Description: "Created via go-sdk-jamfpro-v2",
		SiteID:      "-1",
		Assignments: []static_mobile_device_groups.StaticMobileDeviceGroupAssignment{},
	}

	result, _, err := jamfClient.JamfProAPI.StaticMobileDeviceGroups.Create(context.Background(), req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Created static mobile device group: %+v\n", result)
}
