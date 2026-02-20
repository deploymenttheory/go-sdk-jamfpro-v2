package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_extension_attributes"
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

	enabled := true
	req := &computer_extension_attributes.RequestComputerExtensionAttribute{
		Name:                 "go-sdk-v2-Computer-EA",
		Description:          "Example computer extension attribute",
		DataType:             "String",
		Enabled:              &enabled,
		InventoryDisplayType: "General",
		InputType:            "Text Field",
	}

	result, _, err := jamfClient.ComputerExtensionAttributes.CreateComputerExtensionAttributeV1(context.Background(), req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Created computer extension attribute: %+v\n", result)
}
