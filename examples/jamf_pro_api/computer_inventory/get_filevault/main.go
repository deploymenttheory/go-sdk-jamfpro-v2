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

	computerID := "1"

	result, _, err := jamfClient.ComputerInventory.GetFileVaultByIDV1(context.Background(), computerID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("FileVault Details for Computer: %s\n", result.Name)
	fmt.Printf("  Computer ID: %s\n", result.ComputerId)
	fmt.Printf("  Recovery Key Status: %s\n", result.IndividualRecoveryKeyValidityStatus)
	fmt.Printf("  Institutional Key Present: %v\n", result.InstitutionalRecoveryKeyPresent)
	fmt.Printf("  Partition State: %s\n", result.BootPartitionEncryptionDetails.PartitionFileVault2State)
	fmt.Printf("  Encryption Progress: %d%%\n", result.BootPartitionEncryptionDetails.PartitionFileVault2Percent)
}
