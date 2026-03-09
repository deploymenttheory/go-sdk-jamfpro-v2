package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/cloud_azure"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	authConfig, err := jamfpro.LoadAuthConfigFromFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	req := &cloud_azure.ResourceCloudAzure{
		CloudIdPCommon: cloud_azure.CloudIdPCommon{
			DisplayName:  "Example Azure Cloud IDP",
			ProviderName: "AZURE",
		},
		Server: cloud_azure.CloudAzureServer{
			TenantId:                             "your-tenant-id-here",
			Enabled:                              true,
			TransitiveMembershipEnabled:          true,
			TransitiveDirectoryMembershipEnabled: true,
			SearchTimeout:                        30,
			Mappings: cloud_azure.CloudAzureServerMappings{
				UserId:     "id",
				UserName:   "userPrincipalName",
				RealName:   "displayName",
				Email:      "mail",
				Department: "department",
				Building:   "companyName",
				Room:       "officeLocation",
				Phone:      "mobilePhone",
				Position:   "jobTitle",
				GroupId:    "id",
				GroupName:  "displayName",
			},
		},
	}

	result, _, err := jamfClient.JamfProAPI.CloudAzure.CreateV1(context.Background(), req)
	if err != nil {
		fmt.Printf("Error creating Azure Cloud IDP: %v\n", err)
		return
	}
	out, _ := json.MarshalIndent(result, "", "    ")
	fmt.Printf("Created Azure Cloud IDP:\n%s\n", string(out))
}
