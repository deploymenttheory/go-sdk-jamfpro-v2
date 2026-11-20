package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/jamf_pro_system_initialization"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("provide clientconfig.json path")
	}
	auth, err := jamfpro.LoadAuthConfigFromFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	sdk, err := jamfpro.NewClient(auth)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	body, err := os.ReadFile(os.Getenv("FEDRAMP_REQUEST_FILE"))
	if err != nil {
		log.Fatal(err)
	}
	var request jamf_pro_system_initialization.RequestFedRampInitializeV2
	if err := json.Unmarshal(body, &request); err != nil {
		log.Fatal(err)
	}
	result, err := sdk.JamfProAPI.JamfProSystemInitialization.InitializeFedRampV2(ctx, &request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", result)
}
