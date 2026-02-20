package startup_status

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// StartupStatusServiceInterface defines the interface for startup status operations (read-only).
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_startup-status
	StartupStatusServiceInterface interface {
		// GetStartupStatusV1 returns the Jamf Pro server startup status.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_startup-status
		GetStartupStatusV1(ctx context.Context) (*ResourceStartupStatusV1, *interfaces.Response, error)
	}

	// Service handles communication with the startup status-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_startup-status
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ StartupStatusServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Startup Status Operations
// -----------------------------------------------------------------------------

// GetStartupStatusV1 returns the Jamf Pro server startup status.
// URL: GET /api/startup-status
// https://developer.jamf.com/jamf-pro/reference/get_startup-status
func (s *Service) GetStartupStatusV1(ctx context.Context) (*ResourceStartupStatusV1, *interfaces.Response, error) {
	var result ResourceStartupStatusV1

	endpoint := EndpointStartupStatus

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
