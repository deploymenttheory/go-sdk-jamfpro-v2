package service_discovery_enrollment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ServiceDiscoveryEnrollmentServiceInterface defines the interface for service discovery enrollment operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-service-discovery-enrollment-well-known-settings
	ServiceDiscoveryEnrollmentServiceInterface interface {
		// GetV1 returns all well-known service discovery settings (Get Service Discovery Enrollment Well-Known Settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-service-discovery-enrollment-well-known-settings
		GetV1(ctx context.Context) (*WellKnownSettingsResponseV1, *interfaces.Response, error)

		// UpdateV1 updates the enrollment types for all organizations (Update Service Discovery Enrollment Well-Known Settings).
		// Returns nil for the result when the API responds with 204 No Content.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-service-discovery-enrollment-well-known-settings
		UpdateV1(ctx context.Context, request *WellKnownSettingsResponseV1) (*WellKnownSettingsResponseV1, *interfaces.Response, error)
	}

	// Service handles communication with the service discovery enrollment-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-service-discovery-enrollment-well-known-settings
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ServiceDiscoveryEnrollmentServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Service Discovery Enrollment Operations
// -----------------------------------------------------------------------------

// GetV1 returns all well-known service discovery settings.
// URL: GET /api/v1/service-discovery-enrollment/well-known-settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-service-discovery-enrollment-well-known-settings
func (s *Service) GetV1(ctx context.Context) (*WellKnownSettingsResponseV1, *interfaces.Response, error) {
	var result WellKnownSettingsResponseV1

	endpoint := EndpointWellKnownSettingsV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateV1 updates the enrollment types for all organizations.
// URL: PUT /api/v1/service-discovery-enrollment/well-known-settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-service-discovery-enrollment-well-known-settings
func (s *Service) UpdateV1(ctx context.Context, request *WellKnownSettingsResponseV1) (*WellKnownSettingsResponseV1, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := EndpointWellKnownSettingsV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, nil)
	if err != nil {
		return nil, resp, err
	}

	if resp != nil && resp.StatusCode == http.StatusNoContent {
		return nil, resp, nil
	}

	return nil, resp, nil
}
