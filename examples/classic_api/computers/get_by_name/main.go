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

	computerName := "MacBook-Pro-01" // Replace with the desired computer name
	computer, _, err := jamfClient.ClassicAPI.Computers.GetByName(context.Background(), computerName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	computerXML, err := xml.MarshalIndent(computer, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling computer data: %v", err)
	}
	fmt.Printf("Computer %q:\n%s\n", computerName, string(computerXML))
}
