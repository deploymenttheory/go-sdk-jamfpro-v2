package main

import (
	"context"
	"encoding/xml"
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

	// Replace with a valid user group name from your Jamf Pro instance
	userGroupName := "All Users"

	userGroup, _, err := jamfClient.ClassicUserGroups.GetByName(context.Background(), userGroupName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	userGroupXML, err := xml.MarshalIndent(userGroup, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling user group data: %v", err)
	}
	fmt.Println("User Group:\n" + string(userGroupXML))
}
