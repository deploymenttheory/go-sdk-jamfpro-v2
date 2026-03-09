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

	groupID := 1 // Replace with the desired computer group ID
	group, _, err := jamfClient.ClassicAPI.ComputerGroups.GetByID(context.Background(), groupID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	groupXML, err := xml.MarshalIndent(group, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling computer group data: %v", err)
	}
	fmt.Printf("Computer Group ID %d:\n%s\n", groupID, string(groupXML))
}
