package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	authConfig, err := client.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Example usage of GetAllowedFileExtensionByID
	extensionID := 1 // Replace with the desired allowed file extension ID
	ext, _, err := jamfClient.AllowedFileExtensions.GetAllowedFileExtensionByID(context.Background(), extensionID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Pretty print the allowed file extension details in XML
	extXML, err := xml.MarshalIndent(ext, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling allowed file extension data: %v", err)
	}
	fmt.Println("Allowed File Extension Details (ID " + strconv.Itoa(extensionID) + "):\n" + string(extXML))
}
