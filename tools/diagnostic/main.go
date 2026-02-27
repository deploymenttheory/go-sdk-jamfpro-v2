package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/smart_user_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/static_user_groups"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

func main() {
	authConfig := client.AuthConfigFromEnv()

	jamfClient, err := jamfpro.NewClient(authConfig)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Test 1: Create Smart User Group and check response
	fmt.Println("=== Test 1: Smart User Group Create ===")
	smartGroupReq := &smart_user_groups.RequestSmartUserGroup{
		Name:             fmt.Sprintf("diagnostic-smart-%d", time.Now().Unix()),
		IsSmart:          true,
		IsNotifyOnChange: false,
		Site: &shared.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Criteria: &smart_user_groups.CriteriaContainer{
			Size: 1,
			Criterion: []shared.SharedSubsetCriteria{
				{
					Name:       "Email Address",
					Priority:   0,
					AndOr:      "and",
					SearchType: "like",
					Value:      "@example.com",
				},
			},
		},
	}

	createdSmart, resp1, err := jamfClient.ClassicSmartUserGroups.Create(ctx, smartGroupReq)
	if err != nil {
		fmt.Printf("Smart Create Error: %v\n", err)
	}
	if resp1 != nil {
		fmt.Printf("Smart Create Status: %d\n", resp1.StatusCode)
		fmt.Printf("Smart Create Response Body:\n%s\n\n", string(resp1.Body))
	}
	
	if createdSmart != nil && createdSmart.ID > 0 {
		smartGroupID := createdSmart.ID
		fmt.Printf("Created Smart Group ID: %d\n\n", smartGroupID)
		
		fmt.Println("=== Test 1b: Smart User Group GET ===")
		time.Sleep(1 * time.Second)
		_, resp1b, err := jamfClient.ClassicSmartUserGroups.GetByID(ctx, smartGroupID)
		if err != nil {
			fmt.Printf("Smart GET Error: %v\n", err)
		}
		if resp1b != nil {
			fmt.Printf("Smart GET Status: %d\n", resp1b.StatusCode)
			fmt.Printf("Smart GET Response Body:\n%s\n\n", string(resp1b.Body))
		}
		
		fmt.Println("=== Test 1c: Smart User Group UPDATE ===")
		updateReq := &smart_user_groups.RequestSmartUserGroup{
			Name:             fmt.Sprintf("diagnostic-smart-updated-%d", time.Now().Unix()),
			IsSmart:          true,
			IsNotifyOnChange: true,
			Site: &shared.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
			Criteria: &smart_user_groups.CriteriaContainer{
				Size: 1,
				Criterion: []shared.SharedSubsetCriteria{
					{
						Name:       "Email Address",
						Priority:   0,
						AndOr:      "and",
						SearchType: "like",
						Value:      "@example.com",
					},
				},
			},
		}
		_, resp1c, err := jamfClient.ClassicSmartUserGroups.UpdateByID(ctx, smartGroupID, updateReq)
		if err != nil {
			fmt.Printf("Smart UPDATE Error: %v\n", err)
		}
		if resp1c != nil {
			fmt.Printf("Smart UPDATE Status: %d\n", resp1c.StatusCode)
			fmt.Printf("Smart UPDATE Response Body:\n%s\n\n", string(resp1c.Body))
		}
		
		jamfClient.ClassicSmartUserGroups.DeleteByID(ctx, smartGroupID)
	}

	// Test 2: Create Static User Group and check response
	fmt.Println("=== Test 2: Static User Group Create ===")
	staticGroupReq := &static_user_groups.RequestStaticUserGroup{
		Name:             fmt.Sprintf("diagnostic-static-%d", time.Now().Unix()),
		IsSmart:          false,
		IsNotifyOnChange: false,
		Site: &shared.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
	}

	createdStatic, resp2, err := jamfClient.ClassicStaticUserGroups.Create(ctx, staticGroupReq)
	if err != nil {
		fmt.Printf("Static Create Error: %v\n", err)
	}
	if resp2 != nil {
		fmt.Printf("Static Create Status: %d\n", resp2.StatusCode)
		fmt.Printf("Static Create Response Body:\n%s\n\n", string(resp2.Body))
	}
	
	if createdStatic != nil && createdStatic.ID > 0 {
		staticGroupID := createdStatic.ID
		fmt.Printf("Created Static Group ID: %d\n\n", staticGroupID)
		
		fmt.Println("=== Test 2b: Static User Group GET ===")
		time.Sleep(1 * time.Second)
		_, resp2b, err := jamfClient.ClassicStaticUserGroups.GetByID(ctx, staticGroupID)
		if err != nil {
			fmt.Printf("Static GET Error: %v\n", err)
		}
		if resp2b != nil {
			fmt.Printf("Static GET Status: %d\n", resp2b.StatusCode)
			fmt.Printf("Static GET Response Body:\n%s\n\n", string(resp2b.Body))
		}
		
		fmt.Println("=== Test 2c: Static User Group UPDATE ===")
		updateReq := &static_user_groups.RequestStaticUserGroup{
			Name:             fmt.Sprintf("diagnostic-static-updated-%d", time.Now().Unix()),
			IsSmart:          false,
			IsNotifyOnChange: true,
			Site: &shared.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		}
		_, resp2c, err := jamfClient.ClassicStaticUserGroups.UpdateByID(ctx, staticGroupID, updateReq)
		if err != nil {
			fmt.Printf("Static UPDATE Error: %v\n", err)
		}
		if resp2c != nil {
			fmt.Printf("Static UPDATE Status: %d\n", resp2c.StatusCode)
			fmt.Printf("Static UPDATE Response Body:\n%s\n\n", string(resp2c.Body))
		}
		
		jamfClient.ClassicStaticUserGroups.DeleteByID(ctx, staticGroupID)
	}
}
