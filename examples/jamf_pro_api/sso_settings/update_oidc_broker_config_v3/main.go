package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/jamf_pro_api/sso_settings"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
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
	if len(os.Args) != 3 {
		log.Fatal("usage: example clientconfig.json broker-request.json")
	}
	body, err := os.ReadFile(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	var request sso_settings.RequestOidcBrokerConfig
	if err := json.Unmarshal(body, &request); err != nil {
		log.Fatal(err)
	}
	result, _, err := sdk.JamfProAPI.SsoSettings.UpdateOidcBrokerConfigV3(ctx, &request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.ID)
}
