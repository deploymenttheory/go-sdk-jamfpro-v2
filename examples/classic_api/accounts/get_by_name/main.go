package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
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

	accountName := "testuser"
	account, _, err := jamfClient.ClassicAPI.Accounts.GetByName(context.Background(), accountName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	accountXML, err := xml.MarshalIndent(account, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling account data: %v", err)
	}
	fmt.Printf("Account Details (name %q):\n%s\n", accountName, string(accountXML))
}
