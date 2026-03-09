package devices

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

type (
	// DevicesServiceInterface defines the interface for devices operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-devices-id-groups
	DevicesServiceInterface interface {
		// GetGroupsV1 returns a list of groups that the specified device belongs to.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-devices-id-groups
		GetGroupsV1(ctx context.Context, id string) ([]ResourceGroup, *resty.Response, error)
	}

	// Service handles communication with the devices-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-devices-id-groups
	Devices struct {
		client interfaces.HTTPClient
	}
)

var _ DevicesServiceInterface = (*Devices)(nil)

// NewService returns a new devices Service backed by the provided HTTP client.
func NewDevices(client interfaces.HTTPClient) *Devices {
	return &Devices{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Devices Operations
// -----------------------------------------------------------------------------

// GetGroupsV1 returns a list of groups that the specified device belongs to.
// URL: GET /api/v1/devices/{id}/groups
// Response: array of objects with id and name.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-devices-id-groups
func (s *Devices) GetGroupsV1(ctx context.Context, id string) ([]ResourceGroup, *resty.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("id is required")
	}

	endpoint := fmt.Sprintf("%s/%s/groups", EndpointDevicesV1, id)

	var result []ResourceGroup

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get groups for device ID %s: %w", id, err)
	}

	return result, resp, nil
}
