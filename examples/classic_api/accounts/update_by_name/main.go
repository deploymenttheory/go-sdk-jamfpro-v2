package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/accounts"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	ctx := context.Background()

	originalName := fmt.Sprintf("example-update-name-%d", time.Now().UnixMilli())
	createReq := &accounts.RequestAccount{
		Name:         originalName,
		FullName:     "Example Update By Name User",
		Email:        "example-update-name@example.com",
		EmailAddress: "example-update-name@example.com",
		Password:     "TestPassword123!",
		AccessLevel:  "Full Access",
		PrivilegeSet: "Administrator",
		Enabled:      "Enabled",
	}
	created, _, err := client.Accounts.Create(ctx, createReq)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	fmt.Printf("Created account ID: %d name: %s\n", created.ID, created.Name)

	updatedName := fmt.Sprintf("example-updated-name-%d", time.Now().UnixMilli())
	updateReq := &accounts.RequestAccount{
		Name:         updatedName,
		FullName:     "Updated By Name User",
		Email:        "updated-name@example.com",
		EmailAddress: "updated-name@example.com",
		AccessLevel:  "Full Access",
		PrivilegeSet: "Administrator",
		Enabled:      "Enabled",
	}
	updated, resp, err := client.Accounts.UpdateByName(ctx, originalName, updateReq)
	if err != nil {
		_, _ = client.Accounts.DeleteByID(ctx, created.ID)
		log.Fatalf("UpdateByName failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Updated account ID: %d\n", updated.ID)
	fmt.Printf("New name: %s\n", updated.Name)

	_, _ = client.Accounts.DeleteByID(ctx, created.ID)
	fmt.Println("Cleanup: account deleted")
}
