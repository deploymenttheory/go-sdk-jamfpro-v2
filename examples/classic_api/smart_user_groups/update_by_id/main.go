package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/smart_user_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
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

	groupID := 1
	updateReq := &smart_user_groups.RequestSmartUserGroup{
		Name:             "go-sdk-v2-smart-user-group-updated",
		IsSmart:          true,
		IsNotifyOnChange: true,
		Site: &shared.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Criteria: &smart_user_groups.CriteriaContainer{
			Size: 1,
			Criterion: []shared.SharedSubsetCriteria{
				{
					Name:       "Email Address",
					Priority:   0,
					AndOr:      "and",
					SearchType: "like",
					Value:      "@example.com",
				},
			},
		},
	}

	updated, _, err := jamfClient.ClassicSmartUserGroups.UpdateByID(context.Background(), groupID, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Smart User Group Updated: ID=%d\n", updated.ID)
}
