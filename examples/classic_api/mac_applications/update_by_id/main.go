package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/mac_applications"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared/models"
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

	appID := 1 // Replace with the desired Mac application ID to update
	updateReq := &mac_applications.Resource{
		General: mac_applications.SubsetGeneral{
			Name:           "Updated Mac App Name",
			Version:        "1.1",
			BundleID:       "com.apple.Safari",
			URL:            "https://www.apple.com/safari/",
			DeploymentType: "Install Automatically/Prompt Users to Install",
			Site: &models.SharedResourceSite{
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
			SelfServiceDescription: "Safari web browser (updated)",
		},
	}

	updated, _, err := jamfClient.ClassicAPI.MacApplications.UpdateByID(context.Background(), appID, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Mac Application Updated: ID=%d name=%q\n", updated.General.ID, updated.General.Name)
}

func boolPtr(b bool) *bool {
	return &b
}
