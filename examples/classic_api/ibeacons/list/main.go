// Package main demonstrates ListIBeacons — returns all iBeacons from the Classic API.
//
// Run with: go run ./examples/classic_api/ibeacons/list
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	list, resp, err := client.IBeacons.ListIBeacons(context.Background())
	if err != nil {
		log.Fatalf("ListIBeacons failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Total iBeacons: %d\n", list.Size)
	for _, b := range list.Results {
		fmt.Printf("  ID=%-5d  Name=%s\n", b.ID, b.Name)
	}
}
