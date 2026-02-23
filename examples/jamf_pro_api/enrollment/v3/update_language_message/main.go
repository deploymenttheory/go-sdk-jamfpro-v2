package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/jamf_pro_api/enrollment"
)

func main() {
	client, err := jamfpro.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Update language message for English (en)
	languageCode := "en"

	request := &enrollment.ResourceEnrollmentLanguage{
		LanguageCode:      languageCode,
		Name:              "English",
		Title:             "Device Enrollment",
		LoginDescription:  "Please enter your credentials to enroll your device",
		Username:          "Username",
		Password:          "Password",
		LoginButton:       "Sign In",
	}

	result, resp, err := client.Enrollment.UpdateLanguageMessageV3(context.Background(), languageCode, request)
	if err != nil {
		log.Fatalf("Failed to update language message: %v", err)
	}

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Updated Language Code: %s\n", result.LanguageCode)
	fmt.Printf("Updated Title: %s\n", result.Title)
	fmt.Println("Language message updated successfully")
}
