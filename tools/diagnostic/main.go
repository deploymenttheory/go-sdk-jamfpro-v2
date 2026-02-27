package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/classic_api/mobile_device_groups"
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

	// Create test group
	groupName := fmt.Sprintf("diagnostic-mdgroup-%d", time.Now().Unix())
	createReq := &mobile_device_groups.RequestMobileDeviceGroup{
		Name:    groupName,
		IsSmart: true,
		Site: &shared.SharedResourceSite{ID: -1, Name: "None"},
		Criteria: &mobile_device_groups.CriteriaContainer{
			Size: 1,
			Criterion: []shared.SharedSubsetCriteria{
				{Name: "Model", Priority: 0, AndOr: "and", SearchType: "like", Value: "iPhone"},
			},
		},
	}
	created, resp, err := jamfClient.ClassicMobileDeviceGroups.Create(ctx, createReq)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}
	fmt.Printf("Created group ID=%d status=%d\n\n", created.ID, resp.StatusCode)
	groupID := created.ID
	defer jamfClient.ClassicMobileDeviceGroups.DeleteByID(ctx, groupID)

	// Test 1: 2 criteria with Priority 1 and 2 (both non-zero, won't be omitted)
	fmt.Println("=== Test 1: 2 criteria with Priority 1 and 2 ===")
	updateReq1 := &mobile_device_groups.RequestMobileDeviceGroup{
		Name:    groupName + "-v1",
		IsSmart: true,
		Site: &shared.SharedResourceSite{ID: -1, Name: "None"},
		Criteria: &mobile_device_groups.CriteriaContainer{
			Size: 2,
			Criterion: []shared.SharedSubsetCriteria{
				{Name: "Model", Priority: 1, AndOr: "and", SearchType: "like", Value: "iPhone"},
				{Name: "Device Name", Priority: 2, AndOr: "and", SearchType: "like", Value: "test"},
			},
		},
	}
	xmlBytes1, _ := xml.MarshalIndent(updateReq1, "", "  ")
	fmt.Printf("XML:\n%s\n", string(xmlBytes1))
	_, resp1, err1 := jamfClient.ClassicMobileDeviceGroups.UpdateByID(ctx, groupID, updateReq1)
	if err1 != nil {
		fmt.Printf("Test1 FAILED: %v\n\n", err1)
	} else {
		fmt.Printf("Test1 PASSED: status=%d\n\n", resp1.StatusCode)
	}

	// Test 2: 2 criteria with Priority 0 and 1 (first will be omitted due to omitempty)
	// This simulates what the current test does
	fmt.Println("=== Test 2: 2 criteria with Priority 0 and 1 (0 gets omitted) ===")
	updateReq2 := &mobile_device_groups.RequestMobileDeviceGroup{
		Name:    groupName + "-v2",
		IsSmart: true,
		Site: &shared.SharedResourceSite{ID: -1, Name: "None"},
		Criteria: &mobile_device_groups.CriteriaContainer{
			Size: 2,
			Criterion: []shared.SharedSubsetCriteria{
				{Name: "Model", Priority: 0, AndOr: "and", SearchType: "like", Value: "iPhone"},
				{Name: "Device Name", Priority: 1, AndOr: "and", SearchType: "like", Value: "test"},
			},
		},
	}
	xmlBytes2, _ := xml.MarshalIndent(updateReq2, "", "  ")
	fmt.Printf("XML:\n%s\n", string(xmlBytes2))
	_, resp2, err2 := jamfClient.ClassicMobileDeviceGroups.UpdateByID(ctx, groupID, updateReq2)
	if err2 != nil {
		fmt.Printf("Test2 FAILED (expected): %v\n\n", err2)
	} else {
		fmt.Printf("Test2 PASSED: status=%d\n\n", resp2.StatusCode)
	}

	// Reset group back to 1 criterion
	resetReq := &mobile_device_groups.RequestMobileDeviceGroup{
		Name:    groupName,
		IsSmart: true,
		Site: &shared.SharedResourceSite{ID: -1, Name: "None"},
		Criteria: &mobile_device_groups.CriteriaContainer{
			Size: 1,
			Criterion: []shared.SharedSubsetCriteria{
				{Name: "Model", Priority: 0, AndOr: "and", SearchType: "like", Value: "iPhone"},
			},
		},
	}
	jamfClient.ClassicMobileDeviceGroups.UpdateByID(ctx, groupID, resetReq)

	// Test 3: UpdateByName with 2 criteria (Priority 1 and 2)
	fmt.Println("=== Test 3: UpdateByName with 2 criteria (Priority 1 and 2) ===")
	updateReq3 := &mobile_device_groups.RequestMobileDeviceGroup{
		Name:    groupName + "-v3",
		IsSmart: true,
		Site: &shared.SharedResourceSite{ID: -1, Name: "None"},
		Criteria: &mobile_device_groups.CriteriaContainer{
			Size: 2,
			Criterion: []shared.SharedSubsetCriteria{
				{Name: "Model", Priority: 1, AndOr: "and", SearchType: "like", Value: "iPhone"},
				{Name: "Device Name", Priority: 2, AndOr: "and", SearchType: "like", Value: "test"},
			},
		},
	}
	_, resp3, err3 := jamfClient.ClassicMobileDeviceGroups.UpdateByName(ctx, groupName, updateReq3)
	if err3 != nil {
		fmt.Printf("Test3 FAILED: %v\n\n", err3)
	} else {
		fmt.Printf("Test3 PASSED: status=%d\n\n", resp3.StatusCode)
	}
}
