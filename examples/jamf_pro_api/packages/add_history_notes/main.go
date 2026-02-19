// Package main demonstrates AddPackageHistoryNotesV1 - adds notes to a package's history.
//
// Run with: go run ./examples/jamf_pro_api/packages/add_history_notes
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates a package, adds a note, then deletes it.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/packages"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	// Create a package
	createReq := &packages.RequestPackage{
		PackageName:           fmt.Sprintf("example-history-%d", time.Now().UnixMilli()),
		FileName:              "example.pkg",
		CategoryID:            "-1",
		Info:                  "History notes test",
		Priority:              5,
		FillUserTemplate:      packages.BoolPtr(false),
		RebootRequired:        packages.BoolPtr(false),
		OSInstall:            packages.BoolPtr(false),
		SuppressUpdates:       packages.BoolPtr(false),
		SuppressFromDock:      packages.BoolPtr(false),
		SuppressEula:          packages.BoolPtr(false),
		SuppressRegistration:  packages.BoolPtr(false),
	}
	created, _, err := client.Packages.CreatePackageV1(ctx, createReq)
	if err != nil {
		log.Fatalf("CreatePackageV1 failed: %v", err)
	}
	id := created.ID
	fmt.Printf("Created package ID: %s\n", id)

	// Add a history note
	noteReq := &packages.AddHistoryNotesRequest{
		Note: fmt.Sprintf("Example note added at %s", time.Now().Format(time.RFC3339)),
	}
	resp, err := client.Packages.AddPackageHistoryNotesV1(ctx, id, noteReq)
	if err != nil {
		_, _ = client.Packages.DeletePackageByIDV1(ctx, id)
		log.Fatalf("AddPackageHistoryNotesV1 failed: %v", err)
	}

	fmt.Printf("Status: %d (201 = success)\n", resp.StatusCode)
	fmt.Println("History note added")

	// Fetch history to verify
	history, _, err := client.Packages.GetPackageHistoryV1(ctx, id, nil)
	if err == nil {
		fmt.Printf("History entries: %d\n", history.TotalCount)
		for _, e := range history.Results {
			if e.Note != "" {
				fmt.Printf("  Note: %s (by %s)\n", e.Note, e.Username)
			}
		}
	}

	_, _ = client.Packages.DeletePackageByIDV1(ctx, id)
	fmt.Println("Cleanup: package deleted")
}
