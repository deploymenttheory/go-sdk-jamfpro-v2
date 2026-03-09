package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/computers"
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

	createReq := &computers.ResponseComputer{
		General: computers.ComputerSubsetGeneral{
			Name:         "go-sdk-v2-test-computer",
			MacAddress:   "00:11:22:33:44:55",
			SerialNumber: "C02XYZ123456",
			Site: shared.SharedResourceSite{
				ID:   -1,
				Name: "none",
			},
		},
	}

	created, _, err := jamfClient.ClassicAPI.Computers.Create(context.Background(), createReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Computer Created: ID=%d Name=%s\n", created.General.ID, created.General.Name)
}
