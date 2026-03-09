package engage

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/transport"
	"github.com/deploymenttheory/go-sdk-jamfpro-v2/jamfpro/constants"
	"resty.dev/v3"
)

type (
	// EngageServiceInterface defines the interface for Engage operations.
	//
	// Note: This feature is deprecated in Jamf Pro v11.21.0 and later.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/v11.20.0/reference/get_v2-engage
	EngageServiceInterface interface {
		// GetV2 retrieves the Engage settings from the Jamf Pro server.
		//
		// Note: This feature is deprecated in Jamf Pro v11.21.0 and later.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/v11.20.0/reference/get_v2-engage
		GetV2(ctx context.Context) (*ResourceEngageSettings, *resty.Response, error)

		// UpdateV2 updates the Engage settings on the Jamf Pro server.
		//
		// Note: This feature is deprecated in Jamf Pro v11.21.0 and later.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/v11.20.0/reference/put_v2-engage
		UpdateV2(ctx context.Context, settings *ResourceEngageSettings) (*ResourceEngageSettings, *resty.Response, error)

		// GetHistoryV2 returns the history object for Engage settings.
		//
		// Supports optional RSQL filtering and pagination via rsqlQuery
		// (keys: filter, sort, page, page-size).
		//
		// Note: This feature is deprecated in Jamf Pro v11.21.0 and later.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/v11.20.0/reference/get_v2-engage-history
		GetHistoryV2(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error)

		// AddHistoryNotesV2 adds notes to the Engage settings history.
		//
		// Note: This feature is deprecated in Jamf Pro v11.21.0 and later.
		//
		// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/v11.20.0/reference/post_v2-engage-history
		AddHistoryNotesV2(ctx context.Context, req *RequestAddHistoryNotes) (*ResponseAddHistoryNotes, *resty.Response, error)
	}

	// Service handles communication with the Engage-related methods of the Jamf Pro API.
	//
	// Note: This feature is deprecated in Jamf Pro v11.21.0 and later.
	//
	// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/v11.20.0/reference/get_v2-engage
	Engage struct {
		client transport.HTTPClient
	}
)

var _ EngageServiceInterface = (*Engage)(nil)

// NewService returns a new Engage Service backed by the provided HTTP client.
func NewEngage(client transport.HTTPClient) *Engage {
	return &Engage{client: client}
}

// -----------------------------------------------------------------------------
// Jamf Pro API - Engage Operations
// -----------------------------------------------------------------------------

// GetV2 retrieves the Engage settings from the Jamf Pro server.
// URL: GET /api/v2/engage
// Note: This feature is deprecated in Jamf Pro v11.21.0 and later.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/v11.20.0/reference/get_v2-engage
func (s *Engage) GetV2(ctx context.Context) (*ResourceEngageSettings, *resty.Response, error) {
	endpoint := constants.EndpointJamfProEngageV2

	var result ResourceEngageSettings

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, nil, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// UpdateV2 updates the Engage settings on the Jamf Pro server.
// URL: PUT /api/v2/engage
// Note: This feature is deprecated in Jamf Pro v11.21.0 and later.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/v11.20.0/reference/put_v2-engage
func (s *Engage) UpdateV2(ctx context.Context, settings *ResourceEngageSettings) (*ResourceEngageSettings, *resty.Response, error) {
	if settings == nil {
		return nil, nil, fmt.Errorf("settings cannot be nil")
	}

	endpoint := constants.EndpointJamfProEngageV2

	var result ResourceEngageSettings

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
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
// Note: This feature is deprecated in Jamf Pro v11.21.0 and later.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/v11.20.0/reference/get_v2-engage-history
func (s *Engage) GetHistoryV2(ctx context.Context, rsqlQuery map[string]string) (*HistoryResponse, *resty.Response, error) {
	endpoint := fmt.Sprintf("%s/history", constants.EndpointJamfProEngageV2)

	var result HistoryResponse

	headers := map[string]string{
		"Accept": constants.ApplicationJSON,
	}

	resp, err := s.client.Get(ctx, endpoint, rsqlQuery, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}

// AddHistoryNotesV2 adds notes to the Engage settings history.
// URL: POST /api/v2/engage/history
// Note: This feature is deprecated in Jamf Pro v11.21.0 and later.
// Jamf Pro API docs: https://developer.jamf.com/jamf-pro/v11.20.0/reference/post_v2-engage-history
func (s *Engage) AddHistoryNotesV2(ctx context.Context, req *RequestAddHistoryNotes) (*ResponseAddHistoryNotes, *resty.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("request body is required")
	}
	if req.Note == "" {
		return nil, nil, fmt.Errorf("note is required")
	}

	endpoint := fmt.Sprintf("%s/history", constants.EndpointJamfProEngageV2)

	var result ResponseAddHistoryNotes

	headers := map[string]string{
		"Accept":       constants.ApplicationJSON,
		"Content-Type": constants.ApplicationJSON,
	}

	resp, err := s.client.Post(ctx, endpoint, req, headers, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
