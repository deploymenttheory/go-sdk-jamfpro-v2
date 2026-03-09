package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/licensed_software"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/shared"
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

	licensedSoftwareID := 1 // Replace with the desired licensed software ID
	updateReq := &licensed_software.Resource{
		General: licensed_software.SubsetGeneral{
			Name:      "go-sdk-v2-licensed-software-updated",
			Publisher: "SDK Test Publisher",
			Platform:  "Mac",
			Site:      shared.SharedResourceSite{ID: -1, Name: "None"},
		},
	}

	updated, _, err := jamfClient.ClassicAPI.LicensedSoftware.UpdateByID(context.Background(), licensedSoftwareID, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Licensed Software Updated: ID=%d\n", updated.ID)
}
