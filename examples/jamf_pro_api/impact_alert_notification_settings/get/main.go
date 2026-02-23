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

	result, resp, err := client.ImpactAlertNotificationSettings.GetV1(context.Background())
	if err != nil {
		log.Fatalf("Failed to get impact alert notification settings: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Scopeable Objects Alert Enabled: %t\n", result.ScopeableObjectsAlertEnabled)
	fmt.Printf("Scopeable Objects Confirmation Code Enabled: %t\n", result.ScopeableObjectsConfirmationCodeEnabled)
	fmt.Printf("Deployable Objects Alert Enabled: %t\n", result.DeployableObjectsAlertEnabled)
	fmt.Printf("Deployable Objects Confirmation Code Enabled: %t\n", result.DeployableObjectsConfirmationCodeEnabled)
}
