package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/venafi"
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

	revocationEnabled := true
	req := &venafi.ResourceVenafi{
		Name:               "Venafi Certificate Authority",
		ProxyAddress:       "localhost:9443",
		ClientID:           "jamf-pro",
		RefreshToken:       "YOUR_REFRESH_TOKEN",
		RevocationEnabled: &revocationEnabled,
	}

	result, _, err := jamfClient.Venafi.Create(context.Background(), req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Created Venafi PKI configuration: %+v\n", result)
}
