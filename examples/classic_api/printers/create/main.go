package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/printers"
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

	// Example usage of CreatePrinter
	newPrinter := &printers.RequestPrinter{
		Name:     "go-sdk-v2-Printer",
		CUPSName: "example_printer",
		URI:      "ipp://printer.example.com/ipp",
		Location: "Example Lab",
		Model:    "Example Printer Model",
	}

	createdPrinter, _, err := jamfClient.Printers.CreatePrinter(context.Background(), newPrinter)
	if err != nil {
		fmt.Printf("Error creating printer: %v\n", err)
		return
	}
	fmt.Printf("Created Printer: %+v\n", createdPrinter)
}
