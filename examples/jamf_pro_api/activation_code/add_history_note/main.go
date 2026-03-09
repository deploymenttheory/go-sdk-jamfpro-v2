package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/activation_code"
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

	req := &activation_code.HistoryNoteRequest{
		Note: "Manual note added via API",
	}

	result, _, err := jamfClient.JamfProAPI.ActivationCode.AddHistoryNoteV1(context.Background(), req)
	if err != nil {
		fmt.Printf("Error adding history note: %v\n", err)
		return
	}
	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("Added history note:\n%s\n", string(out))
}
