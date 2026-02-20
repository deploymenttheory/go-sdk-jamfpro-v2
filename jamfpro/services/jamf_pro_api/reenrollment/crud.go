package reenrollment

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ReenrollmentServiceInterface defines the interface for re-enrollment settings operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-reenrollment
	ReenrollmentServiceInterface interface {
		// Get retrieves re-enrollment settings (Get Re-enrollment Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-reenrollment
		Get(ctx context.Context) (*ResourceReenrollmentSettings, *interfaces.Response, error)

		// Update updates re-enrollment settings (Update Re-enrollment Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-reenrollment
		Update(ctx context.Context, request *ResourceReenrollmentSettings) (*ResourceReenrollmentSettings, *interfaces.Response, error)
	}

	// Service handles communication with the re-enrollment settings methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-reenrollment
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ReenrollmentServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Re-enrollment Settings Operations
// -----------------------------------------------------------------------------

// Get retrieves re-enrollment settings.
// URL: GET /api/v1/reenrollment
func (s *Service) Get(ctx context.Context) (*ResourceReenrollmentSettings, *interfaces.Response, error) {
	var result ResourceReenrollmentSettings

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, EndpointReenrollmentV1, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// Update updates re-enrollment settings.
// URL: PUT /api/v1/reenrollment
func (s *Service) Update(ctx context.Context, request *ResourceReenrollmentSettings) (*ResourceReenrollmentSettings, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceReenrollmentSettings

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, EndpointReenrollmentV1, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
