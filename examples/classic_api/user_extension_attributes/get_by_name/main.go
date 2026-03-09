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

	attrName := "Department" // Replace with the desired user extension attribute name
	attr, _, err := jamfClient.ClassicAPI.UserExtensionAttributes.GetByName(context.Background(), attrName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	attrXML, err := xml.MarshalIndent(attr, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling user extension attribute data: %v", err)
	}
	fmt.Printf("User Extension Attribute %q:\n%s\n", attrName, string(attrXML))
}
