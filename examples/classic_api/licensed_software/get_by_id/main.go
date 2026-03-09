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

	licensedSoftwareID := 1 // Replace with the desired licensed software ID
	ls, _, err := jamfClient.ClassicAPI.LicensedSoftware.GetByID(context.Background(), licensedSoftwareID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	lsXML, err := xml.MarshalIndent(ls, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling licensed software data: %v", err)
	}
	fmt.Printf("Licensed Software ID %d:\n%s\n", licensedSoftwareID, string(lsXML))
}
