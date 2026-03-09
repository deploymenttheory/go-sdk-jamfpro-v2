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
		OpenGroup().
		EqualTo("hardware.make", "Apple").
		And().
		GreaterThan("hardware.totalRamMegabytes", "16384").
		CloseGroup().
		Or().
		OpenGroup().
		EqualTo("operatingSystem.name", "macOS").
		And().
		Contains("general.name", "MacBook").
		CloseGroup().
		Build()

	rsqlQuery := map[string]string{
		"filter": filter,
		"sort":   "general.name:asc",
	}

	result, _, err := jamfClient.
		JamfProAPI.
		ComputerInventory.
		ListV3(context.Background(), rsqlQuery)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("Found %d computers matching complex filter:\n", result.TotalCount)
	fmt.Printf("Filter: %s\n", filter)
	fmt.Printf("Results:\n%s\n", string(out))
}
