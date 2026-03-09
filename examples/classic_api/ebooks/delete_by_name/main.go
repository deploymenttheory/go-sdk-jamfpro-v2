package main

import (
	"context"
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

	ebookName := "Sample Ebook" // Replace with the desired ebook name to delete
	_, err = jamfClient.ClassicAPI.Ebooks.DeleteByName(context.Background(), ebookName)
	if err != nil {
		fmt.Printf("Error deleting ebook by name: %v\n", err)
		return
	}
	fmt.Printf("Ebook %q deleted successfully\n", ebookName)
}
