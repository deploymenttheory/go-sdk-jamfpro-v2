package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/computer_groups"
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

	req := &computer_groups.RequestSmartGroup{
		Name: "go-sdk-v2-Smart-Group",
		Criteria: []computer_groups.Criterion{
			{Name: "Computer Name", Priority: 0, AndOr: "and", SearchType: "like", Value: "Mac-%"},
		},
	}

	result, _, err := jamfClient.ComputerGroups.CreateSmartV2(context.Background(), req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Created smart group: %+v\n", result)
}
