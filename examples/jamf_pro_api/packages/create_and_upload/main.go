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

	filePath := "/path/to/your/package.pkg"

	req := &packages.RequestPackage{
		PackageName:          "go-sdk-v2-Package-Upload",
		CategoryID:           "-1",
		Info:                 "Example package with file upload",
		Notes:                "Created and uploaded via CreateAndUpload helper",
		Priority:             10,
		FillUserTemplate:     packages.BoolPtr(true),
		FillExistingUsers:    packages.BoolPtr(false),
		RebootRequired:       packages.BoolPtr(false),
		OSInstall:            packages.BoolPtr(false),
		SuppressUpdates:      packages.BoolPtr(false),
		SuppressFromDock:     packages.BoolPtr(false),
		SuppressEula:         packages.BoolPtr(false),
		SuppressRegistration: packages.BoolPtr(false),
	}

	result, _, err := jamfClient.JamfProAPI.Packages.CreateAndUpload(context.Background(), filePath, req)
	if err != nil {
		fmt.Printf("Error creating and uploading package: %v\n", err)
		return
	}
	fmt.Printf("Package created and uploaded successfully with SHA3_512 verification: %+v\n", result)
}
