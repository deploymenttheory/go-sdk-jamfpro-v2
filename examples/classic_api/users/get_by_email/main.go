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

	user, _, err := jamfClient.ClassicAPI.Users.GetByEmail(context.Background(), "admin@example.com")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	userXML, err := xml.MarshalIndent(user, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling user data: %v", err)
	}
	fmt.Println("User:\n" + string(userXML))
}
