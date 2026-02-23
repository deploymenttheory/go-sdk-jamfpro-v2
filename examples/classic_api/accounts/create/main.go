package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/accounts"
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

	createReq := &accounts.RequestAccount{
		Name:         "example-user",
		FullName:     "Example User",
		Email:        "example-user@example.com",
		EmailAddress: "example-user@example.com",
		Password:     "SecurePassword123!",
		AccessLevel:  "Full Access",
		PrivilegeSet: "Administrator",
		Enabled:      "Enabled",
	}

	created, _, err := jamfClient.ClassicAccounts.Create(context.Background(), createReq)
	if err != nil {
		fmt.Printf("Error creating account: %v\n", err)
		return
	}
	fmt.Printf("Created Account: ID=%d Name=%s\n", created.ID, created.Name)
}
