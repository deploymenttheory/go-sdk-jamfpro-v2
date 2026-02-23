package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/jamf_pro_system_initialization"
)

func main() {
	configFilePath := "/path/to/clientconfig.json"
	authConfig, err := client.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	request := &jamf_pro_system_initialization.ResourceSystemInitialize{
		ActivationCode:  "your-activation-code",
		InstitutionName: "My Institution",
		EulaAccepted:    true,
		Username:        "admin",
		Password:        "secure-password",
		Email:           "admin@example.com",
		JssUrl:          "https://your-instance.jamfcloud.com",
	}

	resp, err := jamfClient.JamfProSystemInitialization.Initialize(context.Background(), request)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Jamf Pro initialized successfully (status: %d)\n", resp.StatusCode)
}
