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
		log.Fatalf("Failed to create client: %v", err)
	}

	uuid := "YOUR-PROFILE-UUID-HERE"
	computerIDs := []string{"1", "2", "3"}

	resp, err := client.JamfProAPI.JamfConnect.RetryDeploymentTasksByUUIDV1(context.Background(), uuid, computerIDs)
	if err != nil {
		log.Fatalf("Failed to retry deployment tasks: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode())
	fmt.Println("Deployment tasks retry initiated successfully")
}
