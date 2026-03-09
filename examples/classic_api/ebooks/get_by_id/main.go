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

	ebookID := 1 // Replace with the desired ebook ID
	ebook, _, err := jamfClient.ClassicAPI.Ebooks.GetByID(context.Background(), ebookID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	ebookXML, err := xml.MarshalIndent(ebook, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling ebook data: %v", err)
	}
	fmt.Printf("Ebook ID %d:\n%s\n", ebookID, string(ebookXML))
}
