package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/ebooks"
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

	ebookID := 1 // Replace with the desired ebook ID to update
	updateReq := &ebooks.Resource{
		General: ebooks.SubsetGeneral{
			Name:            "Updated Ebook Name",
			Author:          "Updated Author",
			Version:         "1.1",
			Free:            true,
			URL:             "https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf",
			DeploymentType:  "Install Automatically/Prompt Users to Install",
			FileType:        "PDF",
			DeployAsManaged: false,
			Site:            shared.SharedResourceSite{ID: -1, Name: "None"},
		},
		Scope: ebooks.SubsetScope{
			AllComputers:     true,
			AllMobileDevices: false,
			AllJSSUsers:      false,
		},
		SelfService: ebooks.SubsetSelfService{
			SelfServiceDisplayName: "Updated Ebook Name",
			InstallButtonText:      "Install",
		},
	}

	updated, _, err := jamfClient.ClassicEbooks.UpdateByID(context.Background(), ebookID, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Ebook Updated: ID=%d\n", updated.ID)
}
