package client_checkin

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// ClientCheckinServiceInterface defines the interface for client check-in settings (singleton).
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-check-in
	ClientCheckinServiceInterface interface {
		// GetV3 returns the current client check-in settings (Get Client Check-In settings).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-check-in
		GetV3(ctx context.Context) (*ResourceClientCheckinSettings, *interfaces.Response, error)

		// UpdateV3 updates the client check-in settings (Update Client Check-In object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v3-check-in
		UpdateV3(ctx context.Context, request *ResourceClientCheckinSettings) (*ResourceClientCheckinSettings, *interfaces.Response, error)

		// GetHistoryV3 returns the client check-in history object (Get Client Check-In history object).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-check-in-history
		GetHistoryV3(ctx context.Context) (*ResourceClientCheckinHistory, *interfaces.Response, error)

		// AddHistoryNoteV3 adds a note to the client check-in history (Add a Note to Client Check-In History).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-check-in-history
		AddHistoryNoteV3(ctx context.Context, request *RequestClientCheckinHistoryNote) (*interfaces.Response, error)
	}

	// Service handles communication with the client check-in-related methods of the Jamf Pro API.
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

// -----------------------------------------------------------------------------
// Jamf Pro API - Client Check-In Operations
// -----------------------------------------------------------------------------

// GetV3 returns the current client check-in settings.
// URL: GET /api/v3/check-in
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-check-in
func (s *Service) GetV3(ctx context.Context) (*ResourceClientCheckinSettings, *interfaces.Response, error) {
	var result ResourceClientCheckinSettings

	endpoint := EndpointClientCheckinV3

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

// GetHistoryV3 returns the client check-in history object.
// URL: GET /api/v3/check-in/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v3-check-in-history
func (s *Service) GetHistoryV3(ctx context.Context) (*ResourceClientCheckinHistory, *interfaces.Response, error) {
	var result ResourceClientCheckinHistory
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Get(ctx, EndpointClientCheckinHistoryV3, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}
	return &result, resp, nil
}

// AddHistoryNoteV3 adds a note to the client check-in history.
// URL: POST /api/v3/check-in/history
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v3-check-in-history
func (s *Service) AddHistoryNoteV3(ctx context.Context, request *RequestClientCheckinHistoryNote) (*interfaces.Response, error) {
	if request == nil {
		return nil, fmt.Errorf("request is required")
	}
	headers := map[string]string{"Accept": mime.ApplicationJSON, "Content-Type": mime.ApplicationJSON}
	resp, err := s.client.Post(ctx, EndpointClientCheckinHistoryV3, request, headers, nil)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// UpdateV3 updates the client check-in settings.
// URL: PUT /api/v3/check-in
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v3-check-in
func (s *Service) UpdateV3(ctx context.Context, request *ResourceClientCheckinSettings) (*ResourceClientCheckinSettings, *interfaces.Response, error) {
	if request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}

	var result ResourceClientCheckinSettings

	endpoint := EndpointClientCheckinV3

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, request, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
