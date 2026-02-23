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

	licensedSoftwareName := "Sample Licensed Software" // Replace with the desired licensed software name
	ls, _, err := jamfClient.ClassicLicensedSoftware.GetByName(context.Background(), licensedSoftwareName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	lsXML, err := xml.MarshalIndent(ls, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling licensed software data: %v", err)
	}
	fmt.Printf("Licensed Software %q:\n%s\n", licensedSoftwareName, string(lsXML))
}
