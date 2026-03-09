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

	queryParams := map[string]string{
		"page":      "0",
		"page-size": "100",
		"sort":      "date:desc",
	}

	page := 0
	pageSize := 100
	req := &activation_code.HistoryExportRequest{
		Page:     &page,
		PageSize: &pageSize,
		Sort:     []string{"date:desc"},
	}

	result, resp, err := jamfClient.JamfProAPI.ActivationCode.ExportHistoryV1(context.Background(), queryParams, req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal result: %v", err)
	}

	fmt.Printf("Response (Status: %d):\n%s\n", resp.StatusCode(), string(out))
}
