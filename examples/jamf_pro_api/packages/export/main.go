package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
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

	rsqlQuery := map[string]string{
		"page":           "0",
		"page-size":      "100",
		"sort":           "id:asc",
		"export-fields":  "id,packageName,fileName",
		"export-labels":  "ID,Package Name,File Name",
	}

	data, _, err := jamfClient.JamfProAPI.Packages.ExportV1(context.Background(), rsqlQuery, nil, constants.ApplicationJSON)
	if err != nil {
		fmt.Printf("Error exporting packages: %v\n", err)
		return
	}

	if err := os.WriteFile("packages_export.json", data, 0644); err != nil {
		fmt.Printf("Error writing export file: %v\n", err)
		return
	}

	fmt.Printf("Packages exported successfully to packages_export.json (%d bytes)\n", len(data))
}
