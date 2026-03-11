package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/packages"
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

	packageID := "1"
	filePath := "/path/to/your/updated-package.pkg"

	req := &packages.ResourcePackage{
		ID:                   packageID,
		PackageName:          "go-sdk-v2-Package-Updated",
		CategoryID:           "-1",
		Info:                 "Updated package info",
		Notes:                "Updated via UpdateAndUpload helper",
		Priority:             15,
		FillUserTemplate:     packages.BoolPtr(true),
		FillExistingUsers:    packages.BoolPtr(false),
		RebootRequired:       packages.BoolPtr(true),
		OSInstall:            packages.BoolPtr(false),
		SuppressUpdates:      packages.BoolPtr(false),
		SuppressFromDock:     packages.BoolPtr(false),
		SuppressEula:         packages.BoolPtr(false),
		SuppressRegistration: packages.BoolPtr(false),
	}

	result, _, err := jamfClient.JamfProAPI.Packages.UpdateAndUpload(context.Background(), packageID, filePath, req)
	if err != nil {
		fmt.Printf("Error updating and uploading package: %v\n", err)
		return
	}
	fmt.Printf("Package updated and uploaded successfully with SHA3_512 verification: %+v\n", result)
}
