package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/packages"
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

	req := &packages.RequestPackage{
		PackageName:           "go-sdk-v2-Package",
		FileName:              "example.pkg",
		CategoryID:            "-1",
		Info:                  "Example package created via SDK",
		Notes:                 "Metadata only - use UploadPackageV1 to add file",
		Priority:              10,
		FillUserTemplate:      packages.BoolPtr(true),
		FillExistingUsers:     packages.BoolPtr(false),
		RebootRequired:        packages.BoolPtr(false),
		OSInstall:             packages.BoolPtr(false),
		SuppressUpdates:       packages.BoolPtr(false),
		SuppressFromDock:      packages.BoolPtr(false),
		SuppressEula:          packages.BoolPtr(false),
		SuppressRegistration: packages.BoolPtr(false),
	}

	result, _, err := jamfClient.Packages.CreatePackageV1(context.Background(), req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Created package: %+v\n", result)
}
