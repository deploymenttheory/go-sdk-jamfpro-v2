// Package main demonstrates UpdateDepartmentByIDV1 - updates an existing department.
//
// Run with: go run ./examples/jamf_pro_api/departments/update
// Requires: INSTANCE_DOMAIN, AUTH_METHOD, and auth env vars.
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
		Name: fmt.Sprintf("example-update-%d", time.Now().UnixMilli()),
	}
	created, _, err := client.Departments.CreateDepartmentV1(ctx, createReq)
	if err != nil {
		log.Fatalf("CreateDepartmentV1 failed: %v", err)
	}
	id := created.ID

	updateReq := &departments.RequestDepartment{
		Name: fmt.Sprintf("example-updated-%d", time.Now().UnixMilli()),
	}
	result, resp, err := client.Departments.UpdateDepartmentByIDV1(ctx, id, updateReq)
	if err != nil {
		_, _ = client.Departments.DeleteDepartmentByIDV1(ctx, id)
		log.Fatalf("UpdateDepartmentByIDV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Updated department ID: %s Name=%q\n", result.ID, result.Name)

	fetched, _, _ := client.Departments.GetDepartmentByIDV1(ctx, id)
	if fetched != nil {
		fmt.Printf("Verified: name=%q\n", fetched.Name)
	}

	_, _ = client.Departments.DeleteDepartmentByIDV1(ctx, id)
	fmt.Println("Cleanup: department deleted")
}
