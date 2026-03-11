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

	packageID := "1"

	rsqlQuery := map[string]string{
		"page":          "0",
		"page-size":     "100",
		"sort":          "date:desc",
		"export-fields": "id,username,date,note",
		"export-labels": "ID,User,Date,Note",
	}

	data, _, err := jamfClient.JamfProAPI.Packages.ExportHistoryV1(context.Background(), packageID, rsqlQuery, nil, constants.ApplicationJSON)
	if err != nil {
		fmt.Printf("Error exporting package history: %v\n", err)
		return
	}

	if err := os.WriteFile("package_history_export.json", data, 0644); err != nil {
		fmt.Printf("Error writing export file: %v\n", err)
		return
	}

	fmt.Printf("Package history exported successfully to package_history_export.json (%d bytes)\n", len(data))
}
