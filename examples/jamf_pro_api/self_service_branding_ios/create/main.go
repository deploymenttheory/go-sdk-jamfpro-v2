package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/self_service_branding_ios"
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

	req := &self_service_branding_ios.ResourceSelfServiceBrandingMobile{
		BrandingName:              "go-sdk-v2-Self-Service-Branding-Mobile",
		HeaderBackgroundColorCode: "#FFFFFF",
		MenuIconColorCode:         "#000000",
		BrandingNameColorCode:     "#333333",
		StatusBarTextColor:        "light",
	}

	result, _, err := jamfClient.JamfProAPI.SelfServiceBrandingIos.CreateV1(context.Background(), req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Created self-service branding mobile: ID=%s Href=%s\n", result.ID, result.Href)
}
