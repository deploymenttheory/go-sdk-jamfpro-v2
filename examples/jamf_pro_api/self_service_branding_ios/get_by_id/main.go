package main

import (
	"context"
	"encoding/json"
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

	id := "1"
	if len(os.Args) > 1 {
		id = os.Args[1]
	}

	result, _, err := jamfClient.SelfServiceBrandingIOS.GetByIDV1(context.Background(), id)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Println("Self-Service Branding Mobile (iOS) by ID:")
	fmt.Println(string(out))
}
