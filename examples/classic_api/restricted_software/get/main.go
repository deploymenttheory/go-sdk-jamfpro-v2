// Package main demonstrates GetRestrictedSoftwareByID — retrieves a single restricted software by ID.
//
// Run with: go run ./examples/classic_api/restricted_software/get
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Set RESTRICTED_SOFTWARE_ID or uses first from list.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	var id int
	if raw := os.Getenv("RESTRICTED_SOFTWARE_ID"); raw != "" {
		id, err = strconv.Atoi(raw)
		if err != nil {
			log.Fatalf("invalid RESTRICTED_SOFTWARE_ID %q: %v", raw, err)
		}
	} else {
		list, _, err := client.RestrictedSoftware.ListRestrictedSoftware(ctx)
		if err != nil || len(list.Results) == 0 {
			log.Fatal("Set RESTRICTED_SOFTWARE_ID or ensure at least one restricted software item exists")
		}
		id = list.Results[0].ID
		fmt.Printf("Using first restricted software ID: %d\n", id)
	}

	software, resp, err := client.RestrictedSoftware.GetRestrictedSoftwareByID(ctx, id)
	if err != nil {
		log.Fatalf("GetRestrictedSoftwareByID failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("ID:              %d\n", software.General.ID)
	fmt.Printf("Name:            %s\n", software.General.Name)
	fmt.Printf("Process Name:    %s\n", software.General.ProcessName)
	fmt.Printf("Send Notification: %v\n", software.General.SendNotification)
	fmt.Printf("Kill Process:    %v\n", software.General.KillProcess)
	fmt.Printf("All Computers:   %v\n", software.Scope.AllComputers)
}
