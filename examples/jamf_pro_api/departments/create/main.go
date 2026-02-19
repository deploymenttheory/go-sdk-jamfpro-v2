// Package main demonstrates CreateDepartmentV1 - creates a new department.
//
// Run with: go run ./examples/jamf_pro_api/departments/create
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

	req := &departments.RequestDepartment{
		Name: fmt.Sprintf("example-dept-%d", time.Now().UnixMilli()),
	}

	result, resp, err := client.Departments.CreateDepartmentV1(ctx, req)
	if err != nil {
		log.Fatalf("CreateDepartmentV1 failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Created department ID: %s\n", result.ID)
	fmt.Printf("Href: %s\n", result.Href)

	if _, err := client.Departments.DeleteDepartmentByIDV1(ctx, result.ID); err != nil {
		fmt.Printf("Note: cleanup delete failed: %v\n", err)
	} else {
		fmt.Println("Cleanup: department deleted")
	}
}
