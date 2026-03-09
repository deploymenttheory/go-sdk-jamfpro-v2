package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/users"
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

	updateReq := &users.RequestUser{
		Name:     "go-sdk-v2-user",
		FullName: "Go SDK v2 Test User (Updated by Name)",
		Email:    "go-sdk-v2-user@example.com",
		Sites: []shared.SharedResourceSite{
			{ID: -1, Name: "None"},
		},
	}

	updated, _, err := jamfClient.ClassicAPI.Users.UpdateByName(context.Background(), "go-sdk-v2-user", updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("User Updated: ID=%d FullName=%s\n", updated.ID, updated.FullName)
}
