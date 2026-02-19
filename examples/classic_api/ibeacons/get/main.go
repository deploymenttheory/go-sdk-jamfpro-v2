// Package main demonstrates GetIBeaconByID — retrieves a single iBeacon by ID.
//
// Run with: go run ./examples/classic_api/ibeacons/get
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Set IBEACON_ID or uses first from list.
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
	if raw := os.Getenv("IBEACON_ID"); raw != "" {
		id, err = strconv.Atoi(raw)
		if err != nil {
			log.Fatalf("invalid IBEACON_ID %q: %v", raw, err)
		}
	} else {
		list, _, err := client.IBeacons.ListIBeacons(ctx)
		if err != nil || len(list.Results) == 0 {
			log.Fatal("Set IBEACON_ID or ensure at least one iBeacon exists")
		}
		id = list.Results[0].ID
		fmt.Printf("Using first iBeacon ID: %d\n", id)
	}

	beacon, resp, err := client.IBeacons.GetIBeaconByID(ctx, id)
	if err != nil {
		log.Fatalf("GetIBeaconByID failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("ID:     %d\n", beacon.ID)
	fmt.Printf("Name:   %s\n", beacon.Name)
	fmt.Printf("UUID:   %s\n", beacon.UUID)
	fmt.Printf("Major:  %d\n", beacon.Major)
	fmt.Printf("Minor:  %d\n", beacon.Minor)
}
