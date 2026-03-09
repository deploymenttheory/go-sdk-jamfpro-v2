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

	invitationID := "1" // Replace with the desired computer invitation ID
	invitation, _, err := jamfClient.ClassicAPI.ComputerInvitations.GetByID(context.Background(), invitationID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	invitationXML, err := xml.MarshalIndent(invitation, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling computer invitation data: %v", err)
	}
	fmt.Printf("Computer Invitation ID %s:\n%s\n", invitationID, string(invitationXML))
}
