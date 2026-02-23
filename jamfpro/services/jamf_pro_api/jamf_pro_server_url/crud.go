package jamf_pro_server_url

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// JamfProServerURLServiceInterface defines the interface for Jamf Pro server URL operations.
	//
	// Manages the Jamf Pro server URL and unsecured enrollment URL settings.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-server-url
	JamfProServerURLServiceInterface interface {
		// GetV1 retrieves the Jamf Pro server URL settings.
		//
		// Returns the configured Jamf Pro server URL and unsecured enrollment URL.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-server-url
		GetV1(ctx context.Context) (*ResourceJamfProServerURL, *interfaces.Response, error)

		// UpdateV1 updates the Jamf Pro server URL settings.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v1-jamf-pro-server-url
		UpdateV1(ctx context.Context, request *ResourceJamfProServerURL) (*ResourceJamfProServerURL, *interfaces.Response, error)
	}

	// Service handles communication with the Jamf Pro server URL-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-server-url
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ JamfProServerURLServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetV1 retrieves the Jamf Pro server URL settings.
// URL: GET /api/v1/jamf-pro-server-url
// https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-server-url
func (s *Service) GetV1(ctx context.Context) (*ResourceJamfProServerURL, *interfaces.Response, error) {
	var result ResourceJamfProServerURL

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, EndpointJamfProServerURLV1, nil, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get Jamf Pro server URL settings: %w", err)
	}

	return &result, resp, nil
}

// UpdateV1 updates the Jamf Pro server URL settings.
// URL: PUT /api/v1/jamf-pro-server-url
// https://developer.jamf.com/jamf-pro/reference/put_v1-jamf-pro-server-url
func (s *Service) UpdateV1(ctx context.Context, request *ResourceJamfProServerURL) (*ResourceJamfProServerURL, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceJamfProServerURL

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, EndpointJamfProServerURLV1, request, headers, &result)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update Jamf Pro server URL settings: %w", err)
	}

	return &result, resp, nil
}
