package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	authConfig := jamfpro.AuthConfigFromEnv()

	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	filter := jamfClient.
		GetTransport().
		RSQLBuilder().
		EqualTo("general.name", "MacBook*").
		And().
		GreaterThan("hardware.totalRamMegabytes", "8192").
		Build()

	rsqlQuery := map[string]string{
		"filter":    filter,
		"page":      "0",
		"page-size": "100",
	}

	result, _, err := jamfClient.
		JamfProAPI.
		ComputerInventory.
		ListV3(context.Background(), rsqlQuery)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("Found %d computers matching filter:\n%s\n", result.TotalCount, string(out))
}
