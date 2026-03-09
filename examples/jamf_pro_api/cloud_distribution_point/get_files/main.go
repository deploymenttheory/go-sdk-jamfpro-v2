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
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	authConfig, err := client.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Example 1: Get all files
	result, _, err := jamfClient.JamfProAPI.CloudDistributionPoint.GetFilesV1(context.Background(), nil)
	if err != nil {
		fmt.Printf("Error getting cloud distribution point files: %v\n", err)
		return
	}
	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("All cloud distribution point files:\n%s\n\n", string(out))

	// Example 2: Get files with RSQL filter (e.g. by status or type)
	rsqlQuery := map[string]string{
		"filter": `status==complete`,
		"sort":   "fileName:asc",
	}
	result, _, err = jamfClient.JamfProAPI.CloudDistributionPoint.GetFilesV1(context.Background(), rsqlQuery)
	if err != nil {
		fmt.Printf("Error getting filtered files: %v\n", err)
		return
	}
	out, _ = json.MarshalIndent(result, "", "    ")
	fmt.Printf("Filtered cloud distribution point files:\n%s\n", string(out))

	// Example 3: Pagination with page and page-size
	rsqlQuery = map[string]string{
		"page":      "0",
		"page-size": "25",
	}
	result, _, err = jamfClient.JamfProAPI.CloudDistributionPoint.GetFilesV1(context.Background(), rsqlQuery)
	if err != nil {
		fmt.Printf("Error getting paginated files: %v\n", err)
		return
	}
	out, _ = json.MarshalIndent(result, "", "    ")
	fmt.Printf("Paginated files (page 0, size 25):\n%s\n", string(out))
}
