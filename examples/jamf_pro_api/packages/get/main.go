// Package main demonstrates GetPackageByIDV1 - retrieves a single package by ID.
//
// Run with: go run ./examples/jamf_pro_api/packages/get
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Set PACKAGE_ID or uses first from list.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()
	id := os.Getenv("PACKAGE_ID")
	if id == "" {
		list, _, err := client.Packages.ListPackagesV1(ctx, map[string]string{"page": "0", "pageSize": "1"})
		if err != nil || len(list.Results) == 0 {
			log.Fatal("Set PACKAGE_ID or ensure at least one package exists")
		}
		id = list.Results[0].ID
		fmt.Printf("Using first package ID: %s\n", id)
	}

	pkg, resp, err := client.Packages.GetPackageByIDV1(ctx, id)
	if err != nil {
		log.Fatalf("GetPackageByIDV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("ID: %s\n", pkg.ID)
	fmt.Printf("Name: %s\n", pkg.PackageName)
	fmt.Printf("FileName: %s\n", pkg.FileName)
	fmt.Printf("Category: %s\n", pkg.CategoryID)
	fmt.Printf("Info: %s\n", pkg.Info)
	fmt.Printf("Notes: %s\n", pkg.Notes)
	reboot := false
	if pkg.RebootRequired != nil {
		reboot = *pkg.RebootRequired
	}
	fmt.Printf("Priority: %d RebootRequired: %v\n", pkg.Priority, reboot)
}
