package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/computers"
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

	computerID := "1" // Replace with the desired computer ID

	// Fetch existing computer first to get full structure
	existing, _, err := jamfClient.ClassicAPI.Computers.GetByID(context.Background(), computerID)
	if err != nil {
		fmt.Printf("Error fetching computer: %v\n", err)
		return
	}

	// Update the name (and other fields as needed)
	updateReq := &computers.ResponseComputer{
		General: computers.ComputerSubsetGeneral{
			ID:            existing.General.ID,
			Name:          "go-sdk-v2-test-computer-updated",
			MacAddress:    existing.General.MacAddress,
			SerialNumber:  existing.General.SerialNumber,
			Site:          models.SharedResourceSite{ID: -1, Name: "none"},
		},
		Location:       existing.Location,
		Purchasing:     existing.Purchasing,
		Peripherals:    existing.Peripherals,
		Hardware:       existing.Hardware,
		Security:       existing.Security,
		Software:       existing.Software,
		GroupsAccounts: existing.GroupsAccounts,
	}

	updated, _, err := jamfClient.ClassicAPI.Computers.UpdateByID(context.Background(), computerID, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Computer Updated: ID=%d Name=%s\n", updated.General.ID, updated.General.Name)
}
