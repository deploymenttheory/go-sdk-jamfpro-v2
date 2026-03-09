package startup_status

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// StartupStatusServiceInterface defines the interface for startup status operations (read-only).
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_startup-status
	StartupStatusServiceInterface interface {
		// GetV1 returns the Jamf Pro server startup status.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_startup-status
		GetV1(ctx context.Context) (*ResourceStartupStatusV1, *resty.Response, error)
	}

	// Service handles communication with the startup status-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_startup-status
	StartupStatus struct {
		client transport.HTTPClient
	}
)

var _ StartupStatusServiceInterface = (*StartupStatus)(nil)

func NewStartupStatus(client transport.HTTPClient) *StartupStatus {
	return &StartupStatus{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Startup Status Operations
// -----------------------------------------------------------------------------

// GetV1 returns the Jamf Pro server startup status.
// URL: GET /api/startup-status
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_startup-status
func (s *StartupStatus) GetV1(ctx context.Context) (*ResourceStartupStatusV1, *resty.Response, error) {
	var result ResourceStartupStatusV1

	endpoint := constants.EndpointJamfProStartupStatus

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
