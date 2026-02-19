// Package main demonstrates DeleteIBeaconByID — removes an iBeacon via the Classic API.
//
// Run with: go run ./examples/classic_api/ibeacons/delete
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

	// Create an iBeacon to delete
	createReq := &ibeacons.RequestIBeacon{
		Name:  fmt.Sprintf("example-delete-%d", time.Now().UnixMilli()),
		UUID:  "E2C56DB5-DFFB-48D2-B060-D0F5A71096E0",
		Major: 1,
		Minor: 1,
	}
	created, _, err := client.IBeacons.CreateIBeacon(ctx, createReq)
	if err != nil {
		log.Fatalf("CreateIBeacon failed: %v", err)
	}
	fmt.Printf("Created iBeacon ID: %d\n", created.ID)

	resp, err := client.IBeacons.DeleteIBeaconByID(ctx, created.ID)
	if err != nil {
		log.Fatalf("DeleteIBeaconByID failed: %v", err)
	}

	fmt.Printf("Status: %d (200 = success)\n", resp.StatusCode)
	fmt.Println("iBeacon deleted successfully")
}
