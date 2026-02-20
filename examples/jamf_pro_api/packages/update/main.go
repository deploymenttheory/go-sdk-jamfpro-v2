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

	packageID := "1" // Replace with the desired package ID
	fetched, _, err := jamfClient.Packages.GetPackageByIDV1(context.Background(), packageID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	updateReq := fetched
	updateReq.PackageName = "go-sdk-v2-Package-Updated"
	updateReq.Info = "Updated metadata via SDK"
	updateReq.Notes = "After update"
	updateReq.Priority = 15
	updateReq.FillUserTemplate = packages.BoolPtr(true)
	updateReq.FillExistingUsers = packages.BoolPtr(true)

	result, _, err := jamfClient.Packages.UpdatePackageByIDV1(context.Background(), packageID, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Updated package: %+v\n", result)
}
