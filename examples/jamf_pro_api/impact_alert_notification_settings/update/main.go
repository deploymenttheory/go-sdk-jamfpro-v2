package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/impact_alert_notification_settings"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Example update request
	request := &impact_alert_notification_settings.ResourceImpactAlertNotificationSettings{
		ScopeableObjectsAlertEnabled:             true,
		ScopeableObjectsConfirmationCodeEnabled:  false,
		DeployableObjectsAlertEnabled:            true,
		DeployableObjectsConfirmationCodeEnabled: false,
	}

	resp, err := client.ImpactAlertNotificationSettings.UpdateV1(context.Background(), request)
	if err != nil {
		log.Fatalf("Failed to update impact alert notification settings: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Println("Impact alert notification settings updated successfully")
}
