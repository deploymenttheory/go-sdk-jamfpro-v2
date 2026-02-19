// Package main demonstrates DeletePackageByIDV1 - removes a package by ID.
//
// Run with: go run ./examples/jamf_pro_api/packages/delete
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates a package then deletes it.
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

	createReq := &packages.RequestPackage{
		PackageName:           fmt.Sprintf("example-delete-%d", time.Now().UnixMilli()),
		FileName:              "example.pkg",
		CategoryID:            "-1",
		Info:                  "Package to be deleted",
		Priority:              1,
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

	resp, err := client.Packages.DeletePackageByIDV1(ctx, id)
	if err != nil {
		log.Fatalf("DeletePackageByIDV1 failed: %v", err)
	}

	fmt.Printf("Status: %d (204 = success)\n", resp.StatusCode)
	fmt.Println("Package deleted successfully")
}
