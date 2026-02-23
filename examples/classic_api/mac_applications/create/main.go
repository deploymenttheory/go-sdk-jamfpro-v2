package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mac_applications"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
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

	createReq := &mac_applications.Resource{
		General: mac_applications.SubsetGeneral{
			Name:           "go-sdk-v2-mac-app",
			Version:        "1.0",
			BundleID:       "com.apple.Safari",
			URL:            "https://www.apple.com/safari/",
			DeploymentType: "Install Automatically/Prompt Users to Install",
			Site: &shared.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: mac_applications.SubsetScope{
			AllComputers: boolPtr(true),
			AllJSSUsers:  boolPtr(false),
		},
		SelfService: mac_applications.SubsetSelfService{
			InstallButtonText:      "Install",
			SelfServiceDescription: "Safari web browser",
		},
	}

	created, _, err := jamfClient.ClassicMacApplications.Create(context.Background(), createReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Mac Application Created: ID=%d\n", created.ID)
}

func boolPtr(b bool) *bool {
	return &b
}
