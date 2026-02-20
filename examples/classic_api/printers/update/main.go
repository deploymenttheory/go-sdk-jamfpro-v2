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

	// Example usage of UpdatePrinterByID — replace with the desired printer ID and updated payload
	printerID := 1
	updateReq := &printers.RequestPrinter{
		Name:     "go-sdk-v2-Printer-Updated",
		CUPSName: "example_updated_printer",
		URI:      "ipp://printer.example.com/ipp-updated",
		Location: "Updated Lab",
		Model:    "Updated Printer Model",
	}

	updatedPrinter, _, err := jamfClient.Printers.UpdatePrinterByID(context.Background(), printerID, updateReq)
	if err != nil {
		fmt.Printf("Error updating printer by ID: %v\n", err)
		return
	}
	fmt.Printf("Updated Printer: %+v\n", updatedPrinter)
}
