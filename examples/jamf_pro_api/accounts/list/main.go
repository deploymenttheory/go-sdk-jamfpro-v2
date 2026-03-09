package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
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

	// Example 1: List all accounts
	result, _, err := jamfClient.JamfProAPI.Accounts.ListV1(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error listing accounts: %v\n", err)
		return
	}
	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("All accounts:\n%s\n\n", string(out))

	// Example 2: List with RSQL filter for enabled accounts
	rsqlQuery := map[string]string{
		"filter": `accountStatus==Enabled`,
		"sort":   "username:asc",
	}
	result, _, err = jamfClient.JamfProAPI.Accounts.ListV1(context.Background(), rsqlQuery)
	if err != nil {
		fmt.Printf("Error listing enabled accounts: %v\n", err)
		return
	}
	out, _ = json.MarshalIndent(result, "", "    ")
	fmt.Printf("Enabled accounts:\n%s\n\n", string(out))

	// Example 3: List with complex RSQL filter
	rsqlQuery = map[string]string{
		"filter": `accountStatus==Enabled and privilegeLevel==ADMINISTRATOR and failedLoginAttempts==0`,
		"sort":   "realname:desc",
	}
	result, _, err = jamfClient.JamfProAPI.Accounts.ListV1(context.Background(), rsqlQuery)
	if err != nil {
		fmt.Printf("Error listing administrators: %v\n", err)
		return
	}
	out, _ = json.MarshalIndent(result, "", "    ")
	fmt.Printf("Enabled administrators with no failed logins:\n%s\n", string(out))
}
