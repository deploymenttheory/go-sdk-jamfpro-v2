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

	ebookName := "Sample Ebook" // Replace with the desired ebook name
	ebook, _, err := jamfClient.ClassicEbooks.GetByName(context.Background(), ebookName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	ebookXML, err := xml.MarshalIndent(ebook, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling ebook data: %v", err)
	}
	fmt.Printf("Ebook %q:\n%s\n", ebookName, string(ebookXML))
}
