// Package main demonstrates UpdateIBeaconByID — updates an existing iBeacon via the Classic API.
//
// Run with: go run ./examples/classic_api/ibeacons/update
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
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

	// Create an iBeacon to update
	createReq := &ibeacons.RequestIBeacon{
		Name:  fmt.Sprintf("example-update-%d", time.Now().UnixMilli()),
		UUID:  "E2C56DB5-DFFB-48D2-B060-D0F5A71096E0",
		Major: 1,
		Minor: 1,
	}
	created, _, err := client.IBeacons.CreateIBeacon(ctx, createReq)
	if err != nil {
		log.Fatalf("CreateIBeacon failed: %v", err)
	}
	fmt.Printf("Created iBeacon ID: %d name: %s\n", created.ID, created.Name)

	// Update the iBeacon
	updateReq := &ibeacons.RequestIBeacon{
		Name:  fmt.Sprintf("example-updated-%d", time.Now().UnixMilli()),
		UUID:  "E2C56DB5-DFFB-48D2-B060-D0F5A71096E0",
		Major: 1,
		Minor: 2,
	}
	updated, resp, err := client.IBeacons.UpdateIBeaconByID(ctx, created.ID, updateReq)
	if err != nil {
		_, _ = client.IBeacons.DeleteIBeaconByID(ctx, created.ID)
		log.Fatalf("UpdateIBeaconByID failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Updated iBeacon ID: %d\n", updated.ID)
	fmt.Printf("New name: %s\n", updated.Name)
	fmt.Printf("Minor: %d\n", updated.Minor)

	_, _ = client.IBeacons.DeleteIBeaconByID(ctx, created.ID)
	fmt.Println("Cleanup: iBeacon deleted")
}
