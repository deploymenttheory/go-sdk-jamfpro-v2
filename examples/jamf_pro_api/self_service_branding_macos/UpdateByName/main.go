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

	brandingName := "Corporate Branding" // Replace with the desired branding configuration name
	updateReq := &self_service_branding_macos.ResourceSelfServiceBrandingMacOS{
		ApplicationName:       "Self Service",
		BrandingName:          "Corporate Branding Updated",
		BrandingNameSecondary: "IT Department",
		HomeHeading:           "Welcome Back",
		HomeSubheading:        "Choose an item below",
	}

	result, _, err := jamfClient.SelfServiceBrandingMacOS.UpdateByName(context.Background(), brandingName, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Updated self-service branding macOS: %+v\n", result)
}
