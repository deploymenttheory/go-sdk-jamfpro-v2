package slasa

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/client"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// Service handles communication with the SLASA-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-slasa
	Slasa struct {
		client client.Client
	}
)

// NewService creates a new SLASA service.
func NewSlasa(client client.Client) *Slasa {
	return &Slasa{client: client}
}

// GetStatusV1 retrieves the current SLASA acceptance status.
// URL: GET /api/v1/managed-software-updates/slasa
// https://developer.jamf.com/jamf-pro/reference/get_v1-slasa
func (s *Slasa) GetStatusV1(ctx context.Context) (*ResourceSLASAStatus, *resty.Response, error) {
	endpoint := constants.EndpointJamfProSLASAV1

	var result ResourceSLASAStatus

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetResult(&result).
		Get(endpoint)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get SLASA status: %w", err)
	}

	return &result, resp, nil
}

// AcceptV1 accepts the SLASA (Software License Agreement Service Acceptance).
// URL: POST /api/v1/managed-software-updates/slasa
// https://developer.jamf.com/jamf-pro/reference/post_v1-slasa
func (s *Slasa) AcceptV1(ctx context.Context) (*resty.Response, error) {
	endpoint := constants.EndpointJamfProSLASAV1

	resp, err := s.client.NewRequest(ctx).
		SetHeader("Accept", constants.ApplicationJSON).
		SetHeader("Content-Type", constants.ApplicationJSON).
		Post(endpoint)
	if err != nil {
		return resp, fmt.Errorf("failed to accept SLASA: %w", err)
	}

	return resp, nil
}
