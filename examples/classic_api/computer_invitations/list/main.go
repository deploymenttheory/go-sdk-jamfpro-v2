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

	list, _, err := jamfClient.ClassicComputerInvitations.List(context.Background())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	listXML, err := xml.MarshalIndent(list, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling computer invitations data: %v", err)
	}
	fmt.Println("Computer Invitations List:\n" + string(listXML))
}
