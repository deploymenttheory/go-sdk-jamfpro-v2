package client_checkin

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/services/shared"
)

type (
	// ClientCheckinServiceInterface defines the interface for client check-in settings (singleton).
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-check-in
	ClientCheckinServiceInterface interface {
		// GetV3 returns the current client check-in settings (Get Check-In).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-check-in
		GetV3(ctx context.Context) (*ResourceClientCheckinSettings, *interfaces.Response, error)

		// UpdateV3 updates the client check-in settings (Update Check-In).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v3-check-in
		UpdateV3(ctx context.Context, settings *ResourceClientCheckinSettings) (*ResourceClientCheckinSettings, *interfaces.Response, error)
	}

	// Service handles communication with the client check-in endpoint.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-check-in
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ ClientCheckinServiceInterface = (*Service)(nil)

func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// GetV3 returns the current client check-in settings.
// URL: GET /api/v3/check-in
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-check-in
func (s *Service) GetV3(ctx context.Context) (*ResourceClientCheckinSettings, *interfaces.Response, error) {
	var result ResourceClientCheckinSettings
	resp, err := s.client.Get(ctx, EndpointClientCheckinV3, nil, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// UpdateV3 updates the client check-in settings.
// URL: PUT /api/v3/check-in
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v3-check-in
func (s *Service) UpdateV3(ctx context.Context, settings *ResourceClientCheckinSettings) (*ResourceClientCheckinSettings, *interfaces.Response, error) {
	if settings == nil {
		return nil, nil, fmt.Errorf("settings is required")
	}
	var result ResourceClientCheckinSettings
	resp, err := s.client.Put(ctx, EndpointClientCheckinV3, settings, shared.JSONHeaders(), &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}
