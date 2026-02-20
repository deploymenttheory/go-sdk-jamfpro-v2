package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/scripts"
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

	id := "1" // Replace with the desired script ID
	updateReq := &scripts.RequestScript{
		Name:           "Updated Script Name",
		Priority:       scripts.ScriptPriorityBefore,
		Info:           "Updated by SDK example",
		ScriptContents: "#!/bin/bash\necho 'updated'",
	}
	result, _, err := jamfClient.Scripts.UpdateScriptByIDV1(context.Background(), id, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Updated script: %+v\n", result)
}
