package health_check

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the health check-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-health-check
	HealthCheck struct {
		client client.Client
	}
)

func NewHealthCheck(client client.Client) *HealthCheck {
	return &HealthCheck{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Health Check Operations
// -----------------------------------------------------------------------------

// GetV1 returns whether the Jamf Pro API is healthy.
// URL: GET /api/v1/health-check
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-health-check
func (s *HealthCheck) GetV1(ctx context.Context) (bool, *resty.Response, error) {
	endpoint := constants.EndpointJamfProHealthCheckV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		Get(endpoint)
	if err != nil {
		return false, resp, err
	}

	return resp != nil && resp.StatusCode() == 200, resp, nil
}

// GetHealthStatusV1 returns request acceptance ratios for each concurrency group and time window.
// URL: GET /api/v1/health-status
// Only available in Jamf Cloud; returns 404 on non-cloud nodes.
func (s *HealthCheck) GetHealthStatusV1(ctx context.Context) (*ResourceHealthStatus, *resty.Response, error) {
	endpoint := constants.EndpointJamfProHealthStatusV1

	var result ResourceHealthStatus

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
