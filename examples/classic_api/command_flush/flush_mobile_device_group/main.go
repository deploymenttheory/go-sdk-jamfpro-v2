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

	// Clear failed MDM commands for mobile device group ID 1
	idType := "mobiledevicegroups"
	id := "1"
	status := "Failed"

	resp, err := jamfClient.ClassicAPI.CommandFlush.FlushByIDAndStatus(context.Background(), idType, id, status)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Successfully cleared %s MDM commands for %s %s (Status: %d)\n", status, idType, id, resp.StatusCode())
}
