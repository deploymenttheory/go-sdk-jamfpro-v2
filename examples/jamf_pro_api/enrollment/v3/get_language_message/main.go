package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Get language message for English (en)
	languageCode := "en"

	result, resp, err := client.Enrollment.GetLanguageMessageV3(context.Background(), languageCode)
	if err != nil {
		log.Fatalf("Failed to get language message: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Language Code: %s\n", result.LanguageCode)
	fmt.Printf("Language Name: %s\n", result.Name)
	fmt.Printf("Title: %s\n", result.Title)
	fmt.Printf("Login Description: %s\n", result.LoginDescription)
	fmt.Printf("Username Label: %s\n", result.Username)
	fmt.Printf("Password Label: %s\n", result.Password)
	fmt.Printf("Login Button: %s\n", result.LoginButton)
}
