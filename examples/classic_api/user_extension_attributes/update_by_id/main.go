package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/user_extension_attributes"
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

	attrID := 1 // Replace with the desired user extension attribute ID
	updateReq := &user_extension_attributes.RequestUserExtensionAttribute{
		Name:        "Department",
		Description: "Updated user department",
		DataType:    "String",
		InputType: user_extension_attributes.ResourceUserExtensionAttributeInputType{
			Type: "Text Field",
		},
	}

	updated, _, err := jamfClient.ClassicAPI.UserExtensionAttributes.UpdateByID(context.Background(), attrID, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("User Extension Attribute Updated: ID=%d name=%q\n", updated.ID, updated.Name)
}
