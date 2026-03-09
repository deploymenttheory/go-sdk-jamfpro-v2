package main

import (
	"context"
	"fmt"
	"log"
	"os"

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

	id := "1"
	if len(os.Args) > 1 {
		id = os.Args[1]
	}

	req := &self_service_branding_ios.ResourceSelfServiceBrandingMobile{
		BrandingName:              "go-sdk-v2-Self-Service-Branding-Mobile-Updated",
		HeaderBackgroundColorCode: "#F0F0F0",
		MenuIconColorCode:         "#0066CC",
		BrandingNameColorCode:     "#222222",
		StatusBarTextColor:        "dark",
	}

	result, _, err := jamfClient.JamfProAPI.SelfServiceBrandingIos.UpdateByIDV1(context.Background(), id, req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Updated self-service branding mobile: ID=%s BrandingName=%s\n", result.ID, result.BrandingName)
}
