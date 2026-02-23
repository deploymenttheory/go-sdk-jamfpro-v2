package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

	computerID := "1"
	attachmentID := "100"

	data, _, err := jamfClient.ComputerInventory.GetAttachmentByIDV3(context.Background(), computerID, attachmentID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	outputPath := "downloaded_attachment.pdf"
	err = os.WriteFile(outputPath, data, 0644)
	if err != nil {
		log.Fatalf("Failed to write attachment file: %v", err)
	}

	fmt.Printf("Attachment downloaded successfully to %s (%d bytes)\n", outputPath, len(data))
}
