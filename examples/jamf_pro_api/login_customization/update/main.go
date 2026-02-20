package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/login_customization"
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

	newLoginCustomization := &login_customization.ResourceLoginCustomizationV1{
		RampInstance:            true,
		IncludeCustomDisclaimer: true,
		DisclaimerHeading:       "Updated Disclaimer Header",
		DisclaimerMainText:      "Updated disclaimer main text",
		ActionText:              "Accept",
	}

	updated, _, err := jamfClient.LoginCustomization.UpdateLoginCustomizationV1(context.Background(), newLoginCustomization)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	out, _ := json.MarshalIndent(updated, "", "    ")
	fmt.Println("Updated login customization:\n" + string(out))
}
