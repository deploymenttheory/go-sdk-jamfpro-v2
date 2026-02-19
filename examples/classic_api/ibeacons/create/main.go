// Package main demonstrates CreateIBeacon — creates a new iBeacon via the Classic API.
//
// Run with: go run ./examples/classic_api/ibeacons/create
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates an iBeacon then deletes it.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/ibeacons"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	req := &ibeacons.RequestIBeacon{
		Name:  fmt.Sprintf("example-ibeacon-%d", time.Now().UnixMilli()),
		UUID:  "E2C56DB5-DFFB-48D2-B060-D0F5A71096E0",
		Major: 1,
		Minor: 1,
	}

	created, resp, err := client.IBeacons.CreateIBeacon(ctx, req)
	if err != nil {
		log.Fatalf("CreateIBeacon failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Created iBeacon ID: %d\n", created.ID)
	fmt.Printf("Name: %s\n", created.Name)

	// Cleanup: delete the created iBeacon
	if _, err := client.IBeacons.DeleteIBeaconByID(ctx, created.ID); err != nil {
		fmt.Printf("Note: cleanup delete failed: %v\n", err)
	} else {
		fmt.Println("Cleanup: iBeacon deleted")
	}
}
