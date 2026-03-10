package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/static_user_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/models"
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

	groupName := "Static Test Group"
	updateReq := &static_user_groups.RequestStaticUserGroup{
		Name:             "go-sdk-v2-static-user-group-updated",
		IsSmart:          false,
		IsNotifyOnChange: true,
		Site: &models.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
	}

	updated, _, err := jamfClient.ClassicAPI.StaticUserGroups.UpdateByName(context.Background(), groupName, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Static User Group Updated: ID=%d\n", updated.ID)
}
