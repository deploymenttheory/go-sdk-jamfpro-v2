package health_check

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// HealthCheckServiceInterface defines the interface for health check operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-health-check
	HealthCheckServiceInterface interface {
		// GetV1 returns whether the Jamf Pro API is healthy (Get Health Check).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-health-check
		GetV1(ctx context.Context) (healthy bool, resp *resty.Response, err error)

		// GetHealthStatusV1 returns request acceptance ratios for each concurrency group and time window.
		// Only available in Jamf Cloud; returns 404 on non-cloud nodes.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-health-status
		GetHealthStatusV1(ctx context.Context) (*ResourceHealthStatus, *resty.Response, error)
	}

	// Service handles communication with the health check-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-health-check
	HealthCheck struct {
		client interfaces.HTTPClient
	}
)

var _ HealthCheckServiceInterface = (*HealthCheck)(nil)

func NewHealthCheck(client interfaces.HTTPClient) *HealthCheck {
	return &HealthCheck{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Health Check Operations
// -----------------------------------------------------------------------------

// GetV1 returns whether the Jamf Pro API is healthy.
// URL: GET /api/v1/health-check
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-health-check
func (s *HealthCheck) GetV1(ctx context.Context) (bool, *resty.Response, error) {
	endpoint := EndpointHealthCheckV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return false, resp, err
	}

	return resp != nil && resp.StatusCode() == 200, resp, nil
}

// GetHealthStatusV1 returns request acceptance ratios for each concurrency group and time window.
// URL: GET /api/v1/health-status
// Only available in Jamf Cloud; returns 404 on non-cloud nodes.
func (s *HealthCheck) GetHealthStatusV1(ctx context.Context) (*ResourceHealthStatus, *resty.Response, error) {
	endpoint := EndpointHealthStatusV1

	var result ResourceHealthStatus

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
