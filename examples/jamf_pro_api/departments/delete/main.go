// Package main demonstrates DeleteDepartmentByIDV1 - removes a department by ID.
//
// Run with: go run ./examples/jamf_pro_api/departments/delete
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars. Creates a department then deletes it.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/departments"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	createReq := &departments.RequestDepartment{
		Name: fmt.Sprintf("example-delete-%d", time.Now().UnixMilli()),
	}
	created, _, err := client.Departments.CreateDepartmentV1(ctx, createReq)
	if err != nil {
		log.Fatalf("CreateDepartmentV1 failed: %v", err)
	}
	id := created.ID
	fmt.Printf("Created department ID: %s\n", id)

	resp, err := client.Departments.DeleteDepartmentByIDV1(ctx, id)
	if err != nil {
		log.Fatalf("DeleteDepartmentByIDV1 failed: %v", err)
	}

	fmt.Printf("Status: %d (204 = success)\n", resp.StatusCode)
	fmt.Println("Department deleted successfully")
}
