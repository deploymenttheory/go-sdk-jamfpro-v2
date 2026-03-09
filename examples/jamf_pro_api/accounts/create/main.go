package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/accounts"
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

	newAccount := &accounts.RequestAccount{
		Username:                  "exampleuser",
		Realname:                  "Example User",
		Email:                     "example@example.com",
		Phone:                     "555-0100",
		PlainPassword:             "SecurePassword123!",
		LdapServerID:              -1,
		SiteID:                    -1,
		AccessLevel:               "FullAccess",
		PrivilegeLevel:            "AUDITOR",
		ChangePasswordOnNextLogin: true,
		AccountStatus:             "Enabled",
		AccountType:               "DEFAULT",
	}

	result, _, err := jamfClient.JamfProAPI.Accounts.CreateV1(context.Background(), newAccount)
	if err != nil {
		fmt.Printf("Error creating account: %v\n", err)
		return
	}
	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("Created account:\n%s\n", string(out))
}
