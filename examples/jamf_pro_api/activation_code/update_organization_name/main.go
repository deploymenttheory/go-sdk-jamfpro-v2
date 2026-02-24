package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/activation_code"
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

	req := &activation_code.OrganizationNameRequest{
		OrganizationName: "Example Organization",
	}

	resp, err := jamfClient.ActivationCode.UpdateOrganizationNameV1(context.Background(), req)
	if err != nil {
		fmt.Printf("Error updating organization name: %v\n", err)
		return
	}
	fmt.Printf("Organization name updated successfully (Status: %d)\n", resp.StatusCode)
}
