package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/classic_api/vpp_assignments"
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

	vppAssignmentID := 1 // Replace with the desired VPP assignment ID to update
	vppAdminAccountID := 1
	vppAdminAccountName := "My VPP Account"

	updateReq := &vpp_assignments.Resource{
		General: vpp_assignments.SubsetGeneral{
			Name:                "go-sdk-v2-vpp-assignment-updated",
			VPPAdminAccountID:   vppAdminAccountID,
			VPPAdminAccountName: vppAdminAccountName,
		},
		Scope: vpp_assignments.SubsetScope{
			AllJSSUsers: false,
		},
	}

	_, _, err = jamfClient.ClassicAPI.VppAssignments.UpdateByID(context.Background(), vppAssignmentID, updateReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("VPP Assignment ID %d updated successfully\n", vppAssignmentID)
}
