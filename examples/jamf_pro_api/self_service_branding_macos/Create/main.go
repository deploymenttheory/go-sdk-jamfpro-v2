package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/self_service_branding_macos"
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

	req := &self_service_branding_macos.ResourceSelfServiceBrandingMacOS{
		ApplicationName:       "Self Service",
		BrandingName:          "go-sdk-v2-Self-Service-Branding",
		BrandingNameSecondary: "IT Department",
		HomeHeading:           "Welcome",
		HomeSubheading:        "Choose an item below",
	}

	result, _, err := jamfClient.SelfServiceBrandingMacOS.Create(context.Background(), req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Created self-service branding macOS: %+v\n", result)
}
