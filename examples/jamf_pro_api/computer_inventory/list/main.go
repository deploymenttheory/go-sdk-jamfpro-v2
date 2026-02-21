package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
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

	result, _, err := jamfClient.ComputerInventory.ListV1(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Total Computers: %d\n\n", result.TotalCount)
	for _, computer := range result.Results {
		fmt.Printf("ID: %s\n", computer.ID)
		fmt.Printf("  Name: %s\n", computer.General.Name)
		fmt.Printf("  Serial: %s\n", computer.Hardware.SerialNumber)
		fmt.Printf("  Model: %s\n", computer.Hardware.Model)
		fmt.Printf("  IP: %s\n", computer.General.LastIpAddress)
		fmt.Println()
	}
}
