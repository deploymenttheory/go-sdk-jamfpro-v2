// Package main demonstrates how to retrieve the custom settings schema list
// for macOS configuration profiles using the Jamf Pro SDK.
//
// The schema list describes available custom settings domains and versions
// that can be used when creating configuration profiles.
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
	configFilePath := "/path/to/clientconfig.json"
	authConfig, err := client.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	result, _, err := jamfClient.JamfProAPI.MacosConfigProfileCustomSettings.GetSchemaList(context.Background())
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Println("Custom settings schema list:\n" + string(out))
}
