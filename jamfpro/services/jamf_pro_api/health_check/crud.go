package health_check

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	HealthCheckServiceInterface interface {
		// GetV1 returns whether the Jamf Pro API is healthy (GET returns 200).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-health-check
		GetV1(ctx context.Context) (healthy bool, resp *interfaces.Response, err error)
	}

	Service struct {
		client interfaces.HTTPClient
	}
)

var _ HealthCheckServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetV1 performs GET /api/v1/health-check. healthy is true when status is 200.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-health-check
func (s *Service) GetV1(ctx context.Context) (bool, *interfaces.Response, error) {
	endpoint := EndpointHealthCheckV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return false, resp, err
	}

	return resp != nil && resp.StatusCode == 200, resp, nil
}
