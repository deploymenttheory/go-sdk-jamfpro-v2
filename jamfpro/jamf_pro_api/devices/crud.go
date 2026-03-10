package devices

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the devices-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-devices-id-groups
	Devices struct {
		client client.Client
	}
)

// NewService returns a new devices Service backed by the provided HTTP client.
func NewDevices(client client.Client) *Devices {
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

	endpoint := fmt.Sprintf("%s/%s/groups", constants.EndpointJamfProDevicesV1, id)

	var result []ResourceGroup

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get groups for device ID %s: %w", id, err)
	}

	return result, resp, nil
}
