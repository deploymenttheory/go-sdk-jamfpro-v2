package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mobile_device_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
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

	createReq := &mobile_device_groups.RequestMobileDeviceGroup{
		Name:    "go-sdk-v2-mobile-smart-group",
		IsSmart: true,
		Site: &shared.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Criteria: &mobile_device_groups.CriteriaContainer{
			Size: 2,
			Criterion: []shared.SharedSubsetCriteria{
				{
					Name:       "Model",
					Priority:   0,
					AndOr:      "and",
					SearchType: "like",
					Value:      "iPhone",
				},
				{
					Name:       "Device Name",
					Priority:   1,
					AndOr:      "and",
					SearchType: "like",
					Value:      "test",
				},
			},
		},
	}

	created, _, err := jamfClient.ClassicAPI.MobileDeviceGroups.Create(context.Background(), createReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Mobile Device Group Created: ID=%d\n", created.ID)
}
