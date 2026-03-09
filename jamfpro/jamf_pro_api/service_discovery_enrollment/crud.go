package service_discovery_enrollment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the service discovery enrollment-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-service-discovery-enrollment-well-known-settings
	ServiceDiscoveryEnrollment struct {
		client transport.HTTPClient
	}
)

func NewServiceDiscoveryEnrollment(client transport.HTTPClient) *ServiceDiscoveryEnrollment {
	return &ServiceDiscoveryEnrollment{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Service Discovery Enrollment Operations
// -----------------------------------------------------------------------------

// GetV1 returns all well-known service discovery settings.
// URL: GET /api/v1/service-discovery-enrollment/well-known-settings
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-service-discovery-enrollment-well-known-settings
func (s *ServiceDiscoveryEnrollment) GetV1(ctx context.Context) (*WellKnownSettingsResponseV1, *resty.Response, error) {
	var result WellKnownSettingsResponseV1

	endpoint := constants.EndpointJamfProWellKnownSettingsV1

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
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
func (s *ServiceDiscoveryEnrollment) UpdateV1(ctx context.Context, request *WellKnownSettingsResponseV1) (*WellKnownSettingsResponseV1, *resty.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	endpoint := constants.EndpointJamfProWellKnownSettingsV1

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, nil)
	if err != nil {
		return nil, resp, err
	}

	if resp != nil && resp.StatusCode() == http.StatusNoContent {
		return nil, resp, nil
	}

	return nil, resp, nil
}
