package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
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

	vppAssignmentID := 1 // Replace with the desired VPP assignment ID
	assignment, _, err := jamfClient.ClassicVPPAssignments.GetByID(context.Background(), vppAssignmentID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	assignmentXML, err := xml.MarshalIndent(assignment, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling VPP assignment data: %v", err)
	}
	fmt.Printf("VPP Assignment ID %d:\n%s\n", vppAssignmentID, string(assignmentXML))
}
