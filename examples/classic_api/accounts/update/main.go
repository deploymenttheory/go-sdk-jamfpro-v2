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

	createReq := &accounts.RequestAccount{
		Name:         fmt.Sprintf("example-update-%d", time.Now().UnixMilli()),
		FullName:     "Example Update User",
		Email:        "example-update@example.com",
		EmailAddress: "example-update@example.com",
		Password:     "TestPassword123!",
		AccessLevel:  "Full Access",
		PrivilegeSet: "Administrator",
		Enabled:      "Enabled",
	}
	created, _, err := client.ClassicAccounts.Create(ctx, createReq)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	fmt.Printf("Created account ID: %d name: %s\n", created.ID, created.Name)

	updateReq := &accounts.RequestAccount{
		Name:         fmt.Sprintf("example-updated-%d", time.Now().UnixMilli()),
		FullName:     "Updated User Name",
		Email:        "updated@example.com",
		EmailAddress: "updated@example.com",
		AccessLevel:  "Full Access",
		PrivilegeSet: "Administrator",
		Enabled:      "Enabled",
	}
	updated, resp, err := client.ClassicAccounts.UpdateByID(ctx, created.ID, updateReq)
	if err != nil {
		_, _ = client.ClassicAccounts.DeleteByID(ctx, created.ID)
		log.Fatalf("UpdateByID failed: %v", err)
	}

	fmt.Printf("Status: %d\n", resp.StatusCode)
	fmt.Printf("Updated account ID: %d\n", updated.ID)
	fmt.Printf("New name: %s\n", updated.Name)

	_, _ = client.ClassicAccounts.DeleteByID(ctx, created.ID)
	fmt.Println("Cleanup: account deleted")
}
