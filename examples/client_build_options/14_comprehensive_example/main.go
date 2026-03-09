package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"go.uber.org/zap"
)

func main() {
	authConfig := jamfpro.AuthConfigFromEnv()

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Failed to create logger: %v", err)
	}

	jamfClient, err := jamfpro.NewClient(
		authConfig,
		jamfpro.WithLogger(logger),
		jamfpro.WithTimeout(60*time.Second),
		jamfpro.WithRetryCount(3),
		jamfpro.WithMaxConcurrentRequests(5),
		jamfpro.WithDebug(),
	)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	ctx := context.Background()

	fmt.Println("=== Example 1: List all buildings ===")
	buildings, _, err := jamfClient.
		JamfProAPI.
		Buildings.
		ListV1(ctx, map[string]string{"page": "0", "page-size": "50"})
	if err != nil {
		log.Fatalf("Error listing buildings: %v", err)
	}
	out, _ := json.MarshalIndent(buildings, "", "  ")
	fmt.Printf("Buildings: %s\n\n", string(out))

	fmt.Println("=== Example 2: List computers with RSQL filter ===")
	filter := jamfClient.
		GetTransport().
		RSQLBuilder().
		GreaterThan("hardware.totalRamMegabytes", "8192").
		And().
		EqualTo("operatingSystem.name", "macOS").
		Build()

	rsqlQuery := map[string]string{
		"filter": filter,
		"sort":   "general.name:asc",
	}

	computers, _, err := jamfClient.
		JamfProAPI.
		ComputerInventory.
		ListV3(ctx, rsqlQuery)
	if err != nil {
		log.Fatalf("Error listing computers: %v", err)
	}
	fmt.Printf("Found %d computers with >8GB RAM running macOS\n", computers.TotalCount)
	fmt.Printf("Filter used: %s\n\n", filter)

	fmt.Println("=== Example 3: Token management ===")
	if err := jamfClient.GetTransport().KeepAliveToken(); err != nil {
		log.Printf("Warning: Failed to keep token alive: %v", err)
	} else {
		fmt.Println("Token kept alive successfully")
	}

	if err := jamfClient.GetTransport().InvalidateToken(); err != nil {
		log.Printf("Warning: Failed to invalidate token: %v", err)
	} else {
		fmt.Println("Token invalidated successfully")
	}
}
