package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/volume_purchasing_locations"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	authConfig, err := client.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	serviceToken := "YOUR_VPP_SERVICE_TOKEN" // Replace with a valid Apple VPP service token
	req := &volume_purchasing_locations.RequestVolumePurchasingLocation{
		Name:                                  "go-sdk-v2-VPL",
		ServiceToken:                          serviceToken,
		AutomaticallyPopulatePurchasedContent: true,
		SendNotificationWhenNoLongerAssigned:  false,
		AutoRegisterManagedUsers:              false,
	}

	result, _, err := jamfClient.VolumePurchasingLocations.CreateVolumePurchasingLocationV1(context.Background(), req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Created volume purchasing location: %+v\n", result)
}
