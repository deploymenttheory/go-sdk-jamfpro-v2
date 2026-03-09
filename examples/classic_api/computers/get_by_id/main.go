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

	computerID := "1" // Replace with the desired computer ID
	computer, _, err := jamfClient.ClassicAPI.Computers.GetByID(context.Background(), computerID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	computerXML, err := xml.MarshalIndent(computer, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling computer data: %v", err)
	}
	fmt.Printf("Computer ID %s:\n%s\n", computerID, string(computerXML))
}
