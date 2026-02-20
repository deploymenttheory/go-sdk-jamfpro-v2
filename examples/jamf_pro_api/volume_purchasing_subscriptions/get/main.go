package main

import (
	"context"
	"encoding/json"
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

	subscriptionID := "1" // Replace with the desired volume purchasing subscription ID
	result, _, err := jamfClient.VolumePurchasingSubscriptions.GetByIDV1(context.Background(), subscriptionID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Println("Volume purchasing subscription:\n" + string(out))
}
