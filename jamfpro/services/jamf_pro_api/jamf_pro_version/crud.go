package jamf_pro_version

import (
	"context"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// JamfProVersionServiceInterface defines the interface for Jamf Pro version (read-only).
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-version
	JamfProVersionServiceInterface interface {
		// GetV1 returns the Jamf Pro server version (Get Jamf Pro Version).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-version
		GetV1(ctx context.Context) (*ResourceJamfProVersion, *interfaces.Response, error)
	}

	// Service handles communication with the Jamf Pro version-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-version
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ JamfProVersionServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Jamf Pro Version Operations
// -----------------------------------------------------------------------------

// GetV1 returns the Jamf Pro server version.
// URL: GET /api/v1/jamf-pro-version
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-version
func (s *Service) GetV1(ctx context.Context) (*ResourceJamfProVersion, *interfaces.Response, error) {
	var result ResourceJamfProVersion

	endpoint := EndpointJamfProVersionV1

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
