package time_zones

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// TimeZonesServiceInterface defines the interface for time zone operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-time-zones
	TimeZonesServiceInterface interface {
		// ListV1 returns all available time zones (Get Time Zones).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-time-zones
		ListV1(ctx context.Context) ([]ResourceTimeZone, *interfaces.Response, error)
	}

	// Service handles communication with the time zones-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-time-zones
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ TimeZonesServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Time Zones Operations
// -----------------------------------------------------------------------------

// ListV1 returns all available time zones.
// URL: GET /api/v1/time-zones
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-time-zones
func (s *Service) ListV1(ctx context.Context) ([]ResourceTimeZone, *interfaces.Response, error) {
	var result []ResourceTimeZone

	endpoint := EndpointTimeZonesV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return result, resp, nil
}
