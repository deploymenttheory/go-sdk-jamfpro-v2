package main

import (
	"context"
	"fmt"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
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
	items, _, err := sdk.JamfProAPI.Notifications.ListV1(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if len(items) != 0 {
		log.Fatal("review notifications before dismissing them")
	}
	response, err := sdk.JamfProAPI.Notifications.DeleteV1(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.StatusCode())
}
