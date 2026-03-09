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

	name := "Corporate Branding"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	result, _, err := jamfClient.JamfProAPI.SelfServiceBrandingIos.GetByNameV1(context.Background(), name)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Println("Self-Service Branding Mobile (iOS) by Name:")
	fmt.Println(string(out))
}
