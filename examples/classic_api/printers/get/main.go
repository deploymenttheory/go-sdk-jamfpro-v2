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

	// Example usage of GetPrinterByID
	printerID := 1 // Replace with the desired printer ID
	printer, _, err := jamfClient.Printers.GetPrinterByID(context.Background(), printerID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Pretty print the printer details in XML
	printerXML, err := xml.MarshalIndent(printer, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling printer data: %v", err)
	}
	fmt.Println("Printer Details (ID " + strconv.Itoa(printerID) + "):\n" + string(printerXML))
}
