package engage

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/interfaces"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/mime"
)

type (
	// EngageServiceInterface defines the interface for Engage operations.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-engage
	EngageServiceInterface interface {
		// GetV2 retrieves the Engage settings from the Jamf Pro server.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-engage
		GetV2(ctx context.Context) (*ResourceEngageSettings, *interfaces.Response, error)

		// UpdateV2 updates the Engage settings on the Jamf Pro server.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/put_v2-engage
		UpdateV2(ctx context.Context, settings *ResourceEngageSettings) (*ResourceEngageSettings, *interfaces.Response, error)

		// GetHistoryV2 returns the history object for Engage settings.
		//
		// Supports optional RSQL filtering and pagination via rsqlQuery
		// (keys: filter, sort, page, page-size).
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-engage-history
		GetHistoryV2(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error)

		// AddHistoryNotesV2 adds notes to the Engage settings history.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/post_v2-engage-history
		AddHistoryNotesV2(ctx context.Context, req *RequestAddHistoryNotes) (*ResponseAddHistoryNotes, *interfaces.Response, error)
	}

	// Service handles communication with the Engage-related methods of the Jamf Pro API.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/reference/get_v2-engage
	Service struct {
		client interfaces.HTTPClient
	}
)

var _ EngageServiceInterface = (*Service)(nil)

// NewService returns a new Engage Service backed by the provided HTTP client.
func NewService(client interfaces.HTTPClient) *Service {
	return &Service{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Engage Operations
// -----------------------------------------------------------------------------

// GetV2 retrieves the Engage settings from the Jamf Pro server.
// URL: GET /api/v2/engage
// https://developer.jamf.com/jamf-pro/reference/get_v2-engage
func (s *Service) GetV2(ctx context.Context) (*ResourceEngageSettings, *interfaces.Response, error) {
	var result ResourceEngageSettings

	endpoint := EndpointEngageV2

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

// UpdateV2 updates the Engage settings on the Jamf Pro server.
// URL: PUT /api/v2/engage
// https://developer.jamf.com/jamf-pro/reference/put_v2-engage
func (s *Service) UpdateV2(ctx context.Context, settings *ResourceEngageSettings) (*ResourceEngageSettings, *interfaces.Response, error) {
	if settings == nil {
		return nil, nil, fmt.Errorf("settings cannot be nil")
	}

	endpoint := EndpointEngageV2

	var result ResourceEngageSettings

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Put(ctx, endpoint, settings, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// GetHistoryV2 returns the history object for Engage settings.
// URL: GET /api/v2/engage/history
// rsqlQuery supports: filter (RSQL), sort, page, page-size (all optional).
// https://developer.jamf.com/jamf-pro/reference/get_v2-engage-history
func (s *Service) GetHistoryV2(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *interfaces.Response, error) {
	endpoint := fmt.Sprintf("%s/history", EndpointEngageV2)

	var result HistoryResponse

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// AddHistoryNotesV2 adds notes to the Engage settings history.
// URL: POST /api/v2/engage/history
// https://developer.jamf.com/jamf-pro/reference/post_v2-engage-history
func (s *Service) AddHistoryNotesV2(ctx context.Context, req *RequestAddHistoryNotes) (*ResponseAddHistoryNotes, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request body is required")
	}
	if req.Note == "" {
		return nil, nil, fmt.Errorf("note is required")
	}

	endpoint := fmt.Sprintf("%s/history", EndpointEngageV2)

	var result ResponseAddHistoryNotes

	headers := map[string]string{
		"Accept":       mime.ApplicationJSON,
		"Content-Type": mime.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

