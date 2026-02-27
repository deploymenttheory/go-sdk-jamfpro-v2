package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/static_user_groups"
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

	createReq := &static_user_groups.RequestStaticUserGroup{
		Name:             "go-sdk-v2-static-user-group",
		IsSmart:          false,
		IsNotifyOnChange: false,
		Site: &shared.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
	}

	created, _, err := jamfClient.ClassicStaticUserGroups.Create(context.Background(), createReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Static User Group Created: ID=%d\n", created.ID)
}
