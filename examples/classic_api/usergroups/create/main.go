package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/usergroups"
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

	createReq := &usergroups.RequestUserGroup{
		Name:             "go-sdk-v2-smart-user-group",
		IsSmart:          true,
		IsNotifyOnChange: false,
		Site: &shared.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Criteria: &usergroups.CriteriaContainer{
			Size: 2,
			Criterion: []shared.SharedSubsetCriteria{
				{
					Name:       "Email Address",
					Priority:   0,
					AndOr:      "and",
					SearchType: "like",
					Value:      "@example.com",
				},
				{
					Name:       "Username",
					Priority:   1,
					AndOr:      "and",
					SearchType: "like",
					Value:      "admin",
				},
			},
		},
	}

	created, _, err := jamfClient.ClassicUserGroups.Create(context.Background(), createReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("User Group Created: ID=%d\n", created.ID)
}
