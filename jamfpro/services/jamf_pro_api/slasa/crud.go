package slasa

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
	"resty.dev/v3"
)

// ServiceInterface defines the interface for SLASA (Software License Agreement Service Acceptance) operations.
//
// SLASA is required for managed software updates in Jamf Pro. Administrators must accept the
// Software License Agreement before using managed software update features.
//
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-slasa
type ServiceInterface interface {
	// GetStatusV1 retrieves the current SLASA acceptance status.
	//
	// Returns whether the Software License Agreement has been accepted or not.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-slasa
	GetStatusV1(ctx context.Context) (*ResourceSLASAStatus, *resty.Response, error)

	// AcceptV1 accepts the SLASA (Software License Agreement Service Acceptance).
	//
	// Must be called before managed software updates can be used.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v1-slasa
	AcceptV1(ctx context.Context) (*resty.Response, error)
}

type (
	// Service handles communication with the SLASA-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-slasa
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ServiceInterface = (*Service)(nil)

// NewService creates a new SLASA service.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetStatusV1 retrieves the current SLASA acceptance status.
// URL: GET /api/v1/managed-software-updates/slasa
// https://developer.jamf.com/jamf-pro/reference/get_v1-slasa
func (s *Service) GetStatusV1(ctx context.Context) (*ResourceSLASAStatus, *resty.Response, error) {
	endpoint := EndpointSLASAV1

	headers := map[string]string{
		"Accept": mime.ApplicationJSON,
	}

	var result ResourceSLASAStatus
	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get SLASA status: %w", err)
	}

	return &result, resp, nil
}

// AcceptV1 accepts the SLASA (Software License Agreement Service Acceptance).
// URL: POST /api/v1/managed-software-updates/slasa
// https://developer.jamf.com/jamf-pro/reference/post_v1-slasa
func (s *Service) AcceptV1(ctx context.Context) (*resty.Response, error) {
	endpoint := EndpointSLASAV1

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, nil, headers, nil)
	if err != nil {
		return resp, fmt.Errorf("failed to accept SLASA: %w", err)
	}

	return resp, nil
}
