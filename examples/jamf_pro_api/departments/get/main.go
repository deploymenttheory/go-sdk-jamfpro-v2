// Package main demonstrates GetDepartmentByIDV1 - retrieves a single department by ID.
//
// Run with: go run ./examples/jamf_pro_api/departments/get
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Set DEPARTMENT_ID or uses first from list.
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
	id := os.Getenv("DEPARTMENT_ID")
	if id == "" {
		list, _, err := client.Departments.ListDepartmentsV1(ctx, map[string]string{"page": "0", "pageSize": "1"})
		if err != nil || len(list.Results) == 0 {
			log.Fatal("Set DEPARTMENT_ID or ensure at least one department exists")
		}
		id = list.Results[0].ID
		fmt.Printf("Using first department ID: %s\n", id)
	}

	dept, resp, err := client.Departments.GetDepartmentByIDV1(ctx, id)
	if err != nil {
		log.Fatalf("GetDepartmentByIDV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("ID: %s\n", dept.ID)
	fmt.Printf("Name: %s\n", dept.Name)
}
