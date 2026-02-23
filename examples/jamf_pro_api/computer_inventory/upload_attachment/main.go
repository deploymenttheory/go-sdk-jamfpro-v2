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
	attachmentFilePath := "/path/to/attachment.pdf"

	attachmentData, err := os.ReadFile(attachmentFilePath)
	if err != nil {
		log.Fatalf("Failed to read attachment file: %v", err)
	}

	resp, err := jamfClient.ComputerInventory.UploadAttachmentByIDV3(context.Background(), computerID, attachmentData)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Attachment uploaded successfully (Status: %d)\n", resp.StatusCode)
}
